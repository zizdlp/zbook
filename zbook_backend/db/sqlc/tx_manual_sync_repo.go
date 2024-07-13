package db

import (
	"context"
	"strconv"

	"github.com/zizdlp/zbook/gitsync"
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

		err = gitsync.Pull(repo.GitUsername, repo.GitAccessToken,
			"/tmp/wiki/", strconv.FormatInt(repo.RepoID, 10))
		if err != nil {
			if err.Error() == "already up-to-date" {
				return nil
			}
			return status.Errorf(codes.Internal, "无法同步仓库: %s", err)
		}
		rootPath := "/tmp/wiki/" + strconv.FormatInt(repo.RepoID, 10)
		err = ConvertFile2DB(ctx, q, rootPath, repo.RepoID, repo.UserID)

		if err != nil {
			return status.Errorf(codes.Internal, "无法转换文件数据到db: %s", err)
		}
		return nil
	})
	return err
}
