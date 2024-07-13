package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetFollowerCount(ctx context.Context, req *rpcs.GetFollowerCountRequest) (*rpcs.GetFollowerCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetFollowerCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateGetFollowerCountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.GetQueryFollowerCountParams{
			CurUserID: authPayload.UserID,
			UserID:    user.UserID,
			Query:     req.GetQuery(),
		}
		count, err := server.store.GetQueryFollowerCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query follower count failed: %s", err)
		}
		rsp := &rpcs.GetFollowerCountResponse{
			Count: count,
		}
		return rsp, nil
	} else {
		arg := db.GetListFollowerCountParams{
			CurUserID: authPayload.UserID,
			UserID:    user.UserID,
		}

		count, err := server.store.GetListFollowerCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list follower count failed : %s", err)
		}

		rsp := &rpcs.GetFollowerCountResponse{
			Count: count,
		}
		return rsp, nil
	}

}
func validateGetFollowerCountRequest(req *rpcs.GetFollowerCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
