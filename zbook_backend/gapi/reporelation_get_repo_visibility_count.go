package gapi

import (
	"context"

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
	userCount, err := server.store.GetRepoVisibilityByRepoCount(ctx, req.GetRepoId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get repository visibility count: %s", err)
	}
	rsp := &rpcs.GetRepoVisibilityCountResponse{
		Count: userCount,
	}
	return rsp, nil

}
