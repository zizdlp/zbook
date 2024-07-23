package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestGetUserByUsername(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUserByUsername(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Motto, user2.Motto)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
func TestGetUserByEmail(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUserByEmail(context.Background(), user1.Email)
	require.NoError(t, err)
	require.NotEmpty(t, user2)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.Motto, user2.Motto)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestGetUnreadCount(t *testing.T) {
	user1 := createRandomUser(t)
	testStore.UpdateUnreadCount(context.Background(), user1.UserID)
	unreadCount, err := testStore.GetUnReadCount(context.Background(), user1.Username)
	require.NoError(t, err)
	require.Equal(t, user1.UnreadCount, unreadCount)
}

func TestGetUserInfoByID(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)

	updateRepoVisibility(t, repo11, util.VisibilityChosed)
	updateRepoVisibility(t, repo12, util.VisibilityPrivate)
	updateRepoVisibility(t, repo13, util.VisibilityPublic)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo23 := createUserRandomRepo(t, user2)
	repo24 := createUserRandomRepo(t, user2)

	updateRepoVisibility(t, repo21, util.VisibilityChosed)
	updateRepoVisibility(t, repo22, util.VisibilityPrivate)
	updateRepoVisibility(t, repo23, util.VisibilityPublic)
	updateRepoVisibility(t, repo24, util.VisibilitySigned)

	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)
	repo33 := createUserRandomRepo(t, user3)
	repo34 := createUserRandomRepo(t, user3)

	updateRepoVisibility(t, repo31, util.VisibilityChosed)
	updateRepoVisibility(t, repo32, util.VisibilityPrivate)
	updateRepoVisibility(t, repo33, util.VisibilityPublic)
	updateRepoVisibility(t, repo34, util.VisibilitySigned)

	testCreateRelationUserRepoRelation(t, user1, repo23, util.RelationTypeDislike)
	testCreateRelationUserRepoRelation(t, user1, repo34, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo14, util.RelationTypeLike)
	arg := GetUserInfoParams{
		CurUserID: user2.UserID,
		UserID:    user1.UserID,
		Signed:    true,
		Role:      user2.UserRole,
	}
	user_info, err := testStore.GetUserInfo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, user_info.UserID, user1.UserID)
	require.Equal(t, user_info.RepoCount, int64(2))
	require.Equal(t, user_info.LikeCount, int64(2))
	require.Equal(t, user_info.IsFollowing, false)
}

func TestListUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)

	updateRepoVisibility(t, repo11, util.VisibilityChosed)
	updateRepoVisibility(t, repo12, util.VisibilityPrivate)
	updateRepoVisibility(t, repo13, util.VisibilityPublic)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo23 := createUserRandomRepo(t, user2)
	repo24 := createUserRandomRepo(t, user2)

	updateRepoVisibility(t, repo21, util.VisibilityChosed)
	updateRepoVisibility(t, repo22, util.VisibilityPrivate)
	updateRepoVisibility(t, repo23, util.VisibilityPublic)
	updateRepoVisibility(t, repo24, util.VisibilitySigned)

	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)
	repo33 := createUserRandomRepo(t, user3)
	repo34 := createUserRandomRepo(t, user3)

	updateRepoVisibility(t, repo31, util.VisibilityChosed)
	updateRepoVisibility(t, repo32, util.VisibilityPrivate)
	updateRepoVisibility(t, repo33, util.VisibilityPublic)
	updateRepoVisibility(t, repo34, util.VisibilitySigned)

	testCreateRelationUserRepoRelation(t, user1, repo23, util.RelationTypeDislike)
	testCreateRelationUserRepoRelation(t, user1, repo34, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo14, util.RelationTypeLike)
	arg := ListUserParams{
		Limit:  1000,
		Offset: 0,
	}
	users, err := testStore.ListUser(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(users) >= 3)

}

func TestQueryUser(t *testing.T) {
	user1 := createRandomUser(t)
	arg_update := UpdateUserBasicInfoParams{
		Username: user1.Username,
		Verified: pgtype.Bool{Bool: true, Valid: true},
	}
	_, err := testStore.UpdateUserBasicInfo(context.Background(), arg_update)
	require.NoError(t, err)
	arg := QueryUserParams{
		Limit:  10,
		Offset: 0,
		Query:  user1.Username,
	}
	users, err := testStore.QueryUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, users)
	require.Equal(t, users[0].UserID, user1.UserID)
}
