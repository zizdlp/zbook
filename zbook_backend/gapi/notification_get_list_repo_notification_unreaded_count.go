package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListRepoNotificationUnreadedCount(ctx context.Context, req *rpcs.GetListRepoNotificationUnreadedCountRequest) (*rpcs.GetListRepoNotificationUnreadedCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListRepoNotificationUnreadedCount"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListRepoNotificationUnreadedCount(ctx, authUser.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list repo notification unreaded count failed: %s", err)
	}

	rsp := &rpcs.GetListRepoNotificationUnreadedCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
