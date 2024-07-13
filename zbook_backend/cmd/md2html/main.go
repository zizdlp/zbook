package main

import (
	"fmt"
	"os"

	"github.com/zizdlp/zbook/markdown/convert"
)

func main() {
	// 从命令行获取源目录和目标目录
	if len(os.Args) != 3 {
		fmt.Println("请提供源目录和目标目录参数")
		return
	}
	srcDir := os.Args[1]
	destDir := os.Args[2]
	// convert.ConvertMd2Html(srcDir, destDir) //convert content
	convert.ConvertFolder(srcDir, destDir)
	// convert.ConvertMdTable2Html(srcDir, destDir) // convert table
}
