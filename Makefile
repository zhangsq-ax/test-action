BINARY_NAME=test
OS=linux
ARCH=amd64

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o ./bin/$(BINARY_NAME) main.go