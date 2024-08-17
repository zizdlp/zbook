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

func (server *Server) GetSelectedUserByRepoCount(ctx context.Context, req *rpcs.GetSelectedUserByRepoCountRequest) (*rpcs.GetSelectedUserByRepoCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetSelectedUserByRepoCount"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
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
		arg := db.GetQuerySelectedUserByRepoCountParams{
			RepoID: repo_id,
			Role:   authUser.UserRole,
			Query:  req.GetQuery(),
		}
		userCount, err := server.store.GetQuerySelectedUserByRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to GetQuerySelectedUserByRepoCount: %s", err)
		}
		rsp := &rpcs.GetSelectedUserByRepoCountResponse{
			Count: userCount,
		}
		return rsp, nil
	} else {
		arg := db.GetListSelectedUserByRepoCountParams{
			RepoID: repo_id,
			Role:   authUser.UserRole,
		}
		userCount, err := server.store.GetListSelectedUserByRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to GetListSelectedUserByRepoCount: %s", err)
		}
		rsp := &rpcs.GetSelectedUserByRepoCountResponse{
			Count: userCount,
		}
		return rsp, nil
	}

}
