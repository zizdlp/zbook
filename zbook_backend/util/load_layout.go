package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 定义结构体来表示 JSON 对象的结构
type Anchor struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
	URL  string `json:"url"`
}

type SubGroup struct {
	Group string   `json:"group"`
	Pages []string `json:"pages"`
}

type Navigation struct {
	Group    string     `json:"group"`
	Pages    []string   `json:"pages"`
	SubGroup []SubGroup `json:"subgroup,omitempty"`
}

type FooterSocials struct {
	Discord string `json:"discord"`
	GitHub  string `json:"github"`
}

type JsonObject struct {
	Anchors       []Anchor      `json:"anchors"`
	Navigation    []Navigation  `json:"navigation"`
	FooterSocials FooterSocials `json:"footerSocials"`
}

func ReadJsonFile(filePath string) (*JsonObject, error) {
	// 打开文件
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %s, error: %v", filePath, err)
	}
	defer file.Close()

	// 读取文件内容
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %s, error: %v", filePath, err)
	}

	// 解析 JSON 数据
	var jsonObject JsonObject
	err = json.Unmarshal(bytes, &jsonObject)
	if err != nil {
		return nil, fmt.Errorf("failed to parse json: %v", err)
	}

	return &jsonObject, nil
}

// func main() {
// 	// 示例文件路径
// 	filePath := "path/to/your/jsonfile.json"

// 	// 读取 JSON 文件
// 	jsonObject, err := ReadJsonFile(filePath)
// 	if err != nil {
// 		log.Fatalf("Error reading JSON file: %v", err)
// 	}

// 	// 输出解析后的数据
// 	fmt.Printf("Anchors:\n")
// 	for _, anchor := range jsonObject.Anchors {
// 		fmt.Printf("  Name: %s, Icon: %s, URL: %s\n", anchor.Name, anchor.Icon, anchor.URL)
// 	}

// 	fmt.Printf("\nNavigation:\n")
// 	for _, nav := range jsonObject.Navigation {
// 		fmt.Printf("  Group: %s\n", nav.Group)
// 		fmt.Printf("  Pages: %v\n", nav.Pages)
// 		for _, sub := range nav.SubGroup {
// 			fmt.Printf("  SubGroup: %s, Pages: %v\n", sub.Group, sub.Pages)
// 		}
// 	}

// 	fmt.Printf("\nFooterSocials:\n")
// 	fmt.Printf("  Discord: %s\n", jsonObject.FooterSocials.Discord)
// 	fmt.Printf("  GitHub: %s\n", jsonObject.FooterSocials.GitHub)
// }
