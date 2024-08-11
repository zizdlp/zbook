package util

import (
	"regexp"
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
