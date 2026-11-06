.PHONY: all build test clean install check fmt fmt-check vet test-race

BINARY_NAME=mdpreview-cli
BIN_DIR=bin

all: check build

build:
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/${BINARY_NAME}

install:
	go install ./cmd/${BINARY_NAME}
	@set -e; bin_dir="$$(go env GOBIN)"; \
	if [ -z "$$bin_dir" ]; then gopath="$$(go env GOPATH)"; bin_dir="$${gopath%%:*}/bin"; fi; \
	if command -v mdc >/dev/null 2>&1 || [ -e "$$bin_dir/mdc" ] || [ -L "$$bin_dir/mdc" ]; then \
		echo "Skipping mdc: command or destination already exists"; \
	else \
		ln -s $(BINARY_NAME) "$$bin_dir/mdc"; \
	fi

fmt:
	gofmt -s -w .

fmt-check:
	@test -z "$$(gofmt -l .)" || (echo "Unformatted files found:" && gofmt -l . && exit 1)

vet:
	go vet ./...

test:
	go test -v ./...
	sh tests/install.sh

test-race:
	go test -race ./...

check: fmt-check vet test test-race

clean:
	rm -rf $(BIN_DIR)
