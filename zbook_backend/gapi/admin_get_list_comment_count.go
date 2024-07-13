package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListCommentCount(ctx context.Context, req *rpcs.GetListCommentCountRequest) (*rpcs.GetListCommentCountResponse, error) {

	apiUserDailyLimit := 10000
	apiKey := "GetListCommentCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		user_count, err := server.store.GetQueryCommentCount(ctx, req.GetQuery())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query comment count failed: %s", err)
		}

		rsp := &rpcs.GetListCommentCountResponse{
			Count: user_count,
		}
		return rsp, nil
	} else {
		user_count, err := server.store.GetListCommentCount(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list comment count failed: %s", err)
		}
		rsp := &rpcs.GetListCommentCountResponse{
			Count: user_count,
		}
		return rsp, nil
	}

}
