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

func (server *Server) MarkUserAsDeleted(ctx context.Context, req *rpcs.MarkUserAsDeletedRequest) (*rpcs.MarkUserAsDeletedResponse, error) {
	violations := validateMarkUserAsDeletedRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "MarkUserAsDeleted"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	deletedUser, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err)
	}
	if deletedUser.UserRole != util.UserRole {
		return nil, status.Errorf(codes.PermissionDenied, "admin account cant not be deleted")
	}

	arg := db.UpdateUserBasicInfoParams{
		Username: req.GetUsername(),
		Deleted:  pgtype.Bool{Bool: true, Valid: true},
	}

	user, err := server.store.UpdateUserBasicInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err)
	}

	rsp := &rpcs.MarkUserAsDeletedResponse{Deleted: user.Deleted}
	return rsp, nil
}
func validateMarkUserAsDeletedRequest(req *rpcs.MarkUserAsDeletedRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
