package val

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateRepoName(t *testing.T) {
	reponame := "文档"
	err := ValidateRepoName(reponame)
	require.NoError(t, err)

	reponame = "document"
	err = ValidateRepoName(reponame)
	require.NoError(t, err)

	reponame = "wiki docs"
	err = ValidateRepoName(reponame)
	require.NoError(t, err)

	reponame = "wiki documents this is a good job,well done,wiki documents this is a good job,well donewiki documents this is a good job,well done,wiki documents this is a good job,well done"
	err = ValidateRepoName(reponame)
	require.EqualError(t, err, "repository name length is not within the valid range:[2,64]")

	reponame = "@"
	err = ValidateRepoName(reponame)
	require.Error(t, err)

	reponame = "/"
	err = ValidateRepoName(reponame)
	require.Error(t, err)

}

func TestValidateTimeZone(t *testing.T) {
	time_zone := ""
	err := ValidTimeZone(time_zone)
	require.EqualError(t, err, "timezone cannot be empty")
	time_zone = "Asia/Shanghai"
	err = ValidTimeZone(time_zone)
	require.NoError(t, err)
}
