package command_test

import (
	"strings"
	"testing"

	. "github.com/markdown-linter/markdown-linter/command"
	"github.com/stretchr/testify/assert"
)

func TestLintHelp(t *testing.T) {
	c := &LintCommand{}

	help := c.Help()

	expected := `
Usage: markdown-linter lint [options]

Options:

  -D directory     Search and lint Markdown files in directory
  -R directory     Recursive search and lint Markdown files in directory
  -f file          Lint single file

Examples:

  markdown-linter lint -R . -R ~/Downloads
  markdown-linter lint -D ~/Documents
  markdown-linter lint -f README.md

Or combine all options:

  markdown-linter lint -R ~/Downloads -D ~/Documents -f README.md
`

	assert.Equal(t, strings.TrimSpace(expected), help)
}

func TestLintSynopsis(t *testing.T) {
	c := &LintCommand{}

	synopsis := c.Synopsis()

	assert.Equal(t, "Lints .md files", synopsis)
}

func TestLint_singleDirectory_doesNotExist(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-D", "../not-exist"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_singleDirectory_isNotADirectory(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-D", "../testdata/test.md"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_singleDirectory_valid(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-D", "../testdata/valid"})

	assert.Equal(t, 0, exitCode)
}

func TestLint_recursiveDirectory_doesNotExist(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-R", "../not-exist"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_recursiveDirectory_isNotADirectory(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-R", "../testdata/test.md"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_recursiveDirectory_valid(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-R", "../testdata/valid"})

	assert.Equal(t, 0, exitCode)
}

func TestLint_singleFile_doesNotExist(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-f", "../not-exist.md"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_singleFile_isNotAFile(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-f", "../testdata"})

	assert.Equal(t, 1, exitCode)
}

func TestLint_singleFile_valid(t *testing.T) {
	c := &LintCommand{}

	exitCode := c.Run([]string{"-f", "../testdata/valid/header-one.md"})

	assert.Equal(t, 0, exitCode)
}
