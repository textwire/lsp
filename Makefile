MAX_LINE_LENGTH := 100

OCI ?= podman-compose
RUN = $(OCI) run --rm -T checks
RUN_IT= $(OCI) run --rm checks
STAMP := .github/.checks-image

# Builds the checks image only when the Dockerfile or compose file changes
$(STAMP): Dockerfile compose.yml
	@$(OCI) build-image checks
	@touch $(STAMP)

.PHONY: build-image
build-image: $(STAMP)

.PHONY: test
test: build-image
	@clear || true
	@$(RUN) go test $(RACE) -count=1 -timeout 30s ./...
	@echo "✅ All tests passed!"

.PHONY: fmt
fmt: build-image
	@$(RUN) go fmt ./...
	@echo "✅ Code formatted!"

.PHONY: build
build: build-image
	@$(RUN) go build main.go
	@echo "✅ main executable is ready"
	@ls -hlaF main

.PHONY: line
line: build-image
	@$(RUN) golines -w -m $(MAX_LINE_LENGTH) .
	@echo "✅ Lines limited!"

.PHONY: lint
lint: build-image
	@$(RUN) golangci-lint run
	@echo "✅ Linting passed!"

.PHONY: todo
todo: build-image
	@if $(RUN) grep -I --exclude="Makefile" --exclude-dir=".git" -r TODO .; then \
		echo "❌ Found TODOs" >&2; \
		exit 1; \
	fi

.PHONY: check
check: RACE=-race
check: test line fmt lint todo

.DEFAULT_GOAL := test
