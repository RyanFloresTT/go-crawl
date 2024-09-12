package main

import (
	"fmt"
	"net/url"
)

func NormalizeURL(base string) (string, error) {
	parsed, err := url.Parse(base)
	if err != nil {
		return "", fmt.Errorf("%w", err)
	}

	normalizedURL := parsed.Scheme + "://" + parsed.Hostname() + parsed.Path

	if len(normalizedURL) > 1 && normalizedURL[len(normalizedURL)-1] == '/' {
		normalizedURL = normalizedURL[:len(normalizedURL)-1]
	}

	return normalizedURL, nil
}
