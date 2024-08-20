package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterDiffFilesByExtensions(t *testing.T) {
	files := []string{
		"README.md",
		"image.png",
		"photo.jpg",
		"document.txt",
		"script.js",
		"style.css",
		"diagram.svg",
		"animation.gif",
		"vector.webp",
		"archive.zip",
	}

	allowedExtensions := map[string]bool{
		".md":   true,
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".svg":  true,
		".gif":  true,
		".webp": true,
	}

	expectedFilteredFiles := []string{
		"README.md",
		"image.png",
		"photo.jpg",
		"diagram.svg",
		"animation.gif",
		"vector.webp",
	}

	filteredFiles := FilterDiffFilesByExtensions(files, allowedExtensions)
	require.Equal(t, expectedFilteredFiles, filteredFiles, "Filtered files do not match expected output")
}

func TestExtractLogDetails(t *testing.T) {
	// 测试数据
	logKey := "logvisitor:::1:Mozilla/5.0:2024-08-20"

	// 期望的结果
	expectedIP := "::1"
	expectedUserAgent := "Mozilla/5.0"
	expectedDate := "2024-08-20"

	// 调用函数
	ip, userAgent, date := ExtractLogDetails(logKey)

	// 使用 require 进行断言
	require.Equal(t, expectedIP, ip, "IP 地址不正确")
	require.Equal(t, expectedUserAgent, userAgent, "User Agent 不正确")
	require.Equal(t, expectedDate, date, "日期不正确")
}
