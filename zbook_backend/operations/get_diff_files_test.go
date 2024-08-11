package operations

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestGetDiffFilesShouldOK(t *testing.T) {
	// 使用一个公开的 Git 仓库 URL 进行测试
	gitURL := "https://github.com/zizdlp/zbook-docs.git"

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
	addedFiles, deletedFiles, modifiedFiles, renameFiles, err := GetDiffFiles(oldCommit, lastCommit, cloneDir)
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

	t.Logf("Rename files between commit %s and commit %s:", oldCommit, lastCommit)
	for i := 0; i < len(renameFiles); i += 2 {
		t.Logf("R - %s - %s", renameFiles[i], renameFiles[i+1])
	}
}

func TestGetAllFilesShouldOK(t *testing.T) {
	// 使用一个公开的 Git 仓库 URL 进行测试
	gitURL := "https://github.com/zizdlp/zbook-docs.git"

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
	addedFiles, deletedFiles, modifiedFiles, renameFiles, err := GetDiffFiles(oldCommit, lastCommit, cloneDir)
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

	t.Logf("Rename files between commit %s and commit %s:", oldCommit, lastCommit)
	for i := 0; i < len(renameFiles); i += 2 {
		t.Logf("R - %s,%s", renameFiles[i], renameFiles[i+1])
	}
}

func TestGetDiffFilesRename(t *testing.T) {
	if testing.Short() {
		fmt.Println("***** TestCloneWithPassword is ignored *****")
		t.Skip()
	}
	// 创建一个临时Git仓库并执行一些命令来设置测试环境
	// 请确保在执行测试之前，创建一个临时的Git仓库，模拟提交和文件修改

	gitURL := "https://github.com/zizdlp/frpc.git"

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

	oldCommitID := "cb924727b336e70cab13ebe46d2daee7381f4b17"
	newCommitID := "ef68cac5c79c3e683e90efa4dcca2db25d13375a"

	addedFiles, modifiedFiles, deletedFiles, renameFiles, err := GetDiffFiles(oldCommitID, newCommitID, cloneDir)
	require.NoError(t, err)

	// 检查返回的文件列表是否符合预期
	require.Equal(t, []string{"getting-started/quick-start.mdx"}, addedFiles)
	require.Empty(t, modifiedFiles)
	require.Equal(t, []string{"getting-started/quick-start.md"}, deletedFiles)
	require.Empty(t, renameFiles)
}
