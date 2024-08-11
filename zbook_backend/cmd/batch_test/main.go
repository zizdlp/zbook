package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"

	"github.com/zizdlp/zbook/util"
)

func create_user(config util.Config, jsonStr string) {
	payload := bytes.NewBuffer([]byte(jsonStr))
	// 创建POST请求
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/v1/create_user", payload)
	if err != nil {
		fmt.Println("创建用户请求失败:", err)
		return
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return
	}

	// 输出响应内容
	fmt.Println("响应内容:", string(body))
}

type LoginResponseData struct {
	AccessToken string `json:"access_token"`
}

func login_user(config util.Config, jsonStr string) (string, error) {
	payload := bytes.NewBuffer([]byte(jsonStr))
	// 创建POST请求
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/v1/login_user", payload)
	if err != nil {
		fmt.Println("创建用户请求失败:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return "", err
	}

	// 输出响应内容
	// fmt.Println("响应内容:", string(body))

	var data LoginResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("解码JSON失败:", err)
		return "", err
	}
	return data.AccessToken, nil
}

func update_user(config util.Config, jsonStr string, access_token string) (string, error) {
	payload := bytes.NewBuffer([]byte(jsonStr))
	// 创建POST请求
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/v1/update_user", payload)
	if err != nil {
		fmt.Println("创建用户请求失败:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return "", err
	}

	// 输出响应内容
	fmt.Println("响应内容:", string(body))
	return "", err
}

func create_follow(config util.Config, jsonStr string, access_token string) (string, error) {
	payload := bytes.NewBuffer([]byte(jsonStr))
	// 创建POST请求
	req, err := http.NewRequest("POST", "http://"+config.HTTPServerAddress+"/v1/create_follow", payload)
	if err != nil {
		fmt.Println("创建用户请求失败:", err)
		return "", err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", access_token))

	client := &http.Client{}

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("发送请求失败:", err)
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应的内容
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应失败:", err)
		return "", err
	}

	// 输出响应内容
	fmt.Println("响应内容:", string(body))
	return "", err
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msgf("cannot load config: %s", err)
	}
	numUser := 10
	for i := 0; i < numUser; i++ {
		username := "test_" + fmt.Sprint(i)
		email := username + "@zizdlp.com"
		password := "zzzz1234"
		content := fmt.Sprintf(`{"email":"%s","username":"%s","password":"%s" }`,
			email, username, password)
		loginContent := fmt.Sprintf(`{"email":"%s","password":"%s" }`,
			email, password)
		create_user(config, content)
		access_token, err := login_user(config, loginContent)
		if err == nil {
			fmt.Println("access_toekn is:", access_token)

			updateContent := fmt.Sprintf(`{"motto":"%s","visibiliy_level":"%s"}`,
				username+"_motto", "public")
			update_user(config, updateContent, access_token)

			followContent := fmt.Sprintf(`{"following_id":"%s" }`,
				"1")

			create_follow(config, followContent, access_token)
		}
	}
	// 定义JSON参数

}
