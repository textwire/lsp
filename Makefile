w.PHONY: test
test:
	@echo "🚀 Running tests..."
	go test ./...
	@echo "✅ Tests passed!"

.PHONE: push
push:
	@echo "🚀 Pushing to GitHub..."
	make test
	git push
	@echo "✅ Pushed to GitHub!"

.DEFAULT_GOAL := test
