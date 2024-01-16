.PHONY: run-app
run-app: # Run the application
	go run ./cmd/main.go

.PHONY: fmt-all
fmt-all: # Format all go files
	go install golang.org/x/tools/cmd/goimports@v0.6.0
	goimports -w .

.PHONY: ensure-lint
ensure-lint: # Ensure that the linting binary is the correct version
	@golangci-lint --version | grep '1.[2-9][0-9]' > /dev/null || echo "Please check your version of golangci-lint. Version >= 1.29.0 is required."

.PHONY: lint
lint: ensure-lint # Lint the application
	@golangci-lint --enable gofmt --enable goimports --enable whitespace --enable exportloopref run ./...

.PHONY: run-tests
run-tests: # Run all application tests
	go test ./... -v -short -count=1

.PHONY: ready-pr
ready-pr: fmt-all run-tests lint
	go mod tidy

.PHONY: run-coverage
run-coverage: # Run all library tests and add test coverage
	go test ./... -v -short -count=1 -coverprofile=coverage.out
	cat coverage.out | grep -v "mock_" > coverage_without_mocks.out
	go tool cover -html=coverage_without_mocks.out

.PHONY: run-db
run-db: # Start the database locally with Docker
	docker-compose up -d

.PHONY: generate-outdoorsy-challenge-dtos
generate-outdoorsy-challenge-dtos: # Generates the dto models from API definition.
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.15
	oapi-codegen --config api/config.gen.yaml api/api-definition.yaml > api/api.gen.go
