// Package webhook 은 포트원 웹훅 페이로드 타입(types.gen.go)과 서명 검증 함수를 제공합니다.
package webhook

import "strings"

// WebhookVerificationFailureReason 웹훅 검증 실패 사유
type WebhookVerificationFailureReason string

const (
	WebhookVerificationFailureReasonMISSING_REQUIRED_HEADERS WebhookVerificationFailureReason = "MISSING_REQUIRED_HEADERS"
	WebhookVerificationFailureReasonNO_MATCHING_SIGNATURE    WebhookVerificationFailureReason = "NO_MATCHING_SIGNATURE"
	WebhookVerificationFailureReasonINVALID_SIGNATURE        WebhookVerificationFailureReason = "INVALID_SIGNATURE"
	WebhookVerificationFailureReasonTIMESTAMP_TOO_OLD        WebhookVerificationFailureReason = "TIMESTAMP_TOO_OLD"
	WebhookVerificationFailureReasonTIMESTAMP_TOO_NEW        WebhookVerificationFailureReason = "TIMESTAMP_TOO_NEW"
)

// GetMessage 웹훅 검증 실패 사유 메시지
func (r WebhookVerificationFailureReason) GetMessage() string {
	switch r {
	case WebhookVerificationFailureReasonMISSING_REQUIRED_HEADERS:
		return "필수 헤더가 누락되었습니다."
	case WebhookVerificationFailureReasonNO_MATCHING_SIGNATURE:
		return "올바른 웹훅 시그니처를 찾을 수 없습니다."
	case WebhookVerificationFailureReasonINVALID_SIGNATURE:
		return "웹훅 시그니처가 유효하지 않습니다."
	case WebhookVerificationFailureReasonTIMESTAMP_TOO_OLD:
		return "웹훅 시그니처의 타임스탬프가 만료 기한을 초과했습니다."
	case WebhookVerificationFailureReasonTIMESTAMP_TOO_NEW:
		return "웹훅 시그니처의 타임스탬프가 미래 시간으로 설정되어 있습니다."
	}
	return "알 수 없는 오류"
}

// IsTransaction 결제 관련 웹훅("Transaction.*")인지 확인합니다.
func (w *Webhook) IsTransaction() bool {
	return strings.HasPrefix(w.Type, "Transaction.")
}

// IsBillingKey 빌링키 관련 웹훅("BillingKey.*")인지 확인합니다.
func (w *Webhook) IsBillingKey() bool {
	return strings.HasPrefix(w.Type, "BillingKey.")
}
