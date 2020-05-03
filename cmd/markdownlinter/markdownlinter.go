package markdownlinter

import (
	"io/ioutil"

	"github.com/markdown-linter/markdown-linter/internal/interfaces"
	"github.com/markdown-linter/markdown-linter/internal/plugin"
	"github.com/markdown-linter/markdown-linter/internal/structs"
)

type MarkdownLinter struct{}

func NewMarkdownLinter() *MarkdownLinter {
	return &MarkdownLinter{}
}

func (ml *MarkdownLinter) Lint(plugins []string, files []string) ([]structs.LinterResult, error) {
	var (
		result        []structs.LinterResult
		loadedPlugins []interfaces.Plugin
		err           error
	)

	loader := plugin.NewLoader()

	if loadedPlugins, err = loader.LoadPlugins(plugins); err != nil {
		return result, err
	}

	for _, fileName := range files {
		bytes, err := ioutil.ReadFile(fileName)

		if err != nil {
			return result, err
		}

		content := string(bytes)

		for _, plugin := range loadedPlugins {
			info := plugin.Info()
			errors := plugin.Lint(content)

			for _, offence := range errors {
				lr := structs.LinterResult{
					FileName:         fileName,
					Line:             offence.Line,
					Plugin:           info.Name,
					ErrorDescription: offence.Description,
				}

				result = append(result, lr)
			}
		}
	}

	return result, nil
}
