package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
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
	_ = createFollowByUser(t, user2, user1)
	_ = createFollowByUser(t, user3, user1)
	_ = createFollowByUser(t, user4, user1)
	{
		arg := ListFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 3)
		require.Equal(t, followers[0].IsFollowing, false)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		createFollowByUser(t, cur_user, user2)
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
		require.Equal(t, followers[0].UserID, user2.UserID)
		require.Equal(t, true, followers[0].IsFollowing)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 3)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
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
		}
		followers, err := testStore.QueryFollower(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowerParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user4.Username,
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
		}
		count, err := testStore.GetListFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
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
		}
		count, err := testStore.GetQueryFollowerCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowerCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user4.Username,
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
	_ = createFollowByUser(t, user1, user2)
	_ = createFollowByUser(t, user1, user3)
	_ = createFollowByUser(t, user1, user4)
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user2.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 3)
		require.Equal(t, followers[0].IsFollowing, false)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Blocked:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		createFollowByUser(t, cur_user, user2)
		followers, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
		require.Equal(t, followers[0].UserID, user2.UserID)
		require.Equal(t, true, followers[0].IsFollowing)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user3.Username,
			Blocked:  pgtype.Bool{Bool: false, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 3)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := ListFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
		}
		followers, err := testStore.ListFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 2)
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
		}
		followers, err := testStore.QueryFollowing(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, followers)
		require.Equal(t, len(followers), 1)
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := QueryFollowingParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Limit:     5,
			Offset:    0,
			Query:     user4.Username,
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
		}
		count, err := testStore.GetListFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(3))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetListFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
		}
		count, err := testStore.GetListFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(2))
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
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(1))
	}
	{
		arg_update_user := UpdateUserBasicInfoParams{
			Username: user4.Username,
			Verified: pgtype.Bool{Bool: true, Valid: true},
			Deleted:  pgtype.Bool{Bool: true, Valid: true},
		}
		testStore.UpdateUserBasicInfo(context.Background(), arg_update_user)
		arg := GetQueryFollowingCountParams{
			CurUserID: cur_user.UserID,
			UserID:    user1.UserID,
			Query:     user4.Username,
		}
		count, err := testStore.GetQueryFollowingCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, count, int64(0))
	}
}
