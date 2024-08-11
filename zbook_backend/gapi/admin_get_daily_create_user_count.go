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

	// 补全日期范围并填充数据
	dates, countsArray := fillMissingDatesAndCounts(counts, req.GetNdays(), req.GetTimeZone())

	rsp := &rpcs.GetDailyCreateUserCountResponse{
		Dates:  dates,
		Counts: countsArray,
	}
	return rsp, nil
}
func fillMissingDatesAndCounts(users []db.GetDailyCreateUserCountRow, ndays int32, timezone string) ([]string, []int32) {
	var dates []string
	var countsMap = make(map[string]int32)

	// Load the location
	location, err := time.LoadLocation(timezone)
	if err != nil {
		// Handle the error according to your application's needs
		return nil, nil
	}

	// Calculate start and end dates
	endDate := time.Now().In(location)
	endDate = time.Date(endDate.Year(), endDate.Month(), endDate.Day(), 0, 0, 0, 0, location)
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
		countsMap[formattedDate] = int32(user.NewUsersCount)
	}

	// Fill counts array in the same order as dates
	var counts []int32
	for _, date := range dates {
		counts = append(counts, countsMap[date])
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
