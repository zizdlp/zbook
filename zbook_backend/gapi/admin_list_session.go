package gapi

import (
	"context"

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

func (server *Server) ListSession(ctx context.Context, req *rpcs.ListSessionRequest) (*rpcs.ListSessionResponse, error) {
	violations := validateListSessionRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	apiUserDailyLimit := 10000
	apiKey := "ListSession"
	_, err := server.authUser(ctx, []string{util.AdminRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	if req.GetQuery() != "" {
		arg := db.QuerySessionParams{
			Limit:  req.GetPageSize(),
			Offset: (req.GetPageId() - 1) * req.GetPageSize(),
			Query:  req.GetQuery(),
		}
		sessions, err := server.store.QuerySession(ctx, arg)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "query active session failed: %s", err)
		}
		rsp := &rpcs.ListSessionResponse{
			Elements: convertListSessionQuery(sessions),
		}
		return rsp, nil
	}

	arg := db.ListSessionParams{
		Limit:  req.GetPageSize(),
		Offset: (req.GetPageId() - 1) * req.GetPageSize(),
	}

	sessions, err := server.store.ListSession(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list active session failed: %s", err)
	}
	rsp := &rpcs.ListSessionResponse{
		Elements: convertListSession(sessions),
	}
	return rsp, nil
}
func validateListSessionRequest(req *rpcs.ListSessionRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateInt32ID(req.GetPageId()); err != nil {
		violations = append(violations, fieldViolation("page_id", err))
	}
	if err := val.ValidatePageSize(req.GetPageSize()); err != nil {
		violations = append(violations, fieldViolation("page_size", err))
	}
	return violations
}

func convertListSession(sessions []db.ListSessionRow) []*models.SessionInfo {
	var ret_sessions []*models.SessionInfo
	for i := 0; i < len(sessions); i++ {
		ret_sessions = append(ret_sessions,
			&models.SessionInfo{
				UserAgent: sessions[i].UserAgent,
				Email:     sessions[i].Email,
				ClientIp:  sessions[i].ClientIp,
				Username:  sessions[i].Username,
				ExpiresAt: timestamppb.New(sessions[i].ExpiresAt),
				CreatedAt: timestamppb.New(sessions[i].CreatedAt),
			},
		)
	}
	return ret_sessions
}

func convertListSessionQuery(sessions []db.QuerySessionRow) []*models.SessionInfo {
	var ret_sessions []*models.SessionInfo
	for i := 0; i < len(sessions); i++ {
		ret_sessions = append(ret_sessions,
			&models.SessionInfo{
				Email:     sessions[i].Email,
				UserAgent: sessions[i].UserAgent,
				ClientIp:  sessions[i].ClientIp,
				Username:  sessions[i].Username,
				ExpiresAt: timestamppb.New(sessions[i].ExpiresAt),
				CreatedAt: timestamppb.New(sessions[i].CreatedAt),
			},
		)
	}
	return ret_sessions
}
