# Textwire LSP

## Installation

```bash
go get github.com/textwire/lsp
```

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
