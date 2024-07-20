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

	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			status := parts[0]
			fileName := parts[1]
			switch status {
			case "A":
				addedFiles = append(addedFiles, fileName)
			case "M":
				modifiedFiles = append(modifiedFiles, fileName)
			case "D":
				deletedFiles = append(deletedFiles, fileName)
			}
		} else if len(parts) == 1 && oldCommitID == "" {
			// If we are listing all files in the repository, mark them as "A" (added)
			fileName := parts[0]
			addedFiles = append(addedFiles, fileName)
		}
	}

	return addedFiles, modifiedFiles, deletedFiles, nil
}
