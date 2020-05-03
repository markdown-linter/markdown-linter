package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/markdown-linter/markdown-linter/cmd/markdownlinter"
	"github.com/olekukonko/tablewriter"
)

func main() {
	plugins := []string{"fixme"}
	files := getFiles()

	markdownlinter := markdownlinter.NewMarkdownLinter()

	result, err := markdownlinter.Lint(plugins, files)

	if err != nil {
		log.Fatalf("Error during linting: %v", err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	if len(result) == 0 {
		os.Exit(0)
	}

	table.SetHeader([]string{"File", "Line", "Plugin", "Description"})

	for _, linterResult := range result {
		table.Append([]string{
			linterResult.FileName,
			strconv.Itoa(linterResult.Line),
			linterResult.Plugin,
			linterResult.ErrorDescription})
	}

	table.Render()

	os.Exit(1)
}

func getFiles() []string {
	files := make([]string, 0)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	sort.Strings(files)

	return files
}
