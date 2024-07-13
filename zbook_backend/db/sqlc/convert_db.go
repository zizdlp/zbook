package db

import (
	"context"
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/yuin/goldmark"
	convert "github.com/zizdlp/zbook/markdown/convert"
	md "github.com/zizdlp/zbook/markdown/render"
	"github.com/zizdlp/zbook/util"
)

func renderFile(prefixPath string, absolutePath string, repoID int64, userID int64, versionKey string, markdown goldmark.Markdown,
	queriesMd5 []QueryMd5ForCheckRow, createParams *CreateParams, updateParams *UpdateParams, keyParams *KeyParams) error {
	data, err := os.ReadFile(absolutePath)
	if err != nil {
		return err
	}

	// 获取相对路径
	relativePath := strings.TrimPrefix(absolutePath, prefixPath)
	relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
	relativePath = strings.ToLower(strings.TrimSuffix(relativePath, ".md")) // Remove ".md" suffix

	// 计算MD5
	newmd5 := md5.Sum(data)
	md5str := fmt.Sprintf("%x", newmd5)

	// 查找文件状态
	var fileType int
	shouldKey := BinarySearch(queriesMd5, len(queriesMd5), relativePath)
	if shouldKey < 0 || shouldKey >= len(queriesMd5) || queriesMd5[shouldKey].RelativePath != relativePath {
		fileType = 0 // 文件不存在
	} else if queriesMd5[shouldKey].Md5 != md5str {
		fileType = 1 // 文件存在但需要更新
	} else {
		fileType = 2 // 文件存在且无需更新
	}

	// 处理文件
	if fileType != 2 {
		table, main, err := convert.ConvertMarkdownBuffer(data, markdown)
		if err != nil {
			fmt.Println("convert markdown to buffer error:", err)
			return err
		}
		html := main.String()
		htmlList := table.String()
		if fileType == 1 { //更新
			updateParams.Append(relativePath, repoID, html, htmlList, md5str, versionKey)
		} else { //创建
			createParams.Append(relativePath, userID, repoID, html, htmlList, md5str, versionKey)
		}
	} else { //更新版本号
		keyParams.Append(relativePath, repoID, versionKey)
	}
	return nil
}

// 遍历path下面所有文件，与markdown:[relative_path,md5] 对比
func iterFS(q *Queries, markdown goldmark.Markdown, prefixPath string, absolutePath string, updateky string, RepoID int64, UserID int64, queries_md5 []QueryMd5ForCheckRow,
	createParams *CreateParams, updateParams *UpdateParams, keyParams *KeyParams) error {
	files, err := os.ReadDir(absolutePath)
	if err != nil {
		return err
	}
	for _, f := range files {
		if f.Type().IsRegular() && strings.HasSuffix(f.Name(), ".md") {
			if err := renderFile(prefixPath, absolutePath+"/"+f.Name(), RepoID, UserID, updateky, markdown, queries_md5, createParams, updateParams, keyParams); err != nil {
				return err
			}
		} else if f.IsDir() {
			if err := iterFS(q, markdown, prefixPath, absolutePath+"/"+f.Name(), updateky, RepoID, UserID, queries_md5, createParams, updateParams, keyParams); err != nil {
				return err
			}
		}
	}
	return nil
}
func ConvertFile2DB(ctx context.Context, q *Queries, rootPath string, repoID int64, userID int64) error {
	startTime := time.Now()

	versionKey := util.RandomString(32)
	markdown := md.GetMarkdownConfig()

	queriesMD5, err := q.QueryMd5ForCheck(ctx, repoID)
	if err != nil {
		return fmt.Errorf("error querying MD5: %v", err)
	}

	createParams := &CreateParams{}
	updateParams := &UpdateParams{}
	keyParams := &KeyParams{}

	err = iterFS(q, markdown, rootPath, rootPath, versionKey, repoID, userID, queriesMD5, createParams, updateParams, keyParams)
	if err != nil {
		return fmt.Errorf("error iterating through filesystem: %v", err)
	}

	err = createMarkdownFiles(ctx, q, createParams)
	if err != nil {
		return fmt.Errorf("error creating markdown files: %v", err)
	}

	err = updateMarkdownFiles(ctx, q, updateParams)
	if err != nil {
		return fmt.Errorf("error updating markdown files: %v", err)
	}

	err = updateMarkdownVersionKeys(ctx, q, keyParams)
	if err != nil {
		return fmt.Errorf("error updating markdown version keys: %v", err)
	}

	err = deleteOldMarkdownFiles(ctx, q, repoID, versionKey)
	if err != nil {
		return fmt.Errorf("error deleting old markdown files: %v", err)
	}

	err = RenderLayout(rootPath, repoID, q, versionKey)
	if err != nil {
		return fmt.Errorf("error rendering layout: %v", err)
	}

	endTime := time.Now()
	fmt.Println("mydebug: total execution time:", endTime.Sub(startTime))

	return nil
}

func createMarkdownFiles(ctx context.Context, q *Queries, params *CreateParams) error {
	fmt.Println("Creating markdown files:", len(params.Md5))
	argCreate := CreateMarkdownMultiParams{
		RelativePath: params.RelativePath,
		UserID:       params.UserID,
		RepoID:       params.RepoID,
		MainContent:  params.MainContent,
		TableContent: params.TableContent,
		Md5:          params.Md5,
		VersionKey:   params.VersionKey,
	}
	err := q.CreateMarkdownMulti(ctx, argCreate)
	if err != nil {
		return err
	}
	return nil
}

func updateMarkdownFiles(ctx context.Context, q *Queries, params *UpdateParams) error {
	fmt.Println("Updating markdown files:", len(params.Md5))
	argUpdate := UpdateMarkdownMultiParams{
		RelativePath: params.RelativePath,
		RepoID:       params.RepoID,
		MainContent:  params.MainContent,
		TableContent: params.TableContent,
		Md5:          params.Md5,
		VersionKey:   params.VersionKey,
	}
	err := q.UpdateMarkdownMulti(ctx, argUpdate)
	if err != nil {
		return err
	}
	return nil
}

func updateMarkdownVersionKeys(ctx context.Context, q *Queries, params *KeyParams) error {
	fmt.Println("Updating markdown version keys:", len(params.VersionKey))
	argKey := UpdateMarkdownVersionKeyParams{
		RelativePath: params.RelativePath,
		RepoID:       params.RepoID,
		VersionKey:   params.VersionKey,
	}
	err := q.UpdateMarkdownVersionKey(ctx, argKey)
	if err != nil {
		return err
	}
	return nil
}

func deleteOldMarkdownFiles(ctx context.Context, q *Queries, repoID int64, versionKey string) error {
	fmt.Println("Deleting old markdown files")
	argCleanDB := DeleteOldMarkdownParams{
		RepoID:     repoID,
		VersionKey: versionKey,
	}
	err := q.DeleteOldMarkdown(ctx, argCleanDB)
	if err != nil {
		return err
	}
	return nil
}
