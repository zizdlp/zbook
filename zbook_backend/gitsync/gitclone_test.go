package gitsync

import (
	"strconv"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestCloneShouldOK(t *testing.T) {
	protocol := "http"
	host := "github.com"
	username := "zizdlp"
	repo := "wiki_demo"
	RootPath := "/tmp/wiki/"
	RepoID := util.RandomInt(1, 10000)
	RepoAccessToken := pgtype.Text{
		String: "",
		Valid:  false,
	}
	err := Clone(protocol, host, username, repo, RepoAccessToken,
		RootPath, strconv.FormatInt(RepoID, 10))
	require.NoError(t, err)
}

func TestPullShouldError(t *testing.T) {
	protocol := "http"
	host := "github.com"
	username := "zizdlp"
	repo := "wiki_demo"
	RootPath := "/tmp/wiki/"
	RepoID := util.RandomInt(1, 10000)
	GitAccessToken := pgtype.Text{
		String: "",
		Valid:  false,
	}
	err := Clone(protocol, host, username, repo, GitAccessToken,
		RootPath, strconv.FormatInt(RepoID, 10))
	require.NoError(t, err)
	err = Pull(username, GitAccessToken,
		RootPath, strconv.FormatInt(RepoID, 10))
	require.EqualError(t, err, "already up-to-date")
}

func TestCloneShouldError(t *testing.T) {
	protocol := "http"
	host := "gitee.com"
	username := "zizdlp"
	repo := "wiki"
	RootPath := "/tmp/wiki/"
	RepoID := util.RandomInt(1, 10000)
	RepoAccessToken := pgtype.Text{
		String: "",
		Valid:  false,
	}
	err := Clone(protocol, host, username, repo, RepoAccessToken,
		RootPath, strconv.FormatInt(RepoID, 10))
	require.Error(t, err)
	require.EqualError(t, err, "authentication required")
}

func TestParseGitURL(t *testing.T) {
	protocol := "http"
	host := "gitee.com"
	username := "zizdlp"
	repo := "wiki"

	gitURL := protocol + "://" + host + "/" + username + "/" + repo + ".git"

	res_protocol, res_host, res_username, res_repo, err := ParseGitURL(gitURL)
	require.NoError(t, err)
	require.Equal(t, protocol, res_protocol)
	require.Equal(t, host, res_host)
	require.Equal(t, username, res_username)
	require.Equal(t, repo, res_repo)
}
