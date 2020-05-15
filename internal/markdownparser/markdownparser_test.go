package markdownparser_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/markdownparser"
	"github.com/markdown-linter/markdown-linter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestMarkdownParser_Parse_H1_ReturnsInvalidHeaderLocation(t *testing.T) {
	p := NewMarkdownParser()

	content := getErrorsFileContent(t)

	tags := p.Parse(content)

	assert.Len(t, tags.H1, 1)
	assert.Equal(t, 15, tags.H1[0].Line)
	assert.Equal(t, "# Invalid Header Location", tags.H1[0].Content)
}

func TestMarkdownParser_Parse_H1_ReturnsOneTag(t *testing.T) {
	p := NewMarkdownParser()

	content := getValidFileContent(t)

	tags := p.Parse(content)

	assert.Len(t, tags.H1, 1)
	assert.Equal(t, 1, tags.H1[0].Line)
	assert.Equal(t, "# Header One", tags.H1[0].Content)
}

func TestMarkdownParser_Parse_Links(t *testing.T) {
	p := NewMarkdownParser()

	content := getErrorsFileContent(t)

	tags := p.Parse(content)

	assert.Len(t, tags.Links, 3)

	assert.Equal(t, 7, tags.Links[0].Line)
	assert.Equal(t, "[Broken Link Here](http://domain.tld/)", tags.Links[0].Content)

	assert.Equal(t, 9, tags.Links[1].Line)
	assert.Equal(t, "[Empty Link]()", tags.Links[1].Content)

	assert.Equal(t, 11, tags.Links[2].Line)
	assert.Equal(t, "[Link without anchor](#)", tags.Links[2].Content)
}

func TestMarkdownParser_Parse_Links_ReturnsOneLink(t *testing.T) {
	p := NewMarkdownParser()

	content := getValidFileContent(t)

	tags := p.Parse(content)

	assert.Len(t, tags.Links, 1)
	assert.Equal(t, 3, tags.Links[0].Line)
	assert.Equal(t, "[My Blog](https://kadyrov.dev/)", tags.Links[0].Content)
}

func getErrorsFileContent(t *testing.T) string {
	content, err := utils.ReadFile("../../testdata/errors.md")

	if err != nil {
		t.Fatal(err)
	}

	return content
}

func getValidFileContent(t *testing.T) string {
	content, err := utils.ReadFile("../../testdata/valid/valid.md")

	if err != nil {
		t.Fatal(err)
	}

	return content
}
