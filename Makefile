.PHONY: build build-all clean test samples

VERSION := 1.0.0
BINARY_NAME := activo-forge
LDFLAGS := -ldflags="-s -w -X main.version=$(VERSION)"
DIST_DIR := dist
SAMPLES_DIR := samples

# 현재 OS용 빌드
build:
	go build $(LDFLAGS) -o $(BINARY_NAME)

# Windows 64비트
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-amd64.exe

# Windows 32비트
build-windows-386:
	GOOS=windows GOARCH=386 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-windows-386.exe

# Linux 64비트
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-linux-amd64

# macOS Intel
build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-amd64

# macOS Apple Silicon
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(DIST_DIR)/$(BINARY_NAME)-darwin-arm64

# 모든 플랫폼 빌드
build-all: build-windows-amd64 build-windows-386 build-linux-amd64 build-darwin-amd64 build-darwin-arm64
	@echo "All platforms built in $(DIST_DIR)/"
	@ls -lh $(DIST_DIR)/

# 샘플 JMX 생성
samples: build
	./$(BINARY_NAME) -i $(SAMPLES_DIR)/sample-api.har -o $(SAMPLES_DIR)/sample-api.jmx
	./$(BINARY_NAME) -i $(SAMPLES_DIR)/sample-web.har -o $(SAMPLES_DIR)/sample-web.jmx
	@echo "Samples generated in $(SAMPLES_DIR)/"

# 정리
clean:
	rm -f $(BINARY_NAME) $(BINARY_NAME).exe
	rm -rf $(DIST_DIR)/*

# 테스트
test:
	go test ./...

# 실행 (개발용)
run:
	go run . -i $(SAMPLES_DIR)/sample-api.har
