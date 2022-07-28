PACKAGE_NAME          := github.com/sysdiglabs/sysdigQueryTranslator

all: build lint

build:
	@echo "Building..."
	go build -buildmode c-shared -o library/translator.so cmd/translator.go
	@echo "Building...done"

cross-build:
	@echo "Building Windows binary..."
	CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ CGO_ENABLED=1 GOOS=windows GOARCH=amd64 go build -buildmode c-shared -o library/translator-windows-amd64.so cmd/translator.go
	@echo "Building Windows ...done"
	@echo "Building MacOS binary..."
	GOOS=darwin GOARCH=amd64 go build -buildmode c-shared -o library/translator-darwin-amd64.so cmd/translator.go
	@echo "Building MacOS ...done"
	@echo "Building Linux binary..."
	GOOS=linux GOARCH=amd64 go build -buildmode c-shared -o library/translator-linux-amd64.so cmd/translator.go
	@echo "Building Linux ...done"

darwin-build:
	@docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-e CC=o64-clang \
		-e CXX=o64-clang++ \
		-e GOOS=darwin \
		-e GOARCH=amd64 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:v1.18.3 build