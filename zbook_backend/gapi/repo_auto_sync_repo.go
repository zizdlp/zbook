package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) AutoSyncRepo(ctx context.Context, req *rpcs.AutoSyncRepoRequest) (*rpcs.AutoSyncRepoResponse, error) {
	violations := validateAutoSyncRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	repo, err := server.store.GetRepo(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}
	if !repo.SyncToken.Valid {
		return nil, status.Errorf(codes.PermissionDenied, "this repo not set auto sync token")
	}
	if repo.SyncToken.String != req.GetSyncToken() {
		return nil, status.Errorf(codes.InvalidArgument, "invalied sync token")
	}

	arg := db.ManualSyncRepoTxParams{
		RepoID: req.GetRepoId(),
		AfterCreate: func(cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error {
			return storage.ConvertFile2Storage(server.minioClient, cloneDir, repoID, userID, addedFiles, modifiedFiles, deletedFiles)
		},
	}

	err = server.store.ManualSyncRepoTx(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "manual sync repo failed: %s", err)
	}

	rsp := &rpcs.AutoSyncRepoResponse{
		RepoId: repo.RepoID,
	}
	return rsp, nil
}
func validateAutoSyncRepoRequest(req *rpcs.AutoSyncRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	if err := val.ValidateString(req.GetSyncToken(), 6, 32); err != nil {
		violations = append(violations, fieldViolation("sync_token", err))
	}
	return violations
}
