package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createRandomSystemNotificationByUser(t *testing.T, user User) SystemNotification {
	arg := CreateSystemNotificationParams{
		UserID:      user.UserID,
		Title:       util.RandomString(8),
		Contents:    util.RandomString(32),
		RedirectUrl: pgtype.Text{String: util.RandomString(6), Valid: util.RandomBool()},
	}
	notification, err := testStore.CreateSystemNotification(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, notification)
	require.Equal(t, arg.Contents, notification.Contents)
	require.Equal(t, arg.UserID, notification.UserID)
	return notification
}
func createRandomSystemNotification(t *testing.T) SystemNotification {
	user := createRandomUser(t)
	arg := CreateSystemNotificationParams{
		UserID:      user.UserID,
		Title:       util.RandomString(8),
		Contents:    util.RandomString(32),
		RedirectUrl: pgtype.Text{String: util.RandomString(6), Valid: util.RandomBool()},
	}
	notification, err := testStore.CreateSystemNotification(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, notification)
	require.Equal(t, arg.Contents, notification.Contents)
	require.Equal(t, arg.UserID, notification.UserID)
	return notification
}
func TestCreateRandomSystemNotification(t *testing.T) {
	createRandomSystemNotification(t)
}
func TestMarkSystemNotificationReaded(t *testing.T) {
	notification := createRandomSystemNotification(t)
	require.Equal(t, false, notification.Readed)
	arg := MarkSystemNotificationReadedParams{
		NotiID: notification.NotiID,
		UserID: notification.UserID,
	}
	updated_notification, err := testStore.MarkSystemNotificationReaded(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, true, updated_notification.Readed)
}

func TestListSystemNotification(t *testing.T) {
	user := createRandomUser(t)
	notification_1 := createRandomSystemNotificationByUser(t, user)
	notification_2 := createRandomSystemNotificationByUser(t, user)
	notification_3 := createRandomSystemNotificationByUser(t, user)
	arg := ListSystemNotificationParams{
		UserID: user.UserID,
		Limit:  10,
		Offset: 0,
	}
	notifications, err := testStore.ListSystemNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(notifications))
	require.Equal(t, notification_1.NotiID, notifications[2].NotiID)
	require.Equal(t, notification_2.NotiID, notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, notifications[0].NotiID)

	arg_mark := MarkSystemNotificationReadedParams{
		NotiID: notification_2.NotiID,
		UserID: user.UserID,
	}
	_, err = testStore.MarkSystemNotificationReaded(context.Background(), arg_mark)
	require.NoError(t, err)

	updated_notifications, err := testStore.ListSystemNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(updated_notifications))
	require.Equal(t, notification_2.NotiID, updated_notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, updated_notifications[0].NotiID)
}

func createRandomFollowerNotificationByUser(t *testing.T, user User) FollowerNotification {
	follower := createRandomUser(t)
	arg := CreateFollowerNotificationParams{
		UserID:     user.UserID,
		FollowerID: follower.UserID,
	}
	notification, err := testStore.CreateFollowerNotification(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, notification)
	require.Equal(t, arg.FollowerID, notification.FollowerID)
	require.Equal(t, arg.UserID, notification.UserID)
	return notification
}
func createRandomFollowerNotification(t *testing.T) FollowerNotification {
	user := createRandomUser(t)
	notification := createRandomFollowerNotificationByUser(t, user)
	return notification
}
func TestCreateFollowerSystemNotification(t *testing.T) {
	createRandomFollowerNotification(t)
}

func TestDeleteFollowerSystemNotification(t *testing.T) {
	notification := createRandomFollowerNotification(t)
	arg := DeleteFollowerNotificationParams{
		UserID:     notification.UserID,
		FollowerID: notification.FollowerID,
	}
	updated_notification, err := testStore.DeleteFollowerNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, updated_notification.NotiID, notification.NotiID)
	_, err = testStore.DeleteFollowerNotification(context.Background(), arg)
	require.Error(t, err)
}
func TestMarkFollowerNotificationReaded(t *testing.T) {
	notification := createRandomFollowerNotification(t)
	require.Equal(t, false, notification.Readed)
	arg := MarkFollowerNotificationReadedParams{
		NotiID: notification.NotiID,
		UserID: notification.UserID,
	}
	updated_notification, err := testStore.MarkFollowerNotificationReaded(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, true, updated_notification.Readed)
}

func TestListFollowerNotification(t *testing.T) {
	user := createRandomUser(t)
	notification_1 := createRandomFollowerNotificationByUser(t, user)
	notification_2 := createRandomFollowerNotificationByUser(t, user)
	notification_3 := createRandomFollowerNotificationByUser(t, user)
	arg := ListFollowerNotificationParams{
		UserID: user.UserID,
		Limit:  10,
		Offset: 0,
	}
	notifications, err := testStore.ListFollowerNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(notifications))
	require.Equal(t, notification_1.NotiID, notifications[2].NotiID)
	require.Equal(t, notification_2.NotiID, notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, notifications[0].NotiID)

	arg_delete := DeleteFollowerNotificationParams{
		FollowerID: notification_2.FollowerID,
		UserID:     user.UserID,
	}
	_, err = testStore.DeleteFollowerNotification(context.Background(), arg_delete)
	require.NoError(t, err)

	updated_notifications, err := testStore.ListFollowerNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 2, len(updated_notifications))
	require.Equal(t, notification_1.NotiID, updated_notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, updated_notifications[0].NotiID)
}

func createRandomRepoNotificationByUser(t *testing.T, user User) RepoNotification {
	repo := createUserRandomRepo(t, user)
	arg := CreateRepoNotificationParams{
		UserID: user.UserID,
		RepoID: repo.RepoID,
	}
	notification, err := testStore.CreateRepoNotification(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, notification)
	require.Equal(t, arg.RepoID, notification.RepoID)
	require.Equal(t, arg.UserID, notification.UserID)
	return notification
}
func createRandomRepoNotification(t *testing.T) RepoNotification {
	user := createRandomUser(t)
	notification := createRandomRepoNotificationByUser(t, user)
	return notification
}
func TestCreateRepoNotification(t *testing.T) {
	createRandomRepoNotification(t)
}
func TestMarkRepoNotificationReaded(t *testing.T) {
	notification := createRandomRepoNotification(t)
	require.Equal(t, false, notification.Readed)
	arg := MarkRepoNotificationReadedParams{
		NotiID: notification.NotiID,
		UserID: notification.UserID,
	}
	updated_notification, err := testStore.MarkRepoNotificationReaded(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, true, updated_notification.Readed)
}
func TestListRepoNotification(t *testing.T) {
	user := createRandomUser(t)
	notification_1 := createRandomRepoNotificationByUser(t, user)
	notification_2 := createRandomRepoNotificationByUser(t, user)
	notification_3 := createRandomRepoNotificationByUser(t, user)
	arg := ListRepoNotificationParams{
		UserID: user.UserID,
		Limit:  10,
		Offset: 0,
	}
	notifications, err := testStore.ListRepoNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(notifications))
	require.Equal(t, notification_1.NotiID, notifications[2].NotiID)
	require.Equal(t, notification_2.NotiID, notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, notifications[0].NotiID)

	arg_mark := MarkRepoNotificationReadedParams{
		NotiID: notification_2.NotiID,
		UserID: user.UserID,
	}
	_, err = testStore.MarkRepoNotificationReaded(context.Background(), arg_mark)
	require.NoError(t, err)

	updated_notifications, err := testStore.ListRepoNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(updated_notifications))
	require.Equal(t, notification_2.NotiID, updated_notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, updated_notifications[0].NotiID)
}
func createRandomCommentNotificationByUser(t *testing.T, user User) CommentNotification {
	md := testCreateRandomMarkdown(t)
	comment := testCreateRandomCommentOnMd(t, md)
	arg := CreateCommentNotificationParams{
		UserID:    user.UserID,
		CommentID: comment.CommentID,
	}
	notification, err := testStore.CreateCommentNotification(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, notification)
	require.Equal(t, arg.CommentID, notification.CommentID)
	require.Equal(t, arg.UserID, notification.UserID)
	return notification
}
func createRandomCommentNotification(t *testing.T) CommentNotification {
	user := createRandomUser(t)
	notification := createRandomCommentNotificationByUser(t, user)
	return notification
}
func TestCreateCommentNotification(t *testing.T) {
	createRandomCommentNotification(t)
}
func TestMarkCommentNotificationReaded(t *testing.T) {
	notification := createRandomCommentNotification(t)
	require.Equal(t, false, notification.Readed)
	arg := MarkCommentNotificationReadedParams{
		NotiID: notification.NotiID,
		UserID: notification.UserID,
	}
	updated_notification, err := testStore.MarkCommentNotificationReaded(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, true, updated_notification.Readed)
}

func TestListCommentNotification(t *testing.T) {
	user := createRandomUser(t)
	notification_1 := createRandomCommentNotificationByUser(t, user)
	notification_2 := createRandomCommentNotificationByUser(t, user)
	notification_3 := createRandomCommentNotificationByUser(t, user)
	arg := ListCommentNotificationParams{
		UserID: user.UserID,
		Limit:  10,
		Offset: 0,
	}
	notifications, err := testStore.ListCommentNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(notifications))
	require.Equal(t, notification_1.NotiID, notifications[2].NotiID)
	require.Equal(t, notification_2.NotiID, notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, notifications[0].NotiID)

	arg_mark := MarkCommentNotificationReadedParams{
		NotiID: notification_2.NotiID,
		UserID: user.UserID,
	}
	_, err = testStore.MarkCommentNotificationReaded(context.Background(), arg_mark)
	require.NoError(t, err)

	updated_notifications, err := testStore.ListCommentNotification(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 3, len(updated_notifications))
	require.Equal(t, notification_2.NotiID, updated_notifications[1].NotiID)
	require.Equal(t, notification_3.NotiID, updated_notifications[0].NotiID)
}
