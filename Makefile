run:
	go run .

.PHONY: build

build:
	rm -rf build
	mkdir -p build
	go build -ldflags="-s -w" -o build/gorecord .
	upx --best --lzma build/gorecord