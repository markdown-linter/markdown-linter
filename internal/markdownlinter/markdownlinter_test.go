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

func TestLint_ReturnsTotalErrorsCount(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{}
	files := []string{"../../testdata/errors.md"}

	result, err := ml.Lint(plugins, files)

	assert.NoError(t, err)

	// NOTE: This line has all issues found by plugins in testdata/errors.md
	assert.Len(t, result, 6)
}

func TestLint_Fixme_Plugin(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"fixme"}
	files := []string{"../../testdata/errors.md"}

	linterResult := structs.LinterResult{
		FileName:         "../../testdata/errors.md",
		Line:             3,
		Plugin:           "FixmeTag",
		ErrorDescription: "The line has FIXME tag",
	}

	result, err := ml.Lint(plugins, files)

	assert.NoError(t, err)
	assert.Len(t, result, 1)
	assert.Equal(t, linterResult, result[0])
}

func TestLint_H1_Plugin(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"h1"}
	files := []string{"../../testdata/errors.md"}

	linterResult := []structs.LinterResult{
		{
			FileName:         "../../testdata/errors.md",
			Line:             1,
			Plugin:           "H1Tag",
			ErrorDescription: "The document does not have H1 tag on the 1st line",
		},
		{
			FileName:         "../../testdata/errors.md",
			Line:             15,
			Plugin:           "H1Tag",
			ErrorDescription: "The header must be located at the beginning of the document",
		},
	}

	result, err := ml.Lint(plugins, files)

	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Contains(t, result, linterResult[0])
	assert.Contains(t, result, linterResult[1])
}

func TestLint_BrokenLinks_Plugin(t *testing.T) {
	ml := NewMarkdownLinter()

	plugins := []string{"brokenlinks"}
	files := []string{"../../testdata/errors.md"}

	linterResult := []structs.LinterResult{
		{
			FileName: "../../testdata/errors.md",
			Line:     7,
			Plugin:   "BrokenLinks",
			ErrorDescription: "Broken link \"http://domain.tld/\" found in \"[Broken Link Here](http://domain.tld/)\". " +
				"Error: No such host",
		},
		{
			FileName:         "../../testdata/errors.md",
			Line:             9,
			Plugin:           "BrokenLinks",
			ErrorDescription: "Empty link found in \"[Empty Link]()\"",
		},
		{
			FileName:         "../../testdata/errors.md",
			Line:             11,
			Plugin:           "BrokenLinks",
			ErrorDescription: "Empty internal link found in \"[Link without anchor](#)\"",
		},
	}

	result, err := ml.Lint(plugins, files)

	assert.NoError(t, err)
	assert.Len(t, result, 3)

	assert.Contains(t, result, linterResult[0])
	assert.Contains(t, result, linterResult[1])
	assert.Contains(t, result, linterResult[2])
}
