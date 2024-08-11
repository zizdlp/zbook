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

func (server *Server) ListUser(ctx context.Context, req *rpcs.ListUserRequest) (*rpcs.ListUserResponse, error) {
	violations := validateListUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListUser"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		arg := db.QueryUserParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			Query:  req.GetQuery(),
			Role:   authUser.UserRole,
		}
		users, err := server.store.QueryUser(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query user failed: %s", err)
		}
		rsp := &rpcs.ListUserResponse{
			Elements: convertQueryUser(users),
		}
		return rsp, nil
	} else {
		arg := db.ListUserParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			Role:   authUser.UserRole,
		}

		users, err := server.store.ListUser(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "list user failed: %s", err)
		}

		rsp := &rpcs.ListUserResponse{
			Elements: convertListUser(users),
		}
		return rsp, nil
	}

}
func validateListUserRequest(req *rpcs.ListUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}
func convertListUser(users []db.User) []*models.ListUserInfo {
	var ret_users []*models.ListUserInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListUserInfo{
				Username:   users[i].Username,
				Email:      users[i].Email,
				Blocked:    users[i].Blocked,
				Verified:   users[i].Verified,
				Onboarding: users[i].Onboarding,
				UpdatedAt:  timestamppb.New(users[i].UpdatedAt),
				CreatedAt:  timestamppb.New(users[i].CreatedAt),
				Role:       users[i].UserRole,
			},
		)
	}
	return ret_users
}
func convertQueryUser(users []db.QueryUserRow) []*models.ListUserInfo {
	var ret_users []*models.ListUserInfo
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.ListUserInfo{
				Username:   users[i].Username,
				Email:      users[i].Email,
				Blocked:    users[i].Blocked,
				Verified:   users[i].Verified,
				Onboarding: users[i].Onboarding,
				Role:       users[i].UserRole,
				UpdatedAt:  timestamppb.New(users[i].UpdatedAt),
				CreatedAt:  timestamppb.New(users[i].CreatedAt),
			},
		)
	}
	return ret_users
}
