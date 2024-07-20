package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) ListUserOwnRepo(ctx context.Context, req *rpcs.ListUserOwnRepoRequest) (*rpcs.ListUserOwnRepoResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListUserOwnRepo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateListUserOwnRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.QueryUserOwnRepoParams{
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Signed:    true,
			CurUserID: authPayload.UserID,
			UserID:    user.UserID,
			Role:      authPayload.UserRole,
			Query:     req.GetQuery(),
		}

		repos, err := server.store.QueryUserOwnRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query user own repo failed: %s", err)
		}

		rsp := &rpcs.ListUserOwnRepoResponse{
			Elements: convertQueryUserOwnRepo(repos, req.GetUsername()),
		}
		return rsp, nil
	} else {
		arg := db.ListUserOwnRepoParams{
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Signed:    true,
			CurUserID: authPayload.UserID,
			UserID:    user.UserID,
			Role:      authPayload.UserRole,
		}

		repos, err := server.store.ListUserOwnRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "list user own repo failed: %s", err)
		}

		rsp := &rpcs.ListUserOwnRepoResponse{
			Elements: convertListUserOwnRepo(repos, req.GetUsername()),
		}
		return rsp, nil
	}

}
func validateListUserOwnRepoRequest(req *rpcs.ListUserOwnRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidateInt32ID(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListUserOwnRepo(repos []db.ListUserOwnRepoRow, username string) []*models.ListRepoInfo {
	var ret_repos []*models.ListRepoInfo
	for i := 0; i < len(repos); i++ {
		ret_repos = append(ret_repos,
			&models.ListRepoInfo{
				RepoId:          repos[i].RepoID,
				RepoName:        repos[i].RepoName,
				Username:        username,
				RepoDescription: repos[i].RepoDescription,
				VisibilityLevel: repos[i].VisibilityLevel,
				GitHost:         repos[i].GitHost,
				LikeCount:       int32(repos[i].LikeCount),
				IsLiked:         repos[i].IsLiked,
				UpdatedAt:       timestamppb.New(repos[i].UpdatedAt),
				CreatedAt:       timestamppb.New(repos[i].CreatedAt),
			},
		)
	}
	return ret_repos
}

func convertQueryUserOwnRepo(repos []db.QueryUserOwnRepoRow, username string) []*models.ListRepoInfo {
	var ret_repos []*models.ListRepoInfo
	for i := 0; i < len(repos); i++ {
		ret_repos = append(ret_repos,
			&models.ListRepoInfo{
				RepoId:          repos[i].RepoID,
				RepoName:        repos[i].RepoName,
				Username:        username,
				RepoDescription: repos[i].RepoDescription,
				VisibilityLevel: repos[i].VisibilityLevel,
				GitHost:         repos[i].GitHost,
				IsLiked:         repos[i].IsLiked,
				LikeCount:       int32(repos[i].LikeCount),
				UpdatedAt:       timestamppb.New(repos[i].UpdatedAt),
				CreatedAt:       timestamppb.New(repos[i].CreatedAt),
			},
		)
	}
	return ret_repos
}
