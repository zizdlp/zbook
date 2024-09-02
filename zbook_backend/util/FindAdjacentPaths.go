package util

import (
	"encoding/json"
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
func (config *RepoConfig) GetFirstDocumentMap() (map[string]string, error) {
	// 初始化返回的 map
	firstDocs := make(map[string]string)

	// 遍历所有的语言
	for lang := range config.Layout {
		// 获取当前语言的第一个文档
		doc, err := config.GetFirstDocument(lang)
		if err != nil {
			// 如果找不到文档，跳过这个语言
			continue
		}
		// 将语言和文档路径添加到 map 中
		firstDocs[lang] = doc
	}

	// 如果没有找到任何文档，返回错误
	if len(firstDocs) == 0 {
		return nil, fmt.Errorf("no documents found for any language")
	}

	return firstDocs, nil
}

// GetDocumentPath 根据请求的语言从 JSON 字符串中获取文档路径，如果找不到则使用默认语言。
func GetDocumentPath(homes string, lang string) (string, error) {
	var restoredDocs map[string]string
	err := json.Unmarshal([]byte(homes), &restoredDocs)
	if err != nil {
		return "", fmt.Errorf("parse home error : %s", err)
	}

	// 获取请求的语言文档路径
	path, ok := restoredDocs[lang]
	if !ok {
		// 如果找不到请求的语言，尝试获取默认语言的文档路径
		path, ok = restoredDocs["default"]
		if !ok {
			// 如果还找不到默认语言的文档路径，返回错误
			return "", fmt.Errorf("document not found for language: %s", lang)
		}
	}

	return path, nil
}
