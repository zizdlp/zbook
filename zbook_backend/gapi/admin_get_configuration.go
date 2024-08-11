package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetConfiguration(ctx context.Context, req *rpcs.GetConfigurationRequest) (*rpcs.GetConfigurationResponse, error) {
	violations := validateGetConfigurationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "GetConfiguration"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	config, err := server.store.GetConfiguration(ctx, req.GetConfigName())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get configuration failed: %s", err)
	}

	rsp := &rpcs.GetConfigurationResponse{
		ConfigValue: config.ConfigValue,
	}
	return rsp, nil
}

func validateGetConfigurationRequest(req *rpcs.GetConfigurationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetConfigName(), 1, 32); err != nil {
		violations = append(violations, fieldViolation("config_name", err))
	}
	return violations
}
