package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFilterDiffFilesByExtensions(t *testing.T) {
	files := []string{
		"README.md",
		"image.png",
		"photo.jpg",
		"document.txt",
		"script.js",
		"style.css",
		"diagram.svg",
		"animation.gif",
		"vector.webp",
		"archive.zip",
	}

	allowedExtensions := map[string]bool{
		".md":   true,
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".svg":  true,
		".gif":  true,
		".webp": true,
	}

	expectedFilteredFiles := []string{
		"README.md",
		"image.png",
		"photo.jpg",
		"diagram.svg",
		"animation.gif",
		"vector.webp",
	}

	filteredFiles := FilterDiffFilesByExtensions(files, allowedExtensions)
	require.Equal(t, expectedFilteredFiles, filteredFiles, "Filtered files do not match expected output")
}
