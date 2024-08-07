package util

import (
	"fmt"
)

// FindAdjacentPaths 查找指定 relative_path 的上一个和下一个路径
func (config *RepoConfig) FindAdjacentPaths(relativePath string) (string, string, error) {
	// 将所有 relative_path 扁平化为一个列表
	paths := flattenLayoutPaths(config.Layout)

	for i, path := range paths {
		if path == relativePath {
			prevPath := ""
			nextPath := ""
			if i > 0 {
				prevPath = paths[i-1]
			}
			if i < len(paths)-1 {
				nextPath = paths[i+1]
			}
			return prevPath, nextPath, nil
		}
	}

	return "", "", fmt.Errorf("relative_path not found: %s", relativePath)
}

// flattenLayoutPaths 将所有 Layout 的 relative_path 扁平化为一个列表，仅包括 isdir 为 false 的路径
func flattenLayoutPaths(layouts []Layout) []string {
	var paths []string
	for _, layout := range layouts {
		if !layout.Isdir {
			paths = append(paths, layout.RelativePath)
		}
		if layout.Sublayouts != nil {
			paths = append(paths, flattenLayoutPaths(layout.Sublayouts)...)
		}
	}
	return paths
}
