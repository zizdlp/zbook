package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetQueryUserCount(ctx context.Context, req *rpcs.GetQueryUserCountRequest) (*rpcs.GetQueryUserCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetQueryUserCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	arg := db.GetQueryUserCountParams{
		Role:  authPayload.UserRole,
		Query: req.GetQuery(),
	}
	user_count, err := server.store.GetQueryUserCount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get query user count failed: %s", err)
	}

	rsp := &rpcs.GetQueryUserCountResponse{
		Count: user_count,
	}
	return rsp, nil

}
