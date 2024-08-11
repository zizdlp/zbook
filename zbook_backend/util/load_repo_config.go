package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// 定义结构体来表示 JSON 对象的结构
type Anchor struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}
type FooterSocial struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}

type RepoConfig struct {
	Anchors       []Anchor       `json:"anchors"`
	Layout        []Layout       `json:"layout"`
	FooterSocials []FooterSocial `json:"footerSocials"`
}

func ReadRepoConfig(filePath string) (*RepoConfig, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s, error: %v", filePath, err)
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s, error: %v", filePath, err)
	}

	// 解析 JSON 数据
	var jsonObject RepoConfig
	err = json.Unmarshal(bytes, &jsonObject)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %v", err)
	}

	return &jsonObject, nil
}
