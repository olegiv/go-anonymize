GO ?= go
GOBIN := $(shell $(GO) env GOPATH)/bin
GOVULNCHECK ?= $(GOBIN)/govulncheck

.PHONY: all check test test-race vet fmt fmt-check build vulncheck tidy tools clean help

all: check

check: fmt-check vet test-race

test:
	$(GO) test ./...

test-race:
	$(GO) test -race -count=1 ./...

vet:
	$(GO) vet ./...

fmt:
	gofmt -w .

fmt-check:
	@out=$$(gofmt -l .); \
	if [ -n "$$out" ]; then \
		echo "gofmt: the following files need formatting:"; \
		echo "$$out"; \
		exit 1; \
	fi

build:
	$(GO) build ./...

vulncheck:
	$(GOVULNCHECK) ./...

tidy:
	$(GO) mod tidy

tools:
	$(GO) install golang.org/x/vuln/cmd/govulncheck@latest

clean:
	$(GO) clean ./...

help:
	@echo "Targets:"
	@echo "  check       fmt-check + vet + test-race (default)"
	@echo "  test        run all tests"
	@echo "  test-race   run tests with the race detector"
	@echo "  vet         run go vet"
	@echo "  fmt         format sources with gofmt"
	@echo "  fmt-check   fail if gofmt would change anything"
	@echo "  build       compile all packages"
	@echo "  vulncheck   run govulncheck (requires: make tools)"
	@echo "  tidy        run go mod tidy"
	@echo "  tools       install govulncheck"
	@echo "  clean       run go clean"
