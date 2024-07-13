package gitsync

import (
	"os"

	"net/url"
	"strings"

	git "github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/transport/http" // with go modules enabled (GO111MODULE=on or outside GOPATH)
	"github.com/jackc/pgx/v5/pgtype"
) // with go modules disabled
func Clone(GitProtocol string, GitHost string, GitUsername string, GitRepo string, GitAccessToken pgtype.Text,
	RootPath string, RepoID string) error {

	gitURL := GetGitURL(GitProtocol, GitHost, GitUsername, GitRepo)

	directory := RootPath + RepoID + "/"

	if !GitAccessToken.Valid {
		_, err := git.PlainClone(directory, false, &git.CloneOptions{
			URL:               gitURL,
			RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		})
		return err
	} else {
		_, err := git.PlainClone(directory, false, &git.CloneOptions{
			Auth: &http.BasicAuth{
				Username: GitUsername,
				Password: GitAccessToken.String,
			},
			URL:      gitURL,
			Progress: os.Stdout,
		})
		return err
	}
}

func Pull(GitUsername string, GitAccessToken pgtype.Text,
	RootPath string, RepoID string) error {
	directory := RootPath + "/" + RepoID + "/"

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(directory)
	if err != nil {
		return err
	}
	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		return err
	}
	if !GitAccessToken.Valid {
		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin"})
	} else {
		// Pull the latest changes from the origin remote and merge into the current branch
		err = w.Pull(&git.PullOptions{RemoteName: "origin", Auth: &http.BasicAuth{
			Username: GitUsername,
			Password: GitAccessToken.String,
		}})
	}
	if err != nil {
		return err
	}

	return err
}

// ParseGitURL 解析Git地址并返回协议、主机、用户名和仓库
func ParseGitURL(gitURL string) (string, string, string, string, error) {
	u, err := url.Parse(gitURL)
	if err != nil {
		return "", "", "", "", err
	}

	var protocol, host, username, repo string

	protocol = u.Scheme
	host = u.Host

	// 去除路径前后的斜杠
	path := strings.Trim(u.Path, "/")

	// 按照斜杠进行分割
	parts := strings.Split(path, "/")
	if len(parts) >= 2 {
		username = parts[0]
		repo = strings.TrimSuffix(parts[len(parts)-1], ".git")
	}
	return protocol, host, username, repo, nil
}
func GetGitURL(GitProtocol string, GitHost string, GitUsername string, GitRepo string) string {
	gitURL := GitProtocol + "://" + GitHost + "/" + GitUsername + "/" + GitRepo + ".git"
	return gitURL
}
