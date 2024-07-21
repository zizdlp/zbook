package gapi

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GithubUser struct {
	ID int `json:"id"`
}

func (server *Server) LoginByOAuth(ctx context.Context, req *rpcs.LoginByOAuthRequest) (*rpcs.LoginByOAuthResponse, error) {
	violations := validateLoginByOAuthRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	// Determine oauth party type
	oauthType := req.GetOauthType()
	if oauthType != util.OAuthTypeGoogle && oauthType != util.OAuthTypeGithub {
		return nil, status.Errorf(codes.InvalidArgument, "unsupported oauth party type")
	}

	if oauthType == util.OAuthTypeGithub {
		// check github account

		resp, err := util.FetchGithub(req.GetAccessToken())
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "github account is invalid: %s", err)
		}
		defer resp.Body.Close()

		// 读取响应
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "read github response failed: %s", err)
		}

		// 解析 JSON 响应
		var user GithubUser
		err = json.Unmarshal(body, &user)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "parse github response failed: %s", err)
		}
		// 将 int 转换为 string 进行比较
		userIDStr := strconv.Itoa(user.ID)
		if userIDStr != req.GetAppId() {
			return nil, status.Errorf(codes.PermissionDenied, "github account does not match: %s", userIDStr)
		}
		log.Info().Msg("github account is matched")
	}

	arg := db.GetOAuthUserParams{
		OauthType: oauthType,
		AppID:     req.GetAppId(),
	}
	user, err := server.store.GetOAuthUser(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not exist,please link an account first")
		}
		return nil, status.Errorf(codes.Internal, "login in failed: %s", err)
	}

	err = server.checkUserStatus(ctx, user.Blocked, user.UserRole, user.Verified)
	if err != nil {
		return nil, err
	}
	response, err := server.CreateLoginPart(ctx, user.Username, user.UserRole, user.UserID, "LoginByOAuth")
	if err != nil {
		return nil, err
	}
	return &rpcs.LoginByOAuthResponse{
		Role:                  response.Role,
		AccessToken:           response.AccessToken,
		RefreshToken:          response.RefreshToken,
		Username:              response.Username,
		AccessTokenExpiresAt:  response.AccessTokenExpiresAt,
		RefreshTokenExpiresAt: response.RefreshTokenExpiresAt,
	}, nil
}
func validateLoginByOAuthRequest(req *rpcs.LoginByOAuthRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetAppId(), 1, 256); err != nil {
		violations = append(violations, fieldViolation("app_id", err))
	}
	if req.GetOauthType() != util.OAuthTypeGithub &&
		req.GetOauthType() != util.OAuthTypeGoogle {
		violations = append(violations, fieldViolation("oauth_type", errors.New("invalid oauth_type")))
	}
	if err := val.ValidateString(req.GetAccessToken(), 1, 256); err != nil {
		violations = append(violations, fieldViolation("access_token", err))
	}
	return violations
}
