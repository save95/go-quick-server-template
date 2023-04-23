GO_CMD=go
GO_BUILD=$(GO_CMD) build
GO_BUILD_COMPRESSION=$(GO_CMD) build -ldflags="-s -w"
BINARY_NAME=server-api

build:
	# swag init -g=../main.go -d=app
	#CGO_ENABLED=1 CC=x86_64-linux-musl-gcc CGO_LDFLAGS="-static" GOOS=linux GOARCH=amd64 $(GO_BUILD_COMPRESSION) -o $(BINARY_NAME) && upx $(BINARY_NAME)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GO_BUILD_COMPRESSION) -o $(BINARY_NAME) && upx $(BINARY_NAME)
