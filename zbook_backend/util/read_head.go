package util

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath" // 导入 svgo 包，支持 SVG 格式

	"golang.org/x/image/webp"
)

func ReadImageBytes(path string) ([]byte, error) {
	// 打开图片文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("open file failed:", err)
		return nil, err
	}
	defer file.Close()

	var img image.Image
	var format string

	// 根据文件扩展名选择解码器
	switch filepath.Ext(path) {
	case ".png":
		// 解码 PNG 文件
		img, err = png.Decode(file)
		format = "png"
	case ".jpg", ".jpeg":
		// 解码 JPEG 文件
		img, err = jpeg.Decode(file)
		format = "jpeg"
	case ".webp":
		// 解码 WebP 文件
		img, err = webp.Decode(file)
		format = "webp"
	case ".gif":
		// 解码 GIF 文件
		img, err = gif.Decode(file)
		format = "gif"
	default:
		return nil, fmt.Errorf("unsupported image format")
	}

	if err != nil {
		fmt.Printf("解码 %s 文件失败: %s\n", format, err)
		return nil, err
	}

	// 创建字节缓冲区
	buffer := new(bytes.Buffer)

	// 将图像编码为 base64 字符串
	switch format {
	case "png":
		err = png.Encode(buffer, img)
	case "jpeg":
		err = jpeg.Encode(buffer, img, nil)
	case "webp":
		options := &jpeg.Options{
			Quality: 80,
		}
		err = jpeg.Encode(buffer, img, options)
	case "gif":
		err = gif.Encode(buffer, img, nil)
	}
	if err != nil {
		fmt.Printf("%s 图像编码失败: %s\n", format, err)
		return nil, err
	}

	// 返回 base64 编码的字节数据
	return buffer.Bytes(), nil
}
