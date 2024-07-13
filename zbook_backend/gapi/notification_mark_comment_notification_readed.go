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

func (server *Server) MarkCommentNotificationReaded(ctx context.Context, req *rpcs.MarkCommentNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "MarkCommentNotificationReaded"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateMarkCommentNotificationReadedRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg := db.MarkCommentNotificationReadedParams{
		NotiID: req.GetNotiId(),
		UserID: authPayload.UserID,
	}
	_, err = server.store.MarkCommentNotificationReaded(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "comment notification not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "mark comment notification readed failed: %s", err)
	}
	rsp := &rpcs.SetNotiReadResponse{}
	return rsp, nil
}
func validateMarkCommentNotificationReadedRequest(req *rpcs.MarkCommentNotificationReadedRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetNotiId()); err != nil {
		violations = append(violations, fieldViolation("noti_id", err))
	}
	return violations
}
