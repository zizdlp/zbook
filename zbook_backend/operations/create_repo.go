package operations

import (
	"fmt"
	"os"
	"strings"

	convert "github.com/zizdlp/zbook/markdown/convert"
	md "github.com/zizdlp/zbook/markdown/render"
	"github.com/zizdlp/zbook/util"
)

// Clone clones a git repository from the specified URL into the specified directory.
func CreateRepo(gitURL string, userID int64, repoID int64) error {

	rsg := util.NewRandomStringGenerator()
	randomString := rsg.RandomString(10)
	cloneDir := "/tmp/zbook_repo/" + randomString
	// 删除目标目录以确保每次测试都是从头开始
	if _, err := os.Stat(cloneDir); err == nil {
		os.RemoveAll(cloneDir)
	}

	// 调用 Clone 函数
	err := Clone(gitURL, cloneDir)
	if err != nil {
		return err
	}

	// 调用 GetLatestCommit 函数
	lastCommit, err := GetLatestCommit(cloneDir)
	if err != nil {
		return err
	}
	oldCommit := ""

	// 调用 GetDiffFiles 函数
	addedFiles, deletedFiles, modifiedFiles, err := GetDiffFiles(oldCommit, lastCommit, cloneDir)
	if err != nil {
		return err
	}
	allowedExtensions := map[string]bool{
		".md": true,
	}
	markdown := md.GetMarkdownConfig()
	createParams := &util.CreateParams{}
	updateParams := &util.UpdateParams{}
	deleteParams := &util.DeleteParams{}
	filteredMarkdowns := util.FilterDiffFilesByExtensions(addedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		data, err := os.ReadFile(cloneDir + "/" + filteredMarkdown)
		if err != nil {
			return err
		}

		table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
		if err != nil {
			fmt.Println("convert markdown to buffer error:", err)
			return err
		}
		html := main.String()
		htmlList := table.String()
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		createParams.Append(relativePath, userID, repoID, html, htmlList)

	}

	filteredMarkdowns = util.FilterDiffFilesByExtensions(modifiedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		data, err := os.ReadFile(cloneDir + "/" + filteredMarkdown)
		if err != nil {
			return err
		}

		table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
		if err != nil {
			fmt.Println("convert markdown to buffer error:", err)
			return err
		}
		html := main.String()
		htmlList := table.String()
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		updateParams.Append(relativePath, repoID, html, htmlList)
	}
	filteredMarkdowns = util.FilterDiffFilesByExtensions(deletedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		deleteParams.Append(relativePath, repoID)
	}
	return nil
}
