package gapi

import (
	"context"
	"encoding/json"
	"time"

	"math/rand"

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

		// 尝试从 Redis 缓存中获取地理信息
		record_recover, err := server.redisClient.Get("geocache_" + ipStr).Result()
		if err == nil && record_recover != "" {
			// 如果缓存中存在，则解析缓存数据
			record := &util.GeoInfo{}
			err := json.Unmarshal([]byte(record_recover), record)
			if err != nil {
				log.Error().Err(err).Msgf("can not unmarshal geo info from cache for ip addr: %s", ipStr)
				continue
			}
			// 在这里添加随机扰动
			latWithNoise := record.Latitude
			longWithNoise := record.Longitude

			ret_reports = append(ret_reports,
				&rpcs.Visitor{
					Ip:    ipStr,
					City:  getCityName(record, lang),
					Count: int32(count),
					Lat:   latWithNoise,
					Long:  longWithNoise,
				},
			)
			continue
		}

		// 如果缓存中不存在，则从数据库中获取
		record, err := server.store.GetGeoInfo(context.Background(), ip)
		if err != nil {
			log.Error().Err(err).Msgf("can not get GeoInfo for ip addr: %s", ipStr)
			continue
		}
		cityNameEn := ""
		if record.CityNameEn.Valid {
			cityNameEn = record.CityNameEn.String
		}
		cityNameZhCn := ""
		if record.CityNameZhCn.Valid {
			cityNameZhCn = record.CityNameZhCn.String
		}
		// 存储到 Redis 缓存,
		// 在这里添加随机扰动
		record_cache := util.GeoInfo{
			IpStr:        ipStr,
			Latitude:     record.Latitude.Float64 + (rand.Float64()*2 - 1),
			Longitude:    record.Longitude.Float64 + (rand.Float64()*2 - 1),
			CityNameEn:   cityNameEn,
			CityNameZhCn: cityNameZhCn,
		}
		recordBytes, err := json.Marshal(record_cache)
		if err != nil {
			log.Error().Err(err).Msgf("can not marshal geo info to cache for ip addr: %s", ipStr)
			continue
		}
		server.redisClient.Set("geocache_"+ipStr, recordBytes, 24*7*time.Hour)
		ret_reports = append(ret_reports,
			&rpcs.Visitor{
				Ip:    ipStr,
				City:  getCityName(&record_cache, lang),
				Count: int32(count),
				Lat:   record_cache.Latitude,
				Long:  record_cache.Longitude,
			},
		)
	}

	// 按照 Count 从大到小排序
	sort.Slice(ret_reports, func(i, j int) bool {
		return ret_reports[i].Count > ret_reports[j].Count
	})
	return ret_reports
}

// 获取城市名称的辅助函数
func getCityName(record *util.GeoInfo, lang string) string {
	switch lang {
	case "en":
		return record.CityNameEn
	case "zh_cn":
		return record.CityNameZhCn
	default:
		return record.CityNameEn
	}
}
