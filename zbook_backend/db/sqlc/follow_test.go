package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createFollowByUser(t *testing.T, follower User, following User) Follow {
	arg := CreateFollowParams{
		FollowerID:  follower.UserID,
		FollowingID: following.UserID,
	}

	Follow, err := testStore.CreateFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, Follow)
	require.Equal(t, arg.FollowingID, Follow.FollowingID)
	require.Equal(t, arg.FollowerID, Follow.FollowerID)
	require.NotZero(t, Follow.CreatedAt)
	return Follow
}

func TestCreateFollow(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	createFollowByUser(t, user1, user2)
}
func TestIsFollowing(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	createFollowByUser(t, user1, user2)

	IsFollowing, err := testStore.IsFollowing(context.Background(),
		IsFollowingParams{
			FollowerID:  user1.UserID,
			FollowingID: user2.UserID})
	require.NoError(t, err)
	require.Equal(t, IsFollowing, true)
}

func TestDeleteFollow(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	createFollowByUser(t, user1, user2)
	arg_delete := DeleteFollowParams{
		FollowerID:  user1.UserID,
		FollowingID: user2.UserID,
	}
	_, err := testStore.DeleteFollow(context.Background(), arg_delete)
	require.NoError(t, err)
	IsFollowing, err := testStore.IsFollowing(context.Background(),
		IsFollowingParams{
			FollowerID:  user1.UserID,
			FollowingID: user2.UserID})
	require.NoError(t, err)
	require.Equal(t, IsFollowing, false)
}

func TestListFollower(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	user5 := createRandomUser(t)

	_ = createFollowByUser(t, user2, user1)
	_ = createFollowByUser(t, user3, user1)
	_ = createFollowByUser(t, user4, user1)
	_ = createFollowByUser(t, user5, user1)
	// blocked 2
	arg_update_user := UpdateUserBasicInfoParams{
		Username: user2.Username,
		Verified: pgtype.Bool{Bool: true, Valid: true},
		Blocked:  pgtype.Bool{Bool: true, Valid: true},
	}
	testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
	// deleted 3
	testStore.DeleteUser(context.Background(), user3.Username)

	repo51 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo51, util.VisibilityPrivate)
	repo52 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo52, util.VisibilityChosed)
	arg_repoV := CreateRepoRelationParams{
		RepoID:       repo52.RepoID,
		UserID:       cur_user.UserID,
		RelationType: util.RelationTypeVisi,
	}
	testStore.CreateRepoRelation(context.Background(), arg_repoV)

	repo53 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo53, util.VisibilitySigned)
	repo54 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo54, util.VisibilityPublic)
	{
		arg := ListFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
		require.Equal(t, followers[0].IsFollowing, false)
		require.Equal(t, followers[0].UserID, user5.UserID)
		require.Equal(t, followers[0].RepoCount, int64(3))
	}
	{
		cur_user2 := createRandomUser(t)
		arg := ListFollowerParams{
			CurUserID: cur_user2.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      cur_user2.UserRole,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
		require.Equal(t, followers[0].IsFollowing, false)
		require.Equal(t, followers[0].UserID, user5.UserID)
		require.Equal(t, followers[0].RepoCount, int64(2))
	}
	{
		cur_user3 := createRandomUser(t)
		arg_basic := UpdateUserBasicInfoParams{
			Username: cur_user3.Username,
			UserRole: pgtype.Text{String: util.AdminRole, Valid: true},
		}
		user, err := testStore.UpdateUserBasicInfo(context.Background(), arg_basic)
		require.NoError(t, err)
		require.Equal(t, user.UserID, cur_user3.UserID)
		arg := ListFollowerParams{
			CurUserID: user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      user.UserRole,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 3)
		require.Equal(t, followers[0].IsFollowing, false)
		require.Equal(t, followers[0].UserID, user5.UserID)
		require.Equal(t, int64(4), followers[0].RepoCount)
	}

}

func TestQueryFollower(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user2, user1)
	_ = createFollowByUser(t, user3, user1)
	_ = createFollowByUser(t, user4, user1)
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
		require.Equal(t, followers[0].IsFollowing, false)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		createFollowByUser(t, cur_user, user2)
		followers, err := testStore.QueryFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
		require.Equal(t, followers[0].UserID, user2.UserID)
		require.Equal(t, true, followers[0].IsFollowing)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
	}
	{

		testStore.DeleteUser(context.Background(), user4.Username)
		arg := QueryFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user4.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollower(context.Background(), arg)
		require.NoError(t, err)
		require.Empty(t, followers)
		require.Equal(t, len(followers), 0)
	}
}

func TestGetListFollowerCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user2, user1)
	_ = createFollowByUser(t, user3, user1)
	_ = createFollowByUser(t, user4, user1)
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(2))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)

		arg := GetListFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
	{
		testStore.DeleteUser(context.Background(), user4.Username)
		arg := GetListFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(2))
	}
}

func TestGetQueryFollowerCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user2, user1)
	_ = createFollowByUser(t, user3, user1)
	_ = createFollowByUser(t, user4, user1)
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		testStore.DeleteUser(context.Background(), user4.Username)
		arg := GetQueryFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user4.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(0))
	}
}

func TestListFollowing(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	user5 := createRandomUser(t)

	_ = createFollowByUser(t, user1, user2)
	_ = createFollowByUser(t, user1, user3)
	_ = createFollowByUser(t, user1, user4)
	_ = createFollowByUser(t, user1, user5)
	// blocked 2
	arg_update_user := UpdateUserBasicInfoParams{
		Username: user2.Username,
		Verified: pgtype.Bool{Bool: true, Valid: true},
		Blocked:  pgtype.Bool{Bool: true, Valid: true},
	}
	testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
	// deleted 3
	testStore.DeleteUser(context.Background(), user3.Username)

	repo51 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo51, util.VisibilityPrivate)
	repo52 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo52, util.VisibilityChosed)
	arg_repoV := CreateRepoRelationParams{
		RepoID:       repo52.RepoID,
		UserID:       cur_user.UserID,
		RelationType: util.RelationTypeVisi,
	}
	testStore.CreateRepoRelation(context.Background(), arg_repoV)

	repo53 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo53, util.VisibilitySigned)
	repo54 := createUserRandomRepo(t, user5)
	updateRepoVisibility(t, repo54, util.VisibilityPublic)
	{
		arg := ListFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      cur_user.UserRole,
		}
		followings, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followings)
		require.Equal(t, len(followings), 2)
		require.Equal(t, followings[0].IsFollowing, false)
		require.Equal(t, followings[0].UserID, user5.UserID)
		require.Equal(t, followings[0].RepoCount, int64(3))
	}
	{
		cur_user2 := createRandomUser(t)
		arg := ListFollowingParams{
			CurUserID: cur_user2.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      cur_user2.UserRole,
		}
		followings, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followings)
		require.Equal(t, len(followings), 2)
		require.Equal(t, followings[0].IsFollowing, false)
		require.Equal(t, followings[0].UserID, user5.UserID)
		require.Equal(t, followings[0].RepoCount, int64(2))
	}
	{
		cur_user3 := createRandomUser(t)
		arg_basic := UpdateUserBasicInfoParams{
			Username: cur_user3.Username,
			UserRole: pgtype.Text{String: util.AdminRole, Valid: true},
		}
		user, err := testStore.UpdateUserBasicInfo(context.Background(), arg_basic)
		require.NoError(t, err)
		require.Equal(t, user.UserID, cur_user3.UserID)
		arg := ListFollowingParams{
			CurUserID: user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Role:      user.UserRole,
		}
		followings, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followings)
		require.Equal(t, len(followings), 3)
		require.Equal(t, followings[0].IsFollowing, false)
		require.Equal(t, followings[0].UserID, user5.UserID)
		require.Equal(t, int64(4), followings[0].RepoCount)
	}
}

func TestQueryFollowing(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user1, user2)
	_ = createFollowByUser(t, user1, user3)
	_ = createFollowByUser(t, user1, user4)

	{
		arg := QueryFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
		require.Equal(t, followers[0].IsFollowing, false)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		createFollowByUser(t, cur_user, user2)
		followers, err := testStore.QueryFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
		require.Equal(t, followers[0].UserID, user2.UserID)
		require.Equal(t, true, followers[0].IsFollowing)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
	}
	{

		testStore.DeleteUser(context.Background(), user4.Username)
		arg := QueryFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user4.Username,
			Role:      cur_user.UserRole,
		}
		followers, err := testStore.QueryFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.Empty(t, followers)
		require.Equal(t, len(followers), 0)
	}
}

func TestGetListFollowingCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user1, user2)
	_ = createFollowByUser(t, user1, user3)
	_ = createFollowByUser(t, user1, user4)
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(2))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)

		arg := GetListFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetListFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
}

func TestGetQueryFollowingCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	_ = createFollowByUser(t, user1, user2)
	_ = createFollowByUser(t, user1, user3)
	_ = createFollowByUser(t, user1, user4)
	{
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))

	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user2.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{

		testStore.DeleteUser(context.Background(), user4.Username)
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user4.Username,
			Role:      cur_user.UserRole,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(0))
	}
}
