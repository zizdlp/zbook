package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListUserCount(ctx context.Context, req *rpcs.GetListUserCountRequest) (*rpcs.GetListUserCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListUserCount"
	_, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		user_count, err := server.store.GetQueryUserCount(ctx, req.GetQuery())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query user count failed: %s", err)
		}
		rsp := &rpcs.GetListUserCountResponse{
			Count: user_count,
		}
		return rsp, nil
	} else {
		user_count, err := server.store.GetListUserCount(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list user count failed: %s", err)
		}

		rsp := &rpcs.GetListUserCountResponse{
			Count: user_count,
		}
		return rsp, nil
	}

}
