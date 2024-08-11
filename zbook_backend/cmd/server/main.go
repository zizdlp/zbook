package main

import (
	"C"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	_ "github.com/zizdlp/zbook/statik"
	storage "github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
)
import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

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
	// Parse the timezone from config
	loc, err := time.LoadLocation(config.TIMEZONE)
	if err != nil {
		log.Fatal().Msgf("cannot load timezone: %s", err)
	}
	// Set the logger to use the custom timezone
	zerolog.TimestampFunc = func() time.Time {
		return time.Now().In(loc)
	}
	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "2006-01-02 03:04:05 PM", // 自定义时间格式为包含日期的12小时制
			FormatTimestamp: func(i interface{}) string {
				switch v := i.(type) {
				case time.Time:
					// 使用指定时区格式化时间戳，并应用暗色
					return v.In(loc).Format("2006-01-02 03:04:05 PM")
				case string:
					// 如果时间戳是字符串，可以尝试转换为 time.Time
					if ts, err := time.Parse(time.RFC3339, v); err == nil {
						return ts.In(loc).Format("2006-01-02 03:04:05 PM")
					}
				}
				return fmt.Sprint(i)
			},
		})
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
	runGatewayServer(ctx, waitGroup, config, store, taskDistributor, redisClient, minioClient)
	runGrpcServer(ctx, waitGroup, config, store, taskDistributor, redisClient, minioClient)
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
	waitGroup *errgroup.Group, config util.Config, store db.Store, taskDistributor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client) {

	server, err := gapi.NewServer(config, store, taskDistributor, redisClient, minioClient)
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
	waitGroup *errgroup.Group, config util.Config, store db.Store, taskDistributor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client) {

	server, err := gapi.NewServer(config, store, taskDistributor, redisClient, minioClient)
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
