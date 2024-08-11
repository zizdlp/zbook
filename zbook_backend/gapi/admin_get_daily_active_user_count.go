package gapi

import (
	"context"
	"sort"
	"strconv"
	"time"

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

	// 将日期和计数拆分成两个数组，并补全缺失日期
	dates, countsArray := fillMissingDatesAndCountsForActiveUser(counts, req.GetNdays(), req.GetTimeZone())
	rsp := &rpcs.GetDailyActiveUserCountResponse{
		Dates:  dates,
		Counts: countsArray,
	}
	return rsp, nil
}
func fillMissingDatesAndCountsForActiveUser(users []db.GetDailyActiveUserCountRow, ndays int32, timezone string) ([]string, []int32) {
	var dates []string
	var countsMap = make(map[string]int32)

	// Load the location
	location, err := time.LoadLocation(timezone)
	if err != nil {
		// Handle the error according to your application's needs
		return nil, nil
	}

	// Get the end date in the specified timezone
	endDate := time.Now().In(location)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, location)

	// Calculate the start date
	startDate := endDate.AddDate(0, 0, -int(ndays))

	// Generate date range
	for date := startDate; !date.After(endDate); date = date.AddDate(0, 0, 1) {
		formattedDate := date.Format("2006-01-02")
		dates = append(dates, formattedDate)
		countsMap[formattedDate] = 0
	}

	// Update counts map with actual data
	for _, user := range users {
		formattedDate := user.RegistrationDate.Time.Format("2006-01-02")
		countsMap[formattedDate] = int32(user.ActiveUsersCount)
	}

	// Fill counts array in the same order as dates
	var counts []int32
	for _, date := range dates {
		counts = append(counts, countsMap[date])
	}

	return dates, counts
}

func validateGetDailyActiveUserCountRequest(req *rpcs.GetDailyActiveUserCountRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidTimeZone(req.GetTimeZone()); err != nil {
		violations = append(violations, fieldViolation("time_zone", err))
	}
	if err := val.ValidateInt32ID(req.GetNdays()); err != nil {
		violations = append(violations, fieldViolation("ndays", err))
	}
	return violations
}
