package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"

	"github.com/gruz0/markdown-linter/cmd/markdownlinter"
	"github.com/labstack/gommon/log"
	"github.com/olekukonko/tablewriter"
)

func main() {
	plugins := []string{"fixme"}
	files := getFiles()

	markdownlinter := markdownlinter.NewMarkdownLinter()

	result, err := markdownlinter.Lint(plugins, files)

	if err != nil {
		log.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"File", "Line", "Plugin", "Description"})

	for _, linterResult := range result {
		table.Append([]string{
			linterResult.FileName,
			strconv.Itoa(linterResult.Line),
			linterResult.Plugin,
			linterResult.ErrorDescription})
	}

	table.Render()
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

// if len(errors) == 0 {
// 	os.Exit(0)
// }
//
// for _, e := range errors {
// 	if len(e.errors) == 0 {
// 		continue
// 	}
//
// 	color.Info.Tips("Processing %s on %s", e.linter, e.fileName)
//
// 	var theme *color.Theme
//
// 	switch severity := e.severity; severity {
// 	case WARN:
// 		theme = color.Warn
// 	case ERROR:
// 		theme = color.Error
// 	default:
// 		theme = color.Info
// 	}
//
// 	for _, errorMessage := range e.errors {
// 		theme.Tips(errorMessage)
// 	}
// }
//
// os.Exit(1)
