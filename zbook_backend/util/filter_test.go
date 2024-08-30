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
	// IPv4 Test Case
	logKeyIPv4 := "logvisitor:192.168.1.1:Mozilla/5.0:2024-08-20"
	expectedIPIPv4 := "192.168.1.1"
	expectedUserAgentIPv4 := "Mozilla/5.0"
	expectedDateIPv4 := "2024-08-20"
	ip, userAgent, date := ExtractLogDetails(logKeyIPv4)
	require.Equal(t, expectedIPIPv4, ip, "IPv4 address not correct")
	require.Equal(t, expectedUserAgentIPv4, userAgent, "User Agent not correct")
	require.Equal(t, expectedDateIPv4, date, "Date not correct")

	// IPv6 Test Case
	logKeyIPv6 := "logvisitor:2001:0db8:85a3:0000:0000:8a2e:0370:7334:Mozilla/5.0:2024-08-20"
	expectedIPIPv6 := "2001:0db8:85a3:0000:0000:8a2e:0370:7334"
	expectedUserAgentIPv6 := "Mozilla/5.0"
	expectedDateIPv6 := "2024-08-20"
	ip, userAgent, date = ExtractLogDetails(logKeyIPv6)
	require.Equal(t, expectedIPIPv6, ip, "IPv6 address not correct")
	require.Equal(t, expectedUserAgentIPv6, userAgent, "User Agent not correct")
	require.Equal(t, expectedDateIPv6, date, "Date not correct")

	// Additional cases with colons in UserAgent
	testCases := []struct {
		logKey            string
		expectedIP        string
		expectedUserAgent string
		expectedDate      string
	}{
		{"logvisitor:66.249.66.168:Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.119 Mobile Safari/537.36 (compatible; GoogleOther):2024-08-26", "66.249.66.168", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.119 Mobile Safari/537.36 (compatible; GoogleOther)", "2024-08-26"},
		{"logvisitor:66.249.66.168:Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.119 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http_//www.google.com/bot.html):2024-08-23", "66.249.66.168", "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.6533.119 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http_//www.google.com/bot.html)", "2024-08-23"},
		// Add more test cases as needed
	}

	for _, tc := range testCases {
		ip, userAgent, date := ExtractLogDetails(tc.logKey)
		require.Equal(t, tc.expectedIP, ip, "IP not correct for logKey: "+tc.logKey)
		require.Equal(t, tc.expectedUserAgent, userAgent, "User Agent not correct for logKey: "+tc.logKey)
		require.Equal(t, tc.expectedDate, date, "Date not correct for logKey: "+tc.logKey)
	}
}
