package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// NewRandomStringGenerator creates a new RandomStringGenerator
func NewRandomStringGenerator() *RandomStringGenerator {
	src := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(src)
	return &RandomStringGenerator{rng: rng}
}

// RandomStringGenerator struct to hold the random generator
type RandomStringGenerator struct {
	rng *rand.Rand
}

// RandomString generates a random string of length n
func (rsg *RandomStringGenerator) RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rsg.rng.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomInt generates a random integer between min and max
func RandomInt32(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomInt generates a random integer between min and max
func RandomInts(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random username
func RandomUsername() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}
func RandomPGBool() pgtype.Bool {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的0或1
	value := randGen.Intn(2)
	// 将随机数转换为布尔值
	return pgtype.Bool{Bool: value == 1, Valid: true}
}

func RandomBool() bool {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的0或1
	value := randGen.Intn(2)
	// 将随机数转换为布尔值
	return value == 1
}

func RandomUserRole() string {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的索引值，0或1
	index := randGen.Intn(2)
	// 根据索引值选择用户角色
	var role string
	switch index {
	case 0:
		role = AdminRole
	case 1:
		role = UserRole
	}
	return role
}

func RandomOAuth() string {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的索引值，0或1
	index := randGen.Intn(3)
	// 根据索引值选择用户角色
	var role string
	switch index {
	case 0:
		role = OAuthTypeGithub
	case 1:
		role = OAuthTypeGoogle
	case 2:
		role = OAuthTypeWechat
	}
	return role
}

func RandomVerificationType() string {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的索引值，0或1
	index := randGen.Intn(2)
	// 根据索引值选择用户角色
	var role string
	switch index {
	case 0:
		role = VerifyTypeResetPassword
	case 1:
		role = VerifyTypeVerifyEmail
	}
	return role
}
func RandomRepoVisibility() string {
	// 创建一个新的本地随机生成器
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成一个随机的索引值，0或1
	index := randGen.Intn(4)
	// 根据索引值选择用户角色
	var role string
	switch index {
	case 0:
		role = VisibilityChosed
	case 1:
		role = VisibilityPrivate
	case 2:
		role = VisibilityPublic
	case 3:
		role = VisibilityPublic
	}
	return role
}
