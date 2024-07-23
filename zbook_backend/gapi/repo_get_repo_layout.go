package gapi

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRepoLayout(ctx context.Context, req *rpcs.GetRepoLayoutRequest) (*rpcs.GetRepoLayoutResponse, error) {
	violations := validateGetRepoLayoutRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	arg_repo := db.GetRepoIDParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo_id, err := server.store.GetRepoID(ctx, arg_repo)
	if err != nil {
		log.Info().Msgf("get repo layout get repo id failed:%s,%s", req.GetUsername(), req.GetRepoName())
		return nil, status.Errorf(codes.Internal, "get repo id failed: %s", err)
	}

	err = server.isRepoVisibleToCurrentUser(ctx, repo_id)
	if err != nil {
		return nil, err
	}

	arg := db.GetRepoLayoutParams{Username: req.GetUsername(), RepoName: req.GetRepoName()}
	repo, err := server.store.GetRepoLayout(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo layout not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo layout failed: %s", err)
	}
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get user by id failed: %s", err)
	}

	rsp := &rpcs.GetRepoLayoutResponse{
		RepoId:          repo.RepoID,
		Layout:          repo.Layout,
		Username:        user.Username,
		VisibilityLevel: repo.VisibilityLevel,
	}
	return rsp, nil
}
func validateGetRepoLayoutRequest(req *rpcs.GetRepoLayoutRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateString(req.GetRepoName(), 1, 64)
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	return violations
}
