
all: build lint

build:
	@echo "Building..."
	go build -buildmode c-shared -o library/translator.so cmd/translator.go
	@echo "Building...done"

cross-build:
	@echo "Building Windows binary..."
	GOOS=windows GOARCH=amd64 go build -buildmode c-shared -o library/translator-windows-amd64.so cmd/translator.go
	@echo "Building Windows ...done"
	@echo "Building MacOS binary..."
	GOOS=darwin GOARCH=amd64 go build -buildmode c-shared -o library/translator-darwin-amd64.so cmd/translator.go
	@echo "Building MacOS ...done"
	@echo "Building Linux binary..."
	GOOS=linux GOARCH=amd64 go build -buildmode c-shared -o library/translator-linux-amd64.so cmd/translator.go
	@echo "Building Linux ...done"