package db

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteUserTxParams struct {
	UserID     int64
	Username   string
	AfterDelte func(userID int64, username string) error
}

func (store *SQLStore) DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {

		err := q.DeleteUser(ctx, arg.Username)
		if err != nil {
			return status.Errorf(codes.Internal, "delete repo failed: %s", err)
		}
		return arg.AfterDelte(arg.UserID, arg.Username)
	})

	return err
}
