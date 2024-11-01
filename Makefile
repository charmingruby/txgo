PROJECT_NAME := txgo

.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o ./bin/${PROJECT_NAME} ./cmd/api/main.go
