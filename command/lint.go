package command

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/markdown-linter/markdown-linter/internal/markdownlinter"
	"github.com/olekukonko/tablewriter"
)

type LintCommand struct{}

type lintFlags struct {
	dirs          stringList
	recursiveDirs stringList
	files         stringList
}

func (c *LintCommand) Run(args []string) int {
	var err error

	flags := c.parseFlags(args)

	files, err := c.populateFiles(flags)

	if err != nil {
		log.Printf("Error while getting files: %s", err.Error())

		return 1
	}

	if len(files) == 0 {
		log.Printf("No files found")

		return 1
	}

	plugins := []string{}

	markdownlinter := markdownlinter.NewMarkdownLinter()

	lintErrors, err := markdownlinter.Lint(plugins, files)

	if err != nil {
		log.Fatalf("Error during linting: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	if len(lintErrors) == 0 {
		return 0
	}

	table.SetAutoWrapText(false)
	table.SetHeader([]string{"File", "Line", "Plugin", "Description"})

	for _, linterResult := range lintErrors {
		table.Append([]string{
			linterResult.FileName,
			strconv.Itoa(linterResult.Line),
			linterResult.Plugin,
			linterResult.ErrorDescription})
	}

	table.Render()

	return 1
}

func (c *LintCommand) Help() string {
	helpText := `
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

	return strings.TrimSpace(helpText)
}

func (c *LintCommand) Synopsis() string {
	return "Lints .md files"
}

func (c *LintCommand) parseFlags(args []string) lintFlags {
	var flags lintFlags

	flagSet := parseFlags("lint", args, func(f *flag.FlagSet) {
		f.Var(&flags.dirs, "D", "Search `directory` for .md files")
		f.Var(&flags.recursiveDirs, "R", "Recursive search `directory` for .md files")
		f.Var(&flags.files, "f", "Lint single `file`")
	})

	if len(args) == 0 {
		log.Println("Usage of lint:")
		flagSet.PrintDefaults()
	}

	return flags
}

func (c *LintCommand) populateFiles(flags lintFlags) ([]string, error) {
	var (
		result []string
		err    error
	)

	files := make([]string, 0)

	// dirs
	if result, err = c.getFiles(flags.dirs); err != nil {
		return files, err
	}

	files = append(files, result...)

	// recursive walk
	if result, err = c.getFilesRecursively(flags.recursiveDirs); err != nil {
		return files, err
	}

	files = append(files, result...)

	// files
	for _, path := range flags.files {
		if err := c.checkFile(path); err != nil {
			return files, err
		}

		files = append(files, path)
	}

	files = unique(files)

	return files, nil
}

func (c *LintCommand) getFiles(dirs []string) ([]string, error) {
	files := make([]string, 0)

	for _, path := range dirs {
		result, err := GetMarkdownFilesInDirectory(path)

		if err != nil {
			return files, err
		}

		files = append(files, result...)
	}

	return files, nil
}

func (c *LintCommand) getFilesRecursively(dirs []string) ([]string, error) {
	files := make([]string, 0)

	for _, path := range dirs {
		result, err := GetMarkdownFilesInDirectoryRecursively(path)

		if err != nil {
			return files, err
		}

		files = append(files, result...)
	}

	return files, nil
}

func (c *LintCommand) checkFile(path string) error {
	stat, err := os.Stat(path)

	if err != nil && os.IsNotExist(err) {
		return fmt.Errorf("path %q does not exist", path)
	}

	if !stat.Mode().IsRegular() {
		return fmt.Errorf("%q is not a regular file", path)
	}

	return nil
}

func unique(slice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true

			list = append(list, entry)
		}
	}

	return list
}
