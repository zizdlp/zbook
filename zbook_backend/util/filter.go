package util

import (
	"fmt"
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
	fmt.Println("mydebug: logKey is:", logKey)
	keyWithoutPrefix := strings.TrimPrefix(logKey, "logvisitor:")

	// 查找最后一个冒号的位置，用来确定日期的分隔点
	lastColonIndex := strings.LastIndex(keyWithoutPrefix, ":")

	// 提取日期
	date = keyWithoutPrefix[lastColonIndex+1:]

	// 移除日期后缀部分，剩下的是 IP 和 UserAgent
	keyWithoutDate := keyWithoutPrefix[:lastColonIndex]

	// 确保以 IPv6 的情况来处理 IP
	ipAndUserAgent := strings.SplitN(keyWithoutDate, ":", 2)
	if len(ipAndUserAgent) == 2 {
		ip = ipAndUserAgent[0]
		userAgent = ipAndUserAgent[1]
	} else {
		ip = keyWithoutDate
		userAgent = ""
	}

	return ip, userAgent, date
}
