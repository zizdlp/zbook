package operations

import (
	"fmt"
	"os/exec"
	"strings"
)

// GetLatestCommit retrieves the latest commit hash from the specified repository directory.
func GetLatestCommit(dir string) (string, error) {
	// Change the working directory to the specified directory
	cmd := exec.Command("git", "-C", dir, "rev-parse", "HEAD")

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get latest commit: %v", err)
	}

	// Return the latest commit hash
	return strings.TrimSpace(string(output)), nil
}
