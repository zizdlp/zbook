package gapi

import (
	"context"
	"net/netip"

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

	rsp := &rpcs.GetDailyVisitorsResponse{
		Visitors: convertVisitors(server, visitors, req.GetLang()),
	}
	return rsp, nil
}

func convertVisitors(server *Server, visitors []*VisitorData, lang string) []*rpcs.Visitor {
	var ret_reports []*rpcs.Visitor
	for i := 0; i < len(visitors); i++ {
		ip, err := netip.ParseAddr(visitors[i].IP)
		if err != nil {
			log.Error().Err(err).Msgf("can not parse ip addr: %s", visitors[i].IP)
			continue
		}
		record, err := server.store.GetGeoInfo(context.Background(), ip)
		if err != nil {
			// 如果解析出错，则将错误信息添加到响应中，继续处理下一个 IP
			ret_reports = append(ret_reports,
				&rpcs.Visitor{
					IP:    visitors[i].IP,
					Agent: visitors[i].Agent,
					Count: int32(visitors[i].Count),
				},
			)
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
					IP:    visitors[i].IP,
					Agent: visitors[i].Agent,
					Count: int32(visitors[i].Count),
					City:  city,
					Lat:   record.Latitude.Float64,
					Long:  record.Longitude.Float64,
				},
			)
		}
	}
	return ret_reports
}
