
# Run unit tests
.PHONY: test
test:
	go test ./pkg/... 

.PHONY: integration-test
# Run integration tests
integration-test:
	go test -tags integration ./pkg/...


.PHONY: lint
lint:
	golangci-lint run


.PHONY: cli
# Build CLI
cli:
	go build -o bin/wman ./cmd/wman/main.go

# Install CLI
.PHONY: cli-install
cli-install:
	go install  ./cmd/wman/main.go

.PHONY: run
run:
	go run  ./cmd/wman/main.go
	