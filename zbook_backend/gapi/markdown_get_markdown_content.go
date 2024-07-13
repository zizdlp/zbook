package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetMarkdownContent(ctx context.Context, req *rpcs.GetMarkdownContentRequest) (*rpcs.GetMarkdownContentResponse, error) {
	violations := validateGetMarkdownContentRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err := server.isRepoVisibleToCurrentUser(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}
	arg := db.GetMarkdownContentParams{
		RelativePath: req.GetRelativePath(),
		RepoID:       req.GetRepoId(),
	}
	markdown, err := server.store.GetMarkdownContent(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "GetMarkdownContent not found error: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "GetMarkdownContent error : %s", err)
	}
	rsp := &rpcs.GetMarkdownContentResponse{
		Markdown: convertMarkdown(markdown),
	}
	return rsp, nil
}
func validateGetMarkdownContentRequest(req *rpcs.GetMarkdownContentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	if err := val.ValidateString(req.GetRelativePath(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("relative_path", err))
	}
	return violations
}
func convertMarkdown(markdown db.Markdown) *models.Markdown {
	return &models.Markdown{
		MarkdownId:   markdown.MarkdownID,
		RelativePath: markdown.RelativePath,
		UserId:       markdown.UserID,
		RepoId:       markdown.RepoID,
		MainContent:  markdown.MainContent,
		TableContent: markdown.TableContent,
		Md5:          markdown.Md5,
		VersionKey:   markdown.VersionKey,
		CreatedAt:    timestamppb.New(markdown.CreatedAt),
	}
}
