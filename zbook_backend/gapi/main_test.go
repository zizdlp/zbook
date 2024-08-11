package gapi

import (
	"testing"
	"time"

	"github.com/go-redis/redis"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/require"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/worker"
)

func newTestServer(t *testing.T, store db.Store, taskDistrubutor worker.TaskDistributor, redisClient *redis.Client, minioClient *minio.Client) *Server {
	config := util.Config{
		TokenSymmetricKey:   util.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store, taskDistrubutor, redisClient, minioClient)
	require.NoError(t, err)

	return server
}
