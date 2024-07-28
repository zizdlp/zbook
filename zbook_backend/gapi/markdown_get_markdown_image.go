package gapi

import (
	"context"
	"errors"
	"strconv"

	db "github.com/zizdlp/zbook/db/sqlc"
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

	arg := db.GetRepoBasicInfoParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo, err := server.store.GetRepoBasicInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo basic info failed: %s", err)
	}
	err = server.isRepoVisibleToCurrentUser(ctx, repo.RepoID)
	if err != nil {
		return nil, err
	}

	path := strconv.FormatInt(repo.UserID, 10) + "/" + strconv.FormatInt(repo.RepoID, 10) + "/" + req.GetFilePath()
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
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	return violations
}
