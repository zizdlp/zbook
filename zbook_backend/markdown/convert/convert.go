package convert

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/text"
	md "github.com/zizdlp/zbook/markdown/render"
)

func ConvertMarkdownBuffer(data []byte, markdown goldmark.Markdown) (bytes.Buffer, bytes.Buffer, error) {

	doc := markdown.Parser().Parse(text.NewReader(data))
	var tableBuffer bytes.Buffer
	if doc.HasChildren() && doc.FirstChild().Kind().String() == "Heading" { // 有heading就意味有 content table
		//remove 1. heading,2 lists
		err := markdown.Renderer().Render(&tableBuffer, data, doc.FirstChild().NextSibling()) // 1.渲染目录
		if err != nil {
			return bytes.Buffer{}, bytes.Buffer{}, err
		}
		content := doc.FirstChild()
		doc.RemoveChild(doc, content)
		content = doc.FirstChild() // 目录
		doc.RemoveChild(doc, content)
	}

	var mainBuffer bytes.Buffer
	err := markdown.Renderer().Render(&mainBuffer, data, doc) // 2.渲染正式内容
	if err != nil {
		return bytes.Buffer{}, bytes.Buffer{}, err
	}
	return tableBuffer, mainBuffer, nil

}
func ConvertMd2Html(src_path string, des_path string) {
	markdown := md.GetMarkdownConfig()
	data, err := os.ReadFile(src_path)
	if err != nil {
		log.Error().Err(err).Msgf("failed to open file: %s", src_path)
		return
	}
	_, main, err := ConvertMarkdownBuffer(data, markdown)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert markdown")
		return
	}
	// 写入 JSON 字符串到文件
	err = os.WriteFile(des_path, main.Bytes(), 0644)
	if err != nil {
		log.Error().Err(err).Msgf("failed to write file: %s", des_path)
		return
	}
}
func ConvertMdTable2Html(src_path string, des_path string) {
	markdown := md.GetMarkdownConfig()
	data, err := os.ReadFile(src_path)
	if err != nil {
		log.Error().Err(err).Msgf("failed to open file: %s", src_path)
		return
	}
	table, main, err := ConvertMarkdownBuffer(data, markdown)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert markdown")
		return
	}
	// 写入 JSON 字符串到文件
	err = os.WriteFile(des_path, main.Bytes(), 0644)
	if err != nil {
		log.Error().Err(err).Msgf("failed to write main file: %s", des_path)
		return
	}

	// 写入 JSON 字符串到文件
	err = os.WriteFile(des_path, table.Bytes(), 0644)
	if err != nil {
		log.Error().Err(err).Msgf("failed to write table file: %s", des_path)
		return
	}
}

func ConvertMd2Json(src_path string, des_path string) {
	markdown := md.GetMarkdownConfig()
	data, err := os.ReadFile(src_path)
	if err != nil {
		log.Error().Err(err).Msgf("failed to open file: %s", src_path)
		return
	}
	table, main, err := ConvertMarkdownBuffer(data, markdown)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert markdown")
		return
	}

	// 构建包含 HTML 表格和内容的 JSON 对象
	jsonObject := map[string]string{
		"table":   table.String(),
		"content": main.String(),
	}
	// 将 JSON 对象序列化为 JSON 字符串
	jsonString, err := json.Marshal(jsonObject)
	if err != nil {
		log.Error().Err(err).Msg("failed to parse json")
		return
	}

	// 写入 JSON 字符串到文件
	err = os.WriteFile(des_path, jsonString, 0644)
	if err != nil {
		log.Error().Err(err).Msgf("failed to write html file: %s", des_path)
		return
	}
}

// copyFile 复制文件
func copyFile(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 确保目标目录及其上级目录存在
	err = os.MkdirAll(filepath.Dir(dest), os.ModePerm)
	if err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}
func ConvertFolder(srcDir string, destDir string) {
	// 记录开始时间
	startTime := time.Now()
	// 遍历源目录
	err := filepath.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 检查是否为文件
		if !info.IsDir() {
			// 构建相对路径
			relPath, err := filepath.Rel(srcDir, srcPath)
			if err != nil {
				return err
			}

			// 构建目标路径
			destPath := filepath.Join(destDir, relPath)

			// 如果是以".md"结尾的文件，修改目标路径后缀为".json"
			if strings.HasSuffix(info.Name(), ".md") {
				destPath = filepath.Join(destDir, relPath[:len(relPath)-len(".md")]+".html")
				// 确保目标目录及其上级目录存在
				err = os.MkdirAll(filepath.Dir(destPath), os.ModePerm)
				if err != nil {
					return err
				}
				ConvertMd2Html(srcPath, destPath)
			} else {
				// 复制文件
				err = copyFile(srcPath, destPath)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})

	if err != nil {
		log.Error().Err(err).Msg("failed to iterate folder")
		return
	}

	// 记录结束时间
	endTime := time.Now()

	// 计算耗时并输出
	elapsedTime := endTime.Sub(startTime)
	log.Info().Msgf("convert folder time consume:%s", elapsedTime)
	// RenderLayout(srcDir, destDir)
}
