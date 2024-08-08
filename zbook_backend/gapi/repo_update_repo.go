package gapi

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateRepoInfo(ctx context.Context, req *rpcs.UpdateRepoInfoRequest) (*rpcs.UpdateRepoInfoResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "UpdateRepoInfo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateUpdateRepoInfoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg_get := db.GetRepoBasicInfoParams{
		Username: req.GetUsername(),
		RepoName: req.GetOldRepoName(),
	}
	repo_get, err := server.store.GetRepoBasicInfo(ctx, arg_get)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo failed: %s", err)
	}
	if authPayload.UserID != repo_get.UserID {
		return nil, status.Error(codes.PermissionDenied, "cannot update other user's repo")
	}

	arg := db.UpdateRepoInfoParams{
		RepoID: repo_get.RepoID,
		RepoName: pgtype.Text{
			String: req.GetRepoName(),
			Valid:  len(req.GetRepoName()) != 0,
		},
		RepoDescription: pgtype.Text{String: req.GetRepoDescription(), Valid: req.GetRepoDescription() != ""},

		GitAccessToken: pgtype.Text{
			String: req.GetGitAccessToken(),
			Valid:  len(req.GetGitAccessToken()) != 0,
		},

		VisibilityLevel: pgtype.Text{
			String: req.GetVisibilityLevel(),
			Valid:  len(req.GetVisibilityLevel()) != 0,
		},
		ThemeSidebar: pgtype.Text{String: req.GetThemeSidebar(), Valid: len(req.GetThemeSidebar()) != 0},
		ThemeColor:   pgtype.Text{String: req.GetThemeColor(), Valid: len(req.GetThemeColor()) != 0},
		SyncToken:    pgtype.Text{String: req.GetSyncToken(), Valid: len(req.GetSyncToken()) != 0},
		HomePage:     pgtype.Text{String: strings.ToLower(req.GetHomePage()), Valid: len(req.GetHomePage()) != 0},
	}

	_, err = server.store.UpdateRepoInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "UpdateRepoInfo not found error: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "UpdateRepoInfo error: %s", err)
	}

	rsp := &rpcs.UpdateRepoInfoResponse{}
	return rsp, nil
}
func validateUpdateRepoInfoRequest(req *rpcs.UpdateRepoInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetOldRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}

	if req.GetRepoName() != "" {
		if err = val.ValidateRepoName(req.GetRepoName()); err != nil {
			violations = append(violations, fieldViolation("repo_name", err))
		}
	}
	if req.GetRepoDescription() != "" {
		if err = val.ValidateString(req.GetRepoDescription(), 1, 512); err != nil {
			violations = append(violations, fieldViolation("repo_description", err))
		}
	}
	if req.GetVisibilityLevel() != "" {
		if err := val.ValidateRepoVisibility(req.GetVisibilityLevel()); err != nil {
			violations = append(violations, fieldViolation("visibility_level", err))
		}
	}
	if req.GetThemeSidebar() != "" {
		if err := val.ValidateRepoSideBarTheme(req.GetThemeSidebar()); err != nil {
			violations = append(violations, fieldViolation("theme_sidebar", err))
		}
	}
	if req.GetThemeColor() != "" {
		if err := val.ValidateRepoThemeColor(req.GetThemeColor()); err != nil {
			violations = append(violations, fieldViolation("theme_color", err))
		}
	}

	return violations
}
