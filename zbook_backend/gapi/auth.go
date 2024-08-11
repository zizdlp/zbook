package gapi

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis"

	"github.com/rs/zerolog/log"
	db "github.com/zizdlp/zbook/db/sqlc"
	"github.com/zizdlp/zbook/pb/rpcs"
	"github.com/zizdlp/zbook/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) checkUserStatus(ctx context.Context, Blocked bool, UserRole string, Verified bool) error {

	if Blocked {
		return status.Errorf(codes.PermissionDenied, "user account has been blocked, please contact admin")
	}
	if server.config.REQUIRE_EMAIL_VERIFY && !Verified {
		return status.Errorf(codes.PermissionDenied, "email not verified for this account")
	}

	config, err := server.store.GetConfiguration(ctx, "allow_login")
	if err != nil {
		return status.Errorf(codes.Internal, "failed to get login configuration: %s", err)
	}
	if !config.ConfigValue && UserRole != util.AdminRole {
		return status.Errorf(codes.PermissionDenied, "login is currently disabled")
	}
	return nil
}

func (server *Server) CreateLoginPart(ctx context.Context, Username string, UserRole string, UserID int64, apiKey string) (*rpcs.LoginUserResponse, error) {

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		Username,
		UserRole,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create access token token failed: %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(
		Username,
		UserRole,
		server.config.RefreshTokenDuration,
	)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create refresh token failed: %s", err)
	}
	apiUserDailyLimit := 100
	err = server.checkUserLimit(UserID, apiKey, apiUserDailyLimit)
	if err != nil {
		return nil, err
	}
	mtdt := server.extractMetadata(ctx)
	_, err = server.store.CreateSession(ctx, db.CreateSessionParams{
		SessionID:    refreshPayload.ID,
		UserID:       UserID,
		RefreshToken: refreshToken,
		UserAgent:    mtdt.UserAgent,
		ClientIp:     mtdt.ClientIP,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "create session failed: %s", err)
	}
	rsp := &rpcs.LoginUserResponse{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		Username:              Username,
		AccessTokenExpiresAt:  timestamppb.New(accessPayload.ExpiredAt),
		RefreshTokenExpiresAt: timestamppb.New(refreshPayload.ExpiredAt),
		Role:                  UserRole,
	}
	return rsp, nil
}

func (server *Server) authUser(ctx context.Context, accessibleRoles []string, apiUserDailyLimit int, apiKey string) (*db.User, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing metadata")
	}
	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "missing authorization header")
	}
	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(fields) < 2 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid authorization header format")
	}
	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, status.Errorf(codes.Unauthenticated, "unsupported authorization type: %s", authType)
	}
	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid access token: %s", err)
	}
	user, err := server.store.GetUserByUsername(ctx, payload.Username)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user: %s", err)
	}
	err = server.checkUserStatus(ctx, user.Blocked, user.UserRole, user.Verified)
	if err != nil {
		return nil, err
	}

	if !hasPermission(payload.Role, accessibleRoles) {
		return nil, status.Errorf(codes.PermissionDenied, "permission denied")
	}
	err = server.checkUserLimit(user.UserID, apiKey, apiUserDailyLimit)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func hasPermission(userRole string, accessibleRoles []string) bool {
	for _, role := range accessibleRoles {
		if userRole == role {
			return true
		}
	}
	return false
}

type DailyUniqueKeysCount struct {
	Date  time.Time
	Count int32
}

type VisitorData struct {
	IP    string
	Agent string
	Count int
}

// 定义排序函数，按照 Count 降序排序
type ByCountDesc []*VisitorData

func (a ByCountDesc) Len() int {
	return len(a)
}
func (a ByCountDesc) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a ByCountDesc) Less(i, j int) bool {
	return a[i].Count > a[j].Count // 降序排序
}

func (server *Server) GetDailyVisitorsForLastNDays(ndays int32) ([]*VisitorData, error) {
	// 定义用于存储符合条件的访客数据的切片
	var visitors []*VisitorData
	location, err := time.LoadLocation(server.config.TIMEZONE)
	if err != nil {
		return nil, fmt.Errorf("failed to load location:%v", err)
	}
	for i := 0; i < int(ndays); i++ {
		// 计算当前日期（向前推 ndays 天）
		currentDate := time.Now().In(location).AddDate(0, 0, -i)
		today := currentDate.Format("2006-01-02")

		// 构建 Redis 键的模式，匹配过去一天的所有键
		pattern := fmt.Sprint("logvisitor:*:*:*", today)
		// 获取符合模式的所有 Redis 键
		keys, err := server.redisClient.Keys(pattern).Result()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "error fetching keys from Redis: %s", err)
		}

		// 遍历符合模式的所有键
		for _, key := range keys {
			// 获取键对应的值，即访问次数
			countStr, err := server.redisClient.Get(key).Result()
			if err != nil {
				return nil, status.Errorf(codes.Internal, "error fetching value from Redis: %s", err)
			}
			count, err := strconv.Atoi(countStr)
			if err != nil {
				return nil, status.Errorf(codes.Internal, "error fetching value from Redis: %s", err)
			}
			// 将键解析为 IP 和 UserAgent
			parts := strings.Split(key, ":")
			// 获取 IP 地址部分
			ipParts := strings.Split(parts[1], ",")[0]
			// 构建 VisitorData 结构体并添加到切片中
			visitors = append(visitors, &VisitorData{
				IP:    ipParts,
				Agent: parts[len(parts)-2],
				Count: count,
			})
		}
	}
	sort.Sort(ByCountDesc(visitors))
	return visitors, nil
}

func (server *Server) GetUniqueKeysCountForLastNDays(ndays int32, timezone string) ([]DailyUniqueKeysCount, error) {
	var dailyCounts []DailyUniqueKeysCount

	// 加载时区
	location, err := time.LoadLocation(server.config.TIMEZONE)
	if err != nil {
		return nil, fmt.Errorf("failed to load location:%v", err)
	}
	// 遍历过去 ndays 天内的每一天
	for i := 0; i < int(ndays); i++ {
		// 计算指定时区的当前日期（向前推 ndays 天）

		currentDate := time.Now().In(location).AddDate(0, 0, -i)
		today := currentDate.Format("2006-01-02")

		// 构建 Redis 键的模式，用于匹配该天的所有键
		pattern := fmt.Sprintf("logvisitor:*:%s", today)
		// 获取匹配的键列表
		keys, err := server.redisClient.Keys(pattern).Result()
		if err != nil {
			return nil, err
		}

		// 计算集合的大小，即为该天的唯一键数量，并存储到数组中
		uniqueKeysCount := int32(len(keys))
		dailyCounts = append(dailyCounts, DailyUniqueKeysCount{Date: currentDate, Count: uniqueKeysCount})
	}
	return dailyCounts, nil
}

func (server *Server) insertRedisKey(redisKey string, dailyLimit int, duration time.Duration) error {
	// 获取当前用户的插入次数
	countStr, err := server.redisClient.Get(redisKey).Result()
	if err == redis.Nil {
		// 如果 Redis 键不存在，则说明用户还没有插入过数据，插入次数初始化为 0
		countStr = "0"
	} else if err != nil {
		return status.Errorf(codes.Internal, "failed to retrieve count from Redis: %s", err)
	}

	// 将插入次数字符串转换为整数
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to convert count string '%s' to integer: %s", countStr, err)
	}

	// 检查插入次数是否已经达到限制
	if count >= dailyLimit {
		return status.Errorf(codes.ResourceExhausted, "user has exceeded the daily limit of %d inserts", dailyLimit)
	}

	// 更新 Redis 中用户的插入次数
	count++
	err = server.redisClient.Set(redisKey, strconv.Itoa(count), duration*time.Hour).Err()
	if err != nil {
		return status.Errorf(codes.Internal, "failed to update count in Redis: %s", err)
	}
	return nil
}

func (server *Server) LogRedisVisitor(ctx context.Context) (err error) {
	mtdt := server.extractMetadata(ctx)
	UserAgent := mtdt.UserAgent
	ClientIp := mtdt.ClientIP

	if strings.HasPrefix(ClientIp, "::1") {
		ClientIp = "127.0.0.1" + strings.TrimPrefix(ClientIp, "::1")
	}
	// 分隔ClientIp并取第一个部分
	if ClientIp == "" {
		ClientIp = "unknown"
	} else {
		// 分隔ClientIp并取第一个部分
		ClientIpParts := strings.Split(ClientIp, ",")
		ClientIp = ClientIpParts[0]
	}
	location, err := time.LoadLocation(server.config.TIMEZONE)
	if err != nil {
		return fmt.Errorf("failed to load location:%v", err)
	}
	today := time.Now().In(location).Format("2006-01-02")
	redisKey := fmt.Sprintf("%s:%s:%s:%s", "logvisitor", ClientIp, UserAgent, today)
	log.Info().Msgf("inser to redis: %s", redisKey)
	return server.insertRedisKey(redisKey, 1000000, 24*31) //保留31天，设置较大使用限制
}

func (server *Server) checkUserLimit(userID int64, keytype string, dailyLimit int) error {

	// 获取当前日期字符串，用于构建 Redis 键和设置过期时间
	location, err := time.LoadLocation(server.config.TIMEZONE)
	if err != nil {
		return fmt.Errorf("failed to load location:%v", err)
	}
	today := time.Now().In(location).Format("2006-01-02")
	userIDStr := strconv.FormatInt(userID, 10)
	redisKey := fmt.Sprintf("userlimit:%s:%s:%s", keytype, userIDStr, today)
	return server.insertRedisKey(redisKey, dailyLimit, 24) //保留24h
}

// 1,表示等同于写的用户，2表示其他可见用户，0表示无权限
func (server *Server) getUserPermessionlevel(ctx context.Context, authUsername string, username string) (int, error) {
	if username == authUsername {
		return 1, nil
	}
	checkUser, err := server.store.GetUserByUsername(ctx, authUsername)
	if err != nil {
		return 0, err
	}
	if checkUser.UserRole == util.AdminRole {
		return 1, nil
	}
	datauser, err := server.store.GetUserByUsername(ctx, username)
	if err != nil {
		return 0, err
	}
	// TODO add verified
	if !datauser.Blocked {
		return 2, nil
	}
	return 0, status.Errorf(codes.PermissionDenied, "current account do not have enough permission")
}

func (server *Server) isRepoVisibleToCurrentUser(ctx context.Context, RepoID int64) error {
	repoInfo, err := server.store.GetRepoPermission(ctx, RepoID)
	if err != nil {
		return status.Errorf(codes.Internal, "get repo failed: %s", err)
	}

	if repoInfo.VisibilityLevel == util.VisibilityPublic {
		return nil
	}

	apiUserDailyLimit := 100000
	apiKey := "isRepoVisibleToCurrentUser"
	authUser, err := server.authUser(ctx, []string{util.AdminRole, util.UserRole}, apiUserDailyLimit, apiKey)
	if err != nil {
		return err
	}
	if repoInfo.VisibilityLevel == util.VisibilitySigned {
		return nil
	}

	if repoInfo.VisibilityLevel == util.VisibilityChosed {

		arg := db.GetRepoRelationParams{UserID: authUser.UserID, RepoID: RepoID, RelationType: util.RelationTypeVisi}
		_, err := server.store.GetRepoRelation(ctx, arg)
		if err != nil && !(authUser.UserRole == util.AdminRole || authUser.UserID == repoInfo.UserID) {
			return status.Errorf(codes.PermissionDenied, "current account can not visit this repo")
		}
		return nil
	} else {
		if authUser.UserRole == util.AdminRole {
			return nil
		} else if authUser.UserID == repoInfo.UserID {
			return nil
		} else {
			return status.Errorf(codes.PermissionDenied, "current account can not visit this repo")
		}
	}
}

func (server *Server) isMarkdownVisibleToCurrentUser(ctx context.Context, MarkdownID int64) error {

	RepoID, err := server.store.GetMarkdownRepoID(ctx, MarkdownID)
	if err != nil {
		return status.Errorf(codes.Internal, "get repo failed: %s", err)
	}
	return server.isRepoVisibleToCurrentUser(ctx, RepoID)
}
func (server *Server) isCommentVisibleToCurrentUser(ctx context.Context, CommentID int64) error {
	comment, err := server.store.GetCommentBasicInfo(ctx, CommentID)
	if err != nil {
		return status.Errorf(codes.Internal, "get comment failed: %s", err)
	}
	RepoID, err := server.store.GetMarkdownRepoID(ctx, comment.MarkdownID)
	if err != nil {
		return status.Errorf(codes.Internal, "get repo failed: %s", err)
	}
	return server.isRepoVisibleToCurrentUser(ctx, RepoID)
}
