package main

import (
	"C"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/oschwald/geoip2-golang"
	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	_ "github.com/zizdlp/zbook/statik"
	storage "github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
)
import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-redis/redis"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/minio/minio-go/v7"
	"github.com/rs/zerolog"
	"github.com/zizdlp/zbook/gapi"
	"github.com/zizdlp/zbook/mail"
	"github.com/zizdlp/zbook/pb"
	"github.com/zizdlp/zbook/worker"
	"github.com/zizdlp/zbook/wsserver"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}
	if config.Environment == "development" {
		// 将 log 包的输出设置为 Zerolog 的 ConsoleWriter
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Msgf("cannot connect to db: %s", err)
	}
	go wsserver.ListenWebSocket(connPool)

	runDBMigration(config.MigrationURL, config.DBSource)
	store := db.NewStore(connPool)

	geoClient, err := geoip2.Open("./GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal().Msgf("cannot connect to geoDB: %s", err)
	}
	defer geoClient.Close()

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}
	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	redisClient := redis.NewClient(&redis.Options{
		Addr: config.RedisAddress, // Redis 服务器地址和端口
	})
	minioClient, err := storage.GetMinioClient()
	if err != nil {
		log.Fatal().Msgf("cannot connect to minio: %s", err)
	}
	waitGroup, ctx := errgroup.WithContext(ctx)

	runTaskProcessor(ctx, waitGroup, config, redisOpt, store)
	runGatewayServer(ctx, waitGroup, config, store, taskDistributor, redisClient, minioClient, geoClient)
	runGrpcServer(ctx, waitGroup, config, store, taskDistributor, redisClient, minioClient, geoClient)
	wsserver.WebSocketServer(ctx, waitGroup, config)
	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}

func runTaskProcessor(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config util.Config, redisOpt asynq.RedisClientOpt, store db.Store) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, store, mailer)
	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown task processor")

		taskProcessor.Shutdown()
		log.Info().Msg("task processor is stopped")

		return nil
	})
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Msgf("cannot create new migrate instance: %s", err)
	}
	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Msgf("failed to run migrate up: %s", err)
	}
	log.Info().Msg("db migrate succ")
}
func runGrpcServer(ctx context.Context,
	waitGroup *errgroup.Group, config util.Config, store db.Store, taskDistributor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client, geoClient *geoip2.Reader) {

	server, err := gapi.NewServer(config, store, taskDistributor, redisClient, minioClient, geoClient)
	if err != nil {
		log.Fatal().Msgf("cannot create server: %s", err)
	}
	grpcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
	grpcServer := grpc.NewServer(grpcLogger)
	pb.RegisterZBookVerificationServer(grpcServer, server)
	pb.RegisterZBookRepoRelationServer(grpcServer, server)
	pb.RegisterZBookCommentRelationServer(grpcServer, server)
	pb.RegisterZBookNotificationServer(grpcServer, server)
	pb.RegisterZBookAdminServer(grpcServer, server)
	pb.RegisterZBookCommentServer(grpcServer, server)
	pb.RegisterZBookRepoServer(grpcServer, server)
	pb.RegisterZBookMarkdownServer(grpcServer, server)
	pb.RegisterZBookUserServer(grpcServer, server)
	pb.RegisterZBookOAuthServer(grpcServer, server)
	pb.RegisterZBookFollowServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Msgf("cannot create listener: %s", err)
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start gRPC server at %s", listener.Addr().String())

		err = grpcServer.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown gRPC server")

		grpcServer.GracefulStop()
		log.Info().Msg("gRPC server is stopped")

		return nil
	})
}

func CustomMatcher(key string) (string, bool) {
	switch key {
	case "X-Envoy-External-Address":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}
}
func runGatewayServer(ctx context.Context,
	waitGroup *errgroup.Group, config util.Config, store db.Store, taskDistributor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client, geoClient *geoip2.Reader) {

	server, err := gapi.NewServer(config, store, taskDistributor, redisClient, minioClient, geoClient)
	if err != nil {
		log.Fatal().Msgf("cannot create server: %s", err)
	}

	options := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	})

	grpcMux := runtime.NewServeMux(options, runtime.WithIncomingHeaderMatcher(CustomMatcher))

	err = pb.RegisterZBookVerificationHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler server")
	}
	err = pb.RegisterZBookRepoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_post server")
	}
	err = pb.RegisterZBookRepoRelationHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_post_relation server")
	}
	err = pb.RegisterZBookCommentHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_comment server")
	}
	err = pb.RegisterZBookCommentRelationHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_comment_relation server")
	}

	err = pb.RegisterZBookNotificationHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler post_notification server")
	}
	err = pb.RegisterZBookAdminHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_admin server")
	}

	err = pb.RegisterZBookMarkdownHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_markdown server")
	}

	err = pb.RegisterZBookUserHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_user server")
	}
	err = pb.RegisterZBookOAuthHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_oauth server")
	}
	err = pb.RegisterZBookFollowHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Msg("cannot register handler ZBook_follow server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", allowCORS(grpcMux))

	// WARNING: static file not exposed for security reason
	// statikFS, err := fs.New()
	// if err != nil {
	// 	log.Fatal().Msg("cannot create statik fs")
	// }
	// swaggerHandler := http.StripPrefix("/statik/", http.FileServer(statikFS))
	// mux.Handle("/statik/", allowCORS(swaggerHandler))

	// fileServer := http.FileServer(http.Dir("/tmp/wiki/"))

	// // Use the mux.Handle() function to register the file server as the handler for
	// // all URL paths that start with "/static/". For matching paths, we strip the
	// // "/static" prefix before the request reaches the file server.
	// mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	httpServer := &http.Server{
		Handler: gapi.HttpLogger(mux),
		Addr:    config.HTTPServerAddress,
	}

	waitGroup.Go(func() error {
		log.Info().Msgf("start HTTP gateway server at %s", httpServer.Addr)
		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("HTTP gateway server failed to serve")
			return err
		}
		return nil
	})
	waitGroup.Go(func() error {
		<-ctx.Done()
		log.Info().Msg("graceful shutdown HTTP gateway server")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown HTTP gateway server")
			return err
		}

		log.Info().Msg("HTTP gateway server is stopped")
		return nil
	})

}

func allowCORS(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			if r.Method == "OPTIONS" && r.Header.Get("Access-Control-Request-Method") != "" {
				preflightHandler(w, r)
				return
			}
		}
		h.ServeHTTP(w, r)
	})
}

// preflightHandler adds the necessary headers in order to serve
// CORS from any origin using the methods "GET", "HEAD", "POST", "PUT", "DELETE"
// We insist, don't do this without consideration in production systems.
func preflightHandler(w http.ResponseWriter, r *http.Request) {
	headers := []string{"Content-Type", "Accept", "Authorization"}
	w.Header().Set("Access-Control-Allow-Headers", strings.Join(headers, ","))
	methods := []string{"GET", "HEAD", "POST", "PUT", "DELETE", "PATCH"}
	w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))
	log.Info().Msgf("preflightHandler at %s", r.URL.Path)
}
