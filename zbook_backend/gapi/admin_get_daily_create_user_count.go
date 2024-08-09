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

func (server *Server) GetDailyCreateUserCount(ctx context.Context, req *rpcs.GetDailyCreateUserCountRequest) (*rpcs.GetDailyCreateUserCountResponse, error) {
	violations := validateGetDailyCreateUserCount(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "GetDailyCreateUserCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	arg := db.GetDailyCreateUserCountParams{
		Timezone:     req.GetTimeZone(),
		IntervalDays: pgtype.Text{String: strconv.Itoa(int(req.GetNdays())), Valid: true},
	}

	counts, err := server.store.GetDailyCreateUserCount(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get daily create user count: %s", err)
	}

	// 对结果进行排序，确保顺序从旧到新
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].RegistrationDate.Time.Before(counts[j].RegistrationDate.Time)
	})

	// 将日期和计数拆分成两个数组
	dates, countsArray := convertDailyNewUserCount(counts)

	rsp := &rpcs.GetDailyCreateUserCountResponse{
		Dates:  dates,
		Counts: countsArray,
	}
	return rsp, nil
}

func convertDailyNewUserCount(users []db.GetDailyCreateUserCountRow) ([]string, []int32) {
	var dates []string
	var counts []int32

	for _, user := range users {
		// 只保留年月日
		formattedDate := user.RegistrationDate.Time.Format("2006-01-02")
		dates = append(dates, formattedDate)
		counts = append(counts, int32(user.NewUsersCount))
	}

	return dates, counts
}

func validateGetDailyCreateUserCount(req *rpcs.GetDailyCreateUserCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidTimeZone(req.GetTimeZone()); err != nil {
		violations = append(violations, fieldViolation("time_zone", err))
	}
	if err := val.ValidateInt32ID(req.GetNdays()); err != nil {
		violations = append(violations, fieldViolation("ndays", err))
	}
	return violations
}
