package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) VerifyEmail(ctx context.Context, req *rpcs.VerifyEmailRequest) (*rpcs.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.VerifyEmailTx(ctx, req.GetVerificationUrl())
	if err != nil {
		return nil, err
	}

	rsp := &rpcs.VerifyEmailResponse{
		IsVerified: user.Verified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *rpcs.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetVerificationUrl(), 16, 64); err != nil {
		violations = append(violations, fieldViolation("verification_url", err))
	}
	return violations
}
