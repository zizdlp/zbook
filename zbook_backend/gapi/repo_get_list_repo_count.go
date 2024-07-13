package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetListRepoCount(ctx context.Context, req *rpcs.GetListRepoCountRequest) (*rpcs.GetListRepoCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListRepoCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		// is unsigned

		if req.GetQuery() != "" {
			arg := db.GetQueryRepoCountParams{
				Query:     req.GetQuery(),
				Role:      util.UserRole,
				Signed:    false,
				CurUserID: 0,
			}
			count, err := server.store.GetQueryRepoCount(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "get query repo count failed: %s", err)
			}
			rsp := &rpcs.GetListRepoCountResponse{
				Count: count,
			}
			return rsp, nil
		} else {
			arg := db.GetListRepoCountParams{
				Signed:    false,
				CurUserID: 0,
				Role:      util.UserRole,
			}
			user_count, err := server.store.GetListRepoCount(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "get list repo count failed: %s", err)
			}
			rsp := &rpcs.GetListRepoCountResponse{
				Count: user_count,
			}
			return rsp, nil
		}

	} else {
		if req.GetQuery() != "" {
			arg := db.GetQueryRepoCountParams{
				Query:     req.GetQuery(),
				Role:      authPayload.UserRole,
				Signed:    true,
				CurUserID: authPayload.UserID,
			}
			count, err := server.store.GetQueryRepoCount(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "get query repo count failed: %s", err)
			}
			rsp := &rpcs.GetListRepoCountResponse{
				Count: count,
			}
			return rsp, nil
		} else {
			arg := db.GetListRepoCountParams{
				Signed:    true,
				CurUserID: authPayload.UserID,
				Role:      authPayload.UserRole,
			}
			user_count, err := server.store.GetListRepoCount(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "get list repo count failed: %s", err)
			}

			rsp := &rpcs.GetListRepoCountResponse{
				Count: user_count,
			}
			return rsp, nil
		}

	}

}
