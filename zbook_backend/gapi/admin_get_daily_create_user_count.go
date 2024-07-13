package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetDailyCreateUserCount(ctx context.Context, req *rpcs.GetDailyCreateUserCountRequest) (*rpcs.GetDailyCreateUserCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetDailyCreateUserCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	counts, err := server.store.GetDailyCreateUserCount(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get daily create user count: %s", err)
	}

	rsp := &rpcs.GetDailyCreateUserCountResponse{
		Counts: convertDailyNewUserCount(counts),
	}
	return rsp, nil
}

func convertDailyNewUserCount(users []db.GetDailyCreateUserCountRow) []*models.DailyCount {
	var ret_users []*models.DailyCount
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.DailyCount{
				Date:  timestamppb.New(users[i].RegistrationDate.Time),
				Count: users[i].NewUsersCount,
			},
		)
	}
	return ret_users
}
