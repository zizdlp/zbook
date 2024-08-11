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

func (server *Server) DeleteRepoVisibility(ctx context.Context, req *rpcs.DeleteRepoVisibilityRequest) (*rpcs.DeleteRepoVisibilityResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "DeleteRepoVisibility"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateDeleteRepoVisibilityRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}

	arg_repo := db.GetRepoIDParams{
		Username: req.GetRepoUsername(),
		RepoName: req.GetRepoName(),
	}
	repo_id, err := server.store.GetRepoID(ctx, arg_repo)
	if err != nil {
		log.Info().Msgf("get repo layout get repo id failed:%s,%s", req.GetUsername(), req.GetRepoName())
		return nil, status.Errorf(codes.Internal, "get repo id failed: %s", err)
	}
	repo, err := server.store.GetRepo(ctx, repo_id)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo failed: %s", err)
	}
	if authPayload.UserID != repo.UserID {
		return nil, status.Error(codes.PermissionDenied, "cannot update other user's repo")
	}

	arg := db.DeleteRepoRelationParams{
		RepoID:       repo.RepoID,
		UserID:       user.UserID,
		RelationType: util.RelationTypeVisi,
	}
	server.store.DeleteRepoRelation(ctx, arg)

	rsp := &rpcs.DeleteRepoVisibilityResponse{}
	return rsp, nil
}
func validateDeleteRepoVisibilityRequest(req *rpcs.DeleteRepoVisibilityRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateUsername(req.GetRepoUsername())
	if err != nil {
		violations = append(violations, fieldViolation("repo_username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	return violations
}
