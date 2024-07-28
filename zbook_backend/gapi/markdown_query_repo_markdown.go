package gapi

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) QueryRepoMarkdown(ctx context.Context, req *rpcs.QueryRepoMarkdownRequest) (*rpcs.QueryRepoMarkdownResponse, error) {
	violations := validateQueryRepoMarkdownRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg_get := db.GetRepoByRepoNameParams{
		Username: req.GetUsername(),
		RepoName: req.GetRepoName(),
	}
	repo, err := server.store.GetRepoByRepoName(ctx, arg_get)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo failed : %s", err)
	}
	err = server.isRepoVisibleToCurrentUser(ctx, repo.RepoID)
	if err != nil {
		return nil, err
	}

	arg := db.QueryRepoMarkdownParams{
		Limit:          req.GetPageSize(),
		Offset:         (req.GetPageId() - 1) * req.GetPageSize(),
		PlaintoTsquery: req.GetPlainToTsquery(),
		RepoID:         repo.RepoID,
		UserID:         repo.UserID,
	}
	markdowns, err := server.store.QueryRepoMarkdown(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "markdown not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "query repo markdown failed: %s", err)
	}

	rsp := &rpcs.QueryRepoMarkdownResponse{
		Elements: convertQueryRepoMarkdown(markdowns),
	}
	return rsp, nil
}
func validateQueryRepoMarkdownRequest(req *rpcs.QueryRepoMarkdownRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPlainToTsquery(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("plain_to_tsquery", err))
	}
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	if err := val.ValidateRepoName(req.GetRepoName()); err != nil {
		violations = append(violations, fieldViolation("repo_name", err))
	}
	return violations
}
func convertQueryRepoMarkdown(markdowns []db.QueryRepoMarkdownRow) []*models.Markdown {
	var ret_markdowns []*models.Markdown
	for i := 0; i < len(markdowns); i++ {
		str, ok := markdowns[i].Coalesce.(string)
		if !ok {
			log.Error().Msg("cannot convert coalesce to string")
		}
		ret_markdowns = append(ret_markdowns,
			&models.Markdown{
				MarkdownId:   markdowns[i].MarkdownID,
				RelativePath: markdowns[i].RelativePath,
				UserId:       markdowns[i].UserID,
				RepoId:       markdowns[i].RepoID,
				MainContent:  str,
				Username:     markdowns[i].Username,
				RepoName:     markdowns[i].RepoName,
			},
		)
	}
	return ret_markdowns
}
