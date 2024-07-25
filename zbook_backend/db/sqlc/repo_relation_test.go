package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func testCreateRelationUserRepoRelation(t *testing.T, user User, repo Repo, relation string) {
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: relation,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)
}
func testCreateUserRepoRelation(t *testing.T, user User, repo Repo) {
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeLike,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)
}
func TestCreateRepoRelation(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	testCreateUserRepoRelation(t, user, repo)
}
func TestDelteRepoRelation(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeLike,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)
	arg_delete := DeleteRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeLike,
	}
	err = testStore.DeleteRepoRelation(context.Background(), arg_delete)
	require.NoError(t, err)
}

func TestGetRepoRelation(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeLike,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)

	arg_get := GetRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: arg.RelationType,
	}
	relation, err := testStore.GetRepoRelation(context.Background(), arg_get)
	require.NoError(t, err)
	require.Equal(t, relation.RelationType, arg_get.RelationType)
	require.Equal(t, relation.UserID, arg_get.UserID)
	require.Equal(t, relation.RepoID, arg_get.RepoID)
}

func TestGetRepoVisibilityByRepoCount(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeVisi,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)

	count, err := testStore.GetRepoVisibilityByRepoCount(context.Background(), arg.RepoID)
	require.NoError(t, err)
	require.Equal(t, count, int64(1))
}
func TestListRepoVisibilityByRepo(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeVisi,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)

	arg_list := ListRepoVisibilityByRepoParams{
		Limit:  5,
		Offset: 0,
		RepoID: arg.RepoID,
	}
	rets, err := testStore.ListRepoVisibilityByRepo(context.Background(), arg_list)
	require.NoError(t, err)
	require.Equal(t, len(rets), 1)
	require.Equal(t, rets[0].UserID, user.UserID)
}

func TestQueryRepoVisibilityByRepo(t *testing.T) {
	user := createRandomUser(t)
	repo := createRandomRepo(t)
	arg := CreateRepoRelationParams{
		UserID:       user.UserID,
		RepoID:       repo.RepoID,
		RelationType: util.RelationTypeVisi,
	}
	err := testStore.CreateRepoRelation(context.Background(), arg)
	require.NoError(t, err)

	arg_list := QueryRepoVisibilityByRepoParams{
		Limit:    5,
		Offset:   0,
		RepoID:   arg.RepoID,
		Username: user.Username,
	}
	rets, err := testStore.QueryRepoVisibilityByRepo(context.Background(), arg_list)
	require.NoError(t, err)
	require.Equal(t, len(rets), 1)
	require.Equal(t, rets[0].UserID, user.UserID)
}
