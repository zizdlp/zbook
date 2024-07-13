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
