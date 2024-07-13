package db

import (
	"context"
	"fmt"
	"strconv"

	"github.com/zizdlp/zbook/gitsync"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// with go modules disabled
type CreateRepoTxParams struct {
	CreateRepoParams
}

type CreateRepoTxResult struct {
	Repo Repo
}

func (store *SQLStore) CreateRepoTx(ctx context.Context, arg CreateRepoTxParams) (CreateRepoTxResult, error) {
	var result CreateRepoTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Repo, err = q.CreateRepo(ctx, arg.CreateRepoParams)
		if err != nil {
			return err
		}

		err = gitsync.Clone(result.Repo.GitProtocol, result.Repo.GitHost, result.Repo.GitUsername, result.Repo.GitRepo, result.Repo.GitAccessToken,
			"/tmp/wiki/", strconv.FormatInt(result.Repo.RepoID, 10))
		if err != nil {
			return status.Errorf(codes.Internal, "无法克隆仓库: %s", err)
		}
		rootPath := "/tmp/wiki/" + strconv.FormatInt(result.Repo.RepoID, 10)
		err = ConvertFile2DB(ctx, q, rootPath, result.Repo.RepoID, arg.UserID)
		if err != nil {
			return status.Errorf(codes.Internal, "无法转换文件数据到db: %s", err)
		}

		user, err := q.GetUserByID(ctx, arg.UserID)
		if err != nil {
			return nil
		}
		if !user.Blocked && !user.Deleted && arg.VisibilityLevel == "public" || arg.VisibilityLevel == "signin" {

			arg := GetListFollowerCountParams{
				CurUserID: arg.UserID,
				UserID:    arg.UserID,
			}

			count, err := store.GetListFollowerCount(ctx, arg)
			if err != nil {
				return nil
			}

			for page := 1; page <= int((count+9)/10); page++ {
				arg := ListFollowerParams{
					CurUserID: arg.UserID,
					UserID:    arg.UserID,
					Limit:     10,
					Offset:    int32((page - 1) * 10),
				}
				follows, err := q.ListFollower(ctx, arg)
				if err != nil {
					break
				}
				for i := 0; i < len(follows); i++ {
					arg_noti_follower := CreateRepoNotificationParams{
						UserID: follows[i].UserID,
						RepoID: result.Repo.RepoID,
					}
					_, err = q.CreateRepoNotification(ctx, arg_noti_follower)
					if err != nil {
						fmt.Println("mydebug:create repo noti error:", err)
						return nil
					} else {
						err = q.UpdateUnreadCount(ctx, follows[i].UserID)
						if err != nil {
							fmt.Println("mydebug:update unread count error:", err)
							return nil
						}
					}
				}
			}
		}

		return nil
	})
	return result, err
}
