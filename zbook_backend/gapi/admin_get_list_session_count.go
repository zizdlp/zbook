package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListSessionCount(ctx context.Context, req *rpcs.GetListSessionCountRequest) (*rpcs.GetListSessionCountResponse, error) {

	apiUserDailyLimit := 10000
	apiKey := "GetListSessionCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		user_count, err := server.store.GetQuerySessionCount(ctx, req.GetQuery())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query active session count failed: %s", err)
		}

		rsp := &rpcs.GetListSessionCountResponse{
			Count: user_count,
		}
		return rsp, nil
	} else {
		user_count, err := server.store.GetListSessionCount(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list active session count failed: %s", err)
		}

		rsp := &rpcs.GetListSessionCountResponse{
			Count: user_count,
		}
		return rsp, nil
	}

}
