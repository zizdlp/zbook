package util

import (
	"fmt"
	"regexp"

	"github.com/google/uuid"
)

// UUIDToString 将UUID编码为只包含数字和小写字母的字符串
func UUIDToString(id uuid.UUID) string {
	// 将UUID转换为字符串
	uuidStr := id.String()
	// 将UUID字符串中的"-"去除
	uuidStr = uuidStr[0:8] + uuidStr[9:13] + uuidStr[14:18] + uuidStr[19:23] + uuidStr[24:]
	return uuidStr

}

// StringToUUID 将只包含数字和小写字母的字符串解码为UUID
func StringToUUID(s string) (uuid.UUID, error) {
	// 使用正则表达式检查字符串是否只包含小写字母和数字
	match, err := regexp.MatchString("^[a-z0-9]+$", s)
	if err != nil {
		return uuid.UUID{}, err
	}
	if !match {
		return uuid.UUID{}, fmt.Errorf("invalid characters in string: %s", s)
	}
	// 添加"-"符号，恢复原始的UUID字符串格式
	uuidStr := s[:8] + "-" + s[8:12] + "-" + s[12:16] + "-" + s[16:20] + "-" + s[20:]
	// 解析UUID字符串
	parsedUUID, err := uuid.Parse(uuidStr)
	if err != nil {
		return uuid.UUID{}, err
	}
	return parsedUUID, nil
}
