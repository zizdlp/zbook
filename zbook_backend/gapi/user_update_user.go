package gapi

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *rpcs.UpdateUserRequest) (*rpcs.UpdateUserResponse, error) {
	apiUserDailyLimit := 1000
	apiKey := "UpdateUser"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}

	hashedPassword := ""
	if len(req.GetPassword()) != 0 {
		hashedPassword, err = util.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
		}
	}

	arg := db.UpdateUserBasicInfoParams{
		Username: authPayload.Username,
		Motto: pgtype.Text{
			String: req.GetMotto(),
			Valid:  len(req.GetMotto()) != 0,
		},
		HashedPassword: pgtype.Text{
			String: hashedPassword,
			Valid:  len(req.GetPassword()) != 0,
		},
	}

	if len(req.GetAvatar()) != 0 {
		avatar, err := util.CompressImage(req.GetAvatar())
		if err != nil {
			log.Info().Msgf("compress image for %s failed: %s", authPayload.Username, err)
		}

		err = storage.UploadFileToStorage(server.minioClient, context.Background(), authPayload.Username, "avatar", avatar)
		if err != nil {
			return nil, status.Errorf(codes.NotFound, "upload avatar failed: %s", err)
		}
	}

	_, err = server.store.UpdateUserBasicInfo(ctx, arg)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "user not found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "fail to Update user: %s", err)
	}

	rsp := &rpcs.UpdateUserResponse{}
	return rsp, nil
}
