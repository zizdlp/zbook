package gapi

import (
	"context"

	"github.com/rs/zerolog/log"
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

func (server *Server) ListRepo(ctx context.Context, req *rpcs.ListRepoRequest) (*rpcs.ListRepoResponse, error) {
	violations := validateListRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListRepo"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		// not signed
		if req.GetQuery() != "" {

			arg := db.QueryRepoParams{
				Limit:     req.GetPageSize(),
				Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
				Query:     req.GetQuery(),
				Role:      util.UserRole,
				Signed:    false,
				CurUserID: 0,
			}
			reports, err := server.store.QueryRepo(ctx, arg)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "query repo failed: %s", err)
			}
			rsp := &rpcs.ListRepoResponse{
				Elements: convertQueryRepo(reports, req.GetLang()),
			}
			return rsp, nil
		}
		arg := db.ListRepoParams{
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Signed:    false,
			CurUserID: 0,
			Role:      util.UserRole,
		}

		reports, err := server.store.ListRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "list repo failed: %s", err)
		}

		rsp := &rpcs.ListRepoResponse{
			Elements: convertListRepos(reports, req.GetLang()),
		}
		return rsp, nil
	}

	if req.GetQuery() != "" {

		arg := db.QueryRepoParams{
			Limit:     req.GetPageSize(),
			Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
			Query:     req.GetQuery(),
			Role:      authUser.UserRole,
			Signed:    true,
			CurUserID: authUser.UserID,
		}
		reports, err := server.store.QueryRepo(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query repo failed: %s", err)
		}
		rsp := &rpcs.ListRepoResponse{
			Elements: convertQueryRepo(reports, req.GetLang()),
		}
		return rsp, nil
	}
	arg := db.ListRepoParams{
		Limit:     req.GetPageSize(),
		Offset:    (req.GetPageId() - 1) * req.GetPageSize(),
		Signed:    true,
		CurUserID: authUser.UserID,
		Role:      util.AdminRole,
	}

	reports, err := server.store.ListRepo(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list repo failed: %s", err)
	}
	rsp := &rpcs.ListRepoResponse{
		Elements: convertListRepos(reports, req.GetLang()),
	}
	return rsp, nil
}
func validateListRepoRequest(req *rpcs.ListRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	if err := val.ValidateLang(req.GetLang()); err != nil {
		violations = append(violations, fieldViolation("lang", err))
	}
	return violations
}

func convertListRepos(reports []db.ListRepoRow, lang string) []*models.ListRepoInfo {
	var ret_reports []*models.ListRepoInfo
	for i := 0; i < len(reports); i++ {
		path := ""
		path, err := util.GetDocumentPath(reports[i].Home, lang)
		if err != nil {
			log.Error().Err(err).Msgf("failed to find home for:%s", lang)
		}
		ret_reports = append(ret_reports,
			&models.ListRepoInfo{
				RepoId:          reports[i].RepoID,
				RepoName:        reports[i].RepoName,
				Username:        reports[i].Username,
				RepoDescription: reports[i].RepoDescription,
				VisibilityLevel: reports[i].VisibilityLevel,
				GitHost:         reports[i].GitHost,
				LikeCount:       int32(reports[i].LikeCount),
				IsLiked:         reports[i].IsLiked,
				UpdatedAt:       timestamppb.New(reports[i].UpdatedAt),
				CreatedAt:       timestamppb.New(reports[i].CreatedAt),
				Home:            path,
			},
		)
	}
	return ret_reports
}

func convertQueryRepo(reports []db.QueryRepoRow, lang string) []*models.ListRepoInfo {
	var ret_reports []*models.ListRepoInfo
	for i := 0; i < len(reports); i++ {
		path := ""
		path, err := util.GetDocumentPath(reports[i].Home, lang)
		if err != nil {
			log.Error().Err(err).Msgf("failed to find home for:%s", lang)
		}
		ret_reports = append(ret_reports,
			&models.ListRepoInfo{
				RepoId:          reports[i].RepoID,
				Username:        reports[i].Username,
				RepoName:        reports[i].RepoName,
				RepoDescription: reports[i].RepoDescription,
				VisibilityLevel: reports[i].VisibilityLevel,
				GitHost:         reports[i].GitHost,
				// LikeCount:       int32(reports[i].LikeCount),
				UpdatedAt: timestamppb.New(reports[i].UpdatedAt),
				CreatedAt: timestamppb.New(reports[i].CreatedAt),
				Home:      path,
			},
		)
	}
	return ret_reports
}
