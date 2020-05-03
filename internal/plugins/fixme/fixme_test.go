package fixme_test

import (
	"testing"

	"github.com/markdown-linter/markdown-linter/internal/entity"
	. "github.com/markdown-linter/markdown-linter/internal/plugins/fixme"
	"github.com/stretchr/testify/assert"
)

func TestLoadReturnsPluginInfo(t *testing.T) {
	parent := entity.Plugin{}
	plugin := Plugin{Plugin: &parent}
	info := plugin.Info()

	assert.Equal(t, "FixmeTag", info.Name)
	assert.Len(t, info.Contributors, 1)
	assert.Equal(t, "Alexander Kadyrov <alexander@kadyrov.dev>", info.Contributors[0])
	assert.Equal(t, "The line has FIXME tag", info.ErrorDescription)
}

func TestLintReturnsNoErrorsIfFIXMENotFound(t *testing.T) {
	parent := entity.Plugin{}
	plugin := Plugin{Plugin: &parent}

	result := plugin.Lint("content")

	assert.Len(t, result, 0)
}

func TestLintSkipsForFIXMEWithoutColon(t *testing.T) {
	parent := entity.Plugin{}
	plugin := Plugin{Plugin: &parent}

	result := plugin.Lint("FIXME Test")

	assert.Len(t, result, 0)
}

func TestLintReturnsErrorredLine(t *testing.T) {
	parent := entity.Plugin{}
	plugin := Plugin{Plugin: &parent}

	result := plugin.Lint("content\n\nFIXME: Test\n\nTest2")

	assert.Len(t, result, 1)

	offence := result[0]
	assert.Equal(t, 3, offence.Line)
	assert.Equal(t, "The line has FIXME tag", offence.Description)
}

func TestLintReturnsErrorredLineForNonCapitalizedFIXME(t *testing.T) {
	parent := entity.Plugin{}
	plugin := Plugin{Plugin: &parent}

	result := plugin.Lint("content\n\n\nfiXme: Test\n\nTest2")

	assert.Len(t, result, 1)

	offence := result[0]
	assert.Equal(t, 4, offence.Line)
	assert.Equal(t, "The line has FIXME tag", offence.Description)
}
