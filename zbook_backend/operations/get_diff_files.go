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
		cmd = exec.Command("git", "ls-files")
	} else {
		// Construct the git diff command
		cmd = exec.Command("git", "diff", "--name-status", oldCommitID, newCommitID)
	}
	cmd.Dir = repoDir // 设置命令执行的工作目录为指定的仓库目录

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to run git command: %v", err)
	}

	// Split output into lines and trim whitespace
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	// Initialize slices for added, deleted, and modified files
	var addedFiles []string
	var deletedFiles []string
	var modifiedFiles []string

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			status := parts[0]
			fileName := parts[1]
			switch status {
			case "A":
				addedFiles = append(addedFiles, fileName)
			case "D":
				deletedFiles = append(deletedFiles, fileName)
			case "M":
				modifiedFiles = append(modifiedFiles, fileName)
			}
		} else if len(parts) == 1 && oldCommitID == "" {
			// If we are listing all files in the repository, mark them as "A" (added)
			fileName := parts[0]
			addedFiles = append(addedFiles, fileName)
		}
	}

	return addedFiles, deletedFiles, modifiedFiles, nil
}
