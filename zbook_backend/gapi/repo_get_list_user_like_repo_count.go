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

func (server *Server) GetListUserLikeRepoCount(ctx context.Context, req *rpcs.GetListUserLikeRepoCountRequest) (*rpcs.GetListUserLikeRepoCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListUserLikeRepoCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateGetListUserLikeRepoCountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.GetQueryUserLikeRepoCountParams{
			UserID:    user.UserID,
			Signed:    true,
			CurUserID: authPayload.UserID,
			Role:      authPayload.UserRole,
			Query:     req.GetQuery(),
		}
		count, err := server.store.GetQueryUserLikeRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query user like repo count failed: %s", err)
		}

		rsp := &rpcs.GetListUserLikeRepoCountResponse{
			Count: count,
		}
		return rsp, nil
	} else {
		arg := db.GetListUserLikeRepoCountParams{
			UserID:    user.UserID,
			Signed:    true,
			CurUserID: authPayload.UserID,
			Role:      authPayload.UserRole,
		}
		count, err := server.store.GetListUserLikeRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list user like repo count failed: %s", err)
		}

		rsp := &rpcs.GetListUserLikeRepoCountResponse{
			Count: count,
		}
		return rsp, nil
	}

}
func validateGetListUserLikeRepoCountRequest(req *rpcs.GetListUserLikeRepoCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
