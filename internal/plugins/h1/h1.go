package h1

import (
	"strings"

	"github.com/markdown-linter/markdown-linter/internal/entity"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser"
	"github.com/markdown-linter/markdown-linter/internal/structs"
)

type Plugin struct {
	entity.Plugin
}

func (p *Plugin) Info() *structs.PluginInfo {
	return &structs.PluginInfo{
		Name:             "H1Tag",
		Contributors:     []string{"Alexander Kadyrov <alexander@kadyrov.dev>"},
		ErrorDescription: "The document does not have valid H1 tag",
	}
}

func (p *Plugin) Lint(content string) []structs.Offence {
	result := make([]structs.Offence, 0)

	if !isHeaderExistsOnTheFirstLine(content) {
		result = append(result, buildError(1, "The document does not have H1 tag on the 1st line"))
	}

	tags := markdownparser.NewMarkdownParser().Parse(content)

	for _, tag := range tags.H1 {
		if tag.Line > 1 {
			if len(tags.H1) > 1 {
				result = append(result, buildError(tag.Line, "The document cannot contain more than one H1 tag"))
			} else {
				result = append(result, buildError(tag.Line, "The header must be located at the beginning of the document"))
			}
		}

		if len(tag.Content) == 0 || strings.TrimSpace(tag.Content)[0:1] != "#" {
			result = append(result, buildError(tag.Line, "Tag found, but value is empty"))
		}
	}

	return result
}

func isHeaderExistsOnTheFirstLine(content string) bool {
	lines := strings.Split(content, "\n")
	firstLine := strings.TrimSpace(lines[0])

	return len(firstLine) > 2 && firstLine[0:2] == "# "
}

func buildError(line int, description string) structs.Offence {
	return structs.Offence{Line: line, Description: description}
}
