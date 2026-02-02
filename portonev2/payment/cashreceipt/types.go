package cashreceipt

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// CashReceiptStatus 현금영수증 상태
type CashReceiptStatus string

const (
	CashReceiptStatusISSUED       CashReceiptStatus = "ISSUED"
	CashReceiptStatusCANCELLED    CashReceiptStatus = "CANCELLED"
	CashReceiptStatusISSUE_FAILED CashReceiptStatus = "ISSUE_FAILED"
)

// CashReceipt 현금영수증 내역
type CashReceipt struct {
	// 현금영수증 상태
	Status CashReceiptStatus `json:"status"`
	// 고객사 아이디
	MerchantId string `json:"merchantId"`
	// 상점 아이디
	StoreId string `json:"storeId"`
	// 결제 건 아이디
	PaymentId *string `json:"paymentId,omitempty"`
	// 현금영수증 발급에 사용된 채널
	Channel *common.SelectedChannel `json:"channel,omitempty"`
	// 현금영수증 발급 금액
	Amount int64 `json:"amount"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 부가세
	VatAmount *int64 `json:"vatAmount,omitempty"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 주문명
	OrderName string `json:"orderName"`
	// 수동 발급 여부
	IsManual bool `json:"isManual"`
	// 현금영수증 유형
	Type *common.CashReceiptType `json:"type,omitempty"`
	// PG사 영수증 발급 아이디
	PgReceiptId *string `json:"pgReceiptId,omitempty"`
	// 승인 번호
	IssueNumber *string `json:"issueNumber,omitempty"`
	// 현금영수증 URL
	Url *string `json:"url,omitempty"`
	// 발급 시점 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
	// 취소 시점 (RFC 3339)
	CancelledAt *string `json:"cancelledAt,omitempty"`
}

// IsIssued 발급 완료 상태인지 확인
func (c *CashReceipt) IsIssued() bool {
	return c.Status == CashReceiptStatusISSUED
}

// IsCancelled 취소 상태인지 확인
func (c *CashReceipt) IsCancelled() bool {
	return c.Status == CashReceiptStatusCANCELLED
}

// CashReceiptSummary 현금영수증 요약 정보
type CashReceiptSummary struct {
	// 현금영수증 발급 번호
	IssueNumber *string `json:"issueNumber,omitempty"`
	// 현금영수증 URL
	Url *string `json:"url,omitempty"`
	// PG사 영수증 발급 아이디
	PgReceiptId *string `json:"pgReceiptId,omitempty"`
}

// IssueCashReceiptBody 현금영수증 발급 요청
type IssueCashReceiptBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 채널 키
	ChannelKey string `json:"channelKey"`
	// 현금영수증 유형
	Type common.CashReceiptType `json:"type"`
	// 주문명
	OrderName string `json:"orderName"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 현금영수증 발급 금액
	Amount common.PaymentAmountInput `json:"amount"`
	// 상품 유형
	ProductType *common.PaymentProductType `json:"productType,omitempty"`
	// 고객 정보
	Customer CashReceiptCustomerInput `json:"customer"`
	// 결제일 (yyyyMMdd 형식)
	PaidAt *string `json:"paidAt,omitempty"`
}

// CashReceiptCustomerInput 현금영수증 고객 입력
type CashReceiptCustomerInput struct {
	// 식별 번호 (휴대폰 번호, 사업자등록번호 등)
	IdentityNumber string `json:"identityNumber"`
	// 이름
	Name *string `json:"name,omitempty"`
	// 이메일
	Email *string `json:"email,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// IssueCashReceiptResponse 현금영수증 발급 응답
type IssueCashReceiptResponse struct {
	// 현금영수증 정보
	CashReceipt CashReceipt `json:"cashReceipt"`
}

// CancelCashReceiptByPaymentIdResponse 현금영수증 취소 응답
type CancelCashReceiptByPaymentIdResponse struct {
	// 취소 금액
	CancelledAmount int64 `json:"cancelledAmount"`
	// 취소 시점 (RFC 3339)
	CancelledAt string `json:"cancelledAt"`
}

// GetCashReceiptResponse 현금영수증 조회 응답
type GetCashReceiptResponse = CashReceipt

// GetCashReceiptsResponse 현금영수증 목록 조회 응답
type GetCashReceiptsResponse struct {
	// 현금영수증 목록
	Items []CashReceipt `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// CashReceiptSortBy 현금영수증 정렬 기준
type CashReceiptSortBy string

const (
	CashReceiptSortByISSUED_AT    CashReceiptSortBy = "ISSUED_AT"
	CashReceiptSortByCANCELLED_AT CashReceiptSortBy = "CANCELLED_AT"
)

// CashReceiptSortInput 현금영수증 정렬 입력
type CashReceiptSortInput struct {
	// 정렬 기준
	By *CashReceiptSortBy `json:"by,omitempty"`
	// 정렬 순서
	Order *common.SortOrder `json:"order,omitempty"`
}

// CashReceiptTimeRangeField 현금영수증 시간 범위 필드
type CashReceiptTimeRangeField string

const (
	CashReceiptTimeRangeFieldISSUED_AT    CashReceiptTimeRangeField = "ISSUED_AT"
	CashReceiptTimeRangeFieldCANCELLED_AT CashReceiptTimeRangeField = "CANCELLED_AT"
)

// CashReceiptFilterInput 현금영수증 필터 입력
type CashReceiptFilterInput struct {
	// 상점 아이디 목록
	StoreIds []string `json:"storeIds,omitempty"`
	// 시간 범위 필드
	TimeRangeField *CashReceiptTimeRangeField `json:"timeRangeField,omitempty"`
	// 시작 시간 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 종료 시간 (RFC 3339)
	Until *string `json:"until,omitempty"`
	// 현금영수증 상태 목록
	Statuses []CashReceiptStatus `json:"statuses,omitempty"`
	// 현금영수증 유형 목록
	Types []common.CashReceiptType `json:"types,omitempty"`
	// PG사 목록
	PgProviders []common.PgProvider `json:"pgProviders,omitempty"`
	// 채널 그룹 아이디 목록
	ChannelGroupIds []string `json:"channelGroupIds,omitempty"`
	// 결제 건 아이디
	PaymentId *string `json:"paymentId,omitempty"`
}

// GetCashReceiptsBody 현금영수증 목록 조회 요청
type GetCashReceiptsBody struct {
	// 페이지 정보
	Page *common.PageInput `json:"page,omitempty"`
	// 정렬 정보
	Sort *CashReceiptSortInput `json:"sort,omitempty"`
	// 필터 정보
	Filter *CashReceiptFilterInput `json:"filter,omitempty"`
}
