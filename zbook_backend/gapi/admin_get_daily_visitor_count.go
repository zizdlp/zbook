package gapi

import (
	"context"
	"sort"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetDailyVisitorCount(ctx context.Context, req *rpcs.GetDailyVisitorCountRequest) (*rpcs.GetDailyVisitorCountResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetDailyVisitorCount"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	counts, err := server.GetUniqueKeysCountForLastNDays(req.GetNdays(), req.GetTimeZone())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get unique visitor count failed: %s", err)
	}

	// 对结果进行排序，确保顺序从旧到新
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Date.Before(counts[j].Date)
	})

	// 分离日期和计数
	dates, countsList := convertDailyVisitorCount(counts)
	rsp := &rpcs.GetDailyVisitorCountResponse{
		Dates:  dates,
		Counts: countsList,
	}
	return rsp, nil
}

func convertDailyVisitorCount(users []DailyUniqueKeysCount) ([]string, []int32) {
	var dates []string
	var counts []int32
	for i := 0; i < len(users); i++ {
		// 只保留年月日
		formattedDate := users[i].Date.Format("2006-01-02")
		dates = append(dates, formattedDate)
		counts = append(counts, int32(users[i].Count))
	}
	return dates, counts
}
