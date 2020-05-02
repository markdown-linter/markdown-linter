package entity

import (
	"strings"

	"github.com/gruz0/markdown-linter/internal/interfaces"
	"github.com/gruz0/markdown-linter/internal/structs"
)

type Plugin struct {
	interfaces.Plugin
}

func (p Plugin) LintByLine(content, errorDescription string, f func(string) bool) []structs.Offence {
	lines := strings.Split(content, "\n")
	result := make([]structs.Offence, 0)

	for idx, line := range lines {
		if len(line) == 0 || f(line) {
			continue
		}

		result = append(result, structs.Offence{Line: idx + 1, Description: errorDescription})
	}

	return result
}
