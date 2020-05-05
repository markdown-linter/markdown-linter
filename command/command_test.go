package command_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/command"
	"github.com/stretchr/testify/assert"
)

func TestGetMarkdownFilesInDirectory_shouldReturnErrorIfDirectoryDoesNotExist(t *testing.T) {
	files, err := GetMarkdownFilesInDirectory("../not-exist")

	assert.EqualError(t, err, "directory \"../not-exist\" does not exist")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectory_shouldReturnErrorIfItIsNotADirectory(t *testing.T) {
	files, err := GetMarkdownFilesInDirectory("../testdata/test.md")

	assert.EqualError(t, err, "\"../testdata/test.md\" is not a directory")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectory_shouldReturnErrorIfPathContainsEndingSpaces(t *testing.T) {
	files, err := GetMarkdownFilesInDirectory("../testdata    ")

	assert.EqualError(t, err, "directory \"../testdata    \" does not exist")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectory_shouldReturnOnlyMarkdownFiles(t *testing.T) {
	files, err := GetMarkdownFilesInDirectory("../testdata")

	assert.NoError(t, err)
	assert.Len(t, files, 2)
	assert.Contains(t, files, "../testdata/test.md")
	assert.Contains(t, files, "../testdata/another.md")
}

func TestGetMarkdownFilesInDirectoryRecursively_shouldReturnErrorIfDirectoryDoesNotExist(t *testing.T) {
	files, err := GetMarkdownFilesInDirectoryRecursively("../not-exist")

	assert.EqualError(t, err, "directory \"../not-exist\" does not exist")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectoryRecursively_shouldReturnErrorIfItIsNotADirectory(t *testing.T) {
	files, err := GetMarkdownFilesInDirectoryRecursively("../testdata/test.md")

	assert.EqualError(t, err, "\"../testdata/test.md\" is not a directory")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectoryRecursively_shouldReturnErrorIfPathContainsEndingSpaces(t *testing.T) {
	files, err := GetMarkdownFilesInDirectoryRecursively("../testdata    ")

	assert.EqualError(t, err, "directory \"../testdata    \" does not exist")
	assert.Empty(t, files)
}

func TestGetMarkdownFilesInDirectoryRecursively_shouldReturnOnlyMarkdownFiles(t *testing.T) {
	files, err := GetMarkdownFilesInDirectoryRecursively("../testdata")

	assert.NoError(t, err)
	assert.Len(t, files, 5)
	assert.Contains(t, files, "../testdata/test.md")
	assert.Contains(t, files, "../testdata/another.md")
	assert.Contains(t, files, "../testdata/another/test.md")
	assert.Contains(t, files, "../testdata/markdown.md/test.md")
	assert.Contains(t, files, "../testdata/valid/header-one.md")
}
