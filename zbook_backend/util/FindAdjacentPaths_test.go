package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindAdjacentPaths(t *testing.T) {
	// 创建测试配置
	config := &RepoConfig{
		Layout: []Layout{
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
	}

	tests := []struct {
		relativePath  string
		expectedPrev  string
		expectedNext  string
		expectedError bool
	}{
		{"b/c", "", "b/e", false},
		{"b/e", "b/c", "a", false},
		{"a", "b/e", "", false},
		{"nonexistent", "", "", true},
	}

	for _, tt := range tests {
		prev, next, err := config.FindAdjacentPaths(tt.relativePath)
		if tt.expectedError {
			require.Error(t, err)
		} else {
			require.NoError(t, err)
		}
		require.Equal(t, tt.expectedPrev, prev)
		require.Equal(t, tt.expectedNext, next)
	}
}
