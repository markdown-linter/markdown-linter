package h1

import (
	"strings"

	"github.com/markdown-linter/markdown-linter/internal/entity"
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

	if len(content) == 0 {
		return append(result, structs.Offence{Line: 1, Description: "Empty file could not be linted"})
	}

	lines := strings.Split(content, "\n")

	line := strings.TrimSpace(lines[0])

	if len(line) == 0 || line[0:1] != "#" {
		return append(result, structs.Offence{Line: 1, Description: p.Info().ErrorDescription})
	}

	header := strings.TrimSpace(line[1:])

	if len(header) == 0 {
		return append(result, structs.Offence{Line: 1, Description: "Tag found, but value is empty"})
	}

	return result
}
