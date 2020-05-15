package parsers_test

import "github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"

type TestData struct {
	in  string
	out []structs.Tag
}

func NoTags() []structs.Tag {
	return []structs.Tag{}
}
