package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store interface {
	Querier
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
	VerifyEmailTx(ctx context.Context, verification_url string) (User, error)
	CreateFollowTx(ctx context.Context, arg CreateFollowTxParams) (CreateFollowTxResult, error)
	DeleteFollowTx(ctx context.Context, arg DeleteFollowTxParams) (DeleteFollowTxResult, error)
	CreateCommentTx(ctx context.Context, arg CreateCommentTxParams) (CreateCommentTxResult, error)
	DeleteRepoTx(ctx context.Context, arg DeleteRepoTxParams) error
	DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error
	CreateSystemNotificationTx(ctx context.Context, arg CreateSystemNotificationTxParams) error
	CreateRepoTx(ctx context.Context, arg CreateRepoTxParams) (CreateRepoTxResult, error)
	ManualSyncRepoTx(ctx context.Context, arg ManualSyncRepoTxParams) error
	ResetPasswordTx(ctx context.Context, arg ResetPasswordTxParams) error
}

// SQLstore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) Store {
	return &SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
