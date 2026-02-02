package common

// InvalidRequestError 요청된 입력 정보가 유효하지 않은 경우
type InvalidRequestError struct {
	Type    string  `json:"type"` // "INVALID_REQUEST"
	Message *string `json:"message,omitempty"`
}

// UnauthorizedError 인증 정보가 올바르지 않은 경우
type UnauthorizedError struct {
	Type    string  `json:"type"` // "UNAUTHORIZED"
	Message *string `json:"message,omitempty"`
}

// ForbiddenError 요청이 거절된 경우
type ForbiddenError struct {
	Type    string  `json:"type"` // "FORBIDDEN"
	Message *string `json:"message,omitempty"`
}

// ChannelNotFoundError 요청된 채널을 찾을 수 없는 경우
type ChannelNotFoundError struct {
	Type    string  `json:"type"` // "CHANNEL_NOT_FOUND"
	Message *string `json:"message,omitempty"`
}

// BillingKeyNotFoundError 빌링키가 존재하지 않는 경우
type BillingKeyNotFoundError struct {
	Type    string  `json:"type"` // "BILLING_KEY_NOT_FOUND"
	Message *string `json:"message,omitempty"`
}

// BillingKeyAlreadyDeletedError 빌링키가 이미 삭제된 경우
type BillingKeyAlreadyDeletedError struct {
	Type    string  `json:"type"` // "BILLING_KEY_ALREADY_DELETED"
	Message *string `json:"message,omitempty"`
}

// PgProviderError PG사에서 오류를 전달한 경우
type PgProviderError struct {
	Type       string  `json:"type"` // "PG_PROVIDER_ERROR"
	Message    *string `json:"message,omitempty"`
	PgCode     *string `json:"pgCode,omitempty"`
	PgMessage  *string `json:"pgMessage,omitempty"`
}

// PaymentScheduleAlreadyExistsError 결제 예약건이 이미 존재하는 경우
type PaymentScheduleAlreadyExistsError struct {
	Type    string  `json:"type"` // "PAYMENT_SCHEDULE_ALREADY_EXISTS"
	Message *string `json:"message,omitempty"`
}

// MaxTransactionCountReachedError 해당 결제건의 최대 거래 횟수가 초과된 경우
type MaxTransactionCountReachedError struct {
	Type    string  `json:"type"` // "MAX_TRANSACTION_COUNT_REACHED"
	Message *string `json:"message,omitempty"`
}

// SumOfPartsExceedsTotalAmountError 면세 금액 등 합계가 결제 금액을 초과한 경우
type SumOfPartsExceedsTotalAmountError struct {
	Type    string  `json:"type"` // "SUM_OF_PARTS_EXCEEDS_TOTAL_AMOUNT"
	Message *string `json:"message,omitempty"`
}

// B2bNotEnabledError B2B 기능이 활성화되지 않은 경우
type B2bNotEnabledError struct {
	Type    string  `json:"type"` // "B2B_NOT_ENABLED"
	Message *string `json:"message,omitempty"`
}

// B2bExternalServiceError 외부 B2B 서비스에서 에러가 발생한 경우
type B2bExternalServiceError struct {
	Type    string  `json:"type"` // "B2B_EXTERNAL_SERVICE_ERROR"
	Message *string `json:"message,omitempty"`
}

// InformationMismatchError 정보 불일치 에러
type InformationMismatchError struct {
	Type    string  `json:"type"` // "INFORMATION_MISMATCH"
	Message *string `json:"message,omitempty"`
}
