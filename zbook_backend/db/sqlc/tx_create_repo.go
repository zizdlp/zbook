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

// with go modules disabled
type CreateRepoTxParams struct {
	CreateRepoParams
	Username    string
	AfterCreate func(cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error
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
		gitURL := util.GetGitURL(result.Repo.GitProtocol, result.Repo.GitHost, result.Repo.GitUsername, result.Repo.GitRepo)
		if arg.GitAccessToken.Valid {
			if result.Repo.GitHost == "github" {
				err = operations.CloneWithToken(gitURL, cloneDir, arg.GitAccessToken.String, arg.Branch)
				if err != nil {
					return status.Errorf(codes.Internal, "clone repo failed: %s", err)
				}
			} else {
				err = operations.CloneWithPassword(gitURL, cloneDir, arg.GitUsername, arg.GitAccessToken.String, arg.Branch)
				if err != nil {
					return status.Errorf(codes.Internal, "clone repo failed: %s", err)
				}
			}
		} else {
			err = operations.Clone(gitURL, cloneDir, arg.Branch)
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
		addedFiles, modifiedFiles, deletedFiles, renameFiles, err := operations.GetDiffFiles("", lastCommit, cloneDir)
		if err != nil {
			return err
		}

		err = ConvertFile2DB(ctx, q, cloneDir, result.Repo.RepoID, arg.UserID, lastCommit, addedFiles, modifiedFiles, deletedFiles, renameFiles)
		if err != nil {
			return status.Errorf(codes.Internal, "无法转换文件数据到db: %s", err)
		}

		user, err := q.GetUserByUsername(ctx, arg.Username)
		if err != nil {
			return nil
		}
		if !user.Blocked && arg.VisibilityLevel == "public" || arg.VisibilityLevel == "signin" {

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
						return nil
					} else {
						err = q.UpdateUnreadCount(ctx, follows[i].UserID)
						if err != nil {
							return nil
						}
					}
				}
			}
		}
		return arg.AfterCreate(cloneDir, result.Repo.RepoID, result.Repo.UserID, addedFiles, modifiedFiles, deletedFiles)
	})
	return result, err
}
