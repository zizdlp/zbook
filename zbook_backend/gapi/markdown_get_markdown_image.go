package gapi

import (
	"context"
	"strconv"

	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetMarkdownImage(ctx context.Context, req *rpcs.GetMarkdownImageRequest) (*rpcs.GetMarkdownImageResponse, error) {
	violations := validateGetMarkdownImageRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err := server.isRepoVisibleToCurrentUser(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}
	repo, err := server.store.GetRepoBasicInfo(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}
	path := strconv.FormatInt(repo.UserID, 10) + "/" + strconv.FormatInt(req.GetRepoId(), 10) + "/" + req.GetFilePath()
	path = util.NormalizePath(path)
	avatarData, err := storage.DownloadFileFromStorage(server.minioClient, ctx, path, "git-files")
	if err != nil {
		return nil, status.Errorf(codes.Internal, "GetMarkdownFile failed: %s", err)
	}
	// Create the response with the (potentially compressed) image data
	rsp := &rpcs.GetMarkdownImageResponse{
		File: avatarData,
	}
	return rsp, nil
}
func validateGetMarkdownImageRequest(req *rpcs.GetMarkdownImageRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetFilePath(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("file_path", err))
	}
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	return violations
}
