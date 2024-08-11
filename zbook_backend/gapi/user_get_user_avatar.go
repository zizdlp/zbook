package gapi

import (
	"context"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUserAvatar(ctx context.Context, req *rpcs.GetUserAvatarRequest) (*rpcs.GetUserAvatarResponse, error) {
	violations := validateGetUserAvatarRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	avatarData, err := storage.DownloadFileFromStorage(server.minioClient, ctx, req.GetUsername(), "avatar")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "download avatar failed: %s", err)
	}

	rsp := &rpcs.GetUserAvatarResponse{
		Avatar: avatarData,
	}
	return rsp, nil
}

func validateGetUserAvatarRequest(req *rpcs.GetUserAvatarRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}
