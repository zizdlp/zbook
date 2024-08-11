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

func (server *Server) GetFollowStatus(ctx context.Context, req *rpcs.GetFollowStatusRequest) (*rpcs.GetFollowStatusResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetFollowStatus"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateGetFollowStatusRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}

	arg := db.IsFollowingParams{
		FollowingID: user.UserID,
		FollowerID:  authPayload.UserID,
	}
	IsFollowing, err := server.store.IsFollowing(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get isfollowing status failed: %s", err)
	}

	rsp := &rpcs.GetFollowStatusResponse{
		IsFollowing: IsFollowing,
	}
	return rsp, nil
}
func validateGetFollowStatusRequest(req *rpcs.GetFollowStatusRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
