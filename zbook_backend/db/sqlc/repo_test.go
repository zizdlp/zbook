package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func createRandomRepo(t *testing.T) Repo {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		GitRepo:         util.RandomString(6),
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(16),
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	require.Equal(t, repo.ThemeColor, arg.ThemeColor)
	return repo
}
func TestCreateRepo(t *testing.T) {
	createRandomRepo(t)
}
func TestUpdateRepoLayout(t *testing.T) {
	repo := createRandomRepo(t)
	arg := UpdateRepoConfigParams{
		RepoID: repo.RepoID,
		Config: util.RandomString(32),
	}
	err := testStore.UpdateRepoConfig(context.Background(), arg)
	require.NoError(t, err)
}
func TestUpdateRepoInfo(t *testing.T) {
	repo := createRandomRepo(t)
	arg := UpdateRepoInfoParams{
		RepoID:          repo.RepoID,
		RepoName:        pgtype.Text{String: util.RandomString(6), Valid: true},
		RepoDescription: pgtype.Text{String: util.RandomString(6), Valid: true},
		SyncToken:       pgtype.Text{String: util.RandomString(6), Valid: true},
		GitAccessToken:  pgtype.Text{String: util.RandomString(6), Valid: true},
		VisibilityLevel: pgtype.Text{String: util.RandomRepoVisibility(), Valid: true},
	}
	repo, err := testStore.UpdateRepoInfo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.RepoID, arg.RepoID)
	require.Equal(t, repo.VisibilityLevel, arg.VisibilityLevel.String)
}
func TestDeleteRepo(t *testing.T) {
	repo := createRandomRepo(t)
	err := testStore.DeleteRepo(context.Background(), repo.RepoID)
	require.NoError(t, err)
	repo, err = testStore.GetRepo(context.Background(), repo.RepoID)
	require.EqualError(t, err, "no rows in result set")
	require.Empty(t, repo)
}
func TestGetRepo(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		GitRepo:         util.RandomString(6),
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.VisibilityPublic,
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	repo_ret, err := testStore.GetRepo(context.Background(), repo.RepoID)
	require.NoError(t, err)
	require.Equal(t, repo_ret.RepoID, repo.RepoID)
}
func TestGetRepoID(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		GitRepo:         util.RandomString(6),
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.VisibilityPublic,
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	arg_get := GetRepoIDParams{
		Username: user.Username,
		RepoName: repo.RepoName,
	}
	repo_id, err := testStore.GetRepoID(context.Background(), arg_get)
	require.NoError(t, err)
	require.Equal(t, repo_id, repo.RepoID)
}
func TestGetRepoByRepoName(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		GitRepo:         util.RandomString(6),
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.VisibilityPublic,
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	arg_get := GetRepoByRepoNameParams{
		Username: user.Username,
		RepoName: repo.RepoName,
	}
	repo_ret, err := testStore.GetRepoByRepoName(context.Background(), arg_get)
	require.NoError(t, err)
	require.Equal(t, repo_ret.RepoID, repo.RepoID)
}

func TestGetRepoLayout(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		GitRepo:         util.RandomString(6),
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.VisibilityPublic,
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	arg_get := GetRepoConfigParams{
		Username: user.Username,
		RepoName: repo.RepoName,
	}
	repo_ret, err := testStore.GetRepoConfig(context.Background(), arg_get)
	require.NoError(t, err)
	require.Equal(t, repo_ret.RepoID, repo.RepoID)
	require.Equal(t, repo_ret.Config, repo.Config)
}

func TestGetRepoPermission(t *testing.T) {
	user := createRandomUser(t)
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     util.RandomString(6),
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		GitRepo:         util.RandomString(6),
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.VisibilityPublic,
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	repo_ret, err := testStore.GetRepoPermission(context.Background(), repo.RepoID)
	require.NoError(t, err)
	require.Equal(t, repo_ret.RepoID, repo.RepoID)
	require.Equal(t, repo_ret.VisibilityLevel, repo.VisibilityLevel)
}
func TestGetRepoBasicInfo(t *testing.T) {
	user := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	repo := createUserRandomRepo(t, user)
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeLike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user2.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeLike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user3.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeDislike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user4.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeDislike})
	{
		arg := GetRepoBasicInfoParams{
			Username: user.Username,
			RepoName: repo.RepoName,
		}
		repo_info, err := testStore.GetRepoBasicInfo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, repo_info.RepoID, repo.RepoID)
		require.Equal(t, repo_info.VisibilityLevel, repo.VisibilityLevel)
	}

}

func TestGetQueryRepoCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo1 := createUserRandomRepo(t, user1)
	arg_update := UpdateRepoInfoParams{
		RepoID:          repo1.RepoID,
		VisibilityLevel: pgtype.Text{String: util.VisibilityPublic, Valid: true},
	}
	testStore.UpdateRepoInfo(context.Background(), arg_update)
	arg := GetQueryRepoCountParams{
		Query:     repo1.RepoName,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count1, err := testStore.GetQueryRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, count1 > 0)
}

func TestQueryRepo(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo1 := createUserRandomRepo(t, user1)
	arg_update := UpdateRepoInfoParams{
		RepoID:          repo1.RepoID,
		VisibilityLevel: pgtype.Text{String: util.VisibilityPublic, Valid: true},
	}
	testStore.UpdateRepoInfo(context.Background(), arg_update)

	arg := QueryRepoParams{
		Limit:     5,
		Offset:    0,
		Query:     repo1.RepoName,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count1, err := testStore.QueryRepo(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(count1) > 0)
}

func TestGetListRepoCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo := createUserRandomRepo(t, user1)
	arg_update := UpdateRepoInfoParams{
		RepoID:          repo.RepoID,
		VisibilityLevel: pgtype.Text{String: util.VisibilityPublic, Valid: true},
	}
	testStore.UpdateRepoInfo(context.Background(), arg_update)
	arg := GetListRepoCountParams{
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count1, err := testStore.GetListRepoCount(context.Background(), arg)
	require.NoError(t, err)
	repo2 := createUserRandomRepo(t, user1)

	arg_update = UpdateRepoInfoParams{
		RepoID:          repo2.RepoID,
		VisibilityLevel: pgtype.Text{String: util.VisibilityPublic, Valid: true},
	}
	testStore.UpdateRepoInfo(context.Background(), arg_update)
	count2, err := testStore.GetListRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count2, count1+1)

	arg_update = UpdateRepoInfoParams{
		RepoID:          repo2.RepoID,
		VisibilityLevel: pgtype.Text{String: util.VisibilityPrivate, Valid: true},
	}
	testStore.UpdateRepoInfo(context.Background(), arg_update)
	count3, err := testStore.GetListRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count3, count1)

}

func TestListRepo(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)
	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)
	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	updateRepoVisibility(t, repo12, util.VisibilityPublic)
	updateRepoVisibility(t, repo13, util.VisibilityPublic)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	updateRepoVisibility(t, repo21, util.VisibilityPrivate)
	updateRepoVisibility(t, repo22, util.VisibilityPublic)
	updateRepoVisibility(t, repo31, util.VisibilityPublic)
	updateRepoVisibility(t, repo32, util.VisibilitySigned)
	arg := ListRepoParams{Limit: 10, Offset: 0, CurUserID: user1.UserID, Role: util.UserRole, Signed: true}
	repos, err := testStore.ListRepo(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(repos) >= 6)
	require.Equal(t, repos[0].RepoID, repo32.RepoID)
}

func TestGetQueryUserOwnRepoCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	updateRepoVisibility(t, repo11, util.VisibilityPublic)
	arg := GetQueryUserOwnRepoCountParams{
		Query:     repo11.RepoName,
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count1, err := testStore.GetQueryUserOwnRepoCount(context.Background(), arg)
	require.NoError(t, err)
	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	arg = GetQueryUserOwnRepoCountParams{
		Query:     repo11.RepoName,
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count2, err := testStore.GetQueryUserOwnRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count1, int64(1))
	require.Equal(t, count2, int64(0))
}
func TestQueryUserOwnRepo(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	updateRepoVisibility(t, repo11, util.VisibilityPublic)
	arg := QueryUserOwnRepoParams{
		Limit:     5,
		Offset:    0,
		Query:     repo11.RepoName,
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	repos1, err := testStore.QueryUserOwnRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 1, len(repos1))
}

func TestGetListUserOwnRepoCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	updateRepoVisibility(t, repo11, util.VisibilityPublic)
	arg := GetListUserOwnRepoCountParams{
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count1, err := testStore.GetListUserOwnRepoCount(context.Background(), arg)
	require.NoError(t, err)
	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	arg = GetListUserOwnRepoCountParams{
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	count2, err := testStore.GetListUserOwnRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, count1, int64(1))
	require.Equal(t, count2, int64(0))
}
func TestListUserOwnRepo(t *testing.T) {
	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)
	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)
	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	updateRepoVisibility(t, repo12, util.VisibilityPublic)
	updateRepoVisibility(t, repo13, util.VisibilityChosed)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	updateRepoVisibility(t, repo21, util.VisibilityPrivate)
	updateRepoVisibility(t, repo22, util.VisibilityPublic)
	updateRepoVisibility(t, repo31, util.VisibilityPublic)
	updateRepoVisibility(t, repo32, util.VisibilitySigned)

	{
		// other in group
		arg_visiblility := CreateRepoRelationParams{
			UserID:       user2.UserID,
			RepoID:       repo13.RepoID,
			RelationType: util.RelationTypeVisi,
		}
		testStore.CreateRepoRelation(context.Background(), arg_visiblility)
		arg := ListUserOwnRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      util.UserRole,
			Signed:    true,
			CurUserID: user2.UserID,
		}
		repos, err := testStore.ListUserOwnRepo(context.Background(), arg)
		require.NoError(t, err)
		require.True(t, len(repos) == 3)

	}

	{
		// self
		arg := ListUserOwnRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      util.UserRole,
			Signed:    true,
			CurUserID: user1.UserID,
		}
		repos, err := testStore.ListUserOwnRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 4)
	}

	{
		// signed other
		arg := ListUserOwnRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      user4.UserRole,
			Signed:    true,
			CurUserID: user4.UserID,
		}
		repos, err := testStore.ListUserOwnRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 2)
	}

	{
		// sign out
		arg := ListUserOwnRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      util.UserRole,
			Signed:    false,
			CurUserID: 0,
		}
		repos, err := testStore.ListUserOwnRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 1)
	}
}

func TestGetListUserLikeRepoCount(t *testing.T) {

	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)
	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)

	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo12, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo13, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo14, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo21, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo22, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo31, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo32, util.RelationTypeLike)

	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	updateRepoVisibility(t, repo12, util.VisibilityPublic)
	updateRepoVisibility(t, repo13, util.VisibilityChosed)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	updateRepoVisibility(t, repo21, util.VisibilityPrivate)
	updateRepoVisibility(t, repo22, util.VisibilityPublic)
	updateRepoVisibility(t, repo31, util.VisibilityPublic)
	updateRepoVisibility(t, repo32, util.VisibilitySigned)

	{
		// sign out
		arg := GetListUserLikeRepoCountParams{
			UserID:    user1.UserID,
			Role:      util.RandomUserRole(),
			Signed:    false,
			CurUserID: 0,
		}
		counts, err := testStore.GetListUserLikeRepoCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, counts, int64(3))
	}

	{
		// signed
		arg := GetListUserLikeRepoCountParams{
			UserID:    user1.UserID,
			Role:      user4.UserRole,
			Signed:    true,
			CurUserID: user4.UserID,
		}
		counts, err := testStore.GetListUserLikeRepoCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, counts, int64(5))
	}

	{
		// other in group
		arg := GetListUserLikeRepoCountParams{
			UserID:    user1.UserID,
			Role:      user2.UserRole,
			Signed:    true,
			CurUserID: user2.UserID,
		}
		counts, err := testStore.GetListUserLikeRepoCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, counts, int64(6))
	}

	{
		// self
		arg := GetListUserLikeRepoCountParams{
			UserID:    user1.UserID,
			Role:      user1.UserRole,
			Signed:    true,
			CurUserID: user1.UserID,
		}
		counts, err := testStore.GetListUserLikeRepoCount(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, counts, int64(7))
	}
}
func TestListUserLikeRepo(t *testing.T) {

	user1 := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	repo12 := createUserRandomRepo(t, user1)
	repo13 := createUserRandomRepo(t, user1)
	repo14 := createUserRandomRepo(t, user1)
	repo21 := createUserRandomRepo(t, user2)
	repo22 := createUserRandomRepo(t, user2)
	repo31 := createUserRandomRepo(t, user3)
	repo32 := createUserRandomRepo(t, user3)

	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo12, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo13, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo14, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo21, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo22, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo31, util.RelationTypeLike)
	testCreateRelationUserRepoRelation(t, user1, repo32, util.RelationTypeLike)

	updateRepoVisibility(t, repo11, util.VisibilityPrivate)
	updateRepoVisibility(t, repo12, util.VisibilityPublic)
	updateRepoVisibility(t, repo13, util.VisibilityChosed)
	updateRepoVisibility(t, repo14, util.VisibilitySigned)

	updateRepoVisibility(t, repo21, util.VisibilityPrivate)
	updateRepoVisibility(t, repo22, util.VisibilityPublic)
	updateRepoVisibility(t, repo31, util.VisibilityPublic)
	updateRepoVisibility(t, repo32, util.VisibilitySigned)

	{
		// sign out
		arg := ListUserLikeRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      util.RandomUserRole(),
			Signed:    false,
			CurUserID: 0,
		}
		repos, err := testStore.ListUserLikeRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 3)
	}

	{
		// signed
		arg := ListUserLikeRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      user4.UserRole,
			Signed:    true,
			CurUserID: user4.UserID,
		}
		repos, err := testStore.ListUserLikeRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 5)
	}

	{
		// other in group
		arg := ListUserLikeRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      user2.UserRole,
			Signed:    true,
			CurUserID: user2.UserID,
		}
		repos, err := testStore.ListUserLikeRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 6)
	}

	{
		// self
		arg := ListUserLikeRepoParams{
			Limit:     10,
			Offset:    0,
			UserID:    user1.UserID,
			Role:      user1.UserRole,
			Signed:    true,
			CurUserID: user1.UserID,
		}
		repos, err := testStore.ListUserLikeRepo(context.Background(), arg)
		require.NoError(t, err)
		require.Equal(t, len(repos), 7)
	}
}

func TestQueryUserLikeRepoCount(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	updateRepoVisibility(t, repo11, util.VisibilityPublic)
	arg := GetQueryUserLikeRepoCountParams{
		Query:     repo11.RepoName,
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	counts, err := testStore.GetQueryUserLikeRepoCount(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, int64(1), counts)
}

func TestQueryUserLikeRepo(t *testing.T) {
	cur_user := createRandomUser(t)
	user1 := createRandomUser(t)
	repo11 := createUserRandomRepo(t, user1)
	testCreateRelationUserRepoRelation(t, user1, repo11, util.RelationTypeLike)
	updateRepoVisibility(t, repo11, util.VisibilityPublic)
	arg := QueryUserLikeRepoParams{
		Limit:     5,
		Offset:    0,
		Query:     repo11.RepoName,
		UserID:    user1.UserID,
		Role:      cur_user.UserRole,
		Signed:    true,
		CurUserID: cur_user.UserID,
	}
	repos1, err := testStore.QueryUserLikeRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, 1, len(repos1))
}

func createUserRandomRepo(t *testing.T, user User) Repo {
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "zbook-user-guide",
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
		GitAccessToken:  pgtype.Text{String: util.RandomString(32), Valid: true},
		RepoName:        util.RandomString(16),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	return repo
}
func updateRepoVisibility(t *testing.T, repo Repo, visibility string) {
	arg := UpdateRepoInfoParams{
		RepoID:          repo.RepoID,
		VisibilityLevel: pgtype.Text{String: visibility, Valid: true},
	}
	repo, err := testStore.UpdateRepoInfo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.RepoID, arg.RepoID)
	require.Equal(t, repo.VisibilityLevel, arg.VisibilityLevel.String)
}
