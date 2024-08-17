package util

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadRepoConfig(t *testing.T) {
	// 创建一个临时文件并写入测试 JSON 数据
	testJSON := `{
		"anchors": [
			{
				"name": "Google",
				"icon": "google-icon",
				"url": "https://google.com"
			}
		],
		"layout": {
			"en": [
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
			]
		},
		"footerSocials": [
			{
				"name": "Discord",
				"icon": "discord",
				"url": "https://discord.com"
			}
		]
	}`

	tmpFile, err := os.CreateTemp("", "test*.json")
	require.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(testJSON))
	require.NoError(t, err)

	err = tmpFile.Close()
	require.NoError(t, err)

	// 读取 JSON 文件并进行断言
	config, err := ReadRepoConfig(tmpFile.Name())
	require.NoError(t, err)

	expectedConfig := &RepoConfig{
		Anchors: []Anchor{
			{
				Name: "Google",
				Icon: "google-icon",
				URL:  "https://google.com",
			},
		},
		Layout: map[string][]Layout{
			"en": {
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
		},
		FooterSocials: []FooterSocial{
			{
				Name: "Discord",
				Icon: "discord",
				URL:  "https://discord.com",
			},
		},
	}

	require.Equal(t, expectedConfig, config)
}
