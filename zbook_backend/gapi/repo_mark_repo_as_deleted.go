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

func (server *Server) MarkRepoAsDeleted(ctx context.Context, req *rpcs.MarkRepoAsDeletedRequest) (*rpcs.MarkRepoAsDeletedResponse, error) {
	violations := validateMarkRepoAsDeletedRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "MarkRepoAsDeleted"
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
	arg := db.MarkRepoAsDeletedTxParams{
		RepoID: req.GetRepoId(),
		UserID: repo.UserID,
	}

	err = server.store.MarkRepoAsDeletedTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := &rpcs.MarkRepoAsDeletedResponse{}
	return rsp, nil
}
func validateMarkRepoAsDeletedRequest(req *rpcs.MarkRepoAsDeletedRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	return violations
}
