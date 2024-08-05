package util

import (
	"encoding/json"
	"fmt"
)

// ParseRepoConfigFromString 解析 JSON 格式的字符串并返回 RepoConfig 结构体
func ParseRepoConfigFromString(jsonStr string) (*RepoConfig, error) {
	var config RepoConfig
	err := json.Unmarshal([]byte(jsonStr), &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %v", err)
	}
	return &config, nil
}
