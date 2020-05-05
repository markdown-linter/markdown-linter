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
		Name:             "MissingH1Tag",
		Contributors:     []string{"Alexander Kadyrov <alexander@kadyrov.dev>"},
		ErrorDescription: "The document does not have Header tag",
	}
}

func (p *Plugin) Lint(content string) []structs.Offence {
	result := make([]structs.Offence, 0)

	if len(content) == 0 {
		return append(result, structs.Offence{Line: 1, Description: "Empty file could not be linted"})
	}

	lines := strings.Split(content, "\n")

	firstLine := lines[0]

	if len(firstLine) < 3 {
		return append(result, structs.Offence{Line: 1, Description: "The first line is too short"})
	}

	if firstLine[0:2] != "# " {
		return append(result, structs.Offence{Line: 1, Description: "The first line must be started with #"})
	}

	header := strings.TrimSpace(firstLine[3:])

	if len(header) == 0 {
		return append(result, structs.Offence{Line: 1, Description: "Empty header found"})
	}

	return result
}
