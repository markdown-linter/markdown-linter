package markdownlinter_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/markdownlinter"
	"github.com/markdown-linter/markdown-linter/internal/structs"
	"github.com/stretchr/testify/assert"
)

func TestLint_ReturnsErrorIfPluginDoesNotExist(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"not-found"}
	files := []string{"../../testdata/test.md"}

	result, err := ml.Lint(plugins, files)

	assert.EqualError(t, err, "plugin \"not-found\" was not found")
	assert.Len(t, result, 0)
}

func TestLint_ReturnsErrorIfFileDoesNotExist(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{}
	files := []string{"../../testdata/not-found.md"}

	result, err := ml.Lint(plugins, files)

	assert.EqualError(t, err, "open ../../testdata/not-found.md: no such file or directory")
	assert.Len(t, result, 0)
}

func TestLint_FixmeTag(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"fixme"}
	files := []string{"../../testdata/test.md"}

	linterResult := structs.LinterResult{
		FileName:         "../../testdata/test.md",
		Line:             3,
		Plugin:           "FixmeTag",
		ErrorDescription: "The line has FIXME tag",
	}

	result, err := ml.Lint(plugins, files)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, linterResult, result[0])
}
