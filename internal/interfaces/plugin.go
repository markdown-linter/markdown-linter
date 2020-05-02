package interfaces

import "github.com/gruz0/markdown-linter/internal/structs"

type Plugin interface {
	Info() *structs.PluginInfo
	Lint(content string) []structs.Offence
}
