package db

import (
	"context"
)

type CreateSystemNotificationTxParams struct {
	CreateSystemNotificationParams
}

func (store *SQLStore) CreateSystemNotificationTx(ctx context.Context, arg CreateSystemNotificationTxParams) error {

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		_, err = q.CreateSystemNotification(ctx, arg.CreateSystemNotificationParams)
		if err != nil {
			return err
		}
		err = q.UpdateUnreadCount(ctx, arg.CreateSystemNotificationParams.UserID)
		if err != nil {
			return err
		}
		return nil
	})

	return err
}
