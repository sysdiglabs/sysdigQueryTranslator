
all: build lint

build:
	@echo "Building..."
	go build -buildmode c-shared -o library/translator.so cmd/translator.go
	@echo "Building...done"