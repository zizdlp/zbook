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

func testCreateRandomMarkdown(t *testing.T) Markdown {
	repo := createRandomRepo(t)
	arg := CreateMarkdownParams{
		RelativePath: util.RandomString(32),
		UserID:       repo.UserID,
		RepoID:       repo.RepoID,
		MainContent:  util.RandomString(32),
		TableContent: util.RandomString(32),
	}
	markdown, err := testStore.CreateMarkdown(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, markdown.RepoID, arg.RepoID)
	return markdown
}
func TestCreateMarkdown(t *testing.T) {
	testCreateRandomMarkdown(t)
}
func TestCreateMarkdownMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "zbook-user-guide",
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
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

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32000))
		TableContent = append(TableContent, util.RandomString(320))
	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)
}

func TestGetMarkdownContent(t *testing.T) {
	markdown := testCreateRandomMarkdown(t)
	arg := GetMarkdownContentParams{
		RelativePath: markdown.RelativePath,
		RepoID:       markdown.RepoID,
	}
	ret, err := testStore.GetMarkdownContent(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, markdown.MarkdownID, ret.MarkdownID)
	require.Equal(t, markdown.RelativePath, ret.RelativePath)
	require.Equal(t, markdown.RepoID, ret.RepoID)
	require.Equal(t, markdown.MainContent, ret.MainContent)
}

func TestGetMarkdownByID(t *testing.T) {
	markdown := testCreateRandomMarkdown(t)
	ret, err := testStore.GetMarkdownByID(context.Background(), markdown.MarkdownID)
	require.NoError(t, err)
	require.Equal(t, markdown.MarkdownID, ret.MarkdownID)
	require.Equal(t, markdown.RelativePath, ret.RelativePath)
	require.Equal(t, markdown.RepoID, ret.RepoID)
}
func TestGetMarkdownRepoID(t *testing.T) {
	markdown := testCreateRandomMarkdown(t)
	repoID, err := testStore.GetMarkdownRepoID(context.Background(), markdown.MarkdownID)
	require.NoError(t, err)
	require.Equal(t, markdown.RepoID, repoID)
}

func TestUpdateMarkdownMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "zbook-user-guide",
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
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

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32))
		TableContent = append(TableContent, util.RandomString(32))

	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)

	for i := 0; i < lens; i++ {
		MainContent[i] = "newmain_content"
		TableContent[i] = "newtable_content"
	}
	arg_key := UpdateMarkdownMultiParams{
		RelativePath:    RelativePath,
		NewRelativePath: RelativePath,
		MainContent:     MainContent,
		TableContent:    TableContent,
		RepoID:          RepoID,
	}
	s = time.Now()
	err = testStore.UpdateMarkdownMulti(context.Background(), arg_key)
	e = time.Since(s)
	fmt.Println("updatemdkeymulti time:", e)
	require.NoError(t, err)
}

func TestDeleteMarkdownMulti(t *testing.T) {

	user := createRandomUser(t)
	arg_repo := CreateRepoParams{
		UserID:          user.UserID,
		GitProtocol:     "http",
		GitHost:         "github.com",
		GitUsername:     "zizdlp",
		GitRepo:         "zbook-user-guide",
		ThemeSidebar:    util.ThemeSideBarFold,
		ThemeColor:      util.ThemeColorSky,
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

	for i := 0; i < lens; i++ {
		RelativePath = append(RelativePath, util.RandomString(32)+".md")
		UserID = append(UserID, user.UserID)
		RepoID = append(RepoID, repo.RepoID)
		MainContent = append(MainContent, util.RandomString(32000))
		TableContent = append(TableContent, util.RandomString(320))
	}
	arg := CreateMarkdownMultiParams{
		RelativePath: RelativePath,
		UserID:       UserID,
		RepoID:       RepoID,
		MainContent:  MainContent,
		TableContent: TableContent,
	}

	s := time.Now()
	err = testStore.CreateMarkdownMulti(context.Background(), arg)
	e := time.Since(s)
	fmt.Println("createmarkdownmulti time:", e)
	require.NoError(t, err)

	arg_delete := DeleteMarkdownMultiParams{
		RelativePath: RelativePath,
		RepoID:       RepoID,
	}

	err = testStore.DeleteMarkdownMulti(context.Background(), arg_delete)
	require.NoError(t, err)
}
func testCreateRandomMarkdownForQuery(t *testing.T) Markdown {
	repo := createRandomRepo(t)
	arg := CreateMarkdownParams{
		RelativePath: util.RandomString(32),
		UserID:       repo.UserID,
		RepoID:       repo.RepoID,
		MainContent:  "fox",
		TableContent: util.RandomString(32),
	}
	markdown, err := testStore.CreateMarkdown(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, markdown.RepoID, arg.RepoID)
	return markdown
}

func TestQueryUserMarkdown(t *testing.T) {
	m1 := testCreateRandomMarkdownForQuery(t)
	arg := QueryUserMarkdownParams{
		Limit:          10,
		Offset:         0,
		UserID:         m1.UserID,
		PlaintoTsquery: m1.MainContent,
		Role:           util.AdminRole,
		Signed:         true,
		CurUserID:      0,
	}
	rets, err := testStore.QueryUserMarkdown(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(rets) > 0)
	require.Equal(t, rets[0].RepoID, m1.RepoID)
}
func TestQueryRepoMarkdown(t *testing.T) {
	m1 := testCreateRandomMarkdownForQuery(t)
	arg := QueryRepoMarkdownParams{
		Limit:          10,
		Offset:         0,
		UserID:         m1.UserID,
		RepoID:         m1.RepoID,
		PlaintoTsquery: m1.MainContent,
	}
	rets, err := testStore.QueryRepoMarkdown(context.Background(), arg)
	require.NoError(t, err)
	require.True(t, len(rets) > 0)
	require.Equal(t, rets[0].RepoID, m1.RepoID)
}
