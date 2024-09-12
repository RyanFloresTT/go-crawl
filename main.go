package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	baseURL, err := url.Parse(args[0])
	if err != nil {
		fmt.Errorf("Error parsing base URL: %v\n", err)
		os.Exit(1)
	}
	cfg := config{
		pages:              make(map[string]int),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, 10),
		wg:                 &sync.WaitGroup{},
		maxPages:           10,
	}
	if len(args) >= 2 {
		// max concurrency
		number, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("The argument for MAX_CONCURRENCY is not a valid integer.")
			return
		}
		cfg.concurrencyControl = make(chan struct{}, number)
	}
	if len(args) == 3 {
		// max pages
		number, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("The argument for MAX_PAGES is not a valid integer.")
			os.Exit(1)
			return
		}
		cfg.maxPages = number
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cfg.baseURL.String())

	cfg.wg.Wait()

	printReport(cfg.pages, cfg.baseURL.String())
}
