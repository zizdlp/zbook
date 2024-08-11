package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

type TaskDistributor interface {
	DistributeTaskVerifyEmail(
		ctx context.Context,
		payload *PayloadVerifyEmail,
		opts ...asynq.Option,
	) error
	DistributeTaskInviteUser(
		ctx context.Context,
		payload *PayloadInviteUser,
		opts ...asynq.Option,
	) error
	DistributeTaskResetPassword(
		ctx context.Context,
		payload *PayloadResetPassword,
		opts ...asynq.Option,
	) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
}

func NewRedisTaskDistributor(redisOpt asynq.RedisClientOpt) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
	}
}
