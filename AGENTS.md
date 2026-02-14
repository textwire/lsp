## Instructions

- In all interactions, be extremely concise and sacrifice grammar for the sake of concision.
- Don't add any other functionality other than what I've said.

## Project Overview

This is a Go-based LSP implementation for the Textwire templating language.

- Language: Go 1.25.0
- Module: `github.com/textwire/lsp`
- Main dependency: `github.com/textwire/textwire/v3`

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
- Run `gofmt` on all files before committing
- All tests must pass
- Maintain existing code style and patterns

## LSP Capabilities

Current LSP methods implemented:

- `initialize` - Server initialization
- `textDocument/didOpen` - Track opened documents
- `textDocument/didChange` - Track document changes
- `textDocument/hover` - Hover information
- `textDocument/completion` - Autocomplete suggestions
