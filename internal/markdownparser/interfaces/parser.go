package interfaces

import "github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"

type Parser interface {
	Parse() []structs.Tags
}
