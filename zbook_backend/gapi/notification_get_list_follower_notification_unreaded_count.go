package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListFollowerNotificationUnreadedCount(ctx context.Context, req *rpcs.GetListFollowerNotificationUnreadedCountRequest) (*rpcs.GetListFollowerNotificationUnreadedCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListFollowerNotificationUnreadedCount"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListFollowerNotificationUnreadedCount(ctx, authUser.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list follower notification unreaded count failed: %s", err)
	}

	rsp := &rpcs.GetListFollowerNotificationUnreadedCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
