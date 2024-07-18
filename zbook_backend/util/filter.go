package util

import "strings"

func IsInt64InArray(value int64, array []int64) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

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

// getFileExtension 获取文件的后缀名
func getFileExtension(fileName string) string {
	segments := strings.Split(fileName, ".")
	if len(segments) > 1 {
		return "." + segments[len(segments)-1]
	}
	return ""
}
