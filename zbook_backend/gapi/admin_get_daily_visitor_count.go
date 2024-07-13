package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetDailyVisitorCount(ctx context.Context, req *rpcs.GetDailyVisitorCountRequest) (*rpcs.GetDailyVisitorCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetDailyVisitorCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	counts, err := server.GetUniqueKeysCountForLastNDays(req.GetNdays())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get unique visitor count failed: %s", err)
	}
	rsp := &rpcs.GetDailyVisitorCountResponse{
		Counts: convertDailyVisitorCount(counts),
	}
	return rsp, nil
}

func convertDailyVisitorCount(users []DailyUniqueKeysCount) []*models.DailyCount {
	var ret_users []*models.DailyCount
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.DailyCount{
				Date:  timestamppb.New(users[i].Date),
				Count: int64(users[i].Count),
			},
		)
	}
	return ret_users
}
