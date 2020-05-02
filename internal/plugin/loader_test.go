package plugin_test

import (
	"testing"

	"github.com/gruz0/markdown-linter/internal/plugin"
	"github.com/stretchr/testify/assert"
)

func TestReturnsErrorIfNoPluginsSpecified(t *testing.T) {
	loader := plugin.NewLoader()

	result, err := loader.LoadPlugins([]string{})

	assert.Equal(t, "you need to specify plugins", err.Error())
	assert.Len(t, result, 0)
}

func TestReturnsErrorIfPluginIsNotExists(t *testing.T) {
	loader := plugin.NewLoader()

	result, err := loader.LoadPlugins([]string{"test"})

	assert.Equal(t, "plugin test was not found", err.Error())
	assert.Len(t, result, 0)
}

func TestReturnsFixmePlugin(t *testing.T) {
	loader := plugin.NewLoader()

	result, err := loader.LoadPlugins([]string{"fixme"})
	assert.Nil(t, err)
	assert.Len(t, result, 1)
}
