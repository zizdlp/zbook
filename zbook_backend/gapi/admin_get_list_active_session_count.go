package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListActiveSessionCount(ctx context.Context, req *rpcs.GetListActiveSessionCountRequest) (*rpcs.GetListActiveSessionCountResponse, error) {

	apiUserDailyLimit := 10000
	apiKey := "GetListActiveSessionCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		user_count, err := server.store.GetQueryActiveSessionCount(ctx, req.GetQuery())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query active session count failed: %s", err)
		}

		rsp := &rpcs.GetListActiveSessionCountResponse{
			Count: user_count,
		}
		return rsp, nil
	} else {
		user_count, err := server.store.GetListActiveSessionCount(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list active session count failed: %s", err)
		}

		rsp := &rpcs.GetListActiveSessionCountResponse{
			Count: user_count,
		}
		return rsp, nil
	}

}
