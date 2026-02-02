package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	// WebhookToleranceInSeconds 웹훅 타임스탬프 허용 범위 (5분)
	WebhookToleranceInSeconds = 5 * 60

	// WebhookSecretPrefix 웹훅 시크릿 접두사
	WebhookSecretPrefix = "whsec_"
)

// WebhookVerificationError 웹훅 검증 오류
type WebhookVerificationError struct {
	Reason WebhookVerificationFailureReason
}

func (e *WebhookVerificationError) Error() string {
	return e.Reason.GetMessage()
}

// InvalidInputError 잘못된 입력 오류
type InvalidInputError struct {
	Message string
}

func (e *InvalidInputError) Error() string {
	return e.Message
}

// Verify 웹훅 페이로드를 검증합니다.
//
// secret: 웹훅 시크릿 (base64 인코딩된 문자열 또는 "whsec_" 접두사 포함)
// payload: 웹훅 페이로드 (JSON 문자열)
// headers: HTTP 요청 헤더
func Verify(secret string, payload []byte, headers http.Header) (*Webhook, error) {
	return VerifyWithTolerance(secret, payload, headers, time.Duration(WebhookToleranceInSeconds)*time.Second)
}

// VerifyWithTolerance 웹훅 페이로드를 검증합니다 (커스텀 허용 범위 지정).
func VerifyWithTolerance(secret string, payload []byte, headers http.Header, tolerance time.Duration) (*Webhook, error) {
	if len(payload) == 0 {
		return nil, &InvalidInputError{Message: "payload가 비어있습니다."}
	}

	// 필수 헤더 추출
	msgId := findHeaderValue(headers, "webhook-id")
	msgSignature := findHeaderValue(headers, "webhook-signature")
	msgTimestamp := findHeaderValue(headers, "webhook-timestamp")

	if msgId == "" || msgSignature == "" || msgTimestamp == "" {
		return nil, &WebhookVerificationError{Reason: WebhookVerificationFailureReasonMISSING_REQUIRED_HEADERS}
	}

	// 타임스탬프 검증
	if err := verifyTimestamp(msgTimestamp, tolerance); err != nil {
		return nil, err
	}

	// 시크릿 파싱
	secretBytes, err := parseSecret(secret)
	if err != nil {
		return nil, err
	}

	// 서명 생성
	expectedSignature := sign(secretBytes, msgId, msgTimestamp, payload)

	// 서명 검증
	signatures := strings.Split(msgSignature, " ")
	for _, versionedSignature := range signatures {
		parts := strings.SplitN(versionedSignature, ",", 2)
		if len(parts) < 2 {
			continue
		}

		version := parts[0]
		signature := parts[1]

		if version != "v1" {
			continue
		}

		signatureDecoded, err := base64.StdEncoding.DecodeString(signature)
		if err != nil {
			continue
		}

		if timingSafeEqual(signatureDecoded, expectedSignature) {
			var webhook Webhook
			if err := json.Unmarshal(payload, &webhook); err != nil {
				return nil, &InvalidInputError{Message: fmt.Sprintf("페이로드 파싱 실패: %v", err)}
			}
			return &webhook, nil
		}
	}

	return nil, &WebhookVerificationError{Reason: WebhookVerificationFailureReasonNO_MATCHING_SIGNATURE}
}

// findHeaderValue 헤더 값을 찾습니다 (대소문자 구분 없음, 단일 값만).
func findHeaderValue(headers http.Header, name string) string {
	nameLower := strings.ToLower(name)
	for key, values := range headers {
		if strings.ToLower(key) == nameLower {
			if len(values) == 1 {
				return values[0]
			}
			// 중복 값은 무시
			return ""
		}
	}
	return ""
}

// parseSecret 시크릿을 파싱합니다.
func parseSecret(secret string) ([]byte, error) {
	if secret == "" {
		return nil, &InvalidInputError{Message: "시크릿은 비어 있을 수 없습니다."}
	}

	secretBase64 := secret
	if strings.HasPrefix(secret, WebhookSecretPrefix) {
		secretBase64 = secret[len(WebhookSecretPrefix):]
	}

	decoded, err := base64.StdEncoding.DecodeString(secretBase64)
	if err != nil {
		return nil, &InvalidInputError{Message: "secret 파라미터가 올바른 Base64 문자열이 아닙니다."}
	}

	if len(decoded) == 0 {
		return nil, &InvalidInputError{Message: "시크릿은 비어 있을 수 없습니다."}
	}

	return decoded, nil
}

// sign HMAC-SHA256 서명을 생성합니다.
func sign(secret []byte, msgId, msgTimestamp string, payload []byte) []byte {
	toSign := fmt.Sprintf("%s.%s.%s", msgId, msgTimestamp, string(payload))

	h := hmac.New(sha256.New, secret)
	h.Write([]byte(toSign))
	return h.Sum(nil)
}

// verifyTimestamp 타임스탬프를 검증합니다.
func verifyTimestamp(timestampHeader string, tolerance time.Duration) error {
	timestamp, err := strconv.ParseInt(timestampHeader, 10, 64)
	if err != nil {
		return &WebhookVerificationError{Reason: WebhookVerificationFailureReasonINVALID_SIGNATURE}
	}

	now := time.Now().Unix()
	toleranceSec := int64(tolerance.Seconds())

	if now-timestamp > toleranceSec {
		return &WebhookVerificationError{Reason: WebhookVerificationFailureReasonTIMESTAMP_TOO_OLD}
	}
	if timestamp > now+toleranceSec {
		return &WebhookVerificationError{Reason: WebhookVerificationFailureReasonTIMESTAMP_TOO_NEW}
	}

	return nil
}

// timingSafeEqual 타이밍 안전 비교를 수행합니다.
func timingSafeEqual(a, b []byte) bool {
	return subtle.ConstantTimeCompare(a, b) == 1
}

// IsWebhookVerificationError 웹훅 검증 오류인지 확인합니다.
func IsWebhookVerificationError(err error) bool {
	var verifyErr *WebhookVerificationError
	return errors.As(err, &verifyErr)
}

// IsInvalidInputError 잘못된 입력 오류인지 확인합니다.
func IsInvalidInputError(err error) bool {
	var inputErr *InvalidInputError
	return errors.As(err, &inputErr)
}
