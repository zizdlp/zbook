package util

import (
	"fmt"
	"net/http"
)

func FetchGithub(access_token string) (*http.Response, error) {
	// 创建新的请求
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// 设置 Authorization 头部
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	// 创建 HTTP 客户端
	client := &http.Client{}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}

	return resp, nil
}
