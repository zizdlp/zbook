package util

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadLayout(t *testing.T) {
	// Define the expected layout structure
	expectedLayout := Layout{
		Anchors: []Anchor{
			{
				Name: "Discord",
				Icon: "discord",
				URL:  "https://loopholelabs.io/discord",
			},
			{
				Name: "GitHub",
				Icon: "github",
				URL:  "https://github.com/loopholelabs/frpc-go",
			},
		},
		Navigation: []Nav{
			{
				Group: "Welcome",
				Pages: []string{"introduction"},
				Subgroup: []Nav{
					{
						Group: "Welcome",
						Pages: []string{"introduction"},
					},
					{
						Group: "Welcome",
						Pages: []string{"introduction"},
					},
				},
			},
			{
				Group: "Getting Started",
				Pages: []string{
					"getting-started/overview",
					"getting-started/quick-start",
					"getting-started/concepts",
					"getting-started/architecture",
					"getting-started/roadmap",
				},
			},
			{
				Group: "Performance",
				Pages: []string{
					"performance/optimizations",
					"performance/grpc-benchmarks",
					"performance/twirp-benchmarks",
				},
			},
			{
				Group: "Reference",
				Pages: []string{
					"reference/overview",
					"reference/client-methods",
					"reference/server-methods",
				},
			},
		},
		FooterSocials: Socials{
			Discord: "https://loopholelabs.io/discord",
			Github:  "https://github.com/loopholelabs/frpc-go",
		},
	}

	// Create a temporary JSON file for testing
	fileContent, err := json.Marshal(expectedLayout)
	require.NoError(t, err)

	tempFile, err := os.CreateTemp("", "layout*.json")
	require.NoError(t, err)
	defer os.Remove(tempFile.Name())

	_, err = tempFile.Write(fileContent)
	require.NoError(t, err)
	tempFile.Close()

	// Load the layout using the LoadLayout function
	loadedLayout, err := LoadLayout(tempFile.Name())
	require.NoError(t, err)

	// Compare the loaded layout with the expected layout
	require.Equal(t, expectedLayout, loadedLayout, "Loaded layout does not match the expected layout")
}
