package portonev2

import (
	"encoding/json"
	"errors"
	"fmt"
)

// PortOneError 는 포트원 API 오류 응답의 공통 필드입니다.
type PortOneError struct {
	// Type 오류 타입 (예: "PAYMENT_NOT_FOUND")
	Type string `json:"type"`
	// Message 오류 메시지
	Message string `json:"message,omitempty"`
}

func (e *PortOneError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.Type, e.Message)
	}
	return e.Type
}

// APIError 는 2xx 가 아닌 HTTP 응답을 나타냅니다.
//
// 각 API 메서드 문서에 적힌 오류 타입(예: payment.GetPaymentError)으로
// Decode 하면 오류별 추가 필드를 확인할 수 있습니다.
//
//	var apiErr *portonev2.APIError
//	if errors.As(err, &apiErr) && apiErr.Type == "PAYMENT_NOT_FOUND" { ... }
type APIError struct {
	StatusCode int
	PortOneError
	// Body 응답 본문 원문
	Body []byte
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Type, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Type)
}

// Decode 응답 본문을 target 으로 역직렬화합니다.
func (e *APIError) Decode(target any) error {
	return json.Unmarshal(e.Body, target)
}

// NewAPIError HTTP 상태 코드와 응답 본문으로 APIError 를 생성합니다.
func NewAPIError(statusCode int, body []byte) *APIError {
	var portoneErr PortOneError
	if err := json.Unmarshal(body, &portoneErr); err != nil || portoneErr.Type == "" {
		portoneErr = PortOneError{
			Type:    "UNKNOWN_ERROR",
			Message: string(body),
		}
	}
	return &APIError{
		StatusCode:   statusCode,
		PortOneError: portoneErr,
		Body:         body,
	}
}

// AsAPIError err 가 *APIError 이면 이를 반환합니다.
func AsAPIError(err error) (*APIError, bool) {
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr, true
	}
	return nil, false
}
