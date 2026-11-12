// Package portonev2 는 포트원 REST API V2 Go 클라이언트입니다.
//
// 타입과 API 메서드는 server-sdk 서브모듈의 openapi.json 으로부터
// codegen/main.ts 가 생성합니다(*.gen.go). 이 파일은 HTTP 처리 등 수작업 코드입니다.
//
//	client := portonev2.NewClient(os.Getenv("PORTONE_API_SECRET"))
//	p, err := client.Payment().GetPayment(ctx, &payment.GetPaymentOptions{PaymentId: "order-1"})
package portonev2

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const (
	// DefaultBaseURL 기본 API URL
	DefaultBaseURL = "https://api.portone.io"

	// DefaultTimeout 기본 타임아웃
	DefaultTimeout = 30 * time.Second

	// UserAgent 기본 User-Agent
	UserAgent = "portone-go-sdk/2.0.0"
)

// Client 포트원 API 클라이언트입니다. 여러 고루틴에서 동시에 사용해도 안전합니다.
type Client struct {
	secret    string
	baseURL   string
	storeID   string
	userAgent string
	http      *http.Client

	clients clients
}

// ClientOption 클라이언트 옵션
type ClientOption func(*Client)

// WithBaseURL 기본 URL(origin)을 설정합니다.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithStoreID 상점 아이디를 설정합니다.
//
// 하위 상점을 대상으로 포트원 기능을 사용할 때 필수입니다.
// 각 API 의 StoreId 파라미터를 지정하지 않으면 이 값이 사용됩니다.
func WithStoreID(storeID string) ClientOption {
	return func(c *Client) {
		c.storeID = storeID
	}
}

// WithHTTPClient HTTP 클라이언트를 설정합니다.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.http = httpClient
	}
}

// WithTimeout HTTP 타임아웃을 설정합니다.
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.http.Timeout = timeout
	}
}

// WithUserAgent User-Agent 헤더 값을 설정합니다.
func WithUserAgent(userAgent string) ClientOption {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// NewClient 새 클라이언트를 생성합니다. secret 은 포트원 API Secret 입니다.
func NewClient(secret string, opts ...ClientOption) *Client {
	c := &Client{
		secret:    secret,
		baseURL:   DefaultBaseURL,
		userAgent: UserAgent,
		http: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
	for _, opt := range opts {
		opt(c)
	}
	c.initClients()
	return c
}

// GetStoreID 설정된 상점 아이디를 반환합니다.
func (c *Client) GetStoreID() string {
	return c.storeID
}

// storeIDPtr 지정된 값이 없으면 클라이언트의 상점 아이디를 사용합니다.
func (c *Client) storeIDPtr(storeID *string) *string {
	if storeID != nil {
		return storeID
	}
	if c.storeID != "" {
		v := c.storeID
		return &v
	}
	return nil
}

// storeIDValue 지정된 값이 비어 있으면 클라이언트의 상점 아이디를 사용합니다.
func (c *Client) storeIDValue(storeID string) string {
	if storeID != "" {
		return storeID
	}
	return c.storeID
}

// setRequestBodyQuery GET/DELETE 요청의 바디를 requestBody 쿼리 파라미터(JSON)로 설정합니다.
func setRequestBodyQuery(query url.Values, body any) error {
	encoded, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("portone: 요청 본문 직렬화 실패: %w", err)
	}
	query.Set("requestBody", string(encoded))
	return nil
}

func (c *Client) send(ctx context.Context, method, path string, query url.Values, body any) ([]byte, error) {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return nil, fmt.Errorf("portone: URL 파싱 실패: %w", err)
	}
	if len(query) > 0 {
		q := u.Query()
		for k, vs := range query {
			for _, v := range vs {
				q.Add(k, v)
			}
		}
		u.RawQuery = q.Encode()
	}

	var reader io.Reader
	if body != nil {
		encoded, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("portone: 요청 본문 직렬화 실패: %w", err)
		}
		reader = bytes.NewReader(encoded)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), reader)
	if err != nil {
		return nil, fmt.Errorf("portone: HTTP 요청 생성 실패: %w", err)
	}
	req.Header.Set("Authorization", "PortOne "+c.secret)
	req.Header.Set("User-Agent", c.userAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("portone: HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("portone: 응답 읽기 실패: %w", err)
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, NewAPIError(resp.StatusCode, respBody)
	}
	return respBody, nil
}

// do 요청을 보내고 JSON 응답을 result 로 역직렬화합니다. result 가 nil 이면 응답 본문을 무시합니다.
func (c *Client) do(ctx context.Context, method, path string, query url.Values, body any, result any) error {
	respBody, err := c.send(ctx, method, path, query, body)
	if err != nil {
		return err
	}
	if result != nil && len(bytes.TrimSpace(respBody)) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("portone: 응답 파싱 실패: %w", err)
		}
	}
	return nil
}

// doText 요청을 보내고 응답 본문(CSV 등)을 문자열로 반환합니다.
func (c *Client) doText(ctx context.Context, method, path string, query url.Values, body any) (string, error) {
	respBody, err := c.send(ctx, method, path, query, body)
	if err != nil {
		return "", err
	}
	return string(respBody), nil
}

// Request 임의의 API 를 호출합니다. 생성된 메서드가 없는 API 를 호출할 때 사용합니다.
//
// body 가 nil 이 아니면 JSON 으로 직렬화하여 전송하고, 응답은 result 로 역직렬화합니다.
func (c *Client) Request(ctx context.Context, method, path string, query url.Values, body any, result any) error {
	return c.do(ctx, method, path, query, body, result)
}
