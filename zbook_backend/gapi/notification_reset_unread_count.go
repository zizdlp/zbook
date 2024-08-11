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

func (server *Server) ResetUnreadCount(ctx context.Context, req *rpcs.ResetUnreadCountRequest) (*rpcs.ResetUnreadCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ResetUnreadCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	err = server.store.ResetUnreadCount(ctx, authPayload.Username)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "unreaded count not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "reset unreaded count failed: %s", err)
	}

	rsp := &rpcs.ResetUnreadCountResponse{}
	return rsp, nil
}
