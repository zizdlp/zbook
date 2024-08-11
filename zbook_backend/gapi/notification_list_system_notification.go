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

func (server *Server) ListSystemNotification(ctx context.Context, req *rpcs.ListSystemNotificationRequest) (*rpcs.ListSystemNotificationResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListSystemNotification"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateListSystemNotificationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	post_arg := db.ListSystemNotificationParams{
		UserID: authPayload.UserID,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	notifications, err := server.store.ListSystemNotification(ctx, post_arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list system notification failed: %s", err)
	}

	rsp := &rpcs.ListSystemNotificationResponse{
		Notifications: convertListSystemNotification(notifications),
	}
	return rsp, nil

}
func validateListSystemNotificationRequest(req *rpcs.ListSystemNotificationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertListSystemNotification(notifications []db.ListSystemNotificationRow) []*models.ListSystemNotificationInfo {
	var ret_notifications []*models.ListSystemNotificationInfo
	for i := 0; i < len(notifications); i++ {
		ret_notifications = append(ret_notifications,
			&models.ListSystemNotificationInfo{
				Title:       notifications[i].Title,
				Contents:    notifications[i].Contents,
				CreatedAt:   timestamppb.New(notifications[i].CreatedAt),
				Readed:      notifications[i].Readed,
				NotiId:      notifications[i].NotiID,
				RedirectUrl: notifications[i].RedirectUrl.String,
			},
		)
	}
	return ret_notifications
}
