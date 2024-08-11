package gapi

import (
	"context"

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

func (server *Server) ListFollower(ctx context.Context, req *rpcs.ListFollowerRequest) (*rpcs.ListFollowerResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "ListFollower"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateListFollowerRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	_, err = server.getUserPermessionlevel(ctx, authPayload.Username, req.GetUsername())
	if err != nil {
		return nil, err
	}
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if req.GetQuery() != "" {
		arg := db.QueryFollowerParams{
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			UserID:    user.UserID,
			CurUserID: authPayload.UserID,
			Query:     req.GetQuery(),
			Role:      authPayload.UserRole,
		}

		follows, err := server.store.QueryFollower(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query follower failed: %s", err)
		}

		rsp := &rpcs.ListFollowerResponse{
			Elements: convertQueryFollower(follows),
		}
		return rsp, nil
	}
	arg := db.ListFollowerParams{
		UserID:    user.UserID,
		Limit:     req.GetPageSize(),
		Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
		CurUserID: authPayload.UserID,
		Role:      authPayload.UserRole,
	}

	follows, err := server.store.ListFollower(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list follower failed: %s", err)
	}

	rsp := &rpcs.ListFollowerResponse{
		Elements: convertListFollower(follows),
	}
	return rsp, nil
}
func validateListFollowerRequest(req *rpcs.ListFollowerRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}

func convertListFollower(users []db.ListFollowerRow) []*models.ListFollowInfo {
	var ret_users []*models.ListFollowInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListFollowInfo{
				Username:    users[i].Username,
				Email:       users[i].Email,
				IsFollowing: users[i].IsFollowing,
				UpdatedAt:   timestamppb.New(users[i].UpdatedAt),
				CreatedAt:   timestamppb.New(users[i].CreatedAt),
				RepoCount:   int32(users[i].RepoCount),
			},
		)
	}
	return ret_users
}

func convertQueryFollower(users []db.QueryFollowerRow) []*models.ListFollowInfo {
	var ret_users []*models.ListFollowInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListFollowInfo{
				Username:    users[i].Username,
				Email:       users[i].Email,
				IsFollowing: users[i].IsFollowing,
				UpdatedAt:   timestamppb.New(users[i].UpdatedAt),
				CreatedAt:   timestamppb.New(users[i].CreatedAt),
				RepoCount:   int32(users[i].RepoCount),
			},
		)
	}
	return ret_users
}
