package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func GetURLsFromHTML(htmlbody, rawBaseURL string) ([]string, error) {
	var urlResults []string

	doc, err := html.Parse(strings.NewReader(htmlbody))
	if err != nil {
		return urlResults, err
	}

	// get <a> tags and the href values, put them into urlResults

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			tagURL := n.Attr[0].Val

			parsedUrl, err := url.Parse(tagURL)
			if err != nil {
				return
			}

			if parsedUrl.Host == "" {
				urlResults = append(urlResults, rawBaseURL+tagURL)
			} else {
				urlResults = append(urlResults, tagURL)
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urlResults, nil
}

func getHTML(rawURL string) (string, error) {
	var result string

	res, err := http.Get(rawURL)
	if err != nil {
		return result, err
	}

	if res.StatusCode >= 400 {
		return result, fmt.Errorf("400 Status Code")
	}

	contentType := res.Header.Get("Content-Type")

	if strings.Contains(contentType, "Text/Html") {
		return result, fmt.Errorf("wrong content-type header value: %s", res.Header.Get("Content-Type"))
	}

	html, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	res.Body.Close()

	result = string(html)

	return result, nil
}
