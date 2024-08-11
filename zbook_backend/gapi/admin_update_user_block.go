package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUserBlock(ctx context.Context, req *rpcs.UpdateUserBlockRequest) (*rpcs.UpdateUserBlockResponse, error) {
	violations := validateUpdateUserBlockRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "UpdateUserBlock"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	user_set, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not exist: %s", err)
	}
	if user_set.UserRole == util.AdminRole {
		return nil, status.Errorf(codes.PermissionDenied, "admin user can not be blocked")
	}
	arg := db.UpdateUserBasicInfoParams{
		Username: req.GetUsername(),
		Blocked:  pgtype.Bool{Bool: req.GetBlocked(), Valid: true},
	}

	user, err := server.store.UpdateUserBasicInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "update user failed: %s", err)
	}

	rsp := &rpcs.UpdateUserBlockResponse{Blocked: user.Blocked}
	return rsp, nil
}
func validateUpdateUserBlockRequest(req *rpcs.UpdateUserBlockRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
