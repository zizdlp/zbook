package util

import (
	"regexp"
	"strings"
)

// FilterDiffFilesByExtensions 过滤指定后缀名的文件
func FilterDiffFilesByExtensions(files []string, allowedExtensions map[string]bool) []string {
	var filteredFiles []string

	for _, file := range files {
		ext := getFileExtension(file)
		if allowedExtensions[ext] {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles
}

// getFileExtension 使用正则表达式获取文件的后缀名
func getFileExtension(fileName string) string {
	re := regexp.MustCompile(`\.[^\.]+$`)
	match := re.FindString(fileName)
	if match != "" {
		return match
	}
	return ""
}

func ExtractLogDetails(logKey string) (ip, userAgent, date string) {
	// 移除前缀 "logvisitor:"
	keyWithoutPrefix := strings.TrimPrefix(logKey, "logvisitor:")

	// 查找最后一个冒号的位置，用来确定日期和 UserAgent 的分隔点
	lastColonIndex := strings.LastIndex(keyWithoutPrefix, ":")

	// 提取日期
	date = keyWithoutPrefix[lastColonIndex+1:]

	// 移除日期后缀部分，剩下的是 IP 和 UserAgent
	keyWithoutDate := keyWithoutPrefix[:lastColonIndex]

	// 再次查找最后一个冒号的位置，以分离 UserAgent 和 IP
	lastColonIndex = strings.LastIndex(keyWithoutDate, ":")

	// 提取 IP 和 UserAgent
	ip = keyWithoutDate[:lastColonIndex]
	userAgent = keyWithoutDate[lastColonIndex+1:]

	return ip, userAgent, date
}
