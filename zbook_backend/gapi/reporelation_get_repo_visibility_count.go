package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRepoVisibilityCount(ctx context.Context, req *rpcs.GetRepoVisibilityCountRequest) (*rpcs.GetRepoVisibilityCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetRepoVisibilityCount"
	_, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
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
	userCount, err := server.store.GetRepoVisibilityByRepoCount(ctx, repo_id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get repository visibility count: %s", err)
	}
	rsp := &rpcs.GetRepoVisibilityCountResponse{
		Count: userCount,
	}
	return rsp, nil

}
