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

func (server *Server) GetFirstDocument(ctx context.Context, req *rpcs.GetFirstDocumentRequest) (*rpcs.GetFirstDocumentResponse, error) {
	violations := validateGetFirstDocumentRequest(req)
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

	arg_config := db.GetRepoConfigParams{Username: req.GetUsername(), RepoName: req.GetRepoName()}
	repo_config, err := server.store.GetRepoConfig(ctx, arg_config)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "get repo config not found error: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "GetRepoConfig error : %s", err)
	}
	config, err := util.ParseRepoConfigFromString(repo_config.Config)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ParseRepoConfigFromString error : %s", err)
	}
	relative_path, err := config.GetFirstDocument(req.GetLang())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetFirstDocument error : %s", err)
	}

	rsp := &rpcs.GetFirstDocumentResponse{
		RelativePath: relative_path,
	}
	return rsp, nil
}
func validateGetFirstDocumentRequest(req *rpcs.GetFirstDocumentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}

	err = val.ValidateLang(req.GetLang())
	if err != nil {
		violations = append(violations, fieldViolation("lang", err))
	}

	return violations
}
