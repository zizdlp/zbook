package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUserOnBoarding(ctx context.Context, req *rpcs.UpdateUserOnBoardingRequest) (*rpcs.UpdateUserOnBoardingResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "UpdateUserOnBoarding"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	arg := db.UpdateUserBasicInfoParams{
		Username:   authPayload.Username,
		Onboarding: pgtype.Bool{Bool: req.GetOnboarding(), Valid: false},
	}
	user, err := server.store.UpdateUserBasicInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "update user failed: %s", err)
	}

	rsp := &rpcs.UpdateUserOnBoardingResponse{Onboarding: user.Onboarding}
	return rsp, nil
}
