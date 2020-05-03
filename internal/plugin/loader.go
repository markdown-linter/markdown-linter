package plugin

import (
	"errors"
	"fmt"

	"github.com/markdown-linter/markdown-linter/internal/entity"
	"github.com/markdown-linter/markdown-linter/internal/interfaces"
	"github.com/markdown-linter/markdown-linter/internal/plugins/fixme"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (d *Loader) LoadPlugins(plugins []string) ([]interfaces.Plugin, error) {
	var result []interfaces.Plugin

	if len(plugins) == 0 {
		return result, errors.New("you need to specify plugins")
	}

	parent := entity.Plugin{}

	for _, pluginName := range plugins {
		switch pluginName {
		case "fixme":
			result = append(result, &fixme.Plugin{Plugin: &parent})
		default:
			return result, fmt.Errorf("plugin %s was not found", pluginName)
		}
	}

	return result, nil
}
