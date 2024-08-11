package gapi

import (
	"context"

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

func (server *Server) ListComment(ctx context.Context, req *rpcs.ListCommentRequest) (*rpcs.ListCommentResponse, error) {
	violations := validateListCommentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListComment"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	if req.GetQuery() != "" {
		arg := db.QueryCommentParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			Query:  req.GetQuery(),
		}

		reports, err := server.store.QueryComment(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query comment failed: %s", err)
		}
		rsp := &rpcs.ListCommentResponse{
			Elements: convertListCommentByID(reports),
		}
		return rsp, nil
	}
	arg := db.ListCommentParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	reports, err := server.store.ListComment(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list comment failed: %s", err)
	}

	rsp := &rpcs.ListCommentResponse{
		Elements: convertListComment(reports),
	}
	return rsp, nil
}
func validateListCommentRequest(req *rpcs.ListCommentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListComment(reports []db.ListCommentRow) []*models.ListAdminCommentInfo {
	var ret_reports []*models.ListAdminCommentInfo
	for i := 0; i < len(reports); i++ {
		ret_reports = append(ret_reports,
			&models.ListAdminCommentInfo{
				CommentId:      reports[i].CommentID,
				CommentContent: reports[i].CommentContent,
				CreatedAt:      timestamppb.New(reports[i].CreatedAt),
				Username:       reports[i].Username,
				Email:          reports[i].Email,
			},
		)
	}
	return ret_reports
}

func convertListCommentByID(reports []db.QueryCommentRow) []*models.ListAdminCommentInfo {
	var ret_reports []*models.ListAdminCommentInfo
	for i := 0; i < len(reports); i++ {
		ret_reports = append(ret_reports,
			&models.ListAdminCommentInfo{
				CommentId:      reports[i].CommentID,
				CommentContent: reports[i].CommentContent,
				CreatedAt:      timestamppb.New(reports[i].CreatedAt),
				Username:       reports[i].Username,
				Email:          reports[i].Email,
			},
		)
	}
	return ret_reports
}
