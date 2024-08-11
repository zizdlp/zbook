package gapi

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/storage"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"github.com/zizdlp/zbook/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *rpcs.CreateUserRequest) (*rpcs.CreateUserResponse, error) {

	violations := validateCreateUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	config, err := server.store.GetConfiguration(ctx, "allow_registration")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get configuration: %v", err)
	}
	if !config.ConfigValue {
		config_invitation, err := server.store.GetConfiguration(ctx, "allow_invitation")
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get configuration: %v", err)
		}
		if !config_invitation.ConfigValue {
			return nil, status.Errorf(codes.PermissionDenied, "registration is currently not enabled")
		} else {
			arg_invitation := db.GetInvitationParams{
				Email:         req.GetEmail(),
				InvitationUrl: req.GetInvitationUrl(),
			}
			invitation, err := server.store.GetInvitation(ctx, arg_invitation)
			if err != nil {
				return nil, status.Errorf(codes.PermissionDenied, "registration is currently not enabled")
			}
			if invitation.IsUsed || invitation.ExpiredAt.Before(time.Now()) {
				return nil, status.Errorf(codes.PermissionDenied, "invitation is invalid or has expired")
			}
		}
	}

	hashedPassword, err := util.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %v", err)
	}
	path := fmt.Sprintf("icons/v%d.png", util.RandomInt32(0, 9))
	avatar, err := util.ReadImageBytes(path)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to read image bytes: %v", err)
	}
	err = storage.UploadFileToStorage(server.minioClient, ctx, req.GetUsername(), "avatar", avatar)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to upload avatar: %v", err)
	}
	arg := db.CreateUserTxParams{
		CreateUserParams: db.CreateUserParams{
			Username:       req.GetUsername(),
			HashedPassword: hashedPassword,
			Email:          req.GetEmail(),
		},
		AfterCreate: func(user db.User) error {
			if server.config.REQUIRE_EMAIL_VERIFY {
				taskPayload := &worker.PayloadVerifyEmail{
					Email: user.Email,
				}
				opts := []asynq.Option{
					asynq.MaxRetry(10),
					asynq.ProcessIn(10 * time.Second),
					asynq.Queue(worker.QueueCritical),
				}
				return server.taskDistributor.DistributeTaskVerifyEmail(ctx, taskPayload, opts...)
			} else {
				return nil
			}
		},
	}

	result, err := server.store.CreateUserTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username or email already exists: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create user: %v", err)
	}

	CreateSystemNotificationParams := db.CreateSystemNotificationParams{
		UserID:      result.User.UserID,
		Title:       "Welcome to ZBook",
		Contents:    "Welcome to ZBook, please click to view the introduction guide.",
		RedirectUrl: pgtype.Text{String: "https://github.com/zizdlp/zbook-user-guide", Valid: true},
	}
	err = server.store.CreateSystemNotificationTx(ctx, db.CreateSystemNotificationTxParams{CreateSystemNotificationParams: CreateSystemNotificationParams})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create system notification: %v", err)
	}

	rsp := &rpcs.CreateUserResponse{}
	return rsp, nil
}

func validateCreateUserRequest(req *rpcs.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
