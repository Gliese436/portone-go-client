package portonev2

import (
	"encoding/json"
	"fmt"
)

// PortOneError 는 PortOne API 에러의 기본 타입입니다.
type PortOneError struct {
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
}

func (e *PortOneError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s: %s", e.Type, e.Message)
	}
	return e.Type
}

// UnmarshalError JSON 데이터를 PortOneError로 파싱합니다.
func UnmarshalError(data []byte) (*PortOneError, error) {
	var e PortOneError
	if err := json.Unmarshal(data, &e); err != nil {
		return nil, err
	}
	return &e, nil
}

// APIError 는 HTTP 응답에서 발생하는 에러입니다.
type APIError struct {
	StatusCode int
	PortOneError
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("[%d] %s: %s", e.StatusCode, e.Type, e.Message)
	}
	return fmt.Sprintf("[%d] %s", e.StatusCode, e.Type)
}

// NewAPIError HTTP 상태 코드와 응답 본문으로 APIError를 생성합니다.
func NewAPIError(statusCode int, body []byte) *APIError {
	var portoneErr PortOneError
	if err := json.Unmarshal(body, &portoneErr); err != nil {
		portoneErr = PortOneError{
			Type:    "UNKNOWN_ERROR",
			Message: string(body),
		}
	}
	return &APIError{
		StatusCode:   statusCode,
		PortOneError: portoneErr,
	}
}
