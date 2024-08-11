package gapi

import (
	"context"
	"errors"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) QueryUser(ctx context.Context, req *rpcs.QueryUserRequest) (*rpcs.QueryUserResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "QueryUser"
	_, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateQueryUserRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	arg := db.QueryUserParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
		Query:  req.GetQuery(),
	}

	users, err := server.store.QueryUser(ctx, arg)

	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "query user failed: %s", err)
	}

	rsp := &rpcs.QueryUserResponse{
		Elements: convertQueryUser(users),
	}
	return rsp, nil

}
func validateQueryUserRequest(req *rpcs.QueryUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	if err := val.ValidateString(req.GetQuery(), 1, 64); err != nil {
		violations = append(violations, fieldViolation("query", err))
	}
	return violations
}
