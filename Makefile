NAME=cloudstack-metadata

GOOS?=linux
GOARCH?=amd64
GO111MODULE=on

OUTPUT_DIR=./bin


all: init test build

.PHONY: build
build:
	@echo "==> Building the cmd"
	go build -o $(OUTPUT_DIR)/$(NAME) cmd/$(NAME)/main.go

.PHONY: test
test:
	@echo "==> Testing all packages"
	@go test -v ./...

.PHONY: clean
clean:
	@echo "==> Cleaning releases"
	go clean -i -x ./...

.PHONY: init
init: vendor
	@echo "==> Initializing..."
	@mkdir -p $(OUTPUT_DIR)

.PHONY: vendor
vendor:
	@echo "==> Install modules"
	@GO111MODULE=on go mod tidy
	@GO111MODULE=on go mod vendor
