# Agent Guide

This repository is `chronos2-go-base-part4`, one of four Go modules that together carry the shards of Amazon's `amazon/chronos-2` checkpoint for [`chronos2-go-base`](https://github.com/pinkhop/chronos2-go-base). Its one package, `part4`, embeds this part's shards and returns them through `Shards()`. These instructions apply to any agent or developer changing the repository.

`chronos2-go-base` imports all four parts and passes their shards, in order, to `chronos2.Load`. Its [`AGENTS.md`](https://github.com/pinkhop/chronos2-go-base/blob/main/AGENTS.md) holds the rules for the checkpoint as a whole, and [`chronos2-go`](https://github.com/pinkhop/chronos2-go)'s `AGENTS.md` holds the full set of conventions. Keep the four part repositories identical in structure.


## Repository Layout

The module sits at the repository root, and the repository holds nothing else:

| Path | Role |
|---|---|
| `part4.go` | Package `part4`: the `//go:embed` declarations for this part's shards, and `Shards()`. |
| `weights/manifest.json` | The upstream source (repository, revision, file, and SHA-256), `max_shard_bytes`, and each shard's file name, byte count, SHA-256, and tensor count. |
| `weights/*.safetensors` | This part's fp32 safetensors shards. |
| `go.mod` | The module definition, which requires nothing. |
| `README.md` | The overview. |
| `AGENTS.md` | This guide; `CLAUDE.md` is a symbolic link to it. |
| `Makefile` | Formatting, analysis, builds, and the weight and size checks. |
| `mise.toml`, `mise.lock` | The pinned Go toolchain and developer tools. |
| `LICENSE`, `NOTICE` | The Apache-2.0 license and the attribution for Chronos-2. |
| `.editorconfig`, `.gitattributes`, `.gitignore` | Editor and Git settings; `.gitattributes` keeps `*.safetensors` binary and out of Git LFS. |


## Commands

Tool versions, including Go itself, are pinned in `mise.toml` and `mise.lock`. Run `mise install` once, then use the `Makefile`:

| Command | Purpose |
|---|---|
| `make help` | List all targets. |
| `make format` | Format Go files with `goimports`, then `gofumpt`. |
| `make lint` | Check formatting, then run `go vet`, `staticcheck`, `errcheck`, `ineffassign`, `misspell`, and `typos`. |
| `make security` | Run `gosec` and `govulncheck`. |
| `make build` | Compile the package, which embeds this part's shards. |
| `make verify-weights` | Check every shard against `weights/manifest.json`: SHA-256s, byte counts, the `max_shard_bytes` ceiling, and that the shards on disk are exactly the ones listed, each listed once. |
| `make check-size` | Check the working tree, as it would be committed, against Go's two 500 MiB module limits, and print the headroom under each. |
| `make check` | Run `verify-weights`, then `check-size`. |

Before considering a change complete, run `make format`, `make lint`, `make security`, and `make check`.


## Hard Constraints

- Never store weights with Git LFS; `go get` cannot fetch LFS content.
- Never edit, re-encode, or re-shard weight files by hand. Every tensor stays bit-identical to the upstream checkpoint at the revision the manifest pins.
- Keep `weights/manifest.json` accurate; `make verify-weights` must pass.
- Embed shards with `//go:embed` into `[]byte` variables, never `string`, so `chronos2.Load` can view them in place. The embed test in `chronos2-go`'s `parity` module enforces this.
- `Shards()` returns this part's shards in checkpoint order. Never change its signature, and never add another exported symbol.
- The module requires nothing. Never add a dependency, Go tests, `testdata/`, or nested modules. Tests of the checkpoint belong in `chronos2-go`'s `parity` module.
- Never add shards to a part: three shards keep each module small enough for the Go module proxy to serve. `make check` must pass.


## Go Conventions

`chronos2-go`'s `AGENTS.md` ("Go Conventions") governs Go code here. Every declaration has a doc comment; use American spelling and the serial comma; keep the Go version and pinned tools in sync with `chronos2-go`.


## Documentation

- Documents describe the current state of the repository. Keep decision history out of them; commit messages carry history.
- Markdown uses ATX headings in title case, two blank lines before every H1 and H2 (except a first-line title), `-` list markers, four-space nested indents, one paragraph per source line, and a language tag on every fenced code block.


## Commits

Write commit messages in the Conventional Commits format (`feat:`, `fix:`, `docs:`, `refactor:`, `chore:`), with the subject in lowercase imperative mood.
