package gapi

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/minio/minio-go/v7"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb"
	"github.com/zizdlp/zbook/token"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/worker"
)

type Server struct {
	pb.UnimplementedZBookVerificationServer
	pb.UnimplementedZBookRepoRelationServer
	pb.UnimplementedZBookCommentRelationServer
	pb.UnimplementedZBookNotificationServer
	pb.UnimplementedZBookAdminServer
	pb.UnimplementedZBookCommentServer
	pb.UnimplementedZBookMarkdownServer
	pb.UnimplementedZBookRepoServer
	pb.UnimplementedZBookUserServer
	pb.UnimplementedZBookOAuthServer
	pb.UnimplementedZBookFollowServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
	redisClient     *redis.Client
	minioClient     *minio.Client
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
		redisClient:     redisClient,
		minioClient:     minioClient,
	}

	return server, nil
}
