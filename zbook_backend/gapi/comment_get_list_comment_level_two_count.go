package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListCommentLevelTwoCount(ctx context.Context, req *rpcs.GetListCommentLevelTwoCountRequest) (*rpcs.GetListCommentLevelCountResponse, error) {

	apiUserDailyLimit := 10000
	apiKey := "GetListCommentLevelTwoCount"
	_, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user_count, err := server.store.GetListCommentLevelTwoCount(ctx, pgtype.Int8{Int64: req.GetRootId(), Valid: true})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get list comment level two count failed: %s", err)
	}

	rsp := &rpcs.GetListCommentLevelCountResponse{
		Count: user_count,
	}
	return rsp, nil
}
