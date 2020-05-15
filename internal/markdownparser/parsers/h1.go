package parsers

import (
	"strings"

	"github.com/markdown-linter/markdown-linter/internal/markdownparser/interfaces"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
)

type H1Parser struct {
	interfaces.Parser
}

func (p *H1Parser) Parse(content string) []structs.Tag {
	tags := make([]structs.Tag, 0)

	lines := strings.Split(content, "\n")

	for idx, line := range lines {
		trimmedLine := strings.TrimSpace(line)

		if len(trimmedLine) < 3 {
			continue
		}

		if trimmedLine[0:2] != "# " {
			continue
		}

		tags = append(tags, structs.Tag{Line: idx + 1, Content: line})
	}

	return tags
}
