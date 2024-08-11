package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseRepoConfigFromString(t *testing.T) {
	jsonStr := `{
		"anchors": [
			{
				"name": "Google",
				"icon": "google-icon",
				"url": "https://google.com"
			}
		],
		"layout": [
			{
				"title": "b",
				"relative_path": "b",
				"isdir": true,
				"sublayouts": [
					{
						"title": "c",
						"relative_path": "b/c",
						"isdir": false,
						"sublayouts": null
					},
					{
						"title": "e",
						"relative_path": "b/e",
						"isdir": false,
						"sublayouts": null
					}
				]
			},
			{
				"title": "a",
				"relative_path": "a",
				"isdir": false,
				"sublayouts": null
			}
		],
		"footerSocials": [
			{
				"name": "Discord",
				"icon": "discord",
				"url": "https://discord.com"
			}
		]
	}`

	expectedConfig := &RepoConfig{
		Anchors: []Anchor{
			{
				Name: "Google",
				Icon: "google-icon",
				URL:  "https://google.com",
			},
		},
		Layout: []Layout{
			{
				Title:        "b",
				RelativePath: "b",
				Isdir:        true,
				Sublayouts: []Layout{
					{
						Title:        "c",
						RelativePath: "b/c",
						Isdir:        false,
						Sublayouts:   nil,
					},
					{
						Title:        "e",
						RelativePath: "b/e",
						Isdir:        false,
						Sublayouts:   nil,
					},
				},
			},
			{
				Title:        "a",
				RelativePath: "a",
				Isdir:        false,
				Sublayouts:   nil,
			},
		},
		FooterSocials: []FooterSocial{
			{
				Name: "Discord",
				Icon: "discord",
				URL:  "https://discord.com",
			},
		},
	}

	config, err := ParseRepoConfigFromString(jsonStr)
	require.NoError(t, err)
	require.Equal(t, expectedConfig, config)
}
