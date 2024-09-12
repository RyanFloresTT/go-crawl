# go-crawl

`go-crawl` is a simple web crawler written in Go. It allows you to crawl a website, optionally limit the number of pages to crawl and control the maximum number of concurrent goroutines used during the crawl.

## Installation

To install `go-crawl`, use the following command:

```bash
go install github.com/RyanFloresTT/go-crawl@latest
```

## Usage

To use the web crawler, run the command with the required and optional parameters. 

### Command Line Parameters

- **Required:**
  - `URL` (string): The URL of the website you want to crawl.

- **Optional:**
  - `max-pages` (int): The maximum number of pages to crawl. If not specified, this defaults to 50.
  - `max-goroutines` (int): The maximum number of goroutines to use concurrently. If not specified, the default value of 10 will be used.

### Example

```bash
go-crawl https://example.com 100 10
```

## Output

The crawler will output which pages it is crawling and provide a report of internal links found. The report is sorted by the number of times each link is found (highest first) and then alphabetically.

### Example Output

```
Crawling page: https://example.com
Crawling page: https://example.com/about
...
Found 15 internal links to https://example.com/contact
Found 10 internal links to https://example.com/about
Found 5 internal links to https://example.com
```

## Reporting

The report is printed using the following format:

```go
fmt.Printf("Found %v internal links to %s\n", kv.Value, kv.Key)
```

Where `kv` is a key-value pair from the map used internally in the crawler. The map is sorted first by the number of times each link is found and then alphabetically by the link itself.
