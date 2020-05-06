package h1_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/plugins/h1"
	"github.com/stretchr/testify/assert"
)

func Test_Plugin_H1_Load_ReturnsPluginInfo(t *testing.T) {
	plugin := Plugin{}
	info := plugin.Info()

	assert.Equal(t, "H1Tag", info.Name)
	assert.Len(t, info.Contributors, 1)
	assert.Equal(t, "Alexander Kadyrov <alexander@kadyrov.dev>", info.Contributors[0])
	assert.Equal(t, "The document does not have valid H1 tag", info.ErrorDescription)
}

type testData struct {
	in  string
	out string
}

func Test_Plugin_H1_Load_ValidatesErrors(t *testing.T) {
	plugin := Plugin{}

	expected := []testData{
		{in: "", out: "Empty file could not be linted"},
		{in: " Header", out: "The document does not have valid H1 tag"},
		{in: "\n# Header", out: "The document does not have valid H1 tag"},
		{in: "  #  ", out: "Tag found, but value is empty"},
	}

	for _, e := range expected {
		result := plugin.Lint(e.in)

		assert.Len(t, result, 1)
		assert.Equal(t, 1, result[0].Line)
		assert.Equal(t, e.out, result[0].Description)
	}
}

func Test_Plugin_H1_Load_Valid(t *testing.T) {
	plugin := Plugin{}

	headers := []string{"# Header", "#    Header", " # Header"}

	for _, in := range headers {
		result := plugin.Lint(in)

		assert.Len(t, result, 0)
	}
}
