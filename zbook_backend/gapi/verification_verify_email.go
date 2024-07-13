package gapi

import (
	"context"

	"github.com/google/uuid"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) VerifyEmail(ctx context.Context, req *rpcs.VerifyEmailRequest) (*rpcs.VerifyEmailResponse, error) {
	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	verificationIDString := req.GetVerificationId()
	verificationID, err := util.StringToUUID(verificationIDString)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "convert string to uuid failed: %v", err)
	}

	user, err := server.store.VerifyEmailTx(ctx, verificationID)
	if err != nil {
		return nil, err
	}

	rsp := &rpcs.VerifyEmailResponse{
		IsVerified: user.Verified,
	}
	return rsp, nil
}

func validateVerifyEmailRequest(req *rpcs.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	_, err := uuid.Parse(req.GetVerificationId())
	if err != nil {
		violations = append(violations, fieldViolation("verification_id", err))
	}
	return violations
}
