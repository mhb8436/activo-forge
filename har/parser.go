package har

import (
	"encoding/json"
	"os"
)

// HAR 루트 구조체
type HAR struct {
	Log Log `json:"log"`
}

// Log HAR 로그
type Log struct {
	Version string  `json:"version"`
	Creator Creator `json:"creator"`
	Entries []Entry `json:"entries"`
}

// Creator HAR 생성자 정보
type Creator struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

// Entry 각 HTTP 요청/응답 쌍
type Entry struct {
	StartedDateTime string   `json:"startedDateTime"`
	Time            float64  `json:"time"`
	Request         Request  `json:"request"`
	Response        Response `json:"response"`
	ServerIPAddress string   `json:"serverIPAddress,omitempty"`
}

// Request HTTP 요청
type Request struct {
	Method      string   `json:"method"`
	URL         string   `json:"url"`
	HTTPVersion string   `json:"httpVersion"`
	Headers     []Header `json:"headers"`
	QueryString []Query  `json:"queryString"`
	PostData    *Post    `json:"postData,omitempty"`
	Cookies     []Cookie `json:"cookies"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
}

// Response HTTP 응답
type Response struct {
	Status      int      `json:"status"`
	StatusText  string   `json:"statusText"`
	HTTPVersion string   `json:"httpVersion"`
	Headers     []Header `json:"headers"`
	Content     Content  `json:"content"`
	RedirectURL string   `json:"redirectURL"`
	HeadersSize int      `json:"headersSize"`
	BodySize    int      `json:"bodySize"`
}

// Header HTTP 헤더
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Query 쿼리 파라미터
type Query struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Post POST 데이터
type Post struct {
	MimeType string  `json:"mimeType"`
	Text     string  `json:"text"`
	Params   []Param `json:"params,omitempty"`
}

// Param POST 파라미터
type Param struct {
	Name        string `json:"name"`
	Value       string `json:"value,omitempty"`
	FileName    string `json:"fileName,omitempty"`
	ContentType string `json:"contentType,omitempty"`
}

// Cookie 쿠키
type Cookie struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	Path     string `json:"path,omitempty"`
	Domain   string `json:"domain,omitempty"`
	Expires  string `json:"expires,omitempty"`
	HTTPOnly bool   `json:"httpOnly,omitempty"`
	Secure   bool   `json:"secure,omitempty"`
}

// Content 응답 컨텐츠
type Content struct {
	Size     int    `json:"size"`
	MimeType string `json:"mimeType"`
	Text     string `json:"text,omitempty"`
	Encoding string `json:"encoding,omitempty"`
}

// Parse HAR 파일을 파싱합니다
func Parse(filepath string) (*HAR, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var har HAR
	if err := json.Unmarshal(data, &har); err != nil {
		return nil, err
	}

	return &har, nil
}

// FilterByDomain 특정 도메인 요청만 필터링
func (h *HAR) FilterByDomain(domain string) []Entry {
	if domain == "" {
		return h.Log.Entries
	}

	var filtered []Entry
	for _, entry := range h.Log.Entries {
		if containsDomain(entry.Request.URL, domain) {
			filtered = append(filtered, entry)
		}
	}
	return filtered
}

func containsDomain(url, domain string) bool {
	return len(url) > 0 && len(domain) > 0 &&
		(contains(url, "://"+domain+"/") ||
		 contains(url, "://"+domain+":") ||
		 hasSuffix(url, "://"+domain))
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && searchString(s, substr) >= 0
}

func hasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func searchString(s, substr string) int {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}
