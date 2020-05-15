package command_test

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	. "github.com/markdown-linter/markdown-linter/command"
	"github.com/markdown-linter/markdown-linter/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestGenerateConfigHelp(t *testing.T) {
	c := &GenerateConfigCommand{}

	help := c.Help()

	expected := `
Usage: markdown-linter generate-config [options]

Options:

  -config string   Path to config file (default ".markdown-linter.yml")
  -force           Use to overwrite existed config

Examples:

  markdown-linter generate-config -config ~/.markdown-linter.yml -force
`

	assert.Equal(t, strings.TrimSpace(expected), help)
}

func TestGenerateConfigSynopsis(t *testing.T) {
	c := &GenerateConfigCommand{}

	synopsis := c.Synopsis()

	assert.Equal(t, "Generates configuration file from the template", synopsis)
}

func TestGenerateConfig_shouldCreateDefaultConfigIfFileDoesNotExist(t *testing.T) {
	defaultConfigPath := ".markdown-linter.yml"

	defer os.Remove(defaultConfigPath)

	c := &GenerateConfigCommand{}

	exitCode := c.Run([]string{})

	assert.Zero(t, exitCode)
	assert.FileExists(t, ".markdown-linter.yml")
}

func TestGenerateConfig_shouldReturnErrorIfConfigIsSetToDirectory(t *testing.T) {
	tempdir := tempDir(t)

	defer removeTempDir(tempdir)

	args := []string{"-config", tempdir}

	c := &GenerateConfigCommand{}

	exitCode := c.Run(args)

	assert.Equal(t, 1, exitCode)
}

func TestGenerateConfig_shouldReturnErrorIfConfigExistsAndForceIsNotUsed(t *testing.T) {
	tempfile := tempFile(t)
	defer removeTempFile(tempfile)

	args := []string{"-config", tempfile.Name()}

	c := &GenerateConfigCommand{}

	exitCode := c.Run(args)

	assert.Equal(t, 1, exitCode)
}

func TestGenerateConfig_shouldOverwriteConfigIfConfigExistsAndForceIsUsed(t *testing.T) {
	tempfile := tempFile(t)
	defer removeTempFile(tempfile)

	assert.FileExists(t, tempfile.Name())

	args := []string{"-config", tempfile.Name(), "-force"}

	c := &GenerateConfigCommand{}

	exitCode := c.Run(args)

	assert.Equal(t, 0, exitCode)
}

func TestGenerateConfig_hasValidDefaultConfigContent(t *testing.T) {
	defaultConfigPath := ".markdown-linter.yml"

	defer os.Remove(defaultConfigPath)

	c := &GenerateConfigCommand{}

	_ = c.Run([]string{})

	content, err := utils.ReadFile(defaultConfigPath)

	if err != nil {
		t.Fatal(err)
	}

	assert.YAMLEq(t, "plugins: []", content)
}

func tempFile(t *testing.T) *os.File {
	tempfile, err := ioutil.TempFile("", ".markdown-linter.*.test")

	if err != nil {
		t.Fatal(err)
	}

	return tempfile
}

func removeTempFile(tempfile *os.File) {
	os.Remove(tempfile.Name())
}

func tempDir(t *testing.T) string {
	tempdir, err := ioutil.TempDir("", "prefix")

	if err != nil {
		t.Fatal(err)
	}

	return tempdir
}

func removeTempDir(dir string) {
	os.RemoveAll(dir)
}
