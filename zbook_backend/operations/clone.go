package operations

import (
	"fmt"
	"os/exec"
	"strings"
)

// Clone clones a git repository from the specified URL into the specified directory.
func Clone(gitURL string, dir string) error {
	// Create the git clone command with the directory parameter
	cmd := exec.Command("git", "clone", gitURL, dir)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repository: %v, output: %s", err, string(output))
	}

	return nil
}

// Clone clones a git repository from the specified URL into the specified directory.
// It supports cloning private repositories using either a personal access token (token)
// or basic authentication (username and password).
func CloneWithPassword(gitURL, dir, username, password string) error {
	// Construct the clone URL with username and password embedded
	urlWithCredentials := embedCredentialsInURL(gitURL, username, password)
	// Create the git clone command with the directory parameter
	cmd := exec.Command("git", "clone", urlWithCredentials, dir)

	// Run the command and capture its output
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to clone repository: %v, output: %s", err, string(output))
	}

	return nil
}

// embedCredentialsInURL embeds the username and password into the git URL for basic authentication.
func embedCredentialsInURL(gitURL, username, password string) string {
	// Split the URL at "//" to insert username and password
	parts := strings.Split(gitURL, "//")
	if len(parts) < 2 {
		return gitURL // Invalid URL format, return as is
	}

	// Insert username and password after "//"
	return fmt.Sprintf("%s//%s:%s@%s", parts[0], username, password, parts[1])
}
