package fixme

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
		Name:             "FixmeTag",
		Contributors:     []string{"Alexander Kadyrov <alexander@kadyrov.dev>"},
		ErrorDescription: "The line has FIXME tag",
	}
}

func (p *Plugin) Lint(content string) []structs.Offence {
	info := p.Info()

	return p.Plugin.LintByLine(content, info.ErrorDescription, func(line string) bool {
		return !strings.Contains(strings.ToUpper(line), "FIXME:")
	})
}
