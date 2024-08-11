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

func (server *Server) SendEmailToVerifyEmail(ctx context.Context, req *rpcs.SendEmailToVerifyEmailRequest) (*rpcs.SendEmailToVerifyEmailResponse, error) {
	violations := validateSendEmailToVerifyEmailRequest(req)
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
	if user.Verified {
		return nil, status.Errorf(codes.Internal, "email is already verified")
	}
	apiKey := "SendEmailToVerifyEmail"
	apiUserDailyLimit := 100
	err = server.checkUserLimit(user.UserID, apiKey, apiUserDailyLimit)
	if err != nil {
		return nil, err
	}

	taskPayload := &worker.PayloadVerifyEmail{
		Email: req.GetEmail(),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskVerifyEmail(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "distribute verify email task failed: %s", err)
	}
	rsp := &rpcs.SendEmailToVerifyEmailResponse{
		IsSend: true,
	}
	return rsp, nil
}

func validateSendEmailToVerifyEmailRequest(req *rpcs.SendEmailToVerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
