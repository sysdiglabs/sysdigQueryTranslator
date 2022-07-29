PACKAGE_NAME          := github.com/sysdiglabs/sysdigQueryTranslator

all: build

build:
	@echo "Building..."
	go build -buildmode c-shared -o library/translator.so translator.go
	@echo "Building...done"
