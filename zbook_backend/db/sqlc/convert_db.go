package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/operations"
	"github.com/zizdlp/zbook/util"
)

func ConvertFile2DB(ctx context.Context, q *Queries, cloneDir string, repoID int64, userID int64, lastCommit string, addedFiles []string, modifiedFiles []string, deletedFiles []string, renameFiles []string) error {
	startTime := time.Now()
	allowedExtensions := map[string]bool{
		".md": true,
	}

	createParams := &util.CreateParams{}
	updateParams := &util.UpdateParams{}
	deleteParams := &util.DeleteParams{}

	// Helper function to process files
	processFiles := func(files []string, isCreate bool) {
		filteredMarkdowns := util.FilterDiffFilesByExtensions(files, allowedExtensions)
		for _, filteredMarkdown := range filteredMarkdowns {
			data, err := os.ReadFile(cloneDir + "/" + filteredMarkdown)
			if err != nil {
				continue
			}
			// table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
			// if err != nil {
			// 	continue
			// }
			// html := main.String()
			// htmlList := table.String()
			relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
			if isCreate {
				createParams.Append(relativePath, userID, repoID, string(data))
			} else {
				updateParams.Append(relativePath, relativePath, repoID, string(data))
			}
		}
	}

	// Process added and modified files
	processFiles(addedFiles, true)
	processFiles(modifiedFiles, false)
	log.Info().Msgf("addfiles: %s", addedFiles)
	log.Info().Msgf("modifiedFiles: %s", modifiedFiles)
	log.Info().Msgf("deletedFiles: %s", deletedFiles)
	log.Info().Msgf("renameFiles: %s", renameFiles)
	// Process deleted files
	filteredMarkdowns := util.FilterDiffFilesByExtensions(deletedFiles, allowedExtensions)
	for _, filteredMarkdown := range filteredMarkdowns {
		relativePath := strings.ToLower(strings.TrimSuffix(filteredMarkdown, ".md"))
		deleteParams.Append(relativePath, repoID)
	}
	// Process renamed files
	for i := 0; i < len(renameFiles); i += 2 {
		relativePath := strings.ToLower(strings.TrimSuffix(renameFiles[i], ".md"))
		newrelativePath := strings.ToLower(strings.TrimSuffix(renameFiles[i+1], ".md"))
		data, err := os.ReadFile(cloneDir + "/" + renameFiles[i+1])
		if err != nil {
			continue
		}
		// table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
		// if err != nil {
		// 	continue
		// }
		// html := main.String()
		// htmlList := table.String()
		updateParams.Append(relativePath, newrelativePath, repoID, string(data))
	}

	// Execute database operations
	if err := executeDBOperations(ctx, q, createParams, updateParams, deleteParams); err != nil {
		return err
	}

	configFromFile, err := util.ReadRepoConfig(cloneDir + "/" + "zbook.json")
	if err != nil {
		// Generate config
		mdFiles, err := operations.ListMarkdownFiles(cloneDir)
		if err != nil {
			return fmt.Errorf("read layout failed: %v", err)
		}
		config := &util.RepoConfig{}
		layout := util.CreateLayout(mdFiles)

		// 存储 layout 到 "default" key 下
		config.Layout = map[string][]util.Layout{
			"default": layout,
		}
		configJSON, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			return fmt.Errorf("generate repo config failed: %v", err)
		}
		home, err := config.GetFirstDocumentMap()
		if err != nil {
			return fmt.Errorf("generate repo home failed: %v", err)
		}
		// 将 docs 转换为 JSON 字符串
		homeJSON, err := json.Marshal(home)
		if err != nil {
			return fmt.Errorf("marshal home docs to JSON failed: %v", err)
		}
		arg_update_repo_config := UpdateRepoConfigParams{
			RepoID:   repoID,
			Config:   string(configJSON),
			CommitID: lastCommit,
			Home:     string(homeJSON),
		}
		if err := q.UpdateRepoConfig(ctx, arg_update_repo_config); err != nil {
			return fmt.Errorf("update repo config failed: %v", err)
		}
		log.Info().Msgf("convert md repo to db: total execution time:%s", time.Since(startTime))
		return nil
	}
	configJSON, err := json.MarshalIndent(configFromFile, "", "  ")
	if err != nil {
		return fmt.Errorf("generate repo config failed: %v", err)
	}
	home, err := configFromFile.GetFirstDocumentMap()
	if err != nil {
		return fmt.Errorf("generate repo home failed: %v", err)
	}
	// 将 docs 转换为 JSON 字符串
	homeJSON, err := json.Marshal(home)
	if err != nil {
		return fmt.Errorf("marshal home docs to JSON failed: %v", err)
	}
	arg_update_repo_config := UpdateRepoConfigParams{
		RepoID:   repoID,
		Config:   string(configJSON),
		CommitID: lastCommit,
		Home:     string(homeJSON),
	}
	if err := q.UpdateRepoConfig(ctx, arg_update_repo_config); err != nil {
		return fmt.Errorf("update repo config failed: %v", err)
	}
	log.Info().Msgf("convert md repo to db: total execution time:%s", time.Since(startTime))
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
		Content:      params.Content,
	}
	err := q.CreateMarkdownMulti(ctx, argCreate)
	if err != nil {
		return err
	}
	return nil
}

func updateMarkdownFiles(ctx context.Context, q *Queries, params *util.UpdateParams) error {
	argUpdate := UpdateMarkdownMultiParams{
		RelativePath:    params.RelativePath,
		NewRelativePath: params.NewRelativePath,
		RepoID:          params.RepoID,
		Content:         params.Content,
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
