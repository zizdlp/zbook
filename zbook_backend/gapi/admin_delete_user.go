package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteUser(ctx context.Context, req *rpcs.DeleteUserRequest) (*rpcs.DeleteUserResponse, error) {
	violations := validateDeleteUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "DeleteUser"
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

	arg := db.DeleteUserTxParams{
		UserID:   deletedUser.UserID,
		Username: deletedUser.Username,
		AfterDelte: func(userID int64, username string) error {
			err := storage.DeleteAvatarByUsername(server.minioClient, context.Background(), username)
			if err != nil {
				return err
			}
			return storage.DeleteFilesByUserID(server.minioClient, context.Background(), userID, "git-files")
		},
	}

	err = server.store.DeleteUserTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete user: %v", err)
	}

	err = server.store.DeleteUser(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to delete user: %s", err)
	}

	rsp := &rpcs.DeleteUserResponse{}
	return rsp, nil
}
func validateDeleteUserRequest(req *rpcs.DeleteUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
