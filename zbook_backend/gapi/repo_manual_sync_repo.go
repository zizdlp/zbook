package gapi

import (
	"context"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) ManualSyncRepo(ctx context.Context, req *rpcs.ManualSyncRepoRequest) (*rpcs.ManualSyncRepoResponse, error) {
	violations := validateManualSyncRepoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	err := server.isRepoVisibleToCurrentUser(ctx, req.GetRepoId())
	if err != nil {
		return nil, err
	}

	arg := db.ManualSyncRepoTxParams{
		RepoID: req.GetRepoId(),
		AfterCreate: func(cloneDir string, repoID int64, userID int64, addedFiles []string, modifiedFiles []string, deletedFiles []string) error {
			return storage.ConvertFile2Storage(server.minioClient, cloneDir, repoID, userID, addedFiles, modifiedFiles, deletedFiles)
		},
	}

	err = server.store.ManualSyncRepoTx(ctx, arg)
	if err != nil {
		return nil, err
	}

	rsp := &rpcs.ManualSyncRepoResponse{}

	return rsp, nil
}
func validateManualSyncRepoRequest(req *rpcs.ManualSyncRepoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	err := val.ValidateID(req.GetRepoId())
	if err != nil {
		violations = append(violations, fieldViolation("repo_id", err))
	}
	return violations
}
