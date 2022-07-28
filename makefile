PACKAGE_NAME          := github.com/sysdiglabs/sysdigQueryTranslator

all: build

build:
	@echo "Building..."
	go build -buildmode c-shared -o library/translator.so translator.go
	@echo "Building...done"

cross-build:
	@docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:v1.18.3 build