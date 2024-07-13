package db

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DeleteFollowTxParams struct {
	DeleteFollowParams
}

type DeleteFollowTxResult struct {
	Follow Follow
}

func (store *SQLStore) DeleteFollowTx(ctx context.Context, arg DeleteFollowTxParams) (DeleteFollowTxResult, error) {
	var result DeleteFollowTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		_, err = q.DeleteFollow(ctx, arg.DeleteFollowParams)
		if err != nil {
			return status.Errorf(codes.Internal, "delete follow failed: %s", err)
		}

		arg_noti := DeleteFollowerNotificationParams{
			UserID:     arg.FollowingID,
			FollowerID: arg.FollowerID,
		}

		_, err = q.DeleteFollowerNotification(ctx, arg_noti)
		if err != nil {
			if errors.Is(err, ErrRecordNotFound) {
				return status.Errorf(codes.NotFound, "follow notification not found: %s", err)
			}
			return status.Errorf(codes.NotFound, "delete follow notification failed: %s", err)
		}

		err = q.UpdateUnreadCount(ctx, arg.FollowingID)
		if err != nil {
			return status.Errorf(codes.Internal, "update unread count failed: %s", err)
		}

		return nil
	})

	return result, err
}
