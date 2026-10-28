# Textwire LSP

This is a Go-based LSP implementation for the Textwire templating language.

## Installation

```bash
go get github.com/textwire/lsp
```

## Project Overview

- Language: Go 1.25.0
- Module: `github.com/textwire/lsp`
- Main dependency: `github.com/textwire/textwire/v5`

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

## License

The Textwire LSP project is licensed under the [MIT License](LICENSE) and is free to use, modify, and distribute.

## Contribute

### With a Container Engine

#### Build the Image

With Podman:

```bash
podman-compose build
```

With Docker:

```bash
docker compose build
```

#### Run the Container

After the image is build, you can run a container from that image. Run this command to create a container:

With Podman:

```bash
podman-compose run --rm app sh
```

With Docker:

```bash
docker compose run --rm app sh
```

Inside of the container, you will be able to run commands like `make test` and `make build`.

#### Remove the Container

When you are done working on a project, you can remove the container to cleanup after yourself:

With Podman:

```bash
podman-compose down
```

With Docker:

```bash
docker compose down
```

### Public a Version

To publish a new version of LSP, you need to use tags. As soon as you push a new tech to GitHub, it will run a build that will create a release for this new version.

```bash
git tag <version>
```

```bash
git push --tags
```
