package gapi

import (
	"context"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"github.com/zizdlp/zbook/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SendEmailToResetPassword(ctx context.Context, req *rpcs.SendEmailToResetPasswordRequest) (*rpcs.SendEmailToResetPasswordResponse, error) {
	violations := validateSendEmailToResetPasswordRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get user by email failed: %s", err)
	}

	apiKey := "SendEmailToResetPassword"
	apiUserDailyLimit := 100
	err = server.checkUserLimit(user.UserID, apiKey, apiUserDailyLimit)
	if err != nil {
		return nil, err
	}

	taskPayload := &worker.PayloadResetPassword{
		Email: req.GetEmail(),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskResetPassword(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "distribute restpassword task failed: %s", err)
	}
	rsp := &rpcs.SendEmailToResetPasswordResponse{
		IsSend: true,
	}
	return rsp, nil
}

func validateSendEmailToResetPasswordRequest(req *rpcs.SendEmailToResetPasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
