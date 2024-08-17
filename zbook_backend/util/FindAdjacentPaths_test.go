package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAdjacentPaths(t *testing.T) {
	// 创建测试配置，包含多语言布局
	config := &RepoConfig{
		Layout: map[string][]Layout{
			"en": {
				{
					Title:        "b",
					RelativePath: "b",
					Isdir:        true,
					Sublayouts: []Layout{
						{
							Title:        "c",
							RelativePath: "b/c",
							Isdir:        false,
							Sublayouts:   nil,
						},
						{
							Title:        "e",
							RelativePath: "b/e",
							Isdir:        false,
							Sublayouts:   nil,
						},
					},
				},
				{
					Title:        "a",
					RelativePath: "a",
					Isdir:        false,
					Sublayouts:   nil,
				},
			},
			"zh": {
				{
					Title:        "一",
					RelativePath: "yi",
					Isdir:        true,
					Sublayouts: []Layout{
						{
							Title:        "二",
							RelativePath: "yi/er",
							Isdir:        false,
							Sublayouts:   nil,
						},
					},
				},
				{
					Title:        "三",
					RelativePath: "san",
					Isdir:        false,
					Sublayouts:   nil,
				},
			},
		},
	}

	tests := []struct {
		lang          string
		relativePath  string
		expectedPrev  string
		expectedNext  string
		expectedError bool
	}{
		{"en", "b/c", "", "b/e", false},
		{"en", "b/e", "b/c", "a", false},
		{"en", "a", "b/e", "", false},
		{"en", "nonexistent", "", "", true},
		{"zh", "yi/er", "", "san", false},
		{"zh", "san", "yi/er", "", false},
		{"zh", "nonexistent", "", "", true},
	}

	for _, tt := range tests {
		prev, next, err := config.FindAdjacentPaths(tt.lang, tt.relativePath)
		if tt.expectedError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tt.expectedPrev, prev)
		require.Equal(t, tt.expectedNext, next)
	}
}

func TestGetFirstDocument(t *testing.T) {
	// 创建测试配置
	config := &RepoConfig{
		Layout: map[string][]Layout{
			"en": {
				{
					Title:        "doc1",
					RelativePath: "doc1",
					Isdir:        false,
					Sublayouts:   nil,
				},
				{
					Title:        "doc2",
					RelativePath: "doc2",
					Isdir:        false,
					Sublayouts:   nil,
				},
			},
			"zh": {
				{
					Title:        "文档1",
					RelativePath: "doc1-zh",
					Isdir:        false,
					Sublayouts:   nil,
				},
			},
			"default": {
				{
					Title:        "default-doc",
					RelativePath: "default-doc",
					Isdir:        false,
					Sublayouts:   nil,
				},
			},
		},
	}

	tests := []struct {
		lang          string
		expectedDoc   string
		expectedError bool
	}{
		{"en", "doc1", false},
		{"zh", "doc1-zh", false},
		{"fr", "default-doc", false}, // Testing for a non-existent language should fallback to default
		{"", "default-doc", false},   // Testing with an invalid language that should return an error
	}

	for _, tt := range tests {
		doc, err := config.GetFirstDocument(tt.lang)
		if tt.expectedError {
			require.Error(t, err)
			require.Empty(t, doc)
		} else {
			require.NoError(t, err)
			require.Equal(t, tt.expectedDoc, doc)
		}
	}
}
