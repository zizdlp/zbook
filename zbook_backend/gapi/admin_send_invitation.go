package gapi

import (
	"context"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"github.com/zizdlp/zbook/worker"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) SendInvitation(ctx context.Context, req *rpcs.SendInvitationRequest) (*rpcs.SendInvitationResponse, error) {
	violations := validateSendInvitationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	apiUserDailyLimit := 10000
	apiKey := "SendInvitation"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	user, err := server.store.GetUserByEmail(ctx, req.GetEmail())
	if err == nil {
		return nil, fmt.Errorf("use already exist for this email: %s", user.Email)
	}
	taskPayload := &worker.PayloadInviteUser{
		Email: req.GetEmail(),
	}
	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(10 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}
	err = server.taskDistributor.DistributeTaskInviteUser(ctx, taskPayload, opts...)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "distribute invite user task failed: %s", err)
	}
	rsp := &rpcs.SendInvitationResponse{
		IsSend: true,
	}
	return rsp, nil

}

func validateSendInvitationRequest(req *rpcs.SendInvitationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	return violations
}
