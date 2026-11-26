# Textwire LSP

This is a Go-based LSP implementation for the Textwire templating language.

## Installation

```bash
go get github.com/textwire/lsp
```

## Project Overview

- Language: Go 1.25.0
- Module: `github.com/textwire/lsp`
- Main dependency: `github.com/textwire/textwire/v6`

## Architecture

- **main.go**: Entry point, message routing via stdio
- **lsp/**: LSP message types (initialize, hover, completion, document changes)
- **analysis/**: Core logic (state management, hover, completion features)
- **rpc/**: RPC encoding/decoding utilities
- **internal/logger/**: Logging utilities

## Development Commands

All commands are in the Makefile:

- `make test` - Run all tests
- `make build` - Build the binary
- `make check-fmt` - Verify Go code formatting

## Code Standards

- Follow standard Go conventions
- Run `gofmt` on all files
- All tests must pass
- Maintain existing code style and patterns

## LSP Capabilities

Current LSP methods implemented:

- `initialize` - Server initialization
- `textDocument/didOpen` - Track opened documents
- `textDocument/didChange` - Track document changes
- `textDocument/hover` - Hover information
- `textDocument/completion` - Autocomplete suggestions

### Available Commands

All commands run inside a container. The image is built automatically the first time you run any command:

```bash
make test
```

- `make build-image` - Build the checks container image
- `make build` - Builds a `main` binary in the root of the project
- `make test` - Run all tests
- `make fmt` - Format Go code
- `make lint` - Run linter
- `make check` - Combines commands like: test, fmt, lint
- `make` - Defaults to `make check`

Commands use `podman-compose` by default. To use Docker Compose instead, override `OCI`:

```bash
make check OCI='docker compose'
```

## New Tag Release

To publish a new version of LSP, you need to use tags. As soon as you push a new tech to GitHub, it will run a build that will create a release for this new version.

```bash
git tag <version>
```

```bash
git push --tags
```

## License

The Textwire LSP project is licensed under the [MIT License](LICENSE) and is free to use, modify, and distribute.
