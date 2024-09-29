package db

import (
	"context"
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/operations"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ManualSyncRepoTxParams struct {
	RepoID      int64
	AfterCreate func(cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error
}

func (store *SQLStore) ManualSyncRepoTx(ctx context.Context, arg ManualSyncRepoTxParams) error {
	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		repo, err := q.GetRepo(ctx, arg.RepoID)
		if err != nil {
			return status.Errorf(codes.Internal, "get repo failed: %s", err)
		}

		rsg := util.NewRandomStringGenerator()
		randomString := rsg.RandomString(32)
		cloneDir := "/tmp/zbook_repo/" + randomString
		// 删除目标目录以确保每次测试都是从头开始
		if _, err := os.Stat(cloneDir); err == nil {
			os.RemoveAll(cloneDir)
		}
		log.Info().Msgf("clone repo to:%s", cloneDir)
		startTime := time.Now()
		// 调用 Clone 函数
		gitURL := util.GetGitURL(repo.GitProtocol, repo.GitHost, repo.GitUsername, repo.GitRepo)
		if repo.GitAccessToken.Valid {
			if repo.GitHost == "github" {
				err = operations.CloneWithToken(gitURL, cloneDir, repo.GitAccessToken.String, repo.Branch)
				if err != nil {
					return status.Errorf(codes.Internal, "clone repo failed: %s", err)
				}
			} else {
				err = operations.CloneWithPassword(gitURL, cloneDir, repo.GitUsername, repo.GitAccessToken.String, repo.Branch)
				if err != nil {
					return status.Errorf(codes.Internal, "clone repo failed: %s", err)
				}
			}
		} else {
			err = operations.Clone(gitURL, cloneDir, repo.Branch)
			if err != nil {
				return status.Errorf(codes.Internal, "clone repo failed: %s", err)
			}
		}
		endTime := time.Now()

		// 计算耗时并输出
		elapsedTime := endTime.Sub(startTime)
		log.Info().Msgf("clone repo done, time consume:%s", elapsedTime)

		lastCommit, err := operations.GetLatestCommit(cloneDir)
		if err != nil {
			return err
		}

		// 调用 GetDiffFiles 函数
		addedFiles, modifiedFiles, deletedFiles, renameFiles, err := operations.GetDiffFiles(repo.CommitID, lastCommit, cloneDir)
		if err != nil {
			return err
		}
		err = ConvertFile2DB(ctx, q, cloneDir, repo.RepoID, repo.UserID, lastCommit, addedFiles, modifiedFiles, deletedFiles, renameFiles)
		if err != nil {
			return status.Errorf(codes.Internal, "failed to convert file to db: %s", err)
		}
		return arg.AfterCreate(cloneDir, repo.RepoID, repo.UserID, addedFiles, modifiedFiles, deletedFiles)

	})
	return err
}
