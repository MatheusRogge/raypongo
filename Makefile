.PHONY: build_windows
build_windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ go build -ldflags="-extldflags=-static -linkmode external -s -w" cmd/main.go

.PHONY: build_linux
build_linux:
	go build -ldflags="-extldflags=-static" cmd/main.go
