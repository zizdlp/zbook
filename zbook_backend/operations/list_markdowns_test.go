package operations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestListMarkdownFilesShouldOK(t *testing.T) {
	// 使用一个公开的 Git 仓库 URL 进行测试
	gitURL := "https://github.com/zizdlp/zbook-user-guide.git"

	rsg := util.NewRandomStringGenerator()
	randomString := rsg.RandomString(10)
	cloneDir := "/tmp/zbook_repo/" + randomString
	// 删除目标目录以确保每次测试都是从头开始
	if _, err := os.Stat(cloneDir); err == nil {
		os.RemoveAll(cloneDir)
	}

	// 调用 Clone 函数
	err := Clone(gitURL, cloneDir)
	// 验证没有返回错误
	require.NoError(t, err)

	// 验证目标目录已创建
	_, err = os.Stat(cloneDir)
	require.NoError(t, err)

	// 调用 ListMarkdownFiles 函数
	mdFiles, err := ListMarkdownFiles(cloneDir)
	// 验证没有返回错误
	require.NoError(t, err)

	// 验证 mdFiles 不为空
	require.NotEmpty(t, mdFiles)

	// 输出 Markdown 文件列表，便于观察测试结果
	t.Logf("Markdown files in repository:")
	for _, file := range mdFiles {
		t.Logf("- %s", file)
	}
}
