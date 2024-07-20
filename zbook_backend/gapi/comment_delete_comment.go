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

func (server *Server) DeleteComment(ctx context.Context, req *rpcs.DeleteCommentRequest) (*rpcs.DeleteCommentResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "DeleteComment"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateDeleteCommentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	comment, err := server.store.GetCommentBasicInfo(ctx, req.GetCommentId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "comment not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get comment failed: %s", err)
	}

	if comment.UserID != authPayload.UserID && authPayload.UserRole != util.AdminRole {
		return nil, status.Errorf(codes.PermissionDenied, "current account can not delete this comment")
	}

	err = server.store.DeleteComment(ctx, req.GetCommentId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete comment failed: %s", err)
	}

	rsp := &rpcs.DeleteCommentResponse{}
	return rsp, nil
}
func validateDeleteCommentRequest(req *rpcs.DeleteCommentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetCommentId()); err != nil {
		violations = append(violations, fieldViolation("comment_id", err))
	}
	return violations
}
