package main

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name     string
		htmlBody string
		rawURL   string
		expected []string
	}{
		{
			name:   "absolute and relative URLs",
			rawURL: "https://blog.boot.dev",
			htmlBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:   "with ending /",
			rawURL: "https://blog.boot.dev",
			htmlBody: `
			<html>
				<body>
					<a href="/path/one/">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/two">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
		`,
			expected: []string{"https://blog.boot.dev/path/one/", "https://other.com/path/two"},
		},
		{
			name:   "mix and match",
			rawURL: "https://www.trustytea.me/",
			htmlBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://www.trustytea.me/">
						<span>Portfolio</span>
					</a>
					<a href="https://www.trustytea.me/blogs">
						<span>Blogs</span>
					</a>
					<a href="https://www.trustytea.me/projects/">
						<span>Projects</span>
					</a>
					<a href="/contact/linkedin">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
		`,
			expected: []string{"https://www.trustytea.me/path/one", "https://www.trustytea.me/", "https://www.trustytea.me/blogs", "https://www.trustytea.me/projects/", "https://www.trustytea.me/contact/linkedin"},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := GetURLsFromHTML(tc.htmlBody, tc.rawURL)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URLs: %v, actual %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
