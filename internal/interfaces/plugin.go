package interfaces

import "github.com/markdown-linter/markdown-linter/internal/structs"

type Plugin interface {
	Info() *structs.PluginInfo
	Lint(content string) []structs.Offence
}
