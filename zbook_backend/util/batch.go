package util

import (
	"net/url"
	"path/filepath"
	"strings"
)

type CreateParams struct {
	RelativePath []string
	UserID       []int64
	RepoID       []int64
	Content      []string
}

func (params *CreateParams) Append(relativePath string, userID int64, repoID int64, content string) {
	params.RelativePath = append(params.RelativePath, relativePath)
	params.UserID = append(params.UserID, userID)
	params.RepoID = append(params.RepoID, repoID)
	params.Content = append(params.Content, content)

}

type UpdateParams struct {
	RelativePath    []string
	NewRelativePath []string
	RepoID          []int64
	Content         []string
}

func (params *UpdateParams) Append(RelativePath string, NewRelativePath string, RepoID int64, Content string) {
	params.RelativePath = append(params.RelativePath, RelativePath)
	params.NewRelativePath = append(params.NewRelativePath, NewRelativePath)
	params.RepoID = append(params.RepoID, RepoID)
	params.Content = append(params.Content, Content)
}

type DeleteParams struct {
	RelativePath []string
	RepoID       []int64
}

func (params *DeleteParams) Append(RelativePath string, RepoID int64) {
	params.RelativePath = append(params.RelativePath, RelativePath)
	params.RepoID = append(params.RepoID, RepoID)
}
func GetGitURL(GitProtocol string, GitHost string, GitUsername string, GitRepo string) string {
	gitURL := GitProtocol + "://" + GitHost + "/" + GitUsername + "/" + GitRepo + ".git"
	return gitURL
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
func NormalizePath(path string) string {
	return filepath.Clean(path)
}

func ParserGitCloneError(message string) string {
	// Convert output to a string

	// Find the start of the error message (e.g., "fatal: ")
	fatalPrefix := "fatal: "
	startIndex := strings.Index(message, fatalPrefix)
	if startIndex != -1 {
		// Extract the error message part
		errorMessage := message[startIndex+7:]
		return errorMessage
	} else {
		return message
	}
}
