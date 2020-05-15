package plugin_test

import (
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/plugin"
	"github.com/stretchr/testify/assert"
)

func TestLoadPlugins_ReturnsErrorIfPluginDoesNotExist(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"test"})

	assert.EqualError(t, err, "plugin \"test\" was not found")
	assert.Len(t, result, 0)
}

func TestLoadPlugins_ReturnsAllPluginsIfNoPluginsSpecified(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{})

	assert.NoError(t, err)

	// NOTE: Do not forget to add test below for each new plugin
	assert.Len(t, result, 3)
}

func TestLoadPlugins_ReturnsFixmePlugin(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"fixme"})

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestLoadPlugins_ReturnsH1Plugin(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"h1"})

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}

func TestLoadPlugins_ReturnsBrokenLinksPlugin(t *testing.T) {
	loader := NewLoader()

	result, err := loader.LoadPlugins([]string{"brokenlinks"})

	assert.NoError(t, err)
	assert.Len(t, result, 1)
}
