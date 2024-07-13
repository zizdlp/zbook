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

func (server *Server) MarkFollowerNotificationReaded(ctx context.Context, req *rpcs.MarkFollowerNotificationReadedRequest) (*rpcs.SetNotiReadResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "MarkFollowerNotificationReaded"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateMarkFollowerNotificationReadedRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.MarkFollowerNotificationReadedParams{
		NotiID: req.GetNotiId(),
		UserID: authPayload.UserID,
	}
	_, err = server.store.MarkFollowerNotificationReaded(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "follower notification not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "mark follower notification readed failed: %s", err)
	}

	rsp := &rpcs.SetNotiReadResponse{}
	return rsp, nil
}
func validateMarkFollowerNotificationReadedRequest(req *rpcs.MarkFollowerNotificationReadedRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetNotiId()); err != nil {
		violations = append(violations, fieldViolation("noti_id", err))
	}
	return violations
}
