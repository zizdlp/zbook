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

func (server *Server) ListRepoNotification(ctx context.Context, req *rpcs.ListRepoNotificationRequest) (*rpcs.ListRepoNotificationResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListRepoNotification"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateListRepoNotificationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.ListRepoNotificationParams{
		UserID: authPayload.UserID,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	notifications, err := server.store.ListRepoNotification(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list repo notification failed: %s", err)
	}

	rsp := &rpcs.ListRepoNotificationResponse{
		Notifications: convertListRepoNotification(notifications),
	}
	return rsp, nil
}
func validateListRepoNotificationRequest(req *rpcs.ListRepoNotificationRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertListRepoNotification(notifications []db.ListRepoNotificationRow) []*models.ListRepoNotificationInfo {
	var ret_notifications []*models.ListRepoNotificationInfo
	for i := 0; i < len(notifications); i++ {

		ret_notifications = append(ret_notifications,
			&models.ListRepoNotificationInfo{
				Username:  notifications[i].Username, // 需要commenter的username
				Email:     notifications[i].Email,
				Readed:    notifications[i].Readed,
				NotiId:    notifications[i].NotiID,
				CreatedAt: timestamppb.New(notifications[i].CreatedAt), // 需要评论created_at
				RepoId:    notifications[i].RepoID,                     // 需要repid+ href 用于跳转
				RepoName:  notifications[i].RepoName,
			},
		)
	}
	return ret_notifications
}
