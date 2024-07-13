package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 创建结构体
type Layout struct {
	Title        string   `json:"title"`
	RelativePath string   `json:"relative_path"`
	MarkdownID   int64    `json:"markdown_id"`
	Isdir        bool     `json:"isdir"`
	Sublayouts   []Layout `json:"sublayouts"`
}

// 查看本目录下面是否有md文件，有的话true
func hasMdUnder(Layout *Layout) bool {
	if !Layout.Isdir {
		return true
	}
	for _, sublayout := range Layout.Sublayouts {
		if hasMdUnder(&sublayout) {
			return true
		}
	}
	return false
}

// 生成一个 layout json 文档
func GenLayout(prefix int, path string, layout *Layout, q *Queries, RepoID int64) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("读取目录:%s 失败: %v", path, err)
	}
	left_path := ""
	if prefix < len(path) {
		left_path = path[prefix:] + "/"
	}

	for _, f := range files {
		if f.Type().IsRegular() && strings.HasSuffix(f.Name(), ".md") {
			arg := GetMarkdownContentParams{
				RelativePath: strings.ToLower(strings.TrimSuffix(left_path+f.Name(), ".md")),
				RepoID:       RepoID,
			}
			markdown_info, err := q.GetMarkdownContent(context.Background(), arg)
			if err != nil {
				return fmt.Errorf("获取markdown: %s 失败: %v", left_path+f.Name(), err)
			}
			current_layout := Layout{
				Title:        strings.TrimSuffix(f.Name(), ".md"),
				Isdir:        false,
				RelativePath: strings.ToLower(strings.TrimSuffix(left_path+f.Name(), ".md")),
				MarkdownID:   markdown_info.MarkdownID,
			}
			layout.Sublayouts = append(layout.Sublayouts, current_layout)
		} else if f.IsDir() {
			current_layout := Layout{
				Title:        f.Name(),
				Isdir:        true,
				RelativePath: strings.ToLower(left_path + f.Name()),
			}
			err := GenLayout(prefix, path+"/"+f.Name(), &current_layout, q, RepoID)
			if err != nil {
				return err
			}
			if hasMdUnder(&current_layout) {
				layout.Sublayouts = append(layout.Sublayouts, current_layout)
				lens := len(layout.Sublayouts)
				if current_layout.Isdir { //[l,r]
					left := 0
					right := len(layout.Sublayouts) - 1

					for left < right {
						mid := (left + right) / 2
						mid_v := layout.Sublayouts[mid].Isdir
						if mid_v {
							left = mid + 1
						} else {
							right = mid
						}
					}
					if left < lens-1 {
						layout.Sublayouts[left], layout.Sublayouts[lens-1] = layout.Sublayouts[lens-1], layout.Sublayouts[left]
					}
				}
			}
		}
	}
	return nil
}

// 生成仓库对应的layout json,
// 出错则返回："",err
func LayoutToString(path string, q *Queries, RepoID int64) (string, error) {
	layout := Layout{
		Title: "wiki",
		Isdir: true,
	}

	prefix := len(path) + 1
	err := GenLayout(prefix, path, &layout, q, RepoID)
	if err != nil {
		return "", err
	}
	json_layout, err := json.Marshal(layout)
	if err != nil {
		return "", err
	}
	return string(json_layout), err
}

func RenderLayout(rootPath string, RepoID int64, q *Queries, version_key string) error {
	string_layout, err := LayoutToString(rootPath, q, RepoID)
	if err != nil {
		return err
	}
	arg_update_repo_layout := UpdateRepoLayoutParams{
		RepoID: RepoID,
		Layout: string_layout,
	}
	err = q.UpdateRepoLayout(context.Background(), arg_update_repo_layout)
	if err != nil {
		return err
	}
	return nil
}
