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

	arg := db.DeleteRepoRelationParams{
		RepoID:       req.GetRepoId(),
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
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	if req.GetRelationType() != util.RelationTypeDislike &&
		req.GetRelationType() != util.RelationTypeLike &&
		req.GetRelationType() != util.RelationTypeShare {
		violations = append(violations, fieldViolation("relation_type", errors.New("invalid relation_type")))
	}
	return violations
}
