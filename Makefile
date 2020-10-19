BINARY_NAME=test

build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)_linux main.go
	GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)_darwin main.go
	GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)_windows main.go
	chmod +x ./bin/*
