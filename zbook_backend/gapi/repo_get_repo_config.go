package gapi

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRepoConfig(ctx context.Context, req *rpcs.GetRepoConfigRequest) (*rpcs.GetRepoConfigResponse, error) {
	violations := validateGetRepoConfigRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg_repo := db.GetRepoIDParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo_id, err := server.store.GetRepoID(ctx, arg_repo)
	if err != nil {
		log.Info().Msgf("get repo config get repo id failed:%s,%s", req.GetUsername(), req.GetRepoName())
		return nil, status.Errorf(codes.Internal, "get repo id failed: %s", err)
	}

	err = server.isRepoVisibleToCurrentUser(ctx, repo_id)
	if err != nil {
		return nil, err
	}

	arg := db.GetRepoConfigParams{Username: req.GetUsername(), RepoName: req.GetRepoName()}
	repo, err := server.store.GetRepoConfig(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo config not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo config failed: %s", err)
	}
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get user by id failed: %s", err)
	}

	path, err := util.GetDocumentPath(repo.Home, req.GetLang())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "parse home error : %s", err)
	}

	rsp := &rpcs.GetRepoConfigResponse{
		Config:          repo.Config,
		Username:        user.Username,
		VisibilityLevel: repo.VisibilityLevel,
		ThemeSidebar:    repo.ThemeSidebar,
		ThemeColor:      repo.ThemeColor,
		Home:            path,
	}
	return rsp, nil
}
func validateGetRepoConfigRequest(req *rpcs.GetRepoConfigRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	return violations
}
