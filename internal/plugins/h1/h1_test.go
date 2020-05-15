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

func Test_Plugin_H1_Lint_ReturnsErrorIfHeaderDoesNotExistOnTheFirstLine(t *testing.T) {
	plugin := Plugin{}

	headers := []string{"", " Header", "  #  "}

	for _, in := range headers {
		result := plugin.Lint(in)

		assert.Len(t, result, 1)
		assert.Equal(t, 1, result[0].Line)
		assert.Equal(t, "The document does not have H1 tag on the 1st line", result[0].Description)
	}
}

func Test_Plugin_H1_Lint_ReturnsErrorIfHeaderLocatedNotOnTheFirstLine(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("\n# Header")

	assert.Len(t, result, 2)
	assert.Equal(t, 2, result[1].Line)
	assert.Equal(t, "The header must be located at the beginning of the document", result[1].Description)
}

func Test_Plugin_H1_Lint_ReturnsErrorIfDocumentHasMoreThanOneH1Tag(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("# Header\n# Another Header")

	assert.Len(t, result, 1)
	assert.Equal(t, 2, result[0].Line)
	assert.Equal(t, "The document cannot contain more than one H1 tag", result[0].Description)
}

func Test_Plugin_H1_Lint_Valid(t *testing.T) {
	plugin := Plugin{}

	headers := []string{"# Header", "#    Header", " # Header"}

	for _, in := range headers {
		result := plugin.Lint(in)

		assert.Len(t, result, 0)
	}
}
