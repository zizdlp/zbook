package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
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

func (server *Server) ListCommentLevelTwo(ctx context.Context, req *rpcs.ListCommentLevelTwoRequest) (*rpcs.ListCommentLevelResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListCommentLevelTwo"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateListCommentLevelTwoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	err = server.isCommentVisibleToCurrentUser(ctx, req.GetRootId())
	if err != nil {
		return nil, err
	}

	arg := db.ListCommentLevelTwoParams{
		RootID: pgtype.Int8{Int64: req.GetRootId(), Valid: true},
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
		UserID: authUser.UserID,
	}

	comments, err := server.store.ListCommentLevelTwo(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list comment level two failed: %s", err)
	}

	rsp := &rpcs.ListCommentLevelResponse{
		Comments: convertListCommentLevelTwo(comments),
	}
	return rsp, nil
}

func validateListCommentLevelTwoRequest(req *rpcs.ListCommentLevelTwoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetRootId()); err != nil {
		violations = append(violations, fieldViolation("root_id", err))
	}
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListCommentLevelTwo(comments []db.ListCommentLevelTwoRow) []*models.ListCommentInfo {
	var ret_comments []*models.ListCommentInfo
	for i := 0; i < len(comments); i++ {
		ret_comments = append(ret_comments,
			&models.ListCommentInfo{
				MarkdownId:     comments[i].MarkdownID,
				ParentId:       comments[i].ParentID.Int64,
				Username:       comments[i].Username,
				Pusername:      comments[i].Pusername.String,
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
