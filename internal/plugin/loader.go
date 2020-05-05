package plugin

import (
	"fmt"

	"github.com/markdown-linter/markdown-linter/internal/interfaces"
	"github.com/markdown-linter/markdown-linter/internal/plugins/fixme"
	"github.com/markdown-linter/markdown-linter/internal/plugins/h1"
)

type Loader struct{}

func NewLoader() *Loader {
	return &Loader{}
}

func (d *Loader) LoadPlugins(plugins []string) ([]interfaces.Plugin, error) {
	var result []interfaces.Plugin

	pluginsMap := d.pluginsMap()

	if len(plugins) == 0 {
		for pluginName := range pluginsMap {
			plugins = append(plugins, pluginName)
		}
	}

	for _, pluginName := range plugins {
		if plugin, ok := pluginsMap[pluginName]; ok {
			result = append(result, plugin)

			continue
		}

		return result, fmt.Errorf("plugin %q was not found", pluginName)
	}

	return result, nil
}

func (d *Loader) pluginsMap() map[string]interfaces.Plugin {
	return map[string]interfaces.Plugin{
		"fixme": &fixme.Plugin{},
		"h1":    &h1.Plugin{},
	}
}
