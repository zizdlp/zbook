package util

import "strings"

type Layout struct {
	Title        string   `json:"title"`
	RelativePath string   `json:"relative_path"`
	Isdir        bool     `json:"isdir"`
	Sublayouts   []Layout `json:"sublayouts"`
}

// findOrCreateSubLayout 查找或创建子布局
func findOrCreateSubLayout(layout *Layout, title string) *Layout {
	for i := range layout.Sublayouts {
		if layout.Sublayouts[i].Title == title {
			return &layout.Sublayouts[i]
		}
	}
	relativePath := title
	if layout.RelativePath != "" {
		relativePath = layout.RelativePath + "/" + title
	}
	newLayout := Layout{
		Title:        title,
		RelativePath: relativePath,
		Isdir:        true,
		Sublayouts:   []Layout{},
	}
	layout.Sublayouts = append(layout.Sublayouts, newLayout)
	return &layout.Sublayouts[len(layout.Sublayouts)-1]
}

func CreateLayout(files []string) Layout {
	root := Layout{
		Title:        "wiki",
		RelativePath: "",
		Isdir:        true,
		Sublayouts:   []Layout{},
	}

	for _, file := range files {
		if !strings.HasSuffix(file, ".md") {
			continue
		}
		parts := strings.Split(file, "/")
		current := &root
		for i, part := range parts {
			if i == len(parts)-1 {
				// 这是一个文件
				relativePath := part
				if current.RelativePath != "" {
					relativePath = current.RelativePath + "/" + part
				}
				relativePath = strings.ToLower(strings.TrimSuffix(relativePath, ".md"))
				title := strings.TrimSuffix(part, ".md")
				current.Sublayouts = append(current.Sublayouts, Layout{
					Title:        title,
					RelativePath: relativePath,
					Isdir:        false,
					Sublayouts:   nil,
				})
			} else {
				// 这是一个目录
				current = findOrCreateSubLayout(current, part)
			}
		}
	}

	return root
}
