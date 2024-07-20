package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	convert "github.com/zizdlp/zbook/markdown/convert"
	md "github.com/zizdlp/zbook/markdown/render"
	"github.com/zizdlp/zbook/operations"
	"github.com/zizdlp/zbook/util"
)

func ConvertFile2DB(ctx context.Context, q *Queries, cloneDir string, repoID int64, userID int64, lastCommit string, addedFiles []string, modifiedFiles []string, deletedFiles []string) error {
	startTime := time.Now()
	allowedExtensions := map[string]bool{
		".md": true,
	}

	markdown := md.GetMarkdownConfig()
	createParams := &util.CreateParams{}
	updateParams := &util.UpdateParams{}
	deleteParams := &util.DeleteParams{}

	// Helper function to process files
	processFiles := func(files []string, isCreate bool) {
		filteredMarkdowns := util.FilterDiffFilesByExtensions(files, allowedExtensions)
		var wg sync.WaitGroup
		for _, filteredMarkdown := range filteredMarkdowns {
			wg.Add(1)
			go func(f string) {
				defer wg.Done()
				data, err := os.ReadFile(cloneDir + "/" + f)
				if err != nil {
					return
				}
				table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
				if err != nil {
					return
				}
				html := main.String()
				htmlList := table.String()
				relativePath := strings.ToLower(strings.TrimSuffix(f, ".md"))
				if isCreate {
					createParams.Append(relativePath, userID, repoID, html, htmlList)
				} else {
					updateParams.Append(relativePath, repoID, html, htmlList)
				}
			}(filteredMarkdown)
		}
		wg.Wait()
	}

	// Process added and modified files
	processFiles(addedFiles, true)
	processFiles(modifiedFiles, false)

	// Process deleted files
	filteredMarkdowns := util.FilterDiffFilesByExtensions(deletedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		deleteParams.Append(relativePath, repoID)
	}

	// Execute database operations
	if err := executeDBOperations(ctx, q, createParams, updateParams, deleteParams); err != nil {
		return err
	}

	// Generate layout
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
	if err := q.UpdateRepoLayout(ctx, arg_update_repo_layout); err != nil {
		return fmt.Errorf("update repo layout failed: %v", err)
	}

	fmt.Println("convert md repo to db: total execution time:", time.Since(startTime))

	return nil
}

// Helper function to execute database operations
func executeDBOperations(ctx context.Context, q *Queries, createParams *util.CreateParams, updateParams *util.UpdateParams, deleteParams *util.DeleteParams) error {
	if err := createMarkdownFiles(ctx, q, createParams); err != nil {
		return fmt.Errorf("create markdown failed: %v", err)
	}
	if err := updateMarkdownFiles(ctx, q, updateParams); err != nil {
		return fmt.Errorf("update markdown failed: %v", err)
	}
	if err := deleteMarkdownFiles(ctx, q, deleteParams); err != nil {
		return fmt.Errorf("delete markdown failed: %v", err)
	}
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
