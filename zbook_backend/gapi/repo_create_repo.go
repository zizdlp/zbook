package gapi

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	storage "github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateRepo(ctx context.Context, req *rpcs.CreateRepoRequest) (*rpcs.CreateRepoResponse, error) {
	apiUserDailyLimit := 100
	apiKey := "CreateRepo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateCreateRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	GitProtocol, GitHost, GitUsername, GitRepo, err := util.ParseGitURL(req.GetGitAddr())
	if err != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.CreateRepoTxParams{
		CreateRepoParams: db.CreateRepoParams{
			UserID:          authPayload.UserID,
			GitProtocol:     GitProtocol,
			GitHost:         GitHost,
			GitUsername:     GitUsername,
			GitRepo:         GitRepo,
			GitAccessToken:  pgtype.Text{String: req.GetGitAccessToken(), Valid: req.GetGitAccessToken() != ""},
			RepoName:        req.GetRepoName(),
			RepoDescription: req.GetRepoDescription(),
			SyncToken:       pgtype.Text{String: req.GetSyncToken(), Valid: req.GetSyncToken() != ""},
			VisibilityLevel: req.GetVisibilityLevel(),
			ThemeSidebar:    req.GetThemeSidebar(),
			ThemeColor:      req.GetThemeColor(),
			Branch:          req.Branch,
		},
		Username: authPayload.Username,
		AfterCreate: func(cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error {
			return storage.ConvertFile2Storage(server.minioClient, cloneDir, repoID, userID, addedFiles, modifiedFiles, deletedFiles)
		},
	}

	_, err = server.store.CreateRepoTx(ctx, arg)
	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation || db.ErrorCode(err) == db.ForeignKeyViolation {
			return nil, status.Errorf(codes.AlreadyExists, "repo already exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "create repo failed: %s", err)
	}

	rsp := &rpcs.CreateRepoResponse{}
	return rsp, nil
}
func validateCreateRepoRequest(req *rpcs.CreateRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateRepoName(req.GetRepoName()); err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	if err := val.ValidateString(req.GetRepoDescription(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("repo_description", err))
	}
	_, _, _, _, err := util.ParseGitURL(req.GetGitAddr())
	if err != nil {
		violations = append(violations, fieldViolation("git_addr", err))
	}
	if err := val.ValidateRepoVisibility(req.GetVisibilityLevel()); err != nil {
		violations = append(violations, fieldViolation("visibility_level", err))
	}

	if err := val.ValidateRepoSideBarTheme(req.GetThemeSidebar()); err != nil {
		violations = append(violations, fieldViolation("theme_sidebar", err))
	}
	if err := val.ValidateRepoThemeColor(req.GetThemeColor()); err != nil {
		violations = append(violations, fieldViolation("theme_color", err))
	}
	if err := val.ValidateString(req.GetBranch(), 0, 255); err != nil {
		violations = append(violations, fieldViolation("branch", err))
	}

	return violations
}
