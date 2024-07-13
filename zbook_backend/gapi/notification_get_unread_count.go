package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUnReadCount(ctx context.Context, req *rpcs.GetUnReadCountRequest) (*rpcs.GetUnReadCountResponse, error) {
	apiUserDailyLimit := 100000
	apiKey := "GetUnReadCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	unreadCount, err := server.store.GetUnReadCount(ctx, authPayload.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "unread count not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get unread count failed: %s", err)
	}

	rsp := &rpcs.GetUnReadCountResponse{
		UnreadCount: unreadCount,
	}
	return rsp, nil
}
