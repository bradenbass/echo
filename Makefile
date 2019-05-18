go-generate:
	@echo "  >  Generating proto files..."
	cd proto; go generate

go-test:
	@echo "  >  Running tests..."
	go test ./internal/...
