package gapi

import (
	"context"
	"sort"
	"strconv"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	arg := db.GetDailyActiveUserCountParams{
		Timezone:     req.GetTimeZone(),
		IntervalDays: pgtype.Text{String: strconv.Itoa(int(req.GetNdays())), Valid: true},
	}

	counts, err := server.store.GetDailyActiveUserCount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get daily active user count failed: %s", err)
	}

	// 对结果进行排序，确保顺序从旧到新
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].RegistrationDate.Time.Before(counts[j].RegistrationDate.Time)
	})

	// 将日期和计数拆分成两个数组
	dates, countsArray := convertDailyActiveUserCount(counts)

	rsp := &rpcs.GetDailyActiveUserCountResponse{
		Dates:  dates,
		Counts: countsArray,
	}
	return rsp, nil
}

// 校验 GetDailyActiveUserCountRequest 请求中的 timezone 参数
func validateGetDailyActiveUserCountRequest(req *rpcs.GetDailyActiveUserCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidTimeZone(req.GetTimeZone()); err != nil {
		violations = append(violations, fieldViolation("time_zone", err))
	}
	if err := val.ValidateInt32ID(req.GetNdays()); err != nil {
		violations = append(violations, fieldViolation("ndays", err))
	}
	return violations
}

func convertDailyActiveUserCount(users []db.GetDailyActiveUserCountRow) ([]string, []int32) {
	var dates []string
	var counts []int32

	for _, user := range users {
		// 只保留年月日
		formattedDate := user.RegistrationDate.Time.Format("2006-01-02")
		dates = append(dates, formattedDate)
		counts = append(counts, int32(user.ActiveUsersCount))
	}

	return dates, counts
}
