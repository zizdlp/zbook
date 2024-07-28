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

func (server *Server) DeleteRepoRelation(ctx context.Context, req *rpcs.DeleteRepoRelationRequest) (*rpcs.DeleteRepoRelationResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "DeleteRepoRelation"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateDeleteRepoRelationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	relationType := util.RelationTypeShare
	if req.GetRelationType() == util.RelationTypeDislike {
		relationType = util.RelationTypeDislike
	} else if req.GetRelationType() == util.RelationTypeLike {
		relationType = util.RelationTypeLike
	} else if req.GetRelationType() == util.RelationTypeShare {
		relationType = util.RelationTypeShare
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Unknown RelationType")
	}

	arg_get := db.GetRepoBasicInfoParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo, err := server.store.GetRepoBasicInfo(ctx, arg_get)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo basic info failed: %s", err)
	}
	err = server.isRepoVisibleToCurrentUser(ctx, repo.RepoID)
	if err != nil {
		return nil, err
	}
	arg := db.DeleteRepoRelationParams{
		RepoID:       repo.RepoID,
		UserID:       authPayload.UserID,
		RelationType: relationType,
	}
	err = server.store.DeleteRepoRelation(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation || db.ErrorCode(err) == db.ForeignKeyViolation {
			return nil, status.Errorf(codes.AlreadyExists, "repo relation already exists: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "delete repo relation failed: %s", err)
	}

	rsp := &rpcs.DeleteRepoRelationResponse{}
	return rsp, nil
}
func validateDeleteRepoRelationRequest(req *rpcs.DeleteRepoRelationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	if req.GetRelationType() != util.RelationTypeDislike &&
		req.GetRelationType() != util.RelationTypeLike &&
		req.GetRelationType() != util.RelationTypeShare {
		violations = append(violations, fieldViolation("relation_type", errors.New("invalid relation_type")))
	}
	return violations
}
