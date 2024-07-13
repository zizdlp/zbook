package convert

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	md "github.com/zizdlp/zbook/markdown/render"
)

func TestConvertMarkdownBuffer(t *testing.T) {
	// 记录开始时间
	srcPath := "./test_files"
	files, err := os.ReadDir(srcPath)
	markdown := md.GetMarkdownConfig()
	require.NoError(t, err)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".md" {
			fmt.Println("test markdown render:", file.Name())
			filePath := filepath.Join(srcPath, file.Name())
			data, err := os.ReadFile(filePath)
			require.NoError(t, err)
			_, _, err = ConvertMarkdownBuffer(data, markdown)
			require.NoError(t, err)
		}
	}
}
