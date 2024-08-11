package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) ResetPassword(ctx context.Context, req *rpcs.ResetPasswordRequest) (*rpcs.ResetPasswordResponse, error) {

	violations := validateResetPasswordRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg := db.ResetPasswordTxParams{
		VerificationUrl: req.GetVerificationUrl(),
		Email:           req.GetEmail(),
		Password:        req.GetPassword(),
	}
	err := server.store.ResetPasswordTx(ctx, arg)
	if err != nil {
		return nil, err
	}
	rsp := &rpcs.ResetPasswordResponse{
		IsReset: true,
	}
	return rsp, nil
}
func validateResetPasswordRequest(req *rpcs.ResetPasswordRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetVerificationUrl(), 16, 64); err != nil {
		violations = append(violations, fieldViolation("verification_url", err))
	}
	if err := val.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("email", err))
	}
	if err := val.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	return violations
}
