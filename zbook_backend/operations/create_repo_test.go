package operations_test

import (
	"testing"

	"github.com/zizdlp/zbook/operations"
)

func TestCreateRepo(t *testing.T) {
	// 测试用的 Git 仓库 URL
	gitURL := "https://github.com/zizdlp/full-stack-guide.git"

	// 调用 CreateRepo 函数
	err := operations.CreateRepo(gitURL, int64(1), int64(1))
	if err != nil {
		t.Fatalf("Failed to create repo: %v", err)
	}
}
