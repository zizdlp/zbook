package gapi

import (
	"context"

	"github.com/rs/zerolog/log"

	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/models"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/storage"
	"github.com/zizdlp/zbook/util"
	"github.com/zizdlp/zbook/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetUserInfo(ctx context.Context, req *rpcs.GetUserInfoRequest) (*rpcs.GetUserInfoResponse, error) {
	apiUserDailyLimit := 10000
	apiKey := "GetUserInfo"
	authPayload, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return nil, err
	}
	violations := validateGetUserInfoRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	var user_basic_info *models.UserBasicInfo
	var user_image_info *models.UserImageInfo
	var user_count_info *models.UserCountInfo
	user, err := server.store.GetUserByUsername(ctx, req.GetUsername())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user by username failed: %s", err)
	}
	arg := db.GetUserInfoParams{
		CurUserID: authPayload.UserID,
		UserID:    user.UserID,
		Signed:    true,
		Role:      authPayload.UserRole,
	}
	user_row, err := server.store.GetUserInfo(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user info failed: %s", err)
	}
	if req.GetUserBasic() {
		user_basic_info = &models.UserBasicInfo{
			UserId:    user.UserID,
			Username:  user.Username,
			Email:     user.Email,
			Motto:     user.Motto,
			CreatedAt: timestamppb.New(user.CreatedAt),
		}
	}
	if req.GetUserImage() {
		avatarData, err := storage.DownloadFileFromStorage(server.minioClient, ctx, user.Username, "avatar")
		if err != nil {
			log.Error().Msgf("download avatar for %s failed: %s", user.Username, err)
			user_image_info = &models.UserImageInfo{
				UserId: user.UserID,
			}
		} else {
			user_image_info = &models.UserImageInfo{
				UserId: user.UserID,
				Avatar: avatarData,
			}
		}

	}

	if req.GetUserCount() {

		user_count_info = &models.UserCountInfo{
			UserId:         user.UserID,
			CountLikes:     int32(user_row.LikeCount),
			CountRepos:     int32(user_row.RepoCount),
			CountFollowing: int32(user_row.FollowingCount),
			CountFollower:  int32(user_row.FollowerCount),
			Following:      user_row.IsFollowing,
		}

	}

	rsp := &rpcs.GetUserInfoResponse{
		UserBasicInfo: user_basic_info,
		UserImageInfo: user_image_info,
		UserCountInfo: user_count_info,
	}
	return rsp, nil
}
func validateGetUserInfoRequest(req *rpcs.GetUserInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {

	if err := val.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	return violations
}
