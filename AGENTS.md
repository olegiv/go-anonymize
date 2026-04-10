# AGENTS.md

This file provides guidance to Codex (codex.openai.com) when working with code in this repository.

## Project

`github.com/olegiv/go-anonymize` is a small Go library of pure functions for anonymizing analytics data: IP masking, User-Agent parsing, timestamp rounding, and Referer domain extraction. No HTTP middleware, no side effects — each function takes inputs and returns outputs so callers can wire it into any ingest path.

Requires Go 1.26. Single package `anonymize` at the repo root.

## Common commands

Prefer the `Makefile` for standard workflows — it matches what CI should run:

```sh
make            # fmt-check + vet + race-enabled tests (default)
make test-race  # go test -race -count=1 ./...
make vulncheck  # govulncheck ./...  (run `make tools` once to install it)
make help       # list all targets
```

Drop down to raw `go` for anything Make doesn't cover, like running a single test:

```sh
go test -run TestParseUA ./...                # one test function
go test -run TestParseUA/chrome_on_windows    # one table-driven subtest
```

The `validate-go-test.sh` PreToolUse hook reminds you to pass `-race` on ad-hoc `go test` invocations; `validate-go-toolchain.sh` blocks Go commands when the installed toolchain disagrees with the compiler.

## Architecture

### Public API shape

`Anonymize(ip, ua, referer string, ts time.Time) Result` in `anon.go` is the single-call entry point. It delegates to four independently exported pure functions, each in its own file, each safe to call concurrently:

- `MaskIP` (`ip.go`) — zeroes the last IPv4 octet or the low 80 bits of an IPv6 address. Uses only `net` from stdlib.
- `ParseUA` (`useragent.go`) — returns `(browser, version, os, mobile)`. Backed by `github.com/ua-parser/uap-go`.
- `RoundTimestamp` (`timestamp.go`) — `t.Truncate(time.Minute)`.
- `ExtractDomain` (`referer.go`) — `url.Parse` → `Hostname()`.

Everything takes strings/`time.Time` and returns strings/`time.Time`/primitives — no errors, no context, no I/O. Invalid inputs return empty strings / zero values rather than errors, because this code sits on the analytics hot path and must never panic or fail-loud on hostile input.

### Non-obvious: uap-go initialization

The one place to be careful is `useragent.go`. The upstream `uap-go` module **already embeds the regex definitions** as `uaparser.DefinitionYaml`, and `uaparser.New()` consumes them internally — there is no `regexes.yaml` file on disk to load, no `//go:embed` directive, and no need to vendor anything from `uap-core`. Don't reintroduce file-based loading.

The parser is expensive to build (thousands of regex compilations) so it is constructed once behind a `sync.Once`. If construction fails, `uaParser` stays nil and `ParseUA` returns zero values — this is deliberate, so a broken regex set can't crash an analytics pipeline.

### `Mobile` semantics

`mobile` is true only when `Os.Family == "iOS" || Os.Family == "Android"`. This is narrow on purpose: it misses Windows Phone, KaiOS, etc., but it's predictable and won't flip tablets, smart TVs, or consoles as "mobile". If you change this rule, update the doc comment on `ParseUA` and the `useragent_test.go` cases.

### UA test expectations

Every expected value in `useragent_test.go` comes from actually running the parser against real UA strings — do not hand-edit them from memory. If a uap-go version bump changes a family name (e.g., "Mobile Safari" → "Safari Mobile"), re-probe the parser and update the table; don't guess.

## Tooling via `.claude/shared` submodule

`.claude/shared` is a git submodule pointing at `olegiv/claude-code-support-tools`. The `.claude/settings.json`, agent files, command files, and Go hook scripts are all symlinks into that submodule — edit those files upstream, not in-place. After pulling submodule updates, run `go test ./...` to confirm the hooks still accept the project.

Available slash commands (all sourced from the submodule):

- `/code-quality` — run the Go stack quality audit
- `/commit-prepare` then `/commit-do` — two-step commit workflow (draft → approve → commit)
- `/finalize` — review session changes, update tests/docs
- `/security-audit` — full security audit; results land in `.audit/` (gitignored)
- `/update-submodule` — refresh `.claude/shared` to latest
