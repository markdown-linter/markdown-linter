package plugin_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/plugin"
	"github.com/stretchr/testify/assert"
)

func TestLoadPlugins_ReturnsErrorIfNoPluginsSpecified(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{})

	assert.Equal(t, "you need to specify plugins", err.Error())
	assert.Len(t, result, 0)
}

func TestLoadPlugins_ReturnsErrorIfPluginDoesNotExist(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"test"})

	assert.Equal(t, "plugin test was not found", err.Error())
	assert.Len(t, result, 0)
}

func TestLoadPlugins_ReturnsFixmePlugin(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"fixme"})
	assert.Nil(t, err)
	assert.Len(t, result, 1)
}
