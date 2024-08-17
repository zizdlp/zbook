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

func (server *Server) QueryUserByRepo(ctx context.Context, req *rpcs.QueryUserByRepoRequest) (*rpcs.QueryUserByRepoResponse, error) {
	violations := validateQueryUserByRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "QueryUserByRepo"
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

	arg := db.QueryUserByRepoParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
		RepoID: repo_id,
		Role:   authUser.UserRole,
		Query:  req.GetQuery(),
	}

	users, err := server.store.QueryUserByRepo(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "QueryUserByRepo failed: %s", err)
	}

	rsp := &rpcs.QueryUserByRepoResponse{
		Elements: convertQueryUserByRepo(users),
	}
	return rsp, nil

}
func validateQueryUserByRepoRequest(req *rpcs.QueryUserByRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertQueryUserByRepo(users []db.QueryUserByRepoRow) []*models.ListUserRepoVisiblityInfo {
	var ret_users []*models.ListUserRepoVisiblityInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListUserRepoVisiblityInfo{
				Username:      users[i].Username,
				Email:         users[i].Email,
				CreatedAt:     timestamppb.New(users[i].CreatedAt),
				UpdatedAt:     timestamppb.New(users[i].UpdatedAt),
				IsRepoVisible: users[i].IsVisible,
			},
		)
	}
	return ret_users
}
