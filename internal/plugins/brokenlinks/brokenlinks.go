package brokenlinks

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/markdown-linter/markdown-linter/internal/entity"
	"github.com/markdown-linter/markdown-linter/internal/markdownparser"
	markdownparserStructs "github.com/markdown-linter/markdown-linter/internal/markdownparser/structs"
	"github.com/markdown-linter/markdown-linter/internal/structs"
)

type Plugin struct {
	entity.Plugin
}

func (p *Plugin) Info() *structs.PluginInfo {
	return &structs.PluginInfo{
		Name:             "BrokenLinks",
		Contributors:     []string{"Alexander Kadyrov <alexander@kadyrov.dev>"},
		ErrorDescription: "The line has broken link",
	}
}

type Result struct {
	Broken  bool
	Offence structs.Offence
}

func (p *Plugin) Lint(content string) []structs.Offence {
	result := make([]structs.Offence, 0)
	c := make(chan Result)

	tags := markdownparser.NewMarkdownParser().Parse(content)

	for _, link := range tags.Links {
		go checkLink(link, c)
	}

	for range tags.Links {
		r := <-c

		if !r.Broken {
			continue
		}

		result = append(result, r.Offence)
	}

	return result
}

func checkLink(link markdownparserStructs.Tag, c chan Result) {
	linkURL := getURL(link.Content)

	if isStringEmpty(linkURL) {
		c <- buildError(link.Line, fmt.Sprintf("Empty link found in %q", link.Content))

		return
	}

	if isStringHasSpacesAtStartOrEnd(linkURL) {
		c <- buildError(
			link.Line,
			fmt.Sprintf("URL cannot contain spaces at the beginning or ending in %q", link.Content),
		)

		return
	}

	linkURL = strings.TrimSpace(linkURL)

	if isReferToInternalDocumentPart(linkURL) {
		c <- buildError(link.Line, fmt.Sprintf("Empty internal link found in %q", link.Content))

		return
	}

	if isReferToLocalFileOrDirectory(linkURL) {
		c <- Result{Broken: false}

		return
	}

	response, err := httpClient().Get(linkURL)

	if err != nil {
		var message string

		switch err.(type) {
		case *url.Error:
			message = "No such host"
		default:
			message = err.Error()
		}

		c <- buildError(link.Line, fmt.Sprintf("Broken link %q found in %q. Error: %s", linkURL, link.Content, message))

		return
	}

	defer response.Body.Close()

	if response.StatusCode == 200 {
		c <- Result{Broken: false}

		return
	}

	c <- buildError(
		link.Line,
		fmt.Sprintf("Broken link %q found in %q. Error: %d", linkURL, link.Content, response.StatusCode),
	)
}

func getURL(content string) string {
	urlRegexp := regexp.MustCompile(`\[[^\]]*\]\(([^)]*)\)`)

	matches := urlRegexp.FindAllStringSubmatch(content, -1)

	return matches[0][1]
}

func isStringEmpty(value string) bool {
	return len(strings.TrimSpace(value)) == 0
}

func isStringHasSpacesAtStartOrEnd(value string) bool {
	return len(value) != len(strings.TrimSpace(value))
}

func isReferToInternalDocumentPart(url string) bool {
	return url[0:1] == "#" && len(url) == 1
}

func isReferToLocalFileOrDirectory(uri string) bool {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	stat, err := os.Stat(dir + string(os.PathSeparator) + uri)

	return err == nil && (stat.Mode().IsRegular() || stat.Mode().IsDir())
}

func httpClient() *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 3 * time.Second,
		}).Dial,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
}

func buildError(line int, description string) Result {
	var result Result

	offence := structs.Offence{Line: line, Description: description}

	result.Broken = true
	result.Offence = offence

	return result
}
