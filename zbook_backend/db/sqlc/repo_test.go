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
		RepoName:        util.RandomString(36),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, repo.UserID, user.UserID)
	return repo
}
func TestCreateRepo(t *testing.T) {
	createRandomRepo(t)
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
		RepoName:        util.RandomString(36),
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
func TestUpdateRepoLayout(t *testing.T) {
	repo := createRandomRepo(t)
	arg := UpdateRepoLayoutParams{
		RepoID: repo.RepoID,
		Layout: util.RandomString(32),
	}
	err := testStore.UpdateRepoLayout(context.Background(), arg)
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
		arg_visiblility := CreateRepoVisibilityParams{
			UserID: user2.UserID,
			RepoID: repo13.RepoID,
		}
		testStore.CreateRepoVisibility(context.Background(), arg_visiblility)
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
func TestGetRepoInfo(t *testing.T) {
	user := createRandomUser(t)
	user2 := createRandomUser(t)
	user3 := createRandomUser(t)
	user4 := createRandomUser(t)
	repo := createRandomRepo(t)
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeLike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user2.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeLike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user3.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeDislike})
	testStore.CreateRepoRelation(context.Background(), CreateRepoRelationParams{UserID: user4.UserID, RepoID: repo.RepoID, RelationType: util.RelationTypeDislike})
	{

		repo_info, err := testStore.GetRepoBasicInfo(context.Background(), repo.RepoID)
		require.NoError(t, err)
		require.Equal(t, repo_info.RepoID, repo.RepoID)
		require.Equal(t, repo_info.VisibilityLevel, repo.VisibilityLevel)

	}

}

func createUserRandomRepo(t *testing.T, user User) Repo {
	arg := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "zbook-user-guide",
		GitAccessToken:  pgtype.Text{String: util.RandomString(32), Valid: true},
		RepoName:        util.RandomString(36),
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
}
