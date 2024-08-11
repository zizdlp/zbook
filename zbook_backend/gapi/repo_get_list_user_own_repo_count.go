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

func (server *Server) GetListUserOwnRepoCount(ctx context.Context, req *rpcs.GetListUserOwnRepoCountRequest) (*rpcs.GetListUserOwnRepoCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetListUserOwnRepoCount"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateGetListUserOwnRepoCountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.GetQueryUserOwnRepoCountParams{
			UserID:    user.UserID,
			Signed:    true,
			CurUserID: authPayload.UserID,
			Role:      authPayload.UserRole,
			Query:     req.GetQuery(),
		}
		count, err := server.store.GetQueryUserOwnRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get query user own repo count failed: %s", err)
		}
		rsp := &rpcs.GetListUserOwnRepoCountResponse{
			Count: count,
		}
		return rsp, nil
	} else {
		arg := db.GetListUserOwnRepoCountParams{
			UserID:    user.UserID,
			Signed:    true,
			CurUserID: authPayload.UserID,
			Role:      authPayload.UserRole,
		}
		count, err := server.store.GetListUserOwnRepoCount(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get list user own repo count failed: %s", err)
		}
		rsp := &rpcs.GetListUserOwnRepoCountResponse{
			Count: count,
		}
		return rsp, nil
	}

}
func validateGetListUserOwnRepoCountRequest(req *rpcs.GetListUserOwnRepoCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
