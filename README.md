# go-anonymize

A small Go library of pure functions for anonymizing analytics data before
storage.

## What it does

- **IP masking** — zeroes the last octet of IPv4 addresses and the lower 80
  bits of IPv6 addresses so clients can't be uniquely identified by IP.
- **User-Agent parsing** — reduces a raw User-Agent header to coarse browser
  family, browser major version, operating system, and a mobile flag using
  [uap-go](https://github.com/ua-parser/uap-go).
- **Timestamp rounding** — truncates timestamps to the nearest preceding
  minute to remove high-resolution fingerprinting.
- **Referer domain extraction** — reduces an HTTP Referer URL to just its
  hostname, dropping scheme, path, query, and fragment.

## Install

```sh
go get github.com/olegiv/go-anonymize
```

Requires Go 1.26 or newer.

## Usage

```go
package main

import (
	"encoding/json"
	"net/http"
	"time"

	anonymize "github.com/olegiv/go-anonymize"
)

func handler(w http.ResponseWriter, r *http.Request) {
	event := anonymize.Anonymize(
		r.RemoteAddr,
		r.UserAgent(),
		r.Referer(),
		time.Now(),
	)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(event)
}

func main() {
	http.HandleFunc("/track", handler)
	_ = http.ListenAndServe(":8080", nil)
}
```

`Anonymize` returns a `Result` struct:

```go
type Result struct {
	IP             string    `json:"ip"`
	Browser        string    `json:"browser"`
	BrowserVersion string    `json:"browser_version"`
	OS             string    `json:"os"`
	Mobile         bool      `json:"mobile"`
	Timestamp      time.Time `json:"timestamp"`
	Referer        string    `json:"referer"`
}
```

## Individual functions

Each anonymization step is also exposed directly so you can mix and match.

```go
anonymize.MaskIP("192.168.1.100")
// => "192.168.1.0"

anonymize.MaskIP("2001:db8:1234:5678::1")
// => "2001:db8:1234::"

browser, version, os, mobile := anonymize.ParseUA(
	"Mozilla/5.0 (Linux; Android 14) AppleWebKit/537.36 " +
		"(KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36",
)
// => "Chrome Mobile", "120", "Android", true

anonymize.RoundTimestamp(time.Now())
// => time truncated to the current minute

anonymize.ExtractDomain("https://www.example.com/path?q=1")
// => "www.example.com"
```

All functions are safe to call concurrently. The User-Agent parser is
initialized once on first use via `sync.Once`.

## Development

A `Makefile` wraps the common Go commands:

```sh
make            # fmt-check + vet + race-enabled tests (default)
make test       # plain go test
make test-race  # go test -race -count=1
make vet
make fmt        # gofmt -w .
make build
make vulncheck  # govulncheck ./...  (run `make tools` once to install it)
make tidy       # go mod tidy
make help       # list all targets
```

## License

MIT — see [LICENSE](LICENSE).
