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

func (server *Server) CreateCommentRelation(ctx context.Context, req *rpcs.CreateCommentRelationRequest) (*rpcs.CreateCommentRelationResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "CreateCommentRelation"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateCreateCommentRelationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err = server.isCommentVisibleToCurrentUser(ctx, req.GetCommentId())
	if err != nil {
		return nil, err
	}
	relationType := util.RelationTypeShare
	if req.GetRelationType() == util.RelationTypeDislike {
		relationType = util.RelationTypeDislike
	} else if req.GetRelationType() == util.RelationTypeLike {
		relationType = util.RelationTypeLike
	} else if req.GetRelationType() == util.RelationTypeShare {
		relationType = util.RelationTypeShare
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "Unknown RelationType: %s", relationType)
	}
	arg := db.CreateCommentRelationParams{
		CommentID:    req.GetCommentId(),
		UserID:       authPayload.UserID,
		RelationType: relationType,
	}

	err = server.store.CreateCommentRelation(ctx, arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation || db.ErrorCode(err) == db.ForeignKeyViolation {
			return nil, status.Errorf(codes.AlreadyExists, "comment relation already exist: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "failed to create comment relation: %s", err)
	}

	rsp := &rpcs.CreateCommentRelationResponse{}
	return rsp, nil
}
func validateCreateCommentRelationRequest(req *rpcs.CreateCommentRelationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateID(req.GetCommentId()); err != nil {
		violations = append(violations, fieldViolation("comment_id", err))
	}
	if req.GetRelationType() != util.RelationTypeDislike &&
		req.GetRelationType() != util.RelationTypeLike &&
		req.GetRelationType() != util.RelationTypeShare {
		violations = append(violations, fieldViolation("relation_type", errors.New("invalid relation_type")))
	}
	return violations
}
