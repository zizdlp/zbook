package util

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestUUIDConvert(t *testing.T) {
	uuid := uuid.New()
	uuid_str := UUIDToString(uuid)
	uuid_recover, err := StringToUUID(uuid_str)
	require.NoError(t, err)
	require.Equal(t, uuid, uuid_recover)
	valid := isValidUUIDString(uuid_str)
	require.True(t, valid, "UUID string contains invalid characters")
}

// isValidUUIDString 检查字符串是否只包含小写字母和数字
func isValidUUIDString(s string) bool {
	// 使用正则表达式匹配小写字母和数字
	match, err := regexp.MatchString("^[a-z0-9]+$", s)
	if err != nil {
		return false
	}
	return match
}
