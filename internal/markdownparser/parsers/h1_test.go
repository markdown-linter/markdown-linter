package parsers_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/markdownparser/parsers"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
	"github.com/stretchr/testify/assert"
)

func Test_MarkdownParser_H1Parser_Parse(t *testing.T) {
	p := H1Parser{}

	expected := []TestData{
		{in: "", out: NoTags()},
		{in: "\n", out: NoTags()},
		{in: "# H1", out: []structs.Tag{
			{Line: 1, Content: "# H1"},
		}},
		{in: "   #     H1", out: []structs.Tag{
			{Line: 1, Content: "   #     H1"},
		}},
		{in: "   #  H1\n# H1\n\n## H2", out: []structs.Tag{
			{Line: 1, Content: "   #  H1"},
			{Line: 2, Content: "# H1"},
		}},
	}

	for _, e := range expected {
		tags := p.Parse(e.in)

		assert.Equal(t, e.out, tags)
	}
}
