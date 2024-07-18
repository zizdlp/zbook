package db

import (
	"context"
	"os"

	"github.com/zizdlp/zbook/operations"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ManualSyncRepoTxParams struct {
	RepoID int64
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

		// 调用 Clone 函数
		gitURL := util.GetGitURL(repo.GitProtocol, repo.GitHost, repo.GitUsername, repo.GitRepo)
		if repo.GitAccessToken.Valid {
			err = operations.CloneWithPassword(gitURL, cloneDir, repo.GitUsername, repo.GitAccessToken.String)
			if err != nil {
				return status.Errorf(codes.Internal, "clone repo failed: %s", err)
			}
		} else {
			err = operations.Clone(gitURL, cloneDir)
			if err != nil {
				return status.Errorf(codes.Internal, "clone repo failed: %s", err)
			}
		}
		err = ConvertFile2DB(ctx, q, cloneDir, repo.RepoID, repo.UserID, repo.CommitID)
		if err != nil {
			return status.Errorf(codes.Internal, "无法转换文件数据到db: %s", err)
		}

		return nil
	})
	return err
}
