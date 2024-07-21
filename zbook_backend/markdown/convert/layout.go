package convert

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// 创建结构体
type Layout struct {
	Title      string   `json:"title"`
	Href       string   `json:"href"`
	Isdir      bool     `json:"isdir"`
	Sublayouts []Layout `json:"sublayouts"`
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
func GenLayout(prefix int, path string, layout *Layout) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("read path :%s failed: %v", path, err)
	}
	left_path := ""
	if prefix < len(path) {
		left_path = path[prefix:] + "/"
	}

	for _, f := range files {
		if f.Type().IsRegular() && strings.HasSuffix(f.Name(), ".md") {
			current_layout := Layout{
				Title: f.Name(),
				Isdir: false,
				Href:  left_path + strings.TrimSuffix(f.Name(), ".md"),
			}
			layout.Sublayouts = append(layout.Sublayouts, current_layout)
		}
		if f.IsDir() {
			current_layout := Layout{
				Title: f.Name(),
				Isdir: true,
			}
			err := GenLayout(prefix, path+"/"+f.Name(), &current_layout)
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
func LayoutToString(path string) (string, error) {
	layout := Layout{
		Title: "wiki",
		Isdir: true,
	}

	prefix := len(path) + 1
	err := GenLayout(prefix, path, &layout)
	if err != nil {
		return "", err
	}
	json_layout, err := json.Marshal(layout)
	if err != nil {
		return "", err
	}
	return string(json_layout), err
}

// RenderLayout 将 LayoutToString 返回的字符串写入文件
func RenderLayout(path string, outputPath string) error {
	stringLayout, err := LayoutToString(path)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputPath+"/layout.layout", []byte(stringLayout), 0644)
	if err != nil {
		return err
	}

	return nil
}
