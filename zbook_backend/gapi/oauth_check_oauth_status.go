package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CheckOAuthStatus(ctx context.Context, req *rpcs.CheckOAuthStatusRequest) (*rpcs.CheckOAuthStatusResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "CheckOAuthStatus"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	third_status, err := server.store.CheckOAuthStatus(ctx, authPayload.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "check third status failed: %s", err)
	}

	rsp := &rpcs.CheckOAuthStatusResponse{
		Github: third_status.GithubStatus,
		Google: third_status.GoogleStatus,
	}
	return rsp, nil
}
