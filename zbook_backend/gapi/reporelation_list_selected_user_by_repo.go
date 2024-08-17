package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
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

func (server *Server) ListSelectedUserByRepo(ctx context.Context, req *rpcs.ListSelectedUserByRepoRequest) (*rpcs.ListSelectedUserByRepoResponse, error) {
	violations := validateListSelectedUserByRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListSelectedUserByRepo"
	authUser, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	arg_repo := db.GetRepoIDParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo_id, err := server.store.GetRepoID(ctx, arg_repo)
	if err != nil {
		log.Info().Msgf("get repo layout get repo id failed:%s,%s", req.GetUsername(), req.GetRepoName())
		return nil, status.Errorf(codes.Internal, "get repo id failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.QuerySelectedUserByRepoParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			RepoID: repo_id,
			Query:  req.GetQuery(),
			Role:   authUser.UserRole,
		}

		users, err := server.store.QuerySelectedUserByRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query repo visisiblity by repo failed: %s", err)
		}
		rsp := &rpcs.ListSelectedUserByRepoResponse{
			Elements: convertQueryRepoVisibility(users),
		}
		return rsp, nil
	} else {
		arg := db.ListSelectedUserByRepoParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			RepoID: repo_id,
			Role:   authUser.UserRole,
		}

		users, err := server.store.ListSelectedUserByRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "list repo visibilisty by repo failed: %s", err)
		}

		rsp := &rpcs.ListSelectedUserByRepoResponse{
			Elements: convertListSelectedUserByRepo(users),
		}
		return rsp, nil
	}

}
func validateListSelectedUserByRepoRequest(req *rpcs.ListSelectedUserByRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertListSelectedUserByRepo(users []db.User) []*models.ListUserRepoVisiblityInfo {
	var ret_users []*models.ListUserRepoVisiblityInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListUserRepoVisiblityInfo{
				Username:      users[i].Username,
				Email:         users[i].Email,
				CreatedAt:     timestamppb.New(users[i].CreatedAt),
				UpdatedAt:     timestamppb.New(users[i].UpdatedAt),
				IsRepoVisible: true,
			},
		)
	}
	return ret_users
}
func convertQueryRepoVisibility(users []db.QuerySelectedUserByRepoRow) []*models.ListUserRepoVisiblityInfo {
	var ret_users []*models.ListUserRepoVisiblityInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListUserRepoVisiblityInfo{
				Username:      users[i].Username,
				Email:         users[i].Email,
				CreatedAt:     timestamppb.New(users[i].CreatedAt),
				UpdatedAt:     timestamppb.New(users[i].UpdatedAt),
				IsRepoVisible: true,
			},
		)
	}
	return ret_users
}
