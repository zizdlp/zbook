package util

import (
	"fmt"
	"strings"
)

// FindAdjacentPaths 查找指定语言的 relative_path 的上一个和下一个路径
func (config *RepoConfig) FindAdjacentPaths(lang, relativePath string) (string, string, error) {
	// 将指定语言的所有 relative_path 扁平化为一个列表
	paths := flattenLayoutPaths(config.Layout, lang)

	for i, path := range paths {
		if path == strings.ToLower(relativePath) {
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

// flattenLayoutPaths 将指定语言的 Layout 的 relative_path 扁平化为一个列表，仅包括 isdir 为 false 的路径
func flattenLayoutPaths(layoutMap map[string][]Layout, lang string) []string {
	var paths []string

	// 获取指定语言的布局，如果找不到，则尝试获取 "default" 布局
	layouts, ok := layoutMap[lang]
	if !ok {
		layouts, ok = layoutMap["default"]
		if !ok {
			return paths // 如果找不到 "default" 布局，返回空列表
		}
	}

	for _, layout := range layouts {
		if !layout.Isdir {
			paths = append(paths, strings.ToLower(layout.RelativePath))
		}
		if layout.Sublayouts != nil {
			paths = append(paths, flattenLayoutPaths(map[string][]Layout{"": layout.Sublayouts}, "")...)
		}
	}
	return paths
}

// GetFirstDocument 查找指定语言的第一个文档，如果找不到，则返回默认语言中的第一个文档
func (config *RepoConfig) GetFirstDocument(lang string) (string, error) {
	layouts, ok := config.Layout[lang]
	if !ok {
		layouts, ok = config.Layout["default"]
		if !ok {
			return "", fmt.Errorf("no documents found for language: %s", lang)
		}
	}

	for _, layout := range layouts {
		if !layout.Isdir {
			return layout.RelativePath, nil
		}
		if layout.Sublayouts != nil {
			for _, sublayout := range layout.Sublayouts {
				if !sublayout.Isdir {
					return sublayout.RelativePath, nil
				}
			}
		}
	}

	return "", fmt.Errorf("no documents found for language: %s", lang)
}
