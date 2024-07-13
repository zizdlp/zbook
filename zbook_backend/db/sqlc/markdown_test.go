package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestCreateMarkdownMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "wiki_demo",
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(36),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg_repo)
	require.NoError(t, err)

	lens := 10
	RelativePath := make([]string, 0)
	UserID := make([]int64, 0)
	RepoID := make([]int64, 0)
	MainContent := make([]string, 0)
	TableContent := make([]string, 0)
	Md5 := make([]string, 0)
	VersionKey := make([]string, 0)

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32000))
		TableContent = append(TableContent, util.RandomString(320))
		Md5 = append(Md5, util.RandomString(32))
		VersionKey = append(VersionKey, util.RandomString(32))
	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
		Md5:          Md5,
		VersionKey:   VersionKey,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)
}

func TestUpdateMarkdownVersionKeyMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "wiki_demo",
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(36),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg_repo)
	require.NoError(t, err)

	lens := 20000
	RelativePath := []string{}
	UserID := []int64{}
	RepoID := []int64{}
	MainContent := []string{}
	TableContent := []string{}
	Md5 := []string{}
	VersionKey := []string{}

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32))
		TableContent = append(TableContent, util.RandomString(32))
		Md5 = append(Md5, util.RandomString(32))
		VersionKey = append(VersionKey, util.RandomString(32))
	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
		Md5:          Md5,
		VersionKey:   VersionKey,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)

	for i := 0; i < lens; i++ {
		VersionKey[i] = "newversion_key"
	}
	arg_key := UpdateMarkdownVersionKeyParams{
		VersionKey:   VersionKey,
		RelativePath: RelativePath,
		RepoID:       RepoID,
	}
	s = time.Now()
	err = testStore.UpdateMarkdownVersionKey(context.Background(), arg_key)
	e = time.Since(s)
	fmt.Println("updatemdkey time:", e)
	require.NoError(t, err)
}
func testCreateRandomMarkdown(t *testing.T) Markdown {
	repo := createRandomRepo(t)
	arg := CreateMarkdownParams{
		RelativePath: util.RandomString(32),
		UserID:       repo.UserID,
		RepoID:       repo.RepoID,
		MainContent:  util.RandomString(32),
		TableContent: util.RandomString(32),
		Md5:          util.RandomString(32),
		VersionKey:   util.RandomString(32),
	}
	markdown, err := testStore.CreateMarkdown(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, markdown.RepoID, arg.RepoID)
	return markdown
}
func TestCreateMarkdown(t *testing.T) {
	testCreateRandomMarkdown(t)
}

func TestUpdateMarkdownMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "wiki_demo",
		GitAccessToken:  pgtype.Text{String: "", Valid: true},
		RepoName:        util.RandomString(36),
		RepoDescription: util.RandomString(200),
		SyncToken:       pgtype.Text{String: util.RandomString(32), Valid: true},
		VisibilityLevel: util.RandomRepoVisibility(),
	}
	repo, err := testStore.CreateRepo(context.Background(), arg_repo)
	require.NoError(t, err)

	lens := 20000
	RelativePath := []string{}
	UserID := []int64{}
	RepoID := []int64{}
	MainContent := []string{}
	TableContent := []string{}
	Md5 := []string{}
	VersionKey := []string{}

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32))
		TableContent = append(TableContent, util.RandomString(32))
		Md5 = append(Md5, util.RandomString(32))
		VersionKey = append(VersionKey, util.RandomString(32))
	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
		Md5:          Md5,
		VersionKey:   VersionKey,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)

	for i := 0; i < lens; i++ {
		VersionKey[i] = "newversion_key"
		MainContent[i] = "newmain_content"
		TableContent[i] = "newtable_content"
		Md5[i] = "newmd5"
	}
	arg_key := UpdateMarkdownMultiParams{
		VersionKey:   VersionKey,
		RelativePath: RelativePath,
		MainContent:  MainContent,
		TableContent: TableContent,
		Md5:          Md5,
		RepoID:       RepoID,
	}
	s = time.Now()
	err = testStore.UpdateMarkdownMulti(context.Background(), arg_key)
	e = time.Since(s)
	fmt.Println("updatemdkeymulti time:", e)
	require.NoError(t, err)
}
