package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func NormalizeURL(base string) (string, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	normalizedURL := fmt.Sprintf(parsed.Hostname() + parsed.Path)

	if normalizedURL[len(normalizedURL)-1] == '/' {
		normalizedURL = normalizedURL[:len(normalizedURL)-1]
	}

	return normalizedURL, nil
}

func GetURLsFromHTML(htmlbody, rawBaseURL string) ([]string, error) {
	var urlResults []string

	doc, err := html.Parse(strings.NewReader(htmlbody))
	if err != nil {
		return urlResults, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			tagURL := n.Attr[0].Val
			urlResults = append(urlResults, tagURL)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(doc)

	return urlResults, nil
}
