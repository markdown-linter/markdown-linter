package parsers

import (
	"regexp"
	"strings"

	"github.com/markdown-linter/markdown-linter/internal/markdownparser/interfaces"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
)

type LinksParser struct {
	interfaces.Parser
}

func (p *LinksParser) Parse(content string) []structs.Tag {
	tags := make([]structs.Tag, 0)

	lines := strings.Split(content, "\n")

	// NOTE: I don't know how to rewrite two regexp to minify if-conditions below
	commonLinkRegexp := regexp.MustCompile(`(?P<image>!?)(?P<text>\[[^\]]*\])(?P<url>\([^)]*\))`)
	linkedImageRegexp := regexp.MustCompile(`(?P<before>\[)*(?P<image>!?)(?P<text>\[[^\]]*\])(?P<url>\([^)]*\))(?P<after>\])?(?P<nextlink>\([^)]*\))`) //nolint:lll

	for idx, line := range lines {
		matches := linkedImageRegexp.FindAllStringSubmatch(line, -1)

		isLinkedImage := false

		if len(matches) == 0 {
			matches = commonLinkRegexp.FindAllStringSubmatch(line, -1)

			if len(matches) == 0 {
				continue
			}
		} else {
			isLinkedImage = true
		}

		var text, url string

		for _, match := range matches {
			if isLinkedImage && len(match[1]) > 0 {
				// The result should look like this: [![Alt](http://link.to/image/here.png)](http://domain.tld)
				// match[1] == [
				// match[2] == !
				// match[3] == [Alt]
				// match[4] == (http://link.to/image/here.png)
				// match[5] == ]
				text = match[1] + match[2] + match[3] + match[4] + match[5]
				url = match[6]
			} else {
				if len(match[1]) > 0 {
					continue
				}

				text = match[2]
				url = match[3]
			}

			tags = append(tags, structs.Tag{Line: idx + 1, Content: text + url})
		}
	}

	return tags
}
