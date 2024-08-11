package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListCommentNotificationUnreadedCount(ctx context.Context, req *rpcs.GetListCommentNotificationUnreadedCountRequest) (*rpcs.GetListCommentNotificationUnreadedCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListCommentNotificationUnreadedCount"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListCommentNotificationUnreadedCount(ctx, authUser.UserID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list comment notification unreaded count failed: %s", err)
	}

	rsp := &rpcs.GetListCommentNotificationUnreadedCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
