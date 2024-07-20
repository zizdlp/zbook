package operations

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetDiffFiles returns three slices containing added, deleted, and modified file names between two commits.
func GetDiffFiles(oldCommitID, newCommitID, repoDir string) ([]string, []string, []string, error) {
	var cmd *exec.Cmd
	if oldCommitID == "" {
		// If oldCommitID is empty, list all files in the repository
		cmd = exec.Command("sh", "-c", "git ls-files -z | xargs -0 -n1 echo")
	} else {
		// Construct the git diff command
		cmd = exec.Command("sh", "-c", fmt.Sprintf("git diff --name-status -z %s %s | xargs -0 -n2 echo", oldCommitID, newCommitID))
	}
	cmd.Dir = repoDir // Set the working directory to the specified repository directory

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to run git command: %v", err)
	}

	// Split output into lines and trim whitespace
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	// Initialize slices for added, deleted, and modified files
	var addedFiles []string
	var modifiedFiles []string
	var deletedFiles []string
	fmt.Println("liens:", lines)
	for _, line := range lines {
		if oldCommitID == "" {
			addedFiles = append(addedFiles, line)
		} else {
			status := line[0]                       // 第一个字符是状态
			fileName := strings.TrimSpace(line[1:]) // 剩余部分是文件名
			switch status {
			case 'A':
				addedFiles = append(addedFiles, fileName)
			case 'M':
				modifiedFiles = append(modifiedFiles, fileName)
			case 'D':
				deletedFiles = append(deletedFiles, fileName)
			}
		}

	}

	return addedFiles, modifiedFiles, deletedFiles, nil
}
