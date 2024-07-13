package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateSystemNotification(ctx context.Context, req *rpcs.CreateSystemNotificationRequest) (*rpcs.CreateSystemNotificationResponse, error) {
	violations := validateCreateSystemNotificationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	apiUserDailyLimit := 10000
	apiKey := "CreateSystemNotification"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err)
	}

	CreateSystemNotificationParams := db.CreateSystemNotificationParams{
		UserID:      user.UserID,
		Title:       req.GetTitle(),
		Contents:    req.GetContents(),
		RedirectUrl: pgtype.Text{String: req.GetRedirectUrl(), Valid: req.GetRedirectUrl() != ""},
	}
	err = server.store.CreateSystemNotificationTx(ctx, db.CreateSystemNotificationTxParams{CreateSystemNotificationParams: CreateSystemNotificationParams})
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation || db.ErrorCode(err) == db.ForeignKeyViolation {
			return nil, status.Errorf(codes.AlreadyExists, "system notification already exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "create system notification failed: %s", err)
	}
	rsp := &rpcs.CreateSystemNotificationResponse{}
	return rsp, nil
}
func validateCreateSystemNotificationRequest(req *rpcs.CreateSystemNotificationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := val.ValidateString(req.GetContents(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("content", err))
	}

	if err := val.ValidateString(req.GetTitle(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("title", err))
	}

	return violations
}
