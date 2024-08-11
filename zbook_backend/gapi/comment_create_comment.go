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

func (server *Server) CreateComment(ctx context.Context, req *rpcs.CreateCommentRequest) (*rpcs.CreateCommentResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "CreateComment"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateCreateCommentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err = server.isMarkdownVisibleToCurrentUser(ctx, req.GetMarkdownId())
	if err != nil {
		return nil, err
	}
	var RootID int64 = 0
	if req.GetParentId() != 0 {
		pcomment, err := server.store.GetCommentBasicInfo(ctx, req.GetParentId())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "get comment failed: %s", err)
		}
		if pcomment.RootID.Valid {
			RootID = pcomment.RootID.Int64
		} else {
			RootID = pcomment.CommentID
		}
	}

	arg := db.CreateCommentTxParams{
		UserID:         authPayload.UserID,
		MarkdownID:     req.GetMarkdownId(),
		ParentID:       req.GetParentId(),
		RootID:         RootID,
		CommentContent: req.GetCommentContent(),
	}

	comment, err := server.store.CreateCommentTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create comment failed: %s", err)
	}

	rsp := &rpcs.CreateCommentResponse{
		Comment: &models.CommentBasicInfo{
			CommentId:      comment.Comment.CommentID,
			UserId:         comment.Comment.UserID,
			ParentId:       req.GetParentId(),
			RootId:         RootID,
			MarkdownId:     comment.Comment.MarkdownID,
			CommentContent: comment.Comment.CommentContent,
			CreatedAt:      timestamppb.New(comment.Comment.CreatedAt),
		},
	}
	return rsp, nil
}
func validateCreateCommentRequest(req *rpcs.CreateCommentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetMarkdownId()); err != nil {
		violations = append(violations, fieldViolation("markdown_id", err))
	}
	if err := val.ValidateString(req.GetCommentContent(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("comment_content", err))
	}
	return violations
}
