# Activo-Forge

HAR (HTTP Archive) 파일을 JMeter JMX 테스트 플랜으로 변환하는 도구입니다.

폐쇄망 Windows 환경에서 단일 실행 파일로 동작합니다.

## 기능

- ✅ HAR 파일을 JMeter JMX 형식으로 변환
- ✅ 정적 리소스 자동 필터링 (js, css, 이미지 등)
- ✅ 특정 도메인만 포함 필터
- ✅ 스레드 수 및 반복 횟수 설정
- ✅ POST 데이터 및 쿼리 파라미터 지원
- ✅ 단일 실행 파일 (의존성 없음)

## 다운로드

`dist/` 디렉토리에서 OS에 맞는 바이너리를 다운로드하세요.

| 파일 | OS | 아키텍처 |
|------|-----|---------|
| `activo-forge-windows-amd64.exe` | Windows | 64비트 |
| `activo-forge-windows-386.exe` | Windows | 32비트 |
| `activo-forge-linux-amd64` | Linux | 64비트 |
| `activo-forge-darwin-amd64` | macOS | Intel |
| `activo-forge-darwin-arm64` | macOS | Apple Silicon |

## 사용법

### 기본 사용

```bash
activo-forge -i recording.har
```

### 전체 옵션

```bash
activo-forge -i <HAR파일> [옵션]

옵션:
  -i string         HAR 파일 경로 (필수)
  -o string         JMX 출력 경로 (기본값: testplan.jmx)
  -threads int      스레드 수 (기본값: 1)
  -loops int        반복 횟수 (기본값: 1)
  -domain string    특정 도메인만 포함
  -exclude-static   정적 리소스 제외 (기본값: true)
  -headers          HTTP 헤더 포함
  -v                버전 정보
```

### 예시

```bash
# 기본 변환
activo-forge -i recording.har -o mytest.jmx

# 10개 스레드, 5회 반복
activo-forge -i recording.har -threads 10 -loops 5

# 특정 도메인만 필터
activo-forge -i recording.har -domain api.example.com

# 정적 리소스 포함
activo-forge -i recording.har -exclude-static=false
```

## 샘플 파일

`samples/` 디렉토리에 예제 파일이 있습니다.

| 파일 | 설명 |
|------|------|
| `sample-api.har` | REST API 호출 예제 |
| `sample-api.jmx` | 변환 결과 |
| `sample-web.har` | 웹 페이지 탐색 예제 |
| `sample-web.jmx` | 변환 결과 |

## HAR 파일 생성 방법

### Chrome DevTools

1. Chrome에서 F12 (DevTools 열기)
2. Network 탭 선택
3. 녹화 버튼 활성화 (빨간 원)
4. 테스트할 동작 수행
5. 우클릭 → "Save all as HAR with content"

### Firefox

1. F12 (개발자 도구 열기)
2. Network 탭 선택
3. 톱니바퀴 → "Save All As HAR"

## 소스에서 빌드

```bash
# 현재 OS
make build

# 모든 플랫폼
make build-all

# 샘플 JMX 생성
make samples
```

### 개별 플랫폼 빌드

```bash
make build-windows-amd64   # Windows 64비트
make build-windows-386     # Windows 32비트
make build-linux-amd64     # Linux 64비트
make build-darwin-amd64    # macOS Intel
make build-darwin-arm64    # macOS Apple Silicon
```

## 프로젝트 구조

```
activo-forge/
├── main.go              # CLI 진입점
├── har/
│   └── parser.go        # HAR 파싱
├── jmeter/
│   ├── generator.go     # JMX 생성
│   └── types.go         # XML 구조체
├── dist/                # OS별 빌드 결과물
├── samples/             # 샘플 HAR/JMX 파일
├── go.mod
├── Makefile
└── README.md
```

## 지원 환경

- Windows 10/11, Windows Server 2016+
- macOS 10.15+ (Intel, Apple Silicon)
- Linux (glibc 2.17+)

## 라이선스

MIT License
