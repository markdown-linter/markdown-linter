package markdownlinter_test

import (
	"testing"

	. "github.com/gruz0/markdown-linter/cmd/markdownlinter"
	"github.com/gruz0/markdown-linter/internal/structs"
	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorIfPluginIsNotExists(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"not-found"}
	files := []string{"../../testdata/test.md"}

	result, err := ml.Lint(plugins, files)

	assert.NotNil(t, err)
	assert.Equal(t, "plugin not-found was not found", err.Error())
	assert.Len(t, result, 0)
}

func TestReturnsErrorIfFileIsNotExists(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"fixme"}
	files := []string{"../../testdata/not-found.md"}

	result, err := ml.Lint(plugins, files)

	assert.NotNil(t, err)
	assert.Equal(t, "open ../../testdata/not-found.md: no such file or directory", err.Error())
	assert.Len(t, result, 0)
}

func TestFixmeTag(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"fixme"}
	files := []string{"../../testdata/test.md"}

	result, err := ml.Lint(plugins, files)

	assert.Len(t, result, 1)

	linterResult := structs.LinterResult{
		FileName:         "../../testdata/test.md",
		Line:             3,
		Plugin:           "FixmeTag",
		ErrorDescription: "The line has FIXME tag",
	}
	assert.Equal(t, linterResult, result[0])

	assert.Nil(t, err)
}
