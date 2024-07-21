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

func (server *Server) QueryUserMarkdown(ctx context.Context, req *rpcs.QueryUserMarkdownRequest) (*rpcs.QueryUserMarkdownResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "QueryUserMarkdown"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateQueryUserMarkdownRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	permissionLevel, err := server.getUserPermessionlevel(ctx, authPayload.Username, req.GetUsername())
	if err != nil {
		return nil, err
	}
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	if permissionLevel == 2 { //其他用户
		arg := db.QueryUserVisibleMarkdownParams{
			Limit:          req.GetPageSize(),
			Offset:         (req.GetPageId() - 1) * req.GetPageSize(),
			PlaintoTsquery: req.GetPlainToTsquery(),
			UserID:         user.UserID,
		}
		markdowns, err := server.store.QueryUserVisibleMarkdown(ctx, arg)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return nil, status.Errorf(codes.NotFound, "markdown not found: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "query user visible markdown failed: %s", err)
		}

		rsp := &rpcs.QueryUserMarkdownResponse{
			Elements: convertQueryUserVisibleMarkdown(markdowns),
		}
		return rsp, nil
	} else {
		arg := db.QueryUserAllMarkdownParams{
			PlaintoTsquery: req.GetPlainToTsquery(),
			UserID:         user.UserID,
			Limit:          req.GetPageSize(),
			Offset:         (req.GetPageId() - 1) * req.GetPageSize(),
		}
		markdowns, err := server.store.QueryUserAllMarkdown(ctx, arg)
		if err != nil {
			if errors.Is(err, db.ErrRecordNotFound) {
				return nil, status.Errorf(codes.NotFound, "markdown not found: %s", err)
			}
			return nil, status.Errorf(codes.Internal, "query user all markdown failed: %s", err)
		}

		rsp := &rpcs.QueryUserMarkdownResponse{
			Elements: convertQueryUserAllMarkdown(markdowns),
		}
		return rsp, nil
	}

}
func validateQueryUserMarkdownRequest(req *rpcs.QueryUserMarkdownRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateString(req.GetPlainToTsquery(), 1, 512); err != nil {
		violations = append(violations, fieldViolation("plain_to_tsquery", err))
	}
	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}
	return violations
}

func convertQueryUserVisibleMarkdown(markdowns []db.QueryUserVisibleMarkdownRow) []*models.Markdown {
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
func convertQueryUserAllMarkdown(markdowns []db.QueryUserAllMarkdownRow) []*models.Markdown {
	var ret_markdowns []*models.Markdown
	for i := 0; i < len(markdowns); i++ {
		str, ok := markdowns[i].Coalesce.(string)
		if !ok {
			log.Warn().Msg("cannot convert coalesce to string")
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
