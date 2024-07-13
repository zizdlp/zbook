package gapi

import (
	"context"
	"os"
	"strconv"
	"strings"

	"github.com/zizdlp/zbook/pb/rpcs"
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
	path := "/tmp/wiki/" + strconv.FormatInt(req.GetRepoId(), 10) + "/" + req.GetFilePath()

	// Read the image file
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "file not exist: %s", err)
	}
	ext := strings.ToLower(path)
	if strings.HasSuffix(ext, ".png") || strings.HasSuffix(ext, ".jpg") || strings.HasSuffix(ext, ".jpeg") || strings.HasSuffix(ext, ".webp") {
		base64, err := util.ReadImageBytes(path)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to read image to base64: %v", err)
		}
		data, err = util.CompressImage(base64)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to compress image: %v", err)
		}
	}

	// Create the response with the (potentially compressed) image data
	rsp := &rpcs.GetMarkdownImageResponse{
		File: data,
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
