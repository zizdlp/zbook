package gapi

import (
	"context"
	"errors"
	"time"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) RefreshToken(ctx context.Context, req *rpcs.RefreshTokenRequest) (*rpcs.RefreshTokenResponse, error) {
	violations := validateRefreshTokenRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.GetRefreshToken())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "verifiy token failed: %s", err)
	}

	user, err := server.store.GetUserByUsername(ctx, refreshPayload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}

	err = server.checkUserStatus(ctx, user.Blocked, user.UserRole, user.Verified)
	if err != nil {
		return nil, err
	}

	session, err := server.store.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "session not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get session failed: %s", err)
	}

	if session.UserID != user.UserID {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect session user")
	}

	if session.RefreshToken != req.RefreshToken {
		return nil, status.Errorf(codes.Unauthenticated, "incorrect session")
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, status.Errorf(codes.Unauthenticated, "expired session")
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		refreshPayload.Role,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create refresh token failed: %s", err)
	}

	rsp := &rpcs.RefreshTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: timestamppb.New(accessPayload.ExpiredAt),
	}
	return rsp, nil
}
func validateRefreshTokenRequest(req *rpcs.RefreshTokenRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetRefreshToken(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("refresh_token", err))
	}
	return violations
}
