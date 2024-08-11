package gapi

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) QueryMarkdown(ctx context.Context, req *rpcs.QueryMarkdownRequest) (*rpcs.QueryMarkdownResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "QueryMarkdown"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateQueryMarkdownRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.QueryMarkdownParams{
		Limit:          req.GetPageSize(),
		Offset:         (req.GetPageId() - 1) * req.GetPageSize(),
		PlaintoTsquery: req.GetPlainToTsquery(),
		Role:           authPayload.UserRole,
		Signed:         true,
		CurUserID:      authPayload.UserID,
	}
	markdowns, err := server.store.QueryMarkdown(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "markdown not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "query markdown failed: %s", err)
	}
	rsp := &rpcs.QueryMarkdownResponse{
		Elements: convertQueryMarkdown(markdowns),
	}
	return rsp, nil

}
func validateQueryMarkdownRequest(req *rpcs.QueryMarkdownRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPlainToTsquery(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("plain_to_tsquery", err))
	}
	return violations
}

func convertQueryMarkdown(markdowns []db.QueryMarkdownRow) []*models.Markdown {
	var ret_markdowns []*models.Markdown
	for i := 0; i < len(markdowns); i++ {
		str, ok := markdowns[i].Coalesce.(string)
		if !ok {
			if !ok {
				log.Error().Msg("cannot convert coalesce to string")
			}
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
