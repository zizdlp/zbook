package operations

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestGetDiffFilesShouldOK(t *testing.T) {
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
	lastCommit, err := GetLatestCommit(cloneDir)
	// 验证没有返回错误
	require.NoError(t, err)

	// 验证 commitHash 不为空
	require.NotEmpty(t, lastCommit)

	oldCommit := "6eec29b0b8188d4877cac57a90b17fd5fa8f165c"

	// 调用 GetDiffFiles 函数
	addedFiles, deletedFiles, modifiedFiles, err := GetDiffFiles(oldCommit, lastCommit, cloneDir)
	require.NoError(t, err)

	// 输出差异文件列表，便于观察测试结果
	t.Logf("Added files between commit %s and commit %s:", oldCommit, lastCommit)
	for _, file := range addedFiles {
		t.Logf("A - %s", file)
	}

	t.Logf("Deleted files between commit %s and commit %s:", oldCommit, lastCommit)
	for _, file := range deletedFiles {
		t.Logf("D - %s", file)
	}

	t.Logf("Modified files between commit %s and commit %s:", oldCommit, lastCommit)
	for _, file := range modifiedFiles {
		t.Logf("M - %s", file)
	}
}

func TestGetAllFilesShouldOK(t *testing.T) {
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
	lastCommit, err := GetLatestCommit(cloneDir)
	// 验证没有返回错误
	require.NoError(t, err)

	// 验证 commitHash 不为空
	require.NotEmpty(t, lastCommit)

	oldCommit := ""

	// 调用 GetDiffFiles 函数
	addedFiles, deletedFiles, modifiedFiles, err := GetDiffFiles(oldCommit, lastCommit, cloneDir)
	require.NoError(t, err)

	// 输出差异文件列表，便于观察测试结果
	t.Logf("Added files in the repository:")
	for _, file := range addedFiles {
		t.Logf("A - %s", file)
	}

	t.Logf("Deleted files in the repository:")
	for _, file := range deletedFiles {
		t.Logf("D - %s", file)
	}

	t.Logf("Modified files in the repository:")
	for _, file := range modifiedFiles {
		t.Logf("M - %s", file)
	}
}
