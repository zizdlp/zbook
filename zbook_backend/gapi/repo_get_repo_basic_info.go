package gapi

import (
	"context"
	"errors"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetRepoBasicInfo(ctx context.Context, req *rpcs.GetRepoBasicInfoRequest) (*rpcs.GetRepoBasicInfoResponse, error) {
	violations := validateGetRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err := server.isRepoVisibleToCurrentUser(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}

	repo, err := server.store.GetRepoBasicInfo(ctx, req.GetRepoId())
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "repo not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "get repo basic info failed: %s", err)
	}
	avatarData, err := storage.DownloadAvatar(server.minioClient, context.Background(), repo.Username, "avatar")
	if err != nil {
		log.Info().Msgf("download avatar for %s failed: %s", repo.Username, err)
	}
	rsp := &rpcs.GetRepoBasicInfoResponse{
		RepoId:          repo.RepoID,
		Username:        repo.Username,
		Email:           repo.Email,
		UpdatedAt:       timestamppb.New(repo.UpdatedAt),
		Avatar:          avatarData,
		RepoName:        repo.RepoName,
		RepoDescription: repo.RepoDescription,
		HomePage:        repo.HomePage,
	}
	return rsp, nil
}
func validateGetRepoRequest(req *rpcs.GetRepoBasicInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	return violations
}
