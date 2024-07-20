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

func (server *Server) ListUserLikeRepo(ctx context.Context, req *rpcs.ListUserLikeRepoRequest) (*rpcs.ListUserLikeRepoResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListUserLikeRepo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateListUserLikeRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err = server.getUserPermessionlevel(ctx, authPayload.Username, req.GetUsername())
	if err != nil {
		return nil, err
	}
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.QueryUserLikeRepoParams{
			UserID:    user.UserID,
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Signed:    true,
			Role:      authPayload.UserRole,
			CurUserID: authPayload.UserID,
			Query:     req.GetQuery(),
		}

		repos, err := server.store.QueryUserLikeRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query user like repo failed: %s", err)
		}

		rsp := &rpcs.ListUserLikeRepoResponse{
			Elements: convertQueryUserLikeRepo(repos),
		}
		return rsp, nil
	} else {
		arg := db.ListUserLikeRepoParams{
			UserID:    user.UserID,
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Signed:    true,
			Role:      authPayload.UserRole,
			CurUserID: authPayload.UserID,
		}

		repos, err := server.store.ListUserLikeRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "list user like repo failed: %s", err)
		}

		rsp := &rpcs.ListUserLikeRepoResponse{
			Elements: convertListUserLikeRepo(repos),
		}
		return rsp, nil
	}

}
func validateListUserLikeRepoRequest(req *rpcs.ListUserLikeRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidateInt32ID(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListUserLikeRepo(repos []db.ListUserLikeRepoRow) []*models.ListRepoInfo {
	var ret_repos []*models.ListRepoInfo
	for i := 0; i < len(repos); i++ {
		ret_repos = append(ret_repos,
			&models.ListRepoInfo{
				RepoId:          repos[i].RepoID,
				Username:        repos[i].Username,
				RepoName:        repos[i].RepoName,
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
func convertQueryUserLikeRepo(repos []db.QueryUserLikeRepoRow) []*models.ListRepoInfo {
	var ret_repos []*models.ListRepoInfo
	for i := 0; i < len(repos); i++ {
		ret_repos = append(ret_repos,
			&models.ListRepoInfo{
				RepoId:          repos[i].RepoID,
				RepoName:        repos[i].RepoName,
				Username:        repos[i].Username,
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
