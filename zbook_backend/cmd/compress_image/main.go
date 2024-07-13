package main

import (
	"fmt"
	"os"

	"github.com/zizdlp/zbook/util"
)

func main() {
	// 从命令行获取源目录和目标目录
	if len(os.Args) != 3 {
		fmt.Println("请提供源目录和目标目录参数")
		return
	}
	srcImg := os.Args[1]
	desImg := os.Args[2]

	base64, err := util.ReadImageBytes(srcImg)
	if err != nil {
		fmt.Printf("failed to read image to base64: %v\n", err)
		return
	}
	// 调用压缩函数
	compressedImage, err := util.CompressImage(base64)
	if err != nil {
		fmt.Printf("failed to compress image: %v\n", err)
		return
	}

	// 保存压缩后的文件
	err = os.WriteFile(desImg, compressedImage, 0644)
	if err != nil {
		fmt.Printf("failed to write compressed image to file: %v\n", err)
		return
	}

	fmt.Printf("Image compressed and saved to %s\n", desImg)
}
