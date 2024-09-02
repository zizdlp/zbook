// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"net/netip"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Comment struct {
	CommentID      int64       `json:"comment_id"`
	RepoID         int64       `json:"repo_id"`
	MarkdownID     int64       `json:"markdown_id"`
	ParentID       pgtype.Int8 `json:"parent_id"`
	RootID         pgtype.Int8 `json:"root_id"`
	UserID         int64       `json:"user_id"`
	Blocked        bool        `json:"blocked"`
	CommentContent string      `json:"comment_content"`
	CreatedAt      time.Time   `json:"created_at"`
	FtsCommentZh   string      `json:"fts_comment_zh"`
	FtsCommentEn   string      `json:"fts_comment_en"`
}

type CommentNotification struct {
	NotiID    int64     `json:"noti_id"`
	UserID    int64     `json:"user_id"`
	CommentID int64     `json:"comment_id"`
	Readed    bool      `json:"readed"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentRelation struct {
	RelationID   int64     `json:"relation_id"`
	RelationType string    `json:"relation_type"`
	UserID       int64     `json:"user_id"`
	CommentID    int64     `json:"comment_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type CommentReport struct {
	ReportID      int64     `json:"report_id"`
	UserID        int64     `json:"user_id"`
	CommentID     int64     `json:"comment_id"`
	ReportContent string    `json:"report_content"`
	Processed     bool      `json:"processed"`
	CreatedAt     time.Time `json:"created_at"`
	FtsReportZh   string    `json:"fts_report_zh"`
	FtsReportEn   string    `json:"fts_report_en"`
}

type Configuration struct {
	ConfigName  string    `json:"config_name"`
	ConfigValue bool      `json:"config_value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Follow struct {
	FollowID    int64     `json:"follow_id"`
	FollowerID  int64     `json:"follower_id"`
	FollowingID int64     `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type FollowerNotification struct {
	NotiID     int64     `json:"noti_id"`
	UserID     int64     `json:"user_id"`
	FollowerID int64     `json:"follower_id"`
	Readed     bool      `json:"readed"`
	CreatedAt  time.Time `json:"created_at"`
}

type Geoip struct {
	GeoipID      int64         `json:"geoip_id"`
	IpRangeCidr  *netip.Prefix `json:"ip_range_cidr"`
	CityNameEn   pgtype.Text   `json:"city_name_en"`
	CityNameZhCn pgtype.Text   `json:"city_name_zh_cn"`
	Latitude     pgtype.Float8 `json:"latitude"`
	Longitude    pgtype.Float8 `json:"longitude"`
}

type Invitation struct {
	InvitationID  int64     `json:"invitation_id"`
	Email         string    `json:"email"`
	InvitationUrl string    `json:"invitation_url"`
	IsUsed        bool      `json:"is_used"`
	CreatedAt     time.Time `json:"created_at"`
	ExpiredAt     time.Time `json:"expired_at"`
}

type Markdown struct {
	MarkdownID   int64     `json:"markdown_id"`
	RelativePath string    `json:"relative_path"`
	UserID       int64     `json:"user_id"`
	RepoID       int64     `json:"repo_id"`
	MainContent  string    `json:"main_content"`
	TableContent string    `json:"table_content"`
	UpdatedAt    time.Time `json:"updated_at"`
	CreatedAt    time.Time `json:"created_at"`
	FtsZh        string    `json:"fts_zh"`
	FtsEn        string    `json:"fts_en"`
}

type Oauth struct {
	OauthID   int64     `json:"oauth_id"`
	UserID    int64     `json:"user_id"`
	OauthType string    `json:"oauth_type"`
	AppID     string    `json:"app_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Repo struct {
	RepoID          int64       `json:"repo_id"`
	UserID          int64       `json:"user_id"`
	GitProtocol     string      `json:"git_protocol"`
	GitHost         string      `json:"git_host"`
	GitUsername     string      `json:"git_username"`
	GitRepo         string      `json:"git_repo"`
	GitAccessToken  pgtype.Text `json:"git_access_token"`
	RepoName        string      `json:"repo_name"`
	RepoDescription string      `json:"repo_description"`
	SyncToken       pgtype.Text `json:"sync_token"`
	VisibilityLevel string      `json:"visibility_level"`
	CommitID        string      `json:"commit_id"`
	Config          string      `json:"config"`
	Home            string      `json:"home"`
	ThemeSidebar    string      `json:"theme_sidebar"`
	ThemeColor      string      `json:"theme_color"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	FtsRepoEn       string      `json:"fts_repo_en"`
	FtsRepoZh       string      `json:"fts_repo_zh"`
}

type RepoNotification struct {
	NotiID    int64     `json:"noti_id"`
	UserID    int64     `json:"user_id"`
	RepoID    int64     `json:"repo_id"`
	Readed    bool      `json:"readed"`
	CreatedAt time.Time `json:"created_at"`
}

type RepoRelation struct {
	RelationID   int64     `json:"relation_id"`
	RelationType string    `json:"relation_type"`
	UserID       int64     `json:"user_id"`
	RepoID       int64     `json:"repo_id"`
	CreatedAt    time.Time `json:"created_at"`
}

type Session struct {
	SessionID    uuid.UUID `json:"session_id"`
	UserID       int64     `json:"user_id"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type SystemNotification struct {
	NotiID      int64       `json:"noti_id"`
	UserID      int64       `json:"user_id"`
	Title       string      `json:"title"`
	Contents    string      `json:"contents"`
	RedirectUrl pgtype.Text `json:"redirect_url"`
	Readed      bool        `json:"readed"`
	CreatedAt   time.Time   `json:"created_at"`
}

type User struct {
	UserID               int64     `json:"user_id"`
	Username             string    `json:"username"`
	Email                string    `json:"email"`
	HashedPassword       string    `json:"hashed_password"`
	Blocked              bool      `json:"blocked"`
	Verified             bool      `json:"verified"`
	Motto                string    `json:"motto"`
	UserRole             string    `json:"user_role"`
	Onboarding           bool      `json:"onboarding"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
	UnreadCount          int32     `json:"unread_count"`
	UnreadCountUpdatedAt time.Time `json:"unread_count_updated_at"`
	FtsUsername          string    `json:"fts_username"`
}

type Verification struct {
	VerificationID   int64     `json:"verification_id"`
	VerificationUrl  string    `json:"verification_url"`
	VerificationType string    `json:"verification_type"`
	UserID           int64     `json:"user_id"`
	IsUsed           bool      `json:"is_used"`
	CreatedAt        time.Time `json:"created_at"`
	ExpiredAt        time.Time `json:"expired_at"`
}
