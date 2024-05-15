GO ?= go
GOOS ?= $(shell $(GO) env GOOS)
GOARCH ?= $(shell $(GO) env GOARCH)
MODULE_NAME ?= $(shell head -n1 go.mod | cut -f 2 -d ' ')

.PHONY: test
test:
	$(GO) test ./... -parallel 10

.PHONY: build
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o ubiqcd cmd/ubiqcd/main.go
	GOOS=$(GOOS) GOARCH=$(GOARCH) $(GO) build -o ubiqctl cmd/ubiqctl/main.go
