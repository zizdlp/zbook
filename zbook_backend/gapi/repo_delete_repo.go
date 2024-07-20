package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteRepo(ctx context.Context, req *rpcs.DeleteRepoRequest) (*rpcs.DeleteRepoResponse, error) {
	violations := validateDeleteRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "DeleteRepo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	repo, err := server.store.GetRepo(ctx, req.GetRepoId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo failed: %s", err)
	}

	if repo.UserID != authPayload.UserID && authPayload.UserRole != util.AdminRole {
		return nil, status.Errorf(codes.PermissionDenied, "current account do not have enough permission")
	}
	arg := db.DeleteRepoTxParams{
		RepoID: req.GetRepoId(),
		UserID: repo.UserID,
	}

	err = server.store.DeleteRepoTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := &rpcs.DeleteRepoResponse{}
	return rsp, nil
}
func validateDeleteRepoRequest(req *rpcs.DeleteRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	return violations
}
