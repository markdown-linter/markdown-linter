package h1_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/plugins/h1"
	"github.com/stretchr/testify/assert"
)

func Test_Plugin_H1_Load_ReturnsPluginInfo(t *testing.T) {
	plugin := Plugin{}
	info := plugin.Info()

	assert.Equal(t, "MissingH1Tag", info.Name)
	assert.Len(t, info.Contributors, 1)
	assert.Equal(t, "Alexander Kadyrov <alexander@kadyrov.dev>", info.Contributors[0])
	assert.Equal(t, "The document does not have Header tag", info.ErrorDescription)
}

func Test_Plugin_H1_Load_ValidatesErrors(t *testing.T) {
	plugin := Plugin{}

	expected := map[string]string{
		"":        "Empty file could not be linted",
		"# ":      "The first line is too short",
		" Header": "The first line must be started with #",
		"#  ":     "Empty header found",
	}

	for value, description := range expected {
		result := plugin.Lint(value)

		assert.Len(t, result, 1)

		offence := result[0]
		assert.Equal(t, 1, offence.Line)
		assert.Equal(t, description, offence.Description)
	}
}

func Test_Plugin_H1_Load_Valid(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("# Header")

	assert.Len(t, result, 0)
}
