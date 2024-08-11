package operations

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetDiffFiles returns three slices containing added, deleted, and modified file names between two commits.
func GetDiffFiles(oldCommitID, newCommitID, repoDir string) ([]string, []string, []string, []string, error) {
	var cmd *exec.Cmd
	if oldCommitID == "" {
		// If oldCommitID is empty, list all files in the repository
		cmd = exec.Command("sh", "-c", "git ls-files -z")
		cmd.Dir = repoDir // Set the working directory to the specified repository directory

		// Run the command and capture its output
		output, err := cmd.Output()
		if err != nil {
			return nil, nil, nil, nil, fmt.Errorf("failed to run git command: %v", err)
		}
		// Split output into fields by null character
		fields := strings.Split(string(output), "\x00")
		if len(fields) > 0 && fields[len(fields)-1] == "" {
			fields = fields[:len(fields)-1] // Remove the last empty element if it exists
		}
		// Initialize slices for added, deleted, and modified files
		var addedFiles []string
		var modifiedFiles []string
		var deletedFiles []string
		var renameFiles []string
		i := 0
		for {
			if i >= len(fields) {
				break
			}
			addedFiles = append(addedFiles, fields[i])
			i += 1
		}
		return addedFiles, modifiedFiles, deletedFiles, renameFiles, nil
	}
	// Construct the git diff command
	cmd = exec.Command("sh", "-c", fmt.Sprintf("git diff --name-status -z %s %s", oldCommitID, newCommitID))

	cmd.Dir = repoDir // Set the working directory to the specified repository directory

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("failed to run git command: %v", err)
	}

	// Split output into fields by null character
	fields := strings.Split(string(output), "\x00")
	if len(fields) > 0 && fields[len(fields)-1] == "" {
		fields = fields[:len(fields)-1] // Remove the last empty element if it exists
	}
	// Initialize slices for added, deleted, and modified files
	var addedFiles []string
	var modifiedFiles []string
	var deletedFiles []string
	var renameFiles []string

	i := 0
	for {
		if i >= len(fields) {
			break
		}
		status := fields[i][0]
		switch status {
		case 'A':
			fileName := fields[i+1]
			addedFiles = append(addedFiles, fileName)
			i += 2
		case 'M':
			fileName := fields[i+1]
			modifiedFiles = append(modifiedFiles, fileName)
			i += 2
		case 'D':
			fileName := fields[i+1]
			deletedFiles = append(deletedFiles, fileName)
			i += 2
		case 'R':
			// Renamed status, next two fields are old and new filenames
			oldName := fields[i+1]
			newName := fields[i+2]
			if filepath.Ext(oldName) == ".md" && filepath.Ext(newName) == ".md" {
				renameFiles = append(renameFiles, oldName)
				renameFiles = append(renameFiles, newName)
			} else {
				deletedFiles = append(deletedFiles, oldName)
				addedFiles = append(addedFiles, newName)
			}
			i += 3 // Skip the next two fields since they are part of the rename
		case 'C':
			// Copied status, next two fields are old and new filenames
			_ = fields[i+1]
			newName := fields[i+2]
			addedFiles = append(addedFiles, newName)
			i += 3 // Skip the next two fields since they are part of the copy
		default:
			// Unsupported status or unhandled status
			i = len(fields)
			fmt.Printf("Unhandled status: %c\n", status)
		}
	}

	return addedFiles, modifiedFiles, deletedFiles, renameFiles, nil
}
