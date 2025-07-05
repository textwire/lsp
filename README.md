# Textwire LSP

## Installation
```bash
go get github.com/textwire/lsp
```

## License
The Textwire LSP project is licensed under the [MIT License](LICENSE) and is free to use, modify, and distribute.

## Contribute
### Without a Container Engine
If you don't use container engines like [Podman](https://podman.io/) or [Docker](https://www.docker.com/), you need to do a little bit setup. You need to have [Go](https://go.dev/doc/install) installed on your machine.

### With a Container Engine
If you use container engines like [Podman](https://podman.io/) or [Docker](https://www.docker.com/) it's a lot easier for you. You just need to have Podman with Podman Compose or Docker with Docker Compose installed on your machine.

#### Build the Image
To build the image, run this command:
```bash
docker compose build
# for Podman, run this:
podman-compose build
```

#### Run the container
After the image is build, you can run a container from that image. Run this command to create a container and enter it:
```bash
docker compose run --rm app sh
# for Podman, run this:
podman-compose run --rm app sh
```

Inside of the container, you will be able to run commands like `make test` and `make build`.

#### Remove the Container
When you are done working on a project, you can remove the container to cleanup after yourself.
```bash
docker compose down
# for Podman, run this:
podman-compose down
```