# Textwire LSP

## Installation
```bash
go get github.com/textwire/lsp
```

## License
The Textwire LSP project is licensed under the [MIT License](LICENSE) and is free to use, modify, and distribute.

## Contribute
### With a Container Engine
If you use container engines like [Podman](https://podman.io/) or [Docker](https://www.docker.com/) it's a lot easier for you. You just need to have Podman or Docker installed on your machine.

#### Build the Image
To build the image, run this command for Docker:
```bash
docker compose build
```
For Podman, run this:
```bash
podman-compose build
```

#### Run the Container
After the image is build, you can run a container from that image. Run this command to create a container and enter this for Docker:
```bash
docker compose run --rm app sh
```
For Podman, run this:
```bash
podman-compose run --rm app sh
```

Inside of the container, you will be able to run commands like `make test` and `make build`.

#### Remove the Container
When you are done working on a project, you can remove the container to cleanup after yourself. Run this for Docker:
```bash
docker compose down
```
For Podman, run this:
```bash
podman-compose down
```