w.PHONY: test
test:
	@echo "🚀 Running tests..."
	go test ./...
	@echo "✅ Tests passed!"

.PHONE: build
build:
	go build main.go

.DEFAULT_GOAL := test
