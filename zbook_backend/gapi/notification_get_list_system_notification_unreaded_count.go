package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListSystemNotificationUnreadedCount(ctx context.Context, req *rpcs.GetListSystemNotificationUnreadedCountRequest) (*rpcs.GetListSystemNotificationUnreadedCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListSystemNotificationUnreadedCount"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListSystemNotificationUnReadedCount(ctx, authUser.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list system notification unreaded count failed: %s", err)
	}

	rsp := &rpcs.GetListSystemNotificationUnreadedCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
