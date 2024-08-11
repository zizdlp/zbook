package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) DeleteOAuthLink(ctx context.Context, req *rpcs.DeleteOAuthLinkRequest) (*rpcs.DeleteOAuthLinkResponse, error) {
	apiUserDailyLimit := 100
	apiKey := "DeleteOAuthLink"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	violations := validateDeleteOAuthLinkRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	oauthType := util.OAuthTypeGoogle
	if req.GetOauthType() == util.OAuthTypeGithub {
		oauthType = util.OAuthTypeGithub
	}
	arg_oauth := db.DeleteOAuthParams{
		UserID:    authPayload.UserID,
		OauthType: oauthType,
	}
	_, err = server.store.DeleteOAuth(ctx, arg_oauth)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "delete oauth failed: %s", err)
	}

	rsp := &rpcs.DeleteOAuthLinkResponse{}
	return rsp, nil
}
func validateDeleteOAuthLinkRequest(req *rpcs.DeleteOAuthLinkRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.GetOauthType() != util.OAuthTypeGithub &&
		req.GetOauthType() != util.OAuthTypeGoogle {
		violations = append(violations, fieldViolation("oauth_type", errors.New("invalid oauth_type")))
	}
	return violations
}
