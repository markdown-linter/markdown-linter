package brokenlinks_test

import (
	"fmt"
	"testing"

	. "github.com/markdown-linter/markdown-linter/internal/plugins/brokenlinks"
	"github.com/markdown-linter/markdown-linter/internal/structs"
	"github.com/stretchr/testify/assert"
)

const brokenLinkErrorFormat = "Broken link %q found in %q. %s"
const emptyLinkErrorFormat = "Empty link found in %q"

func Test_Plugin_BrokenLinks_Info_ReturnsPluginInfo(t *testing.T) {
	plugin := Plugin{}
	info := plugin.Info()

	assert.Equal(t, "BrokenLinks", info.Name)
	assert.Len(t, info.Contributors, 1)
	assert.Equal(t, "Alexander Kadyrov <alexander@kadyrov.dev>", info.Contributors[0])
	assert.Equal(t, "The line has broken link", info.ErrorDescription)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfBrokenLinksNotFound(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("content")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsOn301Redirect(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](http://kadyrov.dev)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsOn200HTTPStatusCode(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](https://kadyrov.dev)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsBecauseOfCanvaProtectedByCloudFlare(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](https://www.canva.com)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsBecauseOfTimewebRequiresUserAgent(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](http://timeweb.com/ru/?i=23372)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfLinkDoesNotHaveProtocolScheme(t *testing.T) {
	t.Skip("TODO: Check URL without protocol scheme: kadyrov.dev")
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfLinkHasDoubleSlashInProtocolScheme(t *testing.T) {
	t.Skip("TODO: Skip URL with valid protocol scheme: //kadyrov.dev")
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfLinkLooksLikeNonHTTPScheme(t *testing.T) {
	t.Skip("TODO: Skip URL with non-HTTP scheme: ftp://domain.tld")
}

func Test_Plugin_BrokenLinks_Lint_ReturnsErrorIfURLContainsSpacesBeforeOrAfter(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link]( https://kadyrov.dev )")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line:        1,
		Description: "URL cannot contain spaces at the beginning or ending in \"[Link]( https://kadyrov.dev )\"",
	}
	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsErrorOnDocumentLinkName(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](#)")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line: 1,
		Description: fmt.Sprintf(
			"Empty internal link found in %q",
			"[Link](#)",
		),
	}
	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsErrorOn404HTTPStatusCode(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](http://kadyrov.dev/not-found.txt)")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line: 1,
		Description: fmt.Sprintf(
			brokenLinkErrorFormat,
			"http://kadyrov.dev/not-found.txt",
			"[Link](http://kadyrov.dev/not-found.txt)",
			"Error: 404",
		),
	}
	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_Lint_ChecksEmptyLinkURL(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link]()")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line:        1,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link]()"),
	}
	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_Lint_ChecksEmptyLinkURL_WithSpaces(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](  )")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line:        1,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link](  )"),
	}
	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_Lint_ChecksNoHostError(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](http://domain.tld)")

	assert.Len(t, result, 1)

	offence := structs.Offence{
		Line: 1,
		Description: fmt.Sprintf(
			brokenLinkErrorFormat,
			"http://domain.tld",
			"[Link](http://domain.tld)",
			"Error: No such host",
		),
	}

	assert.Contains(t, result, offence)
}

func Test_Plugin_BrokenLinks_ChecksLinksInSingleLine(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](  )[Link]()")

	assert.Len(t, result, 2)

	firstLink := structs.Offence{
		Line:        1,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link](  )"),
	}
	assert.Contains(t, result, firstLink)

	secondLink := structs.Offence{
		Line:        1,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link]()"),
	}
	assert.Contains(t, result, secondLink)
}

func Test_Plugin_BrokenLinks_Lint_ChecksLinksInDifferentLines(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](  )\n[Link]()")

	assert.Len(t, result, 2)

	firstLink := structs.Offence{
		Line:        1,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link](  )"),
	}
	assert.Contains(t, result, firstLink)

	secondLink := structs.Offence{
		Line:        2,
		Description: fmt.Sprintf(emptyLinkErrorFormat, "[Link]()"),
	}
	assert.Contains(t, result, secondLink)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfLinkRefersToLocalFile(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](../../../testdata/valid/valid.md)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsNoErrorsIfLinkRefersToLocalDirectory(t *testing.T) {
	plugin := Plugin{}

	result := plugin.Lint("[Link](../../../testdata/valid/)")

	assert.Len(t, result, 0)
}

func Test_Plugin_BrokenLinks_Lint_ReturnsErrorIfLinkRefersToNonExistentLocalFile(t *testing.T) {
	t.Skip("TODO: It must return error without executing HTTP request: [Link](../../../testdata/valid/not-found.md)")
}
