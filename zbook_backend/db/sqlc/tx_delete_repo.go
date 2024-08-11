package db

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteRepoTxParams struct {
	RepoID     int64
	UserID     int64
	AfterDelte func(repoID int64, userID int64) error
}

func (store *SQLStore) DeleteRepoTx(ctx context.Context, arg DeleteRepoTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {

		err := q.DeleteRepo(ctx, arg.RepoID)
		if err != nil {
			return status.Errorf(codes.Internal, "delete repo failed: %s", err)
		}
		return arg.AfterDelte(arg.RepoID, arg.UserID)
	})

	return err
}
