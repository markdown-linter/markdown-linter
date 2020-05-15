package parsers_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/markdownparser/parsers"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
	"github.com/stretchr/testify/assert"
)

func Test_MarkdownParser_LinksParser_Parse(t *testing.T) {
	p := LinksParser{}

	expected := []TestData{
		{in: "", out: NoTags()},
		{in: "\n", out: NoTags()},
		{in: "![Image](/my/image.png)", out: NoTags()},
		{in: "[]()", out: []structs.Tag{
			{Line: 1, Content: "[]()"},
		}},
		{in: "[no link]()", out: []structs.Tag{
			{Line: 1, Content: "[no link]()"},
		}},
		{in: "[internal link](#)", out: []structs.Tag{
			{Line: 1, Content: "[internal link](#)"},
		}},
		{in: "[first link](http://domain.tld) [second link](https://domain.tld/)", out: []structs.Tag{
			{Line: 1, Content: "[first link](http://domain.tld)"},
			{Line: 1, Content: "[second link](https://domain.tld/)"},
		}},
		{in: "[first link](http://domain.tld)\n\n[second link](https://domain.tld/)", out: []structs.Tag{
			{Line: 1, Content: "[first link](http://domain.tld)"},
			{Line: 3, Content: "[second link](https://domain.tld/)"},
		}},
		{in: "[![Alt](http://link.to/image/here.png)](http://domain.tld)", out: []structs.Tag{
			{Line: 1, Content: "[![Alt](http://link.to/image/here.png)](http://domain.tld)"},
		}},
	}

	for _, e := range expected {
		tags := p.Parse(e.in)

		assert.Equal(t, e.out, tags)
	}
}
