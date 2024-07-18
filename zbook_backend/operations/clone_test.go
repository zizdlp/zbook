package operations

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestCloneShouldOK(t *testing.T) {
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
}

func TestCloneWithPasswordShouldOK(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestCloneWithPasswordShouldOK is ignored *****")
		t.Skip()
	}
	// 使用一个公开的 Git 仓库 URL 进行测试
	gitURL := "https://github.com/zizdlp/full-stack-guide.git"

	rsg := util.NewRandomStringGenerator()
	randomString := rsg.RandomString(10)
	cloneDir := "/tmp/zbook_repo/" + randomString
	// 删除目标目录以确保每次测试都是从头开始
	if _, err := os.Stat(cloneDir); err == nil {
		os.RemoveAll(cloneDir)
	}
	token := "right_token"
	username := "zizdlp"
	// 调用 Clone 函数
	err := CloneWithPassword(gitURL, cloneDir, username, token)

	// 验证没有返回错误
	require.NoError(t, err)

	// 验证目标目录已创建
	_, err = os.Stat(cloneDir)
	require.NoError(t, err)
}
