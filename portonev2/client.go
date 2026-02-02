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

	"github.com/Gliese436/portone-go-client/portonev2/common"
	"github.com/Gliese436/portone-go-client/portonev2/platform"
)

const (
	// DefaultBaseURL 기본 API URL
	DefaultBaseURL = "https://api.portone.io"

	// DefaultTimeout 기본 타임아웃
	DefaultTimeout = 30 * time.Second

	// UserAgent 사용자 에이전트
	UserAgent = "portone-go-sdk/2.0.0"
)

// Client PortOne API 클라이언트
type Client struct {
	// 설정
	secret  string
	baseURL string
	storeID string
	http    *http.Client

	// 서브 클라이언트 (지연 초기화)
	payment              *PaymentClient
	identityVerification *IdentityVerificationClient
	auth                 *AuthClient
	b2b                  *B2BClient
	platformClient       *platform.Client
	pgSpecific           *PgSpecificClient
}

// ClientOption 클라이언트 옵션
type ClientOption func(*Client)

// WithBaseURL 기본 URL 설정
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithStoreID 상점 ID 설정
func WithStoreID(storeID string) ClientOption {
	return func(c *Client) {
		c.storeID = storeID
	}
}

// WithHTTPClient HTTP 클라이언트 설정
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.http = httpClient
	}
}

// WithTimeout 타임아웃 설정
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.http.Timeout = timeout
	}
}

// NewClient 새 클라이언트를 생성합니다.
func NewClient(secret string, opts ...ClientOption) *Client {
	c := &Client{
		secret:  secret,
		baseURL: DefaultBaseURL,
		http: &http.Client{
			Timeout: DefaultTimeout,
		},
	}

	for _, opt := range opts {
		opt(c)
	}

	return c
}

// Payment Payment API 클라이언트를 반환합니다.
func (c *Client) Payment() *PaymentClient {
	if c.payment == nil {
		c.payment = NewPaymentClient(c)
	}
	return c.payment
}

// IdentityVerification 본인인증 API 클라이언트를 반환합니다.
func (c *Client) IdentityVerification() *IdentityVerificationClient {
	if c.identityVerification == nil {
		c.identityVerification = NewIdentityVerificationClient(c)
	}
	return c.identityVerification
}

// Auth 인증 API 클라이언트를 반환합니다.
func (c *Client) Auth() *AuthClient {
	if c.auth == nil {
		c.auth = NewAuthClient(c)
	}
	return c.auth
}

// B2B B2B API 클라이언트를 반환합니다.
func (c *Client) B2B() *B2BClient {
	if c.b2b == nil {
		c.b2b = NewB2BClient(c)
	}
	return c.b2b
}

// Platform Platform API 클라이언트를 반환합니다.
func (c *Client) Platform() *platform.Client {
	if c.platformClient == nil {
		c.platformClient = platform.NewClient(c)
	}
	return c.platformClient
}

// PgSpecific PG사별 API 클라이언트를 반환합니다.
func (c *Client) PgSpecific() *PgSpecificClient {
	if c.pgSpecific == nil {
		c.pgSpecific = NewPgSpecificClient(c)
	}
	return c.pgSpecific
}

// GetStoreID 상점 ID를 반환합니다.
func (c *Client) GetStoreID() string {
	return c.storeID
}

// Request HTTP 요청을 수행합니다.
func (c *Client) Request(ctx context.Context, method, path string, query url.Values, body interface{}, result interface{}) error {
	// URL 생성
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("URL 파싱 실패: %w", err)
	}

	if query != nil {
		u.RawQuery = query.Encode()
	}

	// 요청 본문 생성
	var bodyReader io.Reader
	if body != nil {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("요청 본문 직렬화 실패: %w", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	}

	// HTTP 요청 생성
	req, err := http.NewRequestWithContext(ctx, method, u.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("HTTP 요청 생성 실패: %w", err)
	}

	// 헤더 설정
	req.Header.Set("Authorization", "PortOne "+c.secret)
	req.Header.Set("User-Agent", UserAgent)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 요청 수행
	resp, err := c.http.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP 요청 실패: %w", err)
	}
	defer resp.Body.Close()

	// 응답 읽기
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("응답 읽기 실패: %w", err)
	}

	// 에러 처리
	if resp.StatusCode >= 400 {
		return NewAPIError(resp.StatusCode, respBody)
	}

	// 결과 파싱
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("응답 파싱 실패: %w", err)
		}
	}

	return nil
}

// Get GET 요청을 수행합니다.
func (c *Client) Get(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.Request(ctx, http.MethodGet, path, query, nil, result)
}

// Post POST 요청을 수행합니다.
func (c *Client) Post(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Request(ctx, http.MethodPost, path, nil, body, result)
}

// Delete DELETE 요청을 수행합니다.
func (c *Client) Delete(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.Request(ctx, http.MethodDelete, path, query, nil, result)
}

// Patch PATCH 요청을 수행합니다.
func (c *Client) Patch(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Request(ctx, http.MethodPatch, path, nil, body, result)
}

// Put PUT 요청을 수행합니다.
func (c *Client) Put(ctx context.Context, path string, body interface{}, result interface{}) error {
	return c.Request(ctx, http.MethodPut, path, nil, body, result)
}

// GetBankInfos 은행 정보를 조회합니다.
func (c *Client) GetBankInfos(ctx context.Context) (*common.GetBankInfosResponse, error) {
	var result common.GetBankInfosResponse
	if err := c.Get(ctx, "/banks", nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
