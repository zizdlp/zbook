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

func (server *Server) ListCommentNotification(ctx context.Context, req *rpcs.ListCommentNotificationRequest) (*rpcs.ListCommentNotificationResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListCommentNotification"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateListCommentNotificationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	post_arg := db.ListCommentNotificationParams{
		UserID: authPayload.UserID,
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	notifications, err := server.store.ListCommentNotification(ctx, post_arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list comment notification failed: %s", err)
	}

	rsp := &rpcs.ListCommentNotificationResponse{
		Notifications: convertListCommentNotification(notifications),
	}
	return rsp, nil

}
func validateListCommentNotificationRequest(req *rpcs.ListCommentNotificationRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertListCommentNotification(notifications []db.ListCommentNotificationRow) []*models.ListCommentNotificationInfo {
	var ret_notifications []*models.ListCommentNotificationInfo
	for i := 0; i < len(notifications); i++ {

		ret_notifications = append(ret_notifications,
			&models.ListCommentNotificationInfo{
				Username:       notifications[i].Username, // 需要commenter的username
				Email:          notifications[i].Email,
				Readed:         notifications[i].Readed,
				NotiId:         notifications[i].NotiID,
				CreatedAt:      timestamppb.New(notifications[i].CreatedAt), // 需要评论created_at
				CommentContent: notifications[i].CommentContent,             // 需要评论内容
				RepoId:         notifications[i].RepoID,                     // 需要repid+ href 用于跳转
				RelativePath:   notifications[i].RelativePath,               // 需要repid+ href 用于跳转
				RepoName:       notifications[i].RepoName,
				RepoUsername:   notifications[i].RepoUsername,
			},
		)
	}
	return ret_notifications
}
