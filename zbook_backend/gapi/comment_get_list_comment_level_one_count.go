package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListCommentLevelOneCount(ctx context.Context, req *rpcs.GetListCommentLevelOneCountRequest) (*rpcs.GetListCommentLevelCountResponse, error) {

	apiUserDailyLimit := 10000
	apiKey := "GetListCommentLevelOneCount"
	_, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListCommentLevelOneCount(ctx, req.GetMarkdownId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list level one comment count failed: %s", err)
	}

	rsp := &rpcs.GetListCommentLevelCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
