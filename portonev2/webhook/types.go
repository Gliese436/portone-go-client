package webhook

// WebhookType 웹훅 타입
type WebhookType string

const (
	WebhookTypeTransactionReady                WebhookType = "Transaction.Ready"
	WebhookTypeTransactionPaid                 WebhookType = "Transaction.Paid"
	WebhookTypeTransactionVirtualAccountIssued WebhookType = "Transaction.VirtualAccountIssued"
	WebhookTypeTransactionPartialCancelled     WebhookType = "Transaction.PartialCancelled"
	WebhookTypeTransactionCancelled            WebhookType = "Transaction.Cancelled"
	WebhookTypeTransactionFailed               WebhookType = "Transaction.Failed"
	WebhookTypeTransactionPayPending           WebhookType = "Transaction.PayPending"
	WebhookTypeTransactionDisputeCreated       WebhookType = "Transaction.DisputeCreated"
	WebhookTypeTransactionDisputeResolved      WebhookType = "Transaction.DisputeResolved"
	WebhookTypeTransactionCancelPending        WebhookType = "Transaction.CancelPending"
	WebhookTypeTransactionConfirm              WebhookType = "Transaction.Confirm"
	WebhookTypeBillingKeyReady                 WebhookType = "BillingKey.Ready"
	WebhookTypeBillingKeyIssued                WebhookType = "BillingKey.Issued"
	WebhookTypeBillingKeyFailed                WebhookType = "BillingKey.Failed"
	WebhookTypeBillingKeyDeleted               WebhookType = "BillingKey.Deleted"
	WebhookTypeBillingKeyUpdated               WebhookType = "BillingKey.Updated"
)

// WebhookTransactionData 트랜잭션 웹훅 데이터
type WebhookTransactionData struct {
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 트랜잭션 아이디
	TransactionId string `json:"transactionId"`
	// 취소 아이디 (취소 관련 이벤트시)
	CancellationId *string `json:"cancellationId,omitempty"`
}

// WebhookBillingKeyData 빌링키 웹훅 데이터
type WebhookBillingKeyData struct {
	// 빌링키
	BillingKey string `json:"billingKey"`
}

// Webhook 웹훅 페이로드
type Webhook struct {
	// 웹훅 타입
	Type WebhookType `json:"type"`
	// 타임스탬프 (RFC 3339)
	Timestamp string `json:"timestamp"`
	// 트랜잭션 데이터 (Transaction 타입일 때)
	Data WebhookTransactionData `json:"data,omitempty"`
}

// IsTransaction 트랜잭션 관련 웹훅인지 확인
func (w *Webhook) IsTransaction() bool {
	switch w.Type {
	case WebhookTypeTransactionReady,
		WebhookTypeTransactionPaid,
		WebhookTypeTransactionVirtualAccountIssued,
		WebhookTypeTransactionPartialCancelled,
		WebhookTypeTransactionCancelled,
		WebhookTypeTransactionFailed,
		WebhookTypeTransactionPayPending,
		WebhookTypeTransactionDisputeCreated,
		WebhookTypeTransactionDisputeResolved,
		WebhookTypeTransactionCancelPending,
		WebhookTypeTransactionConfirm:
		return true
	}
	return false
}

// IsBillingKey 빌링키 관련 웹훅인지 확인
func (w *Webhook) IsBillingKey() bool {
	switch w.Type {
	case WebhookTypeBillingKeyReady,
		WebhookTypeBillingKeyIssued,
		WebhookTypeBillingKeyFailed,
		WebhookTypeBillingKeyDeleted,
		WebhookTypeBillingKeyUpdated:
		return true
	}
	return false
}

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
