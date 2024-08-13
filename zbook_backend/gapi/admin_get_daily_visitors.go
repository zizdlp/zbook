package gapi

import (
	"context"
	"net/netip"
	"sort"

	"github.com/rs/zerolog/log"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetDailyVisitors(ctx context.Context, req *rpcs.GetDailyVisitorsRequest) (*rpcs.GetDailyVisitorsResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetDailyVisitors"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	visitors, err := server.GetDailyVisitorsForLastNDays(req.GetNdays())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get daily visitor failed: %s", err)
	}
	unique_visitors := util.AggregateByIP(visitors)
	ret_visitors := ConvertVisitor(server, unique_visitors, req.GetLang())
	agents := util.SumAgentCounts(visitors)
	rsp := &rpcs.GetDailyVisitorsResponse{
		Visitors:   ret_visitors,
		AgentCount: ConvertAgent(agents),
	}
	return rsp, nil
}
func ConvertAgent(agent util.AgentCounts) *rpcs.AgentCount {
	// 初始化 ret_agent
	ret_agent := &rpcs.AgentCount{
		Bot:      int32(agent.Bot),
		Computer: int32(agent.Computer),
		Phone:    int32(agent.Phone), // 修复这里的错配
		Tablet:   int32(agent.Tablet),
		Unknown:  int32(agent.Unknown),
	}
	return ret_agent
}

func ConvertVisitor(server *Server, visitors map[string]int, lang string) []*rpcs.Visitor {
	var ret_reports []*rpcs.Visitor
	for ipStr, count := range visitors {
		ip, err := netip.ParseAddr(ipStr)
		if err != nil {
			log.Error().Err(err).Msgf("can not parse ip addr: %s", ipStr)
			continue
		}
		record, err := server.store.GetGeoInfo(context.Background(), ip)
		if err != nil {
			log.Error().Err(err).Msgf("can not get GeoInfo for ip addr: %s", ipStr)
			continue
		} else {
			// 如果解析成功，则将城市、经度和纬度信息添加到响应中
			city := ""
			if lang == "en" {
				if record.CityNameEn.Valid {
					city = record.CityNameEn.String
				}
			} else if record.CityNameZhCn.Valid {
				city = record.CityNameZhCn.String
			} else {
				if record.CityNameEn.Valid {
					city = record.CityNameEn.String
				}
			}
			ret_reports = append(ret_reports,
				&rpcs.Visitor{
					Ip:    ipStr,
					City:  city,
					Count: int32(count),
					Lat:   record.Latitude.Float64,
					Long:  record.Longitude.Float64,
				},
			)
		}
	}
	// 按照 Count 从大到小排序
	sort.Slice(ret_reports, func(i, j int) bool {
		return ret_reports[i].Count > ret_reports[j].Count
	})
	return ret_reports
}
