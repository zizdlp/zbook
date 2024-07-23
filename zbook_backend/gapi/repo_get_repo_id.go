package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetRepoID(ctx context.Context, req *rpcs.GetRepoIDRequest) (*rpcs.GetRepoIDResponse, error) {
	violations := validateGetRepoIDRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.GetRepoIDParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}

	repo_id, err := server.store.GetRepoID(ctx, arg)
	if err != nil {
		log.Info().Msgf("get repo id failed:%s,%s", req.GetUsername(), req.GetRepoName())
		return nil, status.Errorf(codes.Internal, "get repo id failed: %s", err)
	}
	rsp := &rpcs.GetRepoIDResponse{
		RepoId: repo_id,
	}
	return rsp, nil
}
func validateGetRepoIDRequest(req *rpcs.GetRepoIDRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	err := val.ValidateString(req.GetRepoName(), 1, 32)
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	err = val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
