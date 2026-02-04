.PHONY: build build-windows build-linux build-all clean test

VERSION := 1.0.0
BINARY_NAME := activo-forge
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"

# 현재 OS용 빌드
build:
	go build $(LDFLAGS) -o $(BINARY_NAME)

# Windows 64비트
build-windows:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME).exe

# Windows 32비트
build-windows-32:
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o $(BINARY_NAME)-x86.exe

# Linux 64비트
build-linux:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(BINARY_NAME)-linux

# 모든 플랫폼
build-all: build build-windows build-linux

# 정리
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe $(BINARY_NAME)-linux $(BINARY_NAME)-x86.exe

# 테스트
test:
	go test ./...

# 실행 (개발용)
run:
	go run . -i sample.har
