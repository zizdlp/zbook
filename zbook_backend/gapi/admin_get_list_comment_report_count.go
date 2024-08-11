package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListCommentReportCount(ctx context.Context, req *rpcs.GetListCommentReportCountRequest) (*rpcs.GetListCommentReportCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListCommentReportCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListCommentReportCount(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list comment report count failed: %s", err)
	}

	rsp := &rpcs.GetListCommentReportCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
