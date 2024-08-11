package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateCommentReport(ctx context.Context, req *rpcs.CreateCommentReportRequest) (*rpcs.CreateCommentReportResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "CreateCommentReport"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateCreateCommentReportRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateCommentReportParams{
		ReportContent: req.GetReportContent(),
		CommentID:     req.GetCommentId(),
		UserID:        authPayload.UserID,
	}

	err = server.store.CreateCommentReport(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation || db.ErrorCode(err) == db.ForeignKeyViolation {
			return nil, status.Errorf(codes.AlreadyExists, "comment report already exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create comment report: %s", err)
	}

	rsp := &rpcs.CreateCommentReportResponse{}
	return rsp, nil
}
func validateCreateCommentReportRequest(req *rpcs.CreateCommentReportRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetCommentId()); err != nil {
		violations = append(violations, fieldViolation("comment_id", err))
	}
	if err := val.ValidateString(req.GetReportContent(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("report_content", err))
	}
	return violations
}
