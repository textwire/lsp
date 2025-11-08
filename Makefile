.PHONY: test
test:
	@echo "🚀 Running tests..."
	go test ./...
	@echo "✅ Tests passed!"

.PHONE: build
build:
	go build main.go

.PHONE: check-fmt
check-fmt:
	unformatted=$$(gofmt -l .); \
	if [ -n "$$unformatted" ]; then \
		echo "The following files are not formatted properly:"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

.DEFAULT_GOAL := test
