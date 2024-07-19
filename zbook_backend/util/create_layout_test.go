package util_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/zizdlp/zbook/util"
)

func TestCreateLayout(t *testing.T) {
	files := []string{
		"a.md",
		"b/c.md",
		"b/e.md",
	}

	expectedJSON := `{
		"title": "wiki",
		"relative_path": "",
		"isdir": true,
		"sublayouts": [
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
	}`

	layout := util.CreateLayout(files)

	layoutJSON, err := json.MarshalIndent(layout, "", "  ")
	require.NoError(t, err)

	require.JSONEq(t, expectedJSON, string(layoutJSON))
}
