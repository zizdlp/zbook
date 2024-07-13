package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateOAuthLink(ctx context.Context, req *rpcs.CreateOAuthLinkRequest) (*rpcs.CreateOAuthLinkResponse, error) {
	apiUserDailyLimit := 100
	apiKey := "CreateOAuthLink"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateCreateOAuthLinkRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	oauthType := util.OAuthTypeGoogle
	if req.GetOauthType() == util.OAuthTypeGithub {
		oauthType = util.OAuthTypeGithub
	}
	arg_oauth := db.CreateOAuthParams{
		UserID:    authPayload.UserID,
		OauthType: oauthType,
		AppID:     req.GetAppId(),
	}
	_, err = server.store.CreateOAuth(ctx, arg_oauth)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "check oauth failed: %s", err)
	}

	rsp := &rpcs.CreateOAuthLinkResponse{}
	return rsp, nil
}
func validateCreateOAuthLinkRequest(req *rpcs.CreateOAuthLinkRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetAppId(), 1, 256); err != nil {
		violations = append(violations, fieldViolation("app_id", err))
	}
	if req.GetOauthType() != util.OAuthTypeGithub &&
		req.GetOauthType() != util.OAuthTypeGoogle {
		violations = append(violations, fieldViolation("oauth_type", errors.New("invalid oauth_type")))
	}

	return violations
}
