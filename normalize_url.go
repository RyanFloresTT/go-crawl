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

	normalizedURL := fmt.Sprintf(parsed.Hostname() + parsed.Path)

	if normalizedURL[len(normalizedURL)-1] == '/' {
		normalizedURL = normalizedURL[:len(normalizedURL)-1]
	}

	return normalizedURL, nil
}
