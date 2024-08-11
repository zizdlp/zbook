package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) CreateFollow(ctx context.Context, req *rpcs.CreateFollowRequest) (*rpcs.CreateFollowResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "CreateFollow"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateCreateFollowRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}

	arg := db.CreateFollowTxParams{
		CreateFollowParams: db.CreateFollowParams{
			FollowerID:  authPayload.UserID,
			FollowingID: user.UserID,
		},
	}

	txResult, err := server.store.CreateFollowTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create follow failed: %s", err)
	}

	rsp := &rpcs.CreateFollowResponse{
		Follow: convertCreateFollow(txResult.Follow),
	}
	return rsp, nil
}
func validateCreateFollowRequest(req *rpcs.CreateFollowRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}

func convertCreateFollow(follow db.Follow) *models.Follow {
	return &models.Follow{
		FollowId:    follow.FollowID,
		FollowerId:  follow.FollowerID,
		FollowingId: follow.FollowingID,
		CreatedAt:   timestamppb.New(follow.CreatedAt),
	}
}
