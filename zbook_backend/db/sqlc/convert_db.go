package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/minio/minio-go/v7"
	convert "github.com/zizdlp/zbook/markdown/convert"
	md "github.com/zizdlp/zbook/markdown/render"
	"github.com/zizdlp/zbook/operations"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
)

func uploadGitFile(minioClient *minio.Client, ctx context.Context, cloneDir string, filePath string, repoID int64) error {

	data, err := os.ReadFile(cloneDir + "/" + filePath)
	if err != nil {
		return err
	}
	ext := strings.ToLower(filePath)
	if strings.HasSuffix(ext, ".png") || strings.HasSuffix(ext, ".jpg") || strings.HasSuffix(ext, ".jpeg") || strings.HasSuffix(ext, ".webp") {
		base64, err := util.ReadImageBytes(cloneDir + "/" + filePath)
		if err != nil {
			return err
		}
		data, err = util.CompressImage(base64)
		if err != nil {
			return err
		}
	}

	repoIDStr := strconv.FormatInt(repoID, 10)
	name := repoIDStr + "/" + filePath
	err = storage.UploadFileToStorage(minioClient, ctx, name, "git-files", data)
	if err != nil {
		return err
	}

	return nil
}
func ConvertFile2DB(ctx context.Context, q *Queries, cloneDir string, repoID int64, userID int64, oldCommit string) error {
	startTime := time.Now()
	minioClient, err := storage.GetMinioClient()
	if err != nil {
		return err
	}
	// 调用 GetLatestCommit 函数
	lastCommit, err := operations.GetLatestCommit(cloneDir)
	if err != nil {
		return err
	}

	// 调用 GetDiffFiles 函数
	addedFiles, deletedFiles, modifiedFiles, err := operations.GetDiffFiles(oldCommit, lastCommit, cloneDir)
	if err != nil {
		return err
	}
	allowedExtensions := map[string]bool{
		".md": true,
	}
	allowedGitFileExtensions := map[string]bool{

		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".svg":  true,
		".gif":  true,
		".webp": true,
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

	filteredGitFiles := util.FilterDiffFilesByExtensions(addedFiles, allowedGitFileExtensions)

	for _, filteredGitFile := range filteredGitFiles {

		err = uploadGitFile(minioClient, ctx, cloneDir, filteredGitFile, repoID)

		if err != nil {
			return err
		}
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
	filteredGitFiles = util.FilterDiffFilesByExtensions(modifiedFiles, allowedGitFileExtensions)

	for _, filteredGitFile := range filteredGitFiles {

		err = uploadGitFile(minioClient, ctx, cloneDir, filteredGitFile, repoID)

		if err != nil {
			return err
		}
	}

	filteredMarkdowns = util.FilterDiffFilesByExtensions(deletedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		deleteParams.Append(relativePath, repoID)
	}
	filteredGitFiles = util.FilterDiffFilesByExtensions(deletedFiles, allowedGitFileExtensions)

	for _, filteredGitFile := range filteredGitFiles {
		repoIDStr := strconv.FormatInt(repoID, 10)
		name := repoIDStr + "/" + filteredGitFile
		err = storage.DeleteFileFromStorage(minioClient, ctx, name, "git-files")
		if err != nil {
			return err
		}
	}

	err = createMarkdownFiles(ctx, q, createParams)
	if err != nil {
		return fmt.Errorf("create markdown failed: %v", err)
	}
	err = updateMarkdownFiles(ctx, q, updateParams)
	if err != nil {
		return fmt.Errorf("update markdown failed: %v", err)
	}
	err = deleteMarkdownFiles(ctx, q, deleteParams)
	if err != nil {
		return fmt.Errorf("delete markdown failed: %v", err)
	}
	mdFiles, err := operations.ListMarkdownFiles(cloneDir)
	if err != nil {
		return fmt.Errorf("generate layout failed: %v", err)
	}
	layout := util.CreateLayout(mdFiles)

	layoutJSON, err := json.MarshalIndent(layout, "", "  ")
	if err != nil {
		return fmt.Errorf("generate layout failed: %v", err)
	}

	arg_update_repo_layout := UpdateRepoLayoutParams{
		RepoID:   repoID,
		Layout:   string(layoutJSON),
		CommitID: lastCommit,
	}
	err = q.UpdateRepoLayout(ctx, arg_update_repo_layout)
	if err != nil {
		return fmt.Errorf("update repo layout failed: %v", err)
	}

	endTime := time.Now()
	fmt.Println("mydebug: total execution time:", endTime.Sub(startTime))

	return nil
}

func createMarkdownFiles(ctx context.Context, q *Queries, params *util.CreateParams) error {
	argCreate := CreateMarkdownMultiParams{
		RelativePath: params.RelativePath,
		UserID:       params.UserID,
		RepoID:       params.RepoID,
		MainContent:  params.MainContent,
		TableContent: params.TableContent,
	}
	err := q.CreateMarkdownMulti(ctx, argCreate)
	if err != nil {
		return err
	}
	return nil
}

func updateMarkdownFiles(ctx context.Context, q *Queries, params *util.UpdateParams) error {
	argUpdate := UpdateMarkdownMultiParams{
		RelativePath: params.RelativePath,
		RepoID:       params.RepoID,
		MainContent:  params.MainContent,
		TableContent: params.TableContent,
	}
	err := q.UpdateMarkdownMulti(ctx, argUpdate)
	if err != nil {
		return err
	}
	return nil
}

func deleteMarkdownFiles(ctx context.Context, q *Queries, params *util.DeleteParams) error {

	argDelete := DeleteMarkdownMultiParams{
		RelativePath: params.RelativePath,
		RepoID:       params.RepoID,
	}
	err := q.DeleteMarkdownMulti(ctx, argDelete)
	if err != nil {
		return err
	}

	return nil
}
