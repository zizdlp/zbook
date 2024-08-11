package operations

import (
	"fmt"
	"os/exec"
	"strings"
)

// ListMarkdownFiles lists all .md files in the git repository located at the specified directory.
func ListMarkdownFiles(repoDir string) ([]string, error) {
	// Construct the git ls-files command
	// cmd := exec.Command("git", "-C", repoDir, "ls-files")
	cmd := exec.Command("sh", "-c", "git ls-files -z | xargs -0 -n1 echo")
	cmd.Dir = repoDir // Set the working directory to the specified repository directory
	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list files: %v", err)
	}

	// Split output into lines and filter for .md files
	allFiles := strings.Split(strings.TrimSpace(string(output)), "\n")
	mdFiles := []string{}
	for _, file := range allFiles {
		if strings.HasSuffix(file, ".md") {
			mdFiles = append(mdFiles, file)
		}
	}

	return mdFiles, nil
}
