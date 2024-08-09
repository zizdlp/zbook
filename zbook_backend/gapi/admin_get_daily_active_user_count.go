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

func (server *Server) GetDailyActiveUserCount(ctx context.Context, req *rpcs.GetDailyActiveUserCountRequest) (*rpcs.GetDailyActiveUserCountResponse, error) {
	// 校验 timezone 参数
	violations := validateGetDailyActiveUserCountRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	apiUserDailyLimit := 10000
	apiKey := "GetDailyActiveUserCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	counts, err := server.store.GetDailyActiveUserCount(ctx, req.GetTimeZone())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get daily active user count failed: %s", err)
	}

	rsp := &rpcs.GetDailyActiveUserCountResponse{
		Counts: convertDailyActiveUserCount(counts),
	}
	return rsp, nil
}

// 校验 GetDailyActiveUserCountRequest 请求中的 timezone 参数
func validateGetDailyActiveUserCountRequest(req *rpcs.GetDailyActiveUserCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidTimeZone(req.GetTimeZone()); err != nil {
		violations = append(violations, fieldViolation("time_zone", err))
	}
	return violations
}

func convertDailyActiveUserCount(users []db.GetDailyActiveUserCountRow) []*models.DailyCount {
	var ret_users []*models.DailyCount
	for i := 0; i < len(users); i++ {
		ret_users = append(ret_users,
			&models.DailyCount{
				Date:  timestamppb.New(users[i].RegistrationDate.Time),
				Count: users[i].ActiveUsersCount,
			},
		)
	}
	return ret_users
}
