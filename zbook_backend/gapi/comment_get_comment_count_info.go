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
)

func (server *Server) GetCommentCountInfo(ctx context.Context, req *rpcs.GetCommentCountInfoRequest) (*rpcs.GetCommentCountInfoResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetCommentCountInfo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateGetCommentCountInfoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.GetCommentDetailParams{
		UserID:    authPayload.UserID,
		CommentID: req.GetCommentId(),
	}
	comment, err := server.store.GetCommentDetail(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get comment failed: %s", err)
	}

	comment_count_info := &models.CommentCountInfo{
		CommentId:  comment.CommentID,
		LikeCount:  int32(comment.LikeCount),
		ReplyCount: int32(comment.ReplyCount),
		IsLiked:    comment.IsLiked,
		IsDisliked: comment.IsDisliked,
		IsShared:   comment.IsShared,
		IsReported: comment.IsReported,
	}

	rsp := &rpcs.GetCommentCountInfoResponse{
		CommentCountInfo: comment_count_info,
	}
	return rsp, nil
}
func validateGetCommentCountInfoRequest(req *rpcs.GetCommentCountInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateID(req.GetCommentId()); err != nil {
		violations = append(violations, fieldViolation("comment_id", err))
	}

	return violations
}
