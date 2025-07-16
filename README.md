# Textwire LSP

## Installation
```bash
go get github.com/textwire/lsp
```

## License
The Textwire LSP project is licensed under the [MIT License](LICENSE) and is free to use, modify, and distribute.

## Contribute
### With a Container Engine
> [!NOTE]
> If you use [🐳 Docker](https://app.docker.com/) instead of [🦦 Podman](https://podman.io/), just replace `podman-compose` with `docker compose`, and `podman` with `docker` in code examples below.

#### Build the Image
To build the image, run this command:
```bash
podman-compose build
```

#### Run the Container
After the image is build, you can run a container from that image. Run this command to create a container and enter this:
```bash
podman-compose run --rm app sh
```

Inside of the container, you will be able to run commands like `make test` and `make build`.

#### Remove the Container
When you are done working on a project, you can remove the container to cleanup after yourself, run this:
```bash
podman-compose down
```
