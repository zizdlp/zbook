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
		cmd = exec.Command("sh", "-c", "git ls-files -z")
	} else {
		// Construct the git diff command
		cmd = exec.Command("sh", "-c", fmt.Sprintf("git diff --name-status -z %s %s", oldCommitID, newCommitID))
	}
	cmd.Dir = repoDir // Set the working directory to the specified repository directory

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("failed to run git command: %v", err)
	}

	// Split output into fields by null character
	fields := strings.Split(string(output), "\x00")
	// Initialize slices for added, deleted, and modified files
	var addedFiles []string
	var modifiedFiles []string
	var deletedFiles []string

	for i := 0; i < len(fields)-1; i++ {
		status := fields[i][0]
		fileNames := fields[i+1:]
		switch status {
		case 'A':
			addedFiles = append(addedFiles, fileNames[0])
		case 'M':
			modifiedFiles = append(modifiedFiles, fileNames[0])
		case 'D':
			deletedFiles = append(deletedFiles, fileNames[0])
		case 'R':
			// 重命名状态，文件名部分包含旧文件名和新文件名
			names := fileNames[:2]
			if len(names) == 2 {
				deletedFiles = append(deletedFiles, names[0])
				addedFiles = append(addedFiles, names[1])
				i++ // Skip the next field since it is part of the rename
			}
		}
		i++ // Skip to the next status and fileNames pair
	}

	return addedFiles, modifiedFiles, deletedFiles, nil
}
