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

func (server *Server) UpdateCommentReportStatus(ctx context.Context, req *rpcs.UpdateCommentReportStatusRequest) (*rpcs.UpdateCommentReportStatusResponse, error) {
	violations := validateUpdateCommentReportStatusRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "UpdateCommentReportStatus"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	arg := db.UpdateCommentReportStatusParams{
		ReportID:  req.GetReportId(),
		Processed: req.GetProcessed(),
	}

	err = server.store.UpdateCommentReportStatus(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "comment report not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to update comment report: %s", err)
	}

	rsp := &rpcs.UpdateCommentReportStatusResponse{}
	return rsp, nil
}
func validateUpdateCommentReportStatusRequest(req *rpcs.UpdateCommentReportStatusRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetReportId()); err != nil {
		violations = append(violations, fieldViolation("report_id", err))
	}
	return violations
}
