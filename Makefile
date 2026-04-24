.PHONY: run build test fmt vet clean lint tidy test-race

APP_NAME=lucisay
MAIN_FILE=cmd/lucisay/main.go
BIN=bin/$(APP_NAME)

# Standard Build
run:
	go run $(MAIN_FILE)

build:
	@mkdir -p bin
	go build -o $(BIN) $(MAIN_FILE)


# Quality Gate
test:
	go test -race -cover ./...

lint: fmt vet
	golangci-lint run ./...

fmt:
	go fmt ./...

vet:
	go vet ./...


# Maintenance
tidy:
	go mod tidy

clean:
	rm -rf bin
	go clean -testcache


# Manual Verification
demo: build
	$(BIN) say --subject "cat" --message "Hello from Lucisay!"
