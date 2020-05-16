package brokenlinks

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
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

const brokenLinkFormat = "Broken link %q found in %q. Error: %s"

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

	response, err := doRequest(linkURL)

	if err != nil {
		c <- buildError(link.Line, fmt.Sprintf(brokenLinkFormat, linkURL, link.Content, formatRequestErrorMessage(err)))

		return
	}

	defer response.Body.Close()

	errorMessage := fmt.Sprintf(brokenLinkFormat, linkURL, link.Content, strconv.Itoa(response.StatusCode))

	switch response.StatusCode {
	case 200:
		c <- Result{Broken: false}
	case 403:
		content := getContent(response)

		// If the page secured by Cloudflare, then we decide it is not an error
		if strings.Contains(content, "Cloudflare") {
			c <- Result{Broken: false}
		} else {
			c <- buildError(link.Line, errorMessage)
		}
	default:
		c <- buildError(link.Line, errorMessage)
	}
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

func doRequest(linkURL string) (*http.Response, error) {
	req, err := http.NewRequest("GET", linkURL, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Broken Link Checker/1.0.0")
	req.Header.Add("Accept-Encoding", "gzip")

	return httpClient().Do(req)
}

func formatRequestErrorMessage(err error) string {
	var message string

	switch err.(type) {
	case *url.Error:
		switch {
		// It is needed because Go 1.13 has quoted domain in error message and it is triggered error while running in CI
		case strings.Contains(err.Error(), "no such host"):
			message = "No such host"
		// It is needed because of Telegram blocked in Russia and we want to clean error message content
		case strings.Contains(err.Error(), "connection refused"):
			message = "Connection refused by provider"
		default:
			message = err.Error()
		}
	default:
		message = err.Error()
	}

	return message
}

func httpClient() *http.Client {
	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
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

func getContent(response *http.Response) string {
	var (
		reader io.ReadCloser
		err    error
	)

	switch response.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(response.Body)

		if err != nil {
			log.Fatal(err)
		}

		defer reader.Close()
	default:
		reader = response.Body
	}

	var buf strings.Builder

	_, err = io.Copy(&buf, reader) //nolint:gosec

	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
