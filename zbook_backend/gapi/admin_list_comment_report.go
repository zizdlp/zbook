package gapi

import (
	"context"
	"fmt"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) ListCommentReport(ctx context.Context, req *rpcs.ListCommentReportRequest) (*rpcs.ListCommentReportResponse, error) {
	violations := validateListCommentReportRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListCommentReport"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	if req.GetQuery() != "" {
		arg := db.QueryCommentReportParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			Query:  req.GetQuery(),
		}
		reports, err := server.store.QueryCommentReport(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query comment report failed: %s", err)
		}
		rsp := &rpcs.ListCommentReportResponse{
			Elements: convertQueryCommentReport(reports),
		}
		fmt.Println("---:", req.GetQuery(), reports)
		return rsp, nil
	}
	arg := db.ListCommentReportParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	reports, err := server.store.ListCommentReport(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list comment resport failed: %s", err)
	}

	rsp := &rpcs.ListCommentReportResponse{
		Elements: convertListCommentReport(reports),
	}
	return rsp, nil
}
func validateListCommentReportRequest(req *rpcs.ListCommentReportRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListCommentReport(reports []db.ListCommentReportRow) []*models.ListCommentReportInfo {
	var ret_reports []*models.ListCommentReportInfo
	for i := 0; i < len(reports); i++ {
		ret_reports = append(ret_reports,
			&models.ListCommentReportInfo{
				ReportId:       reports[i].ReportID,
				CommentId:      reports[i].CommentID,
				RepoName:       reports[i].RepoName,
				RepoUsername:   reports[i].RepoUsername,
				RelativePath:   reports[i].RelativePath,
				ReportContent:  reports[i].ReportContent,
				CommentContent: reports[i].CommentContent,
				Processed:      reports[i].Processed,
				CreatedAt:      timestamppb.New(reports[i].CreatedAt),
				Username:       reports[i].Username,
			},
		)
	}
	return ret_reports
}

func convertQueryCommentReport(reports []db.QueryCommentReportRow) []*models.ListCommentReportInfo {
	var ret_reports []*models.ListCommentReportInfo
	for i := 0; i < len(reports); i++ {
		ret_reports = append(ret_reports,
			&models.ListCommentReportInfo{
				ReportId:       reports[i].ReportID,
				CommentId:      reports[i].CommentID,
				RepoName:       reports[i].RepoName,
				RepoUsername:   reports[i].RepoUsername,
				RelativePath:   reports[i].RelativePath,
				ReportContent:  reports[i].ReportContent,
				CommentContent: reports[i].CommentContent,
				Processed:      reports[i].Processed,
				CreatedAt:      timestamppb.New(reports[i].CreatedAt),
				Username:       reports[i].Username,
			},
		)
	}
	return ret_reports
}
