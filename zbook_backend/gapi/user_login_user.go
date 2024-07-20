package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *rpcs.LoginUserRequest) (*rpcs.LoginUserResponse, error) {
	violations := validateLoginuserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get user by email failed: %s", err)
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "password is wrong")
	}

	err = server.checkUserStatus(ctx, user.Blocked, user.UserRole, user.Verified)
	if err != nil {
		return nil, err
	}
	response, err := server.CreateLoginPart(ctx, user.Username, user.UserRole, user.UserID, "LoginUser")
	if err != nil {
		return nil, err
	}
	return response, nil
}
func validateLoginuserRequest(req *rpcs.LoginUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	return violations
}
