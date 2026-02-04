package main

import (
	"flag"
	"fmt"
	"os"

	"activo-forge/har"
	"activo-forge/jmeter"
)

var version = "1.0.0"

func main() {
	input := flag.String("i", "", "HAR 파일 경로 (필수)")
	output := flag.String("o", "testplan.jmx", "JMX 출력 경로")
	threads := flag.Int("threads", 1, "스레드 수")
	loops := flag.Int("loops", 1, "반복 횟수")
	filterDomain := flag.String("domain", "", "특정 도메인만 포함 (예: api.example.com)")
	excludeStatic := flag.Bool("exclude-static", true, "정적 리소스 제외 (js, css, 이미지 등)")
	includeHeaders := flag.Bool("headers", false, "HTTP 헤더 포함")
	showVersion := flag.Bool("v", false, "버전 정보")
	flag.Parse()

	if *showVersion {
		fmt.Printf("Activo-Forge v%s\n", version)
		fmt.Println("HAR to JMeter 변환 도구")
		os.Exit(0)
	}

	if *input == "" {
		printUsage()
		os.Exit(1)
	}

	// HAR 파일 존재 확인
	if _, err := os.Stat(*input); os.IsNotExist(err) {
		fmt.Printf("오류: HAR 파일을 찾을 수 없습니다: %s\n", *input)
		os.Exit(1)
	}

	fmt.Printf("Activo-Forge v%s\n", version)
	fmt.Printf("입력: %s\n", *input)
	fmt.Printf("출력: %s\n", *output)
	fmt.Println("---")

	// HAR 파싱
	harData, err := har.Parse(*input)
	if err != nil {
		fmt.Printf("HAR 파싱 실패: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("총 %d개 요청 발견\n", len(harData.Log.Entries))

	// JMX 생성 옵션
	opts := jmeter.GenerateOptions{
		Threads:        *threads,
		Loops:          *loops,
		FilterDomain:   *filterDomain,
		ExcludeStatic:  *excludeStatic,
		IncludeHeaders: *includeHeaders,
	}

	// JMX 생성
	result, err := jmeter.Generate(harData, *output, opts)
	if err != nil {
		fmt.Printf("JMX 생성 실패: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("변환된 요청: %d개\n", result.RequestCount)
	if result.SkippedCount > 0 {
		fmt.Printf("제외된 요청: %d개 (정적 리소스/필터)\n", result.SkippedCount)
	}
	fmt.Printf("---\n")
	fmt.Printf("완료: %s\n", *output)
}

func printUsage() {
	fmt.Println("Activo-Forge - HAR to JMeter 변환 도구")
	fmt.Println()
	fmt.Println("사용법:")
	fmt.Println("  activo-forge -i <HAR파일> [-o <JMX파일>] [옵션]")
	fmt.Println()
	fmt.Println("예시:")
	fmt.Println("  activo-forge -i recording.har")
	fmt.Println("  activo-forge -i recording.har -o mytest.jmx")
	fmt.Println("  activo-forge -i recording.har -threads 10 -loops 5")
	fmt.Println("  activo-forge -i recording.har -domain api.example.com")
	fmt.Println()
	fmt.Println("옵션:")
	flag.PrintDefaults()
}
