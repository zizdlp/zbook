package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
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
	arg_get := db.GetRepoBasicInfoParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo, err := server.store.GetRepoBasicInfo(ctx, arg_get)
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
	arg := db.GetMarkdownContentParams{
		RelativePath: req.GetRelativePath(),
		RepoID:       repo.RepoID,
	}
	markdown, err := server.store.GetMarkdownContent(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "GetMarkdownContent not found error: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "GetMarkdownContent error : %s", err)
	}

	arg_config := db.GetRepoConfigParams{Username: req.GetUsername(), RepoName: req.GetRepoName()}
	repo_config, err := server.store.GetRepoConfig(ctx, arg_config)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "get repo config not found error: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "GetRepoConfig error : %s", err)
	}
	config, err := util.ParseRepoConfigFromString(repo_config.Config)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ParseRepoConfigFromString error : %s", err)
	}
	prev, next, err := config.FindAdjacentPaths(req.GetRelativePath())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "FindAdjacentPaths error : %s", err)
	}

	rsp := &rpcs.GetMarkdownContentResponse{
		Markdown:   convertMarkdown(markdown),
		Prev:       prev,
		Next:       next,
		Footers:    convertFooters(config.FooterSocials),
		UpdatedAt:  timestamppb.New(markdown.UpdatedAt),
		ThemeColor: repo.ThemeColor,
	}
	return rsp, nil
}
func validateGetMarkdownContentRequest(req *rpcs.GetMarkdownContentRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateUsername(req.GetUsername())
	if err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	err = val.ValidateRepoName(req.GetRepoName())
	if err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
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

		CreatedAt: timestamppb.New(markdown.CreatedAt),
	}
}

func convertFooters(footers []util.FooterSocial) []*models.FooterSocial {

	var ret_footers []*models.FooterSocial
	for i := 0; i < len(footers); i++ {
		ret_footers = append(ret_footers,
			&models.FooterSocial{
				Name: footers[i].Name,
				Icon: footers[i].Icon,
				Url:  footers[i].URL,
			},
		)
	}
	return ret_footers

}
