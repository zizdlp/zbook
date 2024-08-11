package db

import (
	"context"
)

type CreateFollowTxParams struct {
	CreateFollowParams
}

type CreateFollowTxResult struct {
	Follow Follow
}

func (store *SQLStore) CreateFollowTx(ctx context.Context, arg CreateFollowTxParams) (CreateFollowTxResult, error) {
	var result CreateFollowTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Follow, err = q.CreateFollow(ctx, arg.CreateFollowParams)
		if err != nil {
			return err
		}

		arg_noti := CreateFollowerNotificationParams{
			UserID:     result.Follow.FollowingID,
			FollowerID: result.Follow.FollowerID,
		}

		_, err = q.CreateFollowerNotification(ctx, arg_noti)
		if err != nil {
			return err
		} else {
			err = q.UpdateUnreadCount(ctx, result.Follow.FollowingID)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return result, err
}
