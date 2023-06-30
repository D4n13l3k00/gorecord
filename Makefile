run:

	go run -ldflags="-X main.version=$(shell git describe --tags --abbrev=0) -X main.commit=$(shell git rev-parse --short=8 HEAD)" .

.PHONY: build
.PHONY: build-action

build:
	rm -rf build
	mkdir -p build
	go build -ldflags="-s -w -X main.version=$(shell git describe --tags --abbrev=0) -X main.commit=$(shell git rev-parse --short=8 HEAD)" -o build/gorecord .

build-action:
	go build -ldflags="-X main.version=$(shell git describe --tags --abbrev=0) -X main.commit=$(shell git rev-parse --short=8 HEAD)" -o gorecord .