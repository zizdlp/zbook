package operations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestGetLatestCommitShouldOK(t *testing.T) {
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

	// 调用 GetLatestCommit 函数
	commitHash, err := GetLatestCommit(cloneDir)
	// 验证没有返回错误
	require.NoError(t, err)

	// 验证 commitHash 不为空
	require.NotEmpty(t, commitHash)

}
