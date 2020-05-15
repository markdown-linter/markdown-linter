package markdownparser

import (
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/parsers"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
)

type MarkdownParser struct{}

func NewMarkdownParser() *MarkdownParser {
	return &MarkdownParser{}
}

func (mp *MarkdownParser) Parse(content string) structs.Tags {
	tags := structs.Tags{}

	tags.H1 = mp.parseH1Tags(content)
	tags.Links = mp.parseLinks(content)

	return tags
}

func (mp *MarkdownParser) parseH1Tags(content string) []structs.Tag {
	parser := parsers.H1Parser{}

	return parser.Parse(content)
}

func (mp *MarkdownParser) parseLinks(content string) []structs.Tag {
	parser := parsers.LinksParser{}

	return parser.Parse(content)
}
