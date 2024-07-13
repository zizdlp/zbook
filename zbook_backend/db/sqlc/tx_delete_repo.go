package db

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MarkRepoAsDeletedTxParams struct {
	RepoID int64
	UserID int64
}

func (store *SQLStore) MarkRepoAsDeletedTx(ctx context.Context, arg MarkRepoAsDeletedTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {

		err := q.MarkRepoAsDeleted(ctx, arg.RepoID)
		if err != nil {
			return status.Errorf(codes.Internal, "delete repo failed: %s", err)
		}
		return nil
	})
	return err
}
