go-generate:
	@echo "  >  Generating proto files..."
	cd proto; go generate

go-test:
	@echo "  >  Running tests..."
	go test ./internal/... ./e2e/...

go-server:
	@echo "  >  Running server..."
	go run server/main.go

go-client:
	go run client/main.go $$MESSAGE