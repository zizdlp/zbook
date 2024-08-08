package util

import (
	"bytes"
	"encoding/base64"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/image/webp"
	// 导入 svgo 包，支持 SVG 格式
)

// ReadImageBytes 读取图片文件并返回其字节数据，支持的格式统一编码为 PNG，不支持的格式直接返回文件的字节数据
func ReadImageBytes(path string) ([]byte, error) {
	// 打开图片文件
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// 解码图片文件
	var img image.Image
	switch filepath.Ext(path) {
	case ".png", ".jpg", ".jpeg", ".gif":
		img, _, err = image.Decode(file)
	case ".webp":
		img, err = webp.Decode(file)
	default:
		// 如果格式不支持，直接返回文件的字节数据
		fileBytes, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil, readErr
		}
		return fileBytes, nil
	}

	if err != nil {
		return nil, err
	}

	// 创建字节缓冲区
	buffer := new(bytes.Buffer)

	// 将图像编码为 PNG
	err = png.Encode(buffer, img)
	if err != nil {
		return nil, err
	}

	// 返回编码后的字节数据
	return buffer.Bytes(), nil
}
func ReadImageBytesToBase64(imagePath string) (string, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	imageBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	base64Image := base64.StdEncoding.EncodeToString(imageBytes)
	return base64Image, nil
}
