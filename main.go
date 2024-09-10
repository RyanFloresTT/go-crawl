package main

import (
	"fmt"
)

func main() {
	url, err := NormalizeURL("https://blog.boot.dev/path/")
	if err != nil {
		fmt.Errorf("%w", err)
		return
	}
	fmt.Println(url)
}
