// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"net/netip"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CheckOAuthStatus(ctx context.Context, userID int64) (CheckOAuthStatusRow, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreateCommentNotification(ctx context.Context, arg CreateCommentNotificationParams) (CommentNotification, error)
	CreateCommentRelation(ctx context.Context, arg CreateCommentRelationParams) error
	CreateCommentReport(ctx context.Context, arg CreateCommentReportParams) error
	CreateFollow(ctx context.Context, arg CreateFollowParams) (Follow, error)
	CreateFollowerNotification(ctx context.Context, arg CreateFollowerNotificationParams) (FollowerNotification, error)
	CreateInvitation(ctx context.Context, arg CreateInvitationParams) (Invitation, error)
	CreateMarkdown(ctx context.Context, arg CreateMarkdownParams) (Markdown, error)
	CreateMarkdownMulti(ctx context.Context, arg CreateMarkdownMultiParams) error
	CreateOAuth(ctx context.Context, arg CreateOAuthParams) (Oauth, error)
	CreateRepo(ctx context.Context, arg CreateRepoParams) (Repo, error)
	CreateRepoNotification(ctx context.Context, arg CreateRepoNotificationParams) (RepoNotification, error)
	CreateRepoRelation(ctx context.Context, arg CreateRepoRelationParams) error
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateSystemNotification(ctx context.Context, arg CreateSystemNotificationParams) (SystemNotification, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	CreateVerification(ctx context.Context, arg CreateVerificationParams) (Verification, error)
	DeleteComment(ctx context.Context, commentID int64) error
	DeleteCommentRelation(ctx context.Context, arg DeleteCommentRelationParams) error
	DeleteFollow(ctx context.Context, arg DeleteFollowParams) (int64, error)
	DeleteFollowerNotification(ctx context.Context, arg DeleteFollowerNotificationParams) (FollowerNotification, error)
	DeleteMarkdownMulti(ctx context.Context, arg DeleteMarkdownMultiParams) error
	DeleteOAuth(ctx context.Context, arg DeleteOAuthParams) (Oauth, error)
	DeleteRepo(ctx context.Context, repoID int64) error
	DeleteRepoRelation(ctx context.Context, arg DeleteRepoRelationParams) error
	DeleteUser(ctx context.Context, username string) error
	GetCommentBasicInfo(ctx context.Context, commentID int64) (GetCommentBasicInfoRow, error)
	GetCommentDetail(ctx context.Context, arg GetCommentDetailParams) (GetCommentDetailRow, error)
	GetCommentRepoInfo(ctx context.Context, commentID int64) (Repo, error)
	GetConfiguration(ctx context.Context, configName string) (Configuration, error)
	GetDailyActiveUserCount(ctx context.Context, arg GetDailyActiveUserCountParams) ([]GetDailyActiveUserCountRow, error)
	GetDailyCreateUserCount(ctx context.Context, arg GetDailyCreateUserCountParams) ([]GetDailyCreateUserCountRow, error)
	GetGeoInfo(ctx context.Context, dollar_1 netip.Addr) (Geoip, error)
	GetGeoInfoBatch(ctx context.Context, dollar_1 []netip.Addr) ([]Geoip, error)
	GetInvitation(ctx context.Context, arg GetInvitationParams) (Invitation, error)
	GetListCommentCount(ctx context.Context) (int64, error)
	GetListCommentLevelOneCount(ctx context.Context, markdownID int64) (int64, error)
	GetListCommentLevelTwoCount(ctx context.Context, rootID pgtype.Int8) (int64, error)
	GetListCommentNotificationUnreadedCount(ctx context.Context, userID int64) (int64, error)
	GetListCommentReportCount(ctx context.Context) (int64, error)
	GetListFollowerCount(ctx context.Context, arg GetListFollowerCountParams) (int64, error)
	GetListFollowerNotificationUnreadedCount(ctx context.Context, userID int64) (int64, error)
	GetListFollowingCount(ctx context.Context, arg GetListFollowingCountParams) (int64, error)
	GetListRepoCount(ctx context.Context, arg GetListRepoCountParams) (int64, error)
	GetListRepoNotificationUnreadedCount(ctx context.Context, userID int64) (int64, error)
	GetListSelectedUserByRepoCount(ctx context.Context, arg GetListSelectedUserByRepoCountParams) (int64, error)
	GetListSessionCount(ctx context.Context) (int64, error)
	GetListSystemNotificationUnReadedCount(ctx context.Context, userID int64) (int64, error)
	GetListUserCount(ctx context.Context, role string) (int64, error)
	GetListUserLikeRepoCount(ctx context.Context, arg GetListUserLikeRepoCountParams) (int64, error)
	GetListUserOwnRepoCount(ctx context.Context, arg GetListUserOwnRepoCountParams) (int64, error)
	GetMarkdownByID(ctx context.Context, markdownID int64) (Markdown, error)
	GetMarkdownContent(ctx context.Context, arg GetMarkdownContentParams) (Markdown, error)
	GetMarkdownRepoID(ctx context.Context, markdownID int64) (int64, error)
	GetOAuthUser(ctx context.Context, arg GetOAuthUserParams) (GetOAuthUserRow, error)
	GetQueryCommentCount(ctx context.Context, query string) (int64, error)
	GetQueryCommentReportCount(ctx context.Context, query string) (int64, error)
	GetQueryFollowerCount(ctx context.Context, arg GetQueryFollowerCountParams) (int64, error)
	GetQueryFollowingCount(ctx context.Context, arg GetQueryFollowingCountParams) (int64, error)
	GetQueryRepoCount(ctx context.Context, arg GetQueryRepoCountParams) (int64, error)
	GetQuerySelectedUserByRepoCount(ctx context.Context, arg GetQuerySelectedUserByRepoCountParams) (int64, error)
	GetQuerySessionCount(ctx context.Context, query string) (int64, error)
	GetQueryUserCount(ctx context.Context, arg GetQueryUserCountParams) (int64, error)
	GetQueryUserLikeRepoCount(ctx context.Context, arg GetQueryUserLikeRepoCountParams) (int64, error)
	GetQueryUserOwnRepoCount(ctx context.Context, arg GetQueryUserOwnRepoCountParams) (int64, error)
	GetRepo(ctx context.Context, repoID int64) (Repo, error)
	GetRepoBasicInfo(ctx context.Context, arg GetRepoBasicInfoParams) (GetRepoBasicInfoRow, error)
	GetRepoByRepoName(ctx context.Context, arg GetRepoByRepoNameParams) (GetRepoByRepoNameRow, error)
	GetRepoConfig(ctx context.Context, arg GetRepoConfigParams) (GetRepoConfigRow, error)
	GetRepoHome(ctx context.Context, arg GetRepoHomeParams) (string, error)
	GetRepoID(ctx context.Context, arg GetRepoIDParams) (int64, error)
	GetRepoPermission(ctx context.Context, repoID int64) (GetRepoPermissionRow, error)
	GetRepoRelation(ctx context.Context, arg GetRepoRelationParams) (RepoRelation, error)
	GetSession(ctx context.Context, sessionID uuid.UUID) (Session, error)
	GetUnReadCount(ctx context.Context, username string) (int32, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUserInfo(ctx context.Context, arg GetUserInfoParams) (GetUserInfoRow, error)
	GetVerification(ctx context.Context, verificationUrl string) (GetVerificationRow, error)
	IsFollowing(ctx context.Context, arg IsFollowingParams) (bool, error)
	ListComment(ctx context.Context, arg ListCommentParams) ([]ListCommentRow, error)
	ListCommentLevelOne(ctx context.Context, arg ListCommentLevelOneParams) ([]ListCommentLevelOneRow, error)
	ListCommentLevelTwo(ctx context.Context, arg ListCommentLevelTwoParams) ([]ListCommentLevelTwoRow, error)
	ListCommentNotification(ctx context.Context, arg ListCommentNotificationParams) ([]ListCommentNotificationRow, error)
	ListCommentReport(ctx context.Context, arg ListCommentReportParams) ([]ListCommentReportRow, error)
	ListFollower(ctx context.Context, arg ListFollowerParams) ([]ListFollowerRow, error)
	ListFollowerNotification(ctx context.Context, arg ListFollowerNotificationParams) ([]ListFollowerNotificationRow, error)
	ListFollowing(ctx context.Context, arg ListFollowingParams) ([]ListFollowingRow, error)
	ListRepo(ctx context.Context, arg ListRepoParams) ([]ListRepoRow, error)
	ListRepoNotification(ctx context.Context, arg ListRepoNotificationParams) ([]ListRepoNotificationRow, error)
	ListSelectedUserByRepo(ctx context.Context, arg ListSelectedUserByRepoParams) ([]User, error)
	ListSession(ctx context.Context, arg ListSessionParams) ([]ListSessionRow, error)
	ListSystemNotification(ctx context.Context, arg ListSystemNotificationParams) ([]ListSystemNotificationRow, error)
	ListUser(ctx context.Context, arg ListUserParams) ([]User, error)
	ListUserLikeRepo(ctx context.Context, arg ListUserLikeRepoParams) ([]ListUserLikeRepoRow, error)
	ListUserOwnRepo(ctx context.Context, arg ListUserOwnRepoParams) ([]ListUserOwnRepoRow, error)
	MarkCommentNotificationReaded(ctx context.Context, arg MarkCommentNotificationReadedParams) (CommentNotification, error)
	MarkFollowerNotificationReaded(ctx context.Context, arg MarkFollowerNotificationReadedParams) (FollowerNotification, error)
	MarkInvitationAsUsed(ctx context.Context, arg MarkInvitationAsUsedParams) (Invitation, error)
	MarkRepoNotificationReaded(ctx context.Context, arg MarkRepoNotificationReadedParams) (RepoNotification, error)
	MarkSystemNotificationReaded(ctx context.Context, arg MarkSystemNotificationReadedParams) (SystemNotification, error)
	MarkVerificationAsUsed(ctx context.Context, verificationUrl string) (Verification, error)
	QueryComment(ctx context.Context, arg QueryCommentParams) ([]QueryCommentRow, error)
	QueryCommentReport(ctx context.Context, arg QueryCommentReportParams) ([]QueryCommentReportRow, error)
	QueryFollower(ctx context.Context, arg QueryFollowerParams) ([]QueryFollowerRow, error)
	QueryFollowing(ctx context.Context, arg QueryFollowingParams) ([]QueryFollowingRow, error)
	QueryMarkdown(ctx context.Context, arg QueryMarkdownParams) ([]QueryMarkdownRow, error)
	QueryRepo(ctx context.Context, arg QueryRepoParams) ([]QueryRepoRow, error)
	QueryRepoMarkdown(ctx context.Context, arg QueryRepoMarkdownParams) ([]QueryRepoMarkdownRow, error)
	QuerySelectedUserByRepo(ctx context.Context, arg QuerySelectedUserByRepoParams) ([]QuerySelectedUserByRepoRow, error)
	QuerySession(ctx context.Context, arg QuerySessionParams) ([]QuerySessionRow, error)
	QueryUser(ctx context.Context, arg QueryUserParams) ([]QueryUserRow, error)
	QueryUserByRepo(ctx context.Context, arg QueryUserByRepoParams) ([]QueryUserByRepoRow, error)
	QueryUserLikeRepo(ctx context.Context, arg QueryUserLikeRepoParams) ([]QueryUserLikeRepoRow, error)
	QueryUserMarkdown(ctx context.Context, arg QueryUserMarkdownParams) ([]QueryUserMarkdownRow, error)
	QueryUserOwnRepo(ctx context.Context, arg QueryUserOwnRepoParams) ([]QueryUserOwnRepoRow, error)
	ResetUnreadCount(ctx context.Context, username string) error
	UpdateCommentReportStatus(ctx context.Context, arg UpdateCommentReportStatusParams) error
	UpdateConfiguration(ctx context.Context, arg UpdateConfigurationParams) error
	UpdateMarkdownMulti(ctx context.Context, arg UpdateMarkdownMultiParams) error
	UpdateRepoConfig(ctx context.Context, arg UpdateRepoConfigParams) error
	UpdateRepoInfo(ctx context.Context, arg UpdateRepoInfoParams) (Repo, error)
	UpdateUnreadCount(ctx context.Context, userID int64) error
	UpdateUserBasicInfo(ctx context.Context, arg UpdateUserBasicInfoParams) (User, error)
}

var _ Querier = (*Queries)(nil)
