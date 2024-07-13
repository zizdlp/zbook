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

func (server *Server) ListCommentLevelOne(ctx context.Context, req *rpcs.ListCommentLevelOneRequest) (*rpcs.ListCommentLevelResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListCommentLevelOne"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateListCommentLevelOneRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	err = server.isMarkdownVisibleToCurrentUser(ctx, req.GetMarkdownId())
	if err != nil {
		return nil, err
	}
	arg := db.ListCommentLevelOneParams{
		MarkdownID: req.GetMarkdownId(),
		Limit:      req.GetPageSize(),
		Offset:     (req.GetPageId() - 1) * req.GetPageSize(),
		UserID:     authPayload.UserID,
	}

	comments, err := server.store.ListCommentLevelOne(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get comment level one failed: %s", err)
	}

	rsp := &rpcs.ListCommentLevelResponse{
		Comments: convertListCommentLevelOne(comments),
	}
	return rsp, nil
}

func validateListCommentLevelOneRequest(req *rpcs.ListCommentLevelOneRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetMarkdownId()); err != nil {
		violations = append(violations, fieldViolation("markdown_id", err))
	}
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListCommentLevelOne(comments []db.ListCommentLevelOneRow) []*models.ListCommentInfo {
	var ret_comments []*models.ListCommentInfo
	for i := 0; i < len(comments); i++ {
		ret_comments = append(ret_comments,
			&models.ListCommentInfo{
				MarkdownId:     comments[i].MarkdownID,
				ParentId:       int64(0),
				Username:       comments[i].Username,
				CommentContent: comments[i].CommentContent,
				CreatedAt:      timestamppb.New(comments[i].CreatedAt),
				LikeCount:      comments[i].LikeCount,
				ReplyCount:     comments[i].ReplyCount,
				IsLiked:        comments[i].IsLiked,
				IsDisliked:     comments[i].IsDisliked,
				IsShared:       comments[i].IsShared,
				IsReported:     comments[i].IsReported,
				CommentId:      comments[i].CommentID,
			},
		)
	}
	return ret_comments
}
