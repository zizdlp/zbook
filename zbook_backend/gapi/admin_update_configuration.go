package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateConfiguration(ctx context.Context, req *rpcs.UpdateConfigurationRequest) (*rpcs.UpdateConfigurationResponse, error) {
	violations := validateUpdateConfigurationRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "UpdateConfiguration"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	arg := db.UpdateConfigurationParams{
		ConfigName:  req.GetConfigName(),
		ConfigValue: req.GetConfigValue(),
	}
	err = server.store.UpdateConfiguration(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "update configuration failed: %s", err)
	}

	rsp := &rpcs.UpdateConfigurationResponse{}
	return rsp, nil
}

func validateUpdateConfigurationRequest(req *rpcs.UpdateConfigurationRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetConfigName(), 1, 32); err != nil {
		violations = append(violations, fieldViolation("config_name", err))
	}
	return violations
}
