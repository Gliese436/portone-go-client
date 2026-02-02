package payment

import (
	"encoding/json"

	"github.com/gliese436/portone-go-client/portonev2/common"
)

// PaymentStatus 결제 상태
type PaymentStatus string

const (
	PaymentStatusREADY                 PaymentStatus = "READY"
	PaymentStatusPAID                  PaymentStatus = "PAID"
	PaymentStatusFAILED                PaymentStatus = "FAILED"
	PaymentStatusCANCELLED             PaymentStatus = "CANCELLED"
	PaymentStatusPARTIAL_CANCELLED     PaymentStatus = "PARTIAL_CANCELLED"
	PaymentStatusPAY_PENDING           PaymentStatus = "PAY_PENDING"
	PaymentStatusVIRTUAL_ACCOUNT_ISSUED PaymentStatus = "VIRTUAL_ACCOUNT_ISSUED"
)

// PaymentMethodType 결제수단 타입
type PaymentMethodType string

const (
	PaymentMethodTypeCard              PaymentMethodType = "PaymentMethodCard"
	PaymentMethodTypeConvenienceStore  PaymentMethodType = "PaymentMethodConvenienceStore"
	PaymentMethodTypeEasyPay           PaymentMethodType = "PaymentMethodEasyPay"
	PaymentMethodTypeGiftCertificate   PaymentMethodType = "PaymentMethodGiftCertificate"
	PaymentMethodTypeMobile            PaymentMethodType = "PaymentMethodMobile"
	PaymentMethodTypeTransfer          PaymentMethodType = "PaymentMethodTransfer"
	PaymentMethodTypeVirtualAccount    PaymentMethodType = "PaymentMethodVirtualAccount"
)

// PaymentAmount 결제 금액 세부 정보
type PaymentAmount struct {
	// 총 결제금액
	Total int64 `json:"total"`
	// 면세액
	TaxFree int64 `json:"taxFree"`
	// 부가세액
	Vat *int64 `json:"vat,omitempty"`
	// 공급가액
	Supply *int64 `json:"supply,omitempty"`
	// 총 할인금액
	Discount int64 `json:"discount"`
	// 실제 결제금액
	Paid int64 `json:"paid"`
	// 총 취소금액
	Cancelled int64 `json:"cancelled"`
	// 총 취소금액 중 면세액
	CancelledTaxFree int64 `json:"cancelledTaxFree"`
}

// PaymentInstallment 할부 정보
type PaymentInstallment struct {
	// 할부 개월 수
	Month int `json:"month"`
	// 무이자할부 여부
	IsInterestFree bool `json:"isInterestFree"`
}

// PaymentMethodCard 결제수단 카드 정보
type PaymentMethodCard struct {
	Type string `json:"type"` // "PaymentMethodCard"
	// 카드 상세 정보
	Card *common.Card `json:"card,omitempty"`
	// 승인 번호
	ApprovalNumber *string `json:"approvalNumber,omitempty"`
	// 할부 정보
	Installment *PaymentInstallment `json:"installment,omitempty"`
	// 카드 포인트 사용여부
	PointUsed *bool `json:"pointUsed,omitempty"`
}

// PaymentMethodTransfer 계좌 이체 상세 정보
type PaymentMethodTransfer struct {
	Type string `json:"type"` // "PaymentMethodTransfer"
	// 은행
	Bank *common.Bank `json:"bank,omitempty"`
}

// PaymentMethodVirtualAccount 가상계좌 상세 정보
type PaymentMethodVirtualAccount struct {
	Type string `json:"type"` // "PaymentMethodVirtualAccount"
	// 은행
	Bank *common.Bank `json:"bank,omitempty"`
	// 계좌번호
	AccountNumber *string `json:"accountNumber,omitempty"`
	// 계좌 유형
	AccountType *string `json:"accountType,omitempty"`
	// 계좌 소유자명
	RemitteeName *string `json:"remitteeName,omitempty"`
	// 송금인명
	RemitterName *string `json:"remitterName,omitempty"`
	// 입금 만료 시점 (RFC 3339)
	ExpiredAt *string `json:"expiredAt,omitempty"`
	// 입금 시점 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
	// 환불 계좌
	RefundStatus *string `json:"refundStatus,omitempty"`
}

// PaymentMethodMobile 모바일 상세 정보
type PaymentMethodMobile struct {
	Type string `json:"type"` // "PaymentMethodMobile"
	// 통신사
	Carrier *string `json:"carrier,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// PaymentMethodGiftCertificate 상품권 상세 정보
type PaymentMethodGiftCertificate struct {
	Type string `json:"type"` // "PaymentMethodGiftCertificate"
	// 상품권 종류
	GiftCertificateType *string `json:"giftCertificateType,omitempty"`
	// 승인 번호
	ApprovalNumber *string `json:"approvalNumber,omitempty"`
}

// PaymentMethodEasyPay 간편 결제 상세 정보
type PaymentMethodEasyPay struct {
	Type string `json:"type"` // "PaymentMethodEasyPay"
	// 간편결제사
	Provider *common.EasyPayProvider `json:"provider,omitempty"`
	// 간편결제 수단
	EasyPayMethod *json.RawMessage `json:"easyPayMethod,omitempty"`
}

// PaymentMethodConvenienceStore 편의점 결제 상세 정보
type PaymentMethodConvenienceStore struct {
	Type string `json:"type"` // "PaymentMethodConvenienceStore"
	// 편의점 종류
	ConvenienceStoreType *string `json:"convenienceStoreType,omitempty"`
	// 결제 코드
	PaymentCode *string `json:"paymentCode,omitempty"`
}

// PaymentMethod 결제수단 정보 (discriminated union)
type PaymentMethod struct {
	Type string `json:"type"`
	// Card fields
	Card           *common.Card        `json:"card,omitempty"`
	ApprovalNumber *string             `json:"approvalNumber,omitempty"`
	Installment    *PaymentInstallment `json:"installment,omitempty"`
	PointUsed      *bool               `json:"pointUsed,omitempty"`
	// Transfer fields
	Bank *common.Bank `json:"bank,omitempty"`
	// VirtualAccount fields
	AccountNumber *string `json:"accountNumber,omitempty"`
	AccountType   *string `json:"accountType,omitempty"`
	RemitteeName  *string `json:"remitteeName,omitempty"`
	RemitterName  *string `json:"remitterName,omitempty"`
	ExpiredAt     *string `json:"expiredAt,omitempty"`
	IssuedAt      *string `json:"issuedAt,omitempty"`
	RefundStatus  *string `json:"refundStatus,omitempty"`
	// Mobile fields
	Carrier     *string `json:"carrier,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// GiftCertificate fields
	GiftCertificateType *string `json:"giftCertificateType,omitempty"`
	// EasyPay fields
	Provider      *common.EasyPayProvider `json:"provider,omitempty"`
	EasyPayMethod *json.RawMessage        `json:"easyPayMethod,omitempty"`
	// ConvenienceStore fields
	ConvenienceStoreType *string `json:"convenienceStoreType,omitempty"`
	PaymentCode          *string `json:"paymentCode,omitempty"`
}

// IsCard 카드 결제인지 확인
func (m *PaymentMethod) IsCard() bool {
	return m.Type == string(PaymentMethodTypeCard)
}

// IsTransfer 계좌이체인지 확인
func (m *PaymentMethod) IsTransfer() bool {
	return m.Type == string(PaymentMethodTypeTransfer)
}

// IsVirtualAccount 가상계좌인지 확인
func (m *PaymentMethod) IsVirtualAccount() bool {
	return m.Type == string(PaymentMethodTypeVirtualAccount)
}

// PaymentFailure 결제 실패 정보
type PaymentFailure struct {
	// 실패 사유
	Reason *string `json:"reason,omitempty"`
	// PG사 실패 코드
	PgCode *string `json:"pgCode,omitempty"`
	// PG사 실패 메시지
	PgMessage *string `json:"pgMessage,omitempty"`
}

// PaymentWebhookStatus 웹훅 상태
type PaymentWebhookStatus string

const (
	PaymentWebhookStatusSUCCEEDED  PaymentWebhookStatus = "SUCCEEDED"
	PaymentWebhookStatusFAILED    PaymentWebhookStatus = "FAILED"
	PaymentWebhookStatusINPROGRESS PaymentWebhookStatus = "IN_PROGRESS"
)

// PaymentWebhookTrigger 웹훅 트리거
type PaymentWebhookTrigger string

const (
	PaymentWebhookTriggerMANUAL     PaymentWebhookTrigger = "MANUAL"
	PaymentWebhookTriggerAUTOMATIC  PaymentWebhookTrigger = "AUTOMATIC"
	PaymentWebhookTriggerVIRTUAL   PaymentWebhookTrigger = "VIRTUAL"
)

// PaymentWebhookPaymentStatus 웹훅 결제 상태
type PaymentWebhookPaymentStatus string

const (
	PaymentWebhookPaymentStatusREADY                 PaymentWebhookPaymentStatus = "READY"
	PaymentWebhookPaymentStatusPAID                  PaymentWebhookPaymentStatus = "PAID"
	PaymentWebhookPaymentStatusFAILED                PaymentWebhookPaymentStatus = "FAILED"
	PaymentWebhookPaymentStatusCANCELLED             PaymentWebhookPaymentStatus = "CANCELLED"
	PaymentWebhookPaymentStatusPARTIAL_CANCELLED     PaymentWebhookPaymentStatus = "PARTIAL_CANCELLED"
	PaymentWebhookPaymentStatusPAY_PENDING           PaymentWebhookPaymentStatus = "PAY_PENDING"
	PaymentWebhookPaymentStatusVIRTUAL_ACCOUNT_ISSUED PaymentWebhookPaymentStatus = "VIRTUAL_ACCOUNT_ISSUED"
)

// PaymentWebhookRequest 웹훅 요청 정보
type PaymentWebhookRequest struct {
	// 요청 헤더
	Headers *string `json:"headers,omitempty"`
	// 요청 본문
	Body *string `json:"body,omitempty"`
	// 요청 시점 (RFC 3339)
	RequestedAt *string `json:"requestedAt,omitempty"`
}

// PaymentWebhookResponse 웹훅 응답 정보
type PaymentWebhookResponse struct {
	// 응답 코드
	Code *string `json:"code,omitempty"`
	// 응답 헤더
	Headers *string `json:"headers,omitempty"`
	// 응답 본문
	Body *string `json:"body,omitempty"`
	// 응답 시점 (RFC 3339)
	RespondedAt *string `json:"respondedAt,omitempty"`
}

// PaymentWebhook 웹훅 발송 내역
type PaymentWebhook struct {
	// 웹훅 발송 시 결제 건 상태
	PaymentStatus *PaymentWebhookPaymentStatus `json:"paymentStatus,omitempty"`
	// 웹훅 아이디
	ID string `json:"id"`
	// 웹훅 상태
	Status *PaymentWebhookStatus `json:"status,omitempty"`
	// 웹훅이 발송된 url
	URL string `json:"url"`
	// 비동기 웹훅 여부
	IsAsync *bool `json:"isAsync,omitempty"`
	// 현재 발송 횟수
	CurrentExecutionCount *int `json:"currentExecutionCount,omitempty"`
	// 최대 발송 횟수
	MaxExecutionCount *int `json:"maxExecutionCount,omitempty"`
	// 웹훅 실행 맥락
	Trigger *PaymentWebhookTrigger `json:"trigger,omitempty"`
	// 웹훅 요청 정보
	Request *PaymentWebhookRequest `json:"request,omitempty"`
	// 웹훅 응답 정보
	Response *PaymentWebhookResponse `json:"response,omitempty"`
	// 웹훅 처리 시작 시점 (RFC 3339)
	TriggeredAt *string `json:"triggeredAt,omitempty"`
}

// DisputeStatus 분쟁 상태
type DisputeStatus string

const (
	DisputeStatusPENDING  DisputeStatus = "PENDING"
	DisputeStatusRESOLVED DisputeStatus = "RESOLVED"
)

// Dispute 분쟁 내역
type Dispute struct {
	// 분쟁 상태
	Status DisputeStatus `json:"status"`
	// PG사 분쟁 아이디
	PgDisputeId *string `json:"pgDisputeId,omitempty"`
	// 분쟁 사유
	Reason string `json:"reason"`
	// 분쟁 발생 시각 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 분쟁 해소 시각 (RFC 3339)
	ResolvedAt *string `json:"resolvedAt,omitempty"`
}

// PaymentCancellationStatus 취소 상태
type PaymentCancellationStatus string

const (
	PaymentCancellationStatusFAILED    PaymentCancellationStatus = "FAILED"
	PaymentCancellationStatusREQUESTED PaymentCancellationStatus = "REQUESTED"
	PaymentCancellationStatusSUCCEEDED PaymentCancellationStatus = "SUCCEEDED"
)

// Trigger 취소 요청 경로
type Trigger string

const (
	TriggerCUSTOMER Trigger = "CUSTOMER"
	TriggerADMIN    Trigger = "ADMIN"
	TriggerPG       Trigger = "PG"
)

// PaymentCancellation 결제 취소 내역
type PaymentCancellation struct {
	// 결제 취소 내역 상태
	Status PaymentCancellationStatus `json:"status"`
	// 취소 내역 아이디
	ID string `json:"id"`
	// PG사 결제 취소 내역 아이디
	PgCancellationId *string `json:"pgCancellationId,omitempty"`
	// 취소 금액
	TotalAmount int64 `json:"totalAmount"`
	// 취소 금액 중 면세 금액
	TaxFreeAmount int64 `json:"taxFreeAmount"`
	// 취소 금액 중 부가세액
	VatAmount int64 `json:"vatAmount"`
	// 적립형 포인트의 환불 금액
	EasyPayDiscountAmount *int64 `json:"easyPayDiscountAmount,omitempty"`
	// 취소 사유
	Reason string `json:"reason"`
	// 취소 시점 (RFC 3339)
	CancelledAt *string `json:"cancelledAt,omitempty"`
	// 취소 요청 시점 (RFC 3339)
	RequestedAt string `json:"requestedAt"`
	// 취소 영수증 URL
	ReceiptUrl *string `json:"receiptUrl,omitempty"`
	// 취소 요청 경로
	Trigger *Trigger `json:"trigger,omitempty"`
}

// PaymentCashReceiptStatus 현금영수증 상태
type PaymentCashReceiptStatus string

const (
	PaymentCashReceiptStatusISSUED    PaymentCashReceiptStatus = "ISSUED"
	PaymentCashReceiptStatusCANCELLED PaymentCashReceiptStatus = "CANCELLED"
)

// PaymentCashReceipt 현금영수증
type PaymentCashReceipt struct {
	// 현금영수증 상태
	Status PaymentCashReceiptStatus `json:"status"`
	// 현금영수증 타입
	Type common.CashReceiptType `json:"type"`
	// PG사 영수증 발급 아이디
	PgReceiptId *string `json:"pgReceiptId,omitempty"`
	// 승인 번호
	IssueNumber *string `json:"issueNumber,omitempty"`
	// 발급 금액
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 현금영수증 URL
	Url *string `json:"url,omitempty"`
	// 발급 시점 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
	// 취소 시점 (RFC 3339)
	CancelledAt *string `json:"cancelledAt,omitempty"`
}

// PaymentEscrowStatus 에스크로 상태
type PaymentEscrowStatus string

const (
	PaymentEscrowStatusBEFORE_REGISTERED PaymentEscrowStatus = "BEFORE_REGISTERED"
	PaymentEscrowStatusDELIVERED         PaymentEscrowStatus = "DELIVERED"
	PaymentEscrowStatusCONFIRMED         PaymentEscrowStatus = "CONFIRMED"
	PaymentEscrowStatusREJECTED          PaymentEscrowStatus = "REJECTED"
	PaymentEscrowStatusCANCELLED         PaymentEscrowStatus = "CANCELLED"
)

// PaymentEscrow 에스크로 결제 정보
type PaymentEscrow struct {
	// 에스크로 상태
	Status PaymentEscrowStatus `json:"status"`
}

// Payment 결제 건 (상태에 따른 필드 포함)
type Payment struct {
	// 결제 건 상태
	Status PaymentStatus `json:"status"`
	// 결제 건 아이디
	ID string `json:"id"`
	// 결제 건 포트원 채번 아이디
	TransactionId string `json:"transactionId"`
	// 고객사 아이디
	MerchantId string `json:"merchantId"`
	// 상점 아이디
	StoreId string `json:"storeId"`
	// 결제수단 정보
	Method *PaymentMethod `json:"method,omitempty"`
	// 결제 채널
	Channel *common.SelectedChannel `json:"channel,omitempty"`
	// 결제 채널 그룹 정보
	ChannelGroup *common.ChannelGroupSummary `json:"channelGroup,omitempty"`
	// 포트원 버전
	Version common.PortOneVersion `json:"version"`
	// 결제 예약 건 아이디
	ScheduleId *string `json:"scheduleId,omitempty"`
	// 결제 시 사용된 빌링키
	BillingKey *string `json:"billingKey,omitempty"`
	// 웹훅 발송 내역
	Webhooks []PaymentWebhook `json:"webhooks,omitempty"`
	// 결제 요청 시점 (RFC 3339)
	RequestedAt string `json:"requestedAt"`
	// 업데이트 시점 (RFC 3339)
	UpdatedAt string `json:"updatedAt"`
	// 상태 업데이트 시점 (RFC 3339)
	StatusChangedAt string `json:"statusChangedAt"`
	// 주문명
	OrderName string `json:"orderName"`
	// 결제 금액 관련 세부 정보
	Amount PaymentAmount `json:"amount"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 구매자 정보
	Customer common.Customer `json:"customer"`
	// 프로모션 아이디
	PromotionId *string `json:"promotionId,omitempty"`
	// 문화비 지출 여부
	IsCulturalExpense *bool `json:"isCulturalExpense,omitempty"`
	// 에스크로 결제 정보
	Escrow *PaymentEscrow `json:"escrow,omitempty"`
	// 상품 정보
	Products []common.PaymentProduct `json:"products,omitempty"`
	// 상품 갯수
	ProductCount *int `json:"productCount,omitempty"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 국가 코드
	Country *common.Country `json:"country,omitempty"`

	// PAID 상태 필드
	// 결제 완료 시점 (RFC 3339)
	PaidAt *string `json:"paidAt,omitempty"`
	// PG사 거래 아이디
	PgTxId *string `json:"pgTxId,omitempty"`
	// PG사 거래 응답 본문
	PgResponse *string `json:"pgResponse,omitempty"`
	// 현금영수증
	CashReceipt *PaymentCashReceipt `json:"cashReceipt,omitempty"`
	// 거래 영수증 URL
	ReceiptUrl *string `json:"receiptUrl,omitempty"`
	// 분쟁 목록
	Disputes []Dispute `json:"disputes,omitempty"`

	// FAILED 상태 필드
	// 결제 실패 시점 (RFC 3339)
	FailedAt *string `json:"failedAt,omitempty"`
	// 결제 실패 정보
	Failure *PaymentFailure `json:"failure,omitempty"`

	// CANCELLED/PARTIAL_CANCELLED 상태 필드
	// 결제 취소 내역
	Cancellations []PaymentCancellation `json:"cancellations,omitempty"`
	// 결제 취소 시점 (RFC 3339)
	CancelledAt *string `json:"cancelledAt,omitempty"`
}

// IsPaid 결제 완료 상태인지 확인
func (p *Payment) IsPaid() bool {
	return p.Status == PaymentStatusPAID
}

// IsFailed 결제 실패 상태인지 확인
func (p *Payment) IsFailed() bool {
	return p.Status == PaymentStatusFAILED
}

// IsCancelled 결제 취소 상태인지 확인
func (p *Payment) IsCancelled() bool {
	return p.Status == PaymentStatusCANCELLED
}

// IsPartialCancelled 결제 부분 취소 상태인지 확인
func (p *Payment) IsPartialCancelled() bool {
	return p.Status == PaymentStatusPARTIAL_CANCELLED
}

// IsReady 결제 준비 상태인지 확인
func (p *Payment) IsReady() bool {
	return p.Status == PaymentStatusREADY
}

// GetPaymentResponse 결제 조회 응답
type GetPaymentResponse = Payment

// GetPaymentsResponse 결제 목록 조회 응답
type GetPaymentsResponse struct {
	// 결제 목록
	Items []Payment `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// GetAllPaymentsByCursorResponse 커서 기반 결제 목록 조회 응답
type GetAllPaymentsByCursorResponse struct {
	// 결제 목록
	Items []Payment `json:"items"`
	// 다음 페이지 커서
	NextCursor *string `json:"nextCursor,omitempty"`
}

// CancelPaymentResponse 결제 취소 응답
type CancelPaymentResponse struct {
	// 취소 내역
	Cancellation PaymentCancellation `json:"cancellation"`
}

// PayInstantlyResponse 즉시 결제 응답
type PayInstantlyResponse struct {
	// 결제 건
	Payment Payment `json:"payment"`
}

// PayWithBillingKeyResponse 빌링키 결제 응답
type PayWithBillingKeyResponse struct {
	// 결제 건
	Payment Payment `json:"payment"`
}

// ConfirmPaymentResponse 결제 승인 응답
type ConfirmPaymentResponse struct {
	// 결제 건
	Payment Payment `json:"payment"`
}

// PreRegisterPaymentResponse 결제 사전 등록 응답
type PreRegisterPaymentResponse struct{}

// CancelPaymentOptions 결제 취소 옵션
type CancelPaymentOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 취소 금액 (미입력시 전액 취소)
	Amount *int64 `json:"amount,omitempty"`
	// 취소 금액 중 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 취소 금액 중 부가세
	VatAmount *int64 `json:"vatAmount,omitempty"`
	// 취소 사유
	Reason string `json:"reason"`
	// 현금영수증 취소 유형
	CashReceiptCancelType *string `json:"cashReceiptCancelType,omitempty"`
	// 취소 요청자
	Requester *string `json:"requester,omitempty"`
	// 환불 계좌 정보 (가상계좌 취소시)
	RefundAccount *RefundAccount `json:"refundAccount,omitempty"`
}

// RefundAccount 환불 계좌 정보
type RefundAccount struct {
	// 은행
	Bank common.Bank `json:"bank"`
	// 계좌번호
	Number string `json:"number"`
	// 예금주명
	HolderName string `json:"holderName"`
	// 예금주 연락처
	HolderPhoneNumber *string `json:"holderPhoneNumber,omitempty"`
}

// PaymentFilterInput 결제 필터 입력
type PaymentFilterInput struct {
	// 고객사 아이디
	MerchantId *string `json:"merchantId,omitempty"`
	// 상점 아이디 목록
	StoreIds []string `json:"storeIds,omitempty"`
	// 결제 시작일시 (RFC 3339)
	TimestampType *string `json:"timestampType,omitempty"`
	// 결제 시작일시 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 결제 종료일시 (RFC 3339)
	Until *string `json:"until,omitempty"`
	// 결제 상태 목록
	Statuses []PaymentStatus `json:"statuses,omitempty"`
	// 결제 수단 목록
	Methods []common.PaymentMethodType `json:"methods,omitempty"`
	// PG사 목록
	PgProviders []common.PgProvider `json:"pgProviders,omitempty"`
	// 테스트 결제 여부
	IsTest *bool `json:"isTest,omitempty"`
	// 결제 예약 여부
	IsScheduled *bool `json:"isScheduled,omitempty"`
	// 정렬 순서
	SortBy *string `json:"sortBy,omitempty"`
	// 정렬 방향
	SortOrder *common.SortOrder `json:"sortOrder,omitempty"`
	// 포트원 버전
	Version *common.PortOneVersion `json:"version,omitempty"`
	// 통화
	Currencies []common.Currency `json:"currencies,omitempty"`
	// 에스크로 여부
	IsEscrow *bool `json:"isEscrow,omitempty"`
	// 텍스트 검색
	TextSearch *PaymentTextSearch `json:"textSearch,omitempty"`
}

// PaymentTextSearch 결제 텍스트 검색
type PaymentTextSearch struct {
	// 검색 필드
	Field string `json:"field"`
	// 검색어
	Value string `json:"value"`
}

// CapturePaymentResponse 수동 매입 응답
type CapturePaymentResponse struct {
	// 결제 건
	Payment Payment `json:"payment"`
}

// CloseVirtualAccountResponse 가상계좌 말소 응답
type CloseVirtualAccountResponse struct {
	// 가상계좌 말소 시점 (RFC 3339)
	ClosedAt string `json:"closedAt"`
}

// ApplyEscrowLogisticsResponse 에스크로 배송 정보 등록 응답
type ApplyEscrowLogisticsResponse struct {
	// 에스크로 정보
	SendEmail bool `json:"sendEmail"`
	// 등록 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// ModifyEscrowLogisticsResponse 에스크로 배송 정보 수정 응답
type ModifyEscrowLogisticsResponse struct {
	// 에스크로 정보
	SendEmail bool `json:"sendEmail"`
	// 수정 시점 (RFC 3339)
	ModifiedAt string `json:"modifiedAt"`
}

// ConfirmEscrowResponse 에스크로 구매 확정 응답
type ConfirmEscrowResponse struct {
	// 구매 확정 시점 (RFC 3339)
	CompletedAt string `json:"completedAt"`
}

// RegisterStoreReceiptResponse 영수증 내 하위 상점 거래 등록 응답
type RegisterStoreReceiptResponse struct {
	// 결제 건 영수증
	ReceiptUrl *string `json:"receiptUrl,omitempty"`
}

// GetPaymentTransactionsResponse 결제 시도 내역 조회 응답
type GetPaymentTransactionsResponse struct {
	// 결제 시도 내역 목록
	Items []PaymentTransaction `json:"items"`
}

// GetAllPaymentEventsByCursorResponse 결제 이벤트 대용량 다건 조회 응답
type GetAllPaymentEventsByCursorResponse struct {
	// 결제 이벤트 목록
	Items []PaymentEvent `json:"items"`
	// 다음 페이지 커서
	NextCursor *string `json:"nextCursor,omitempty"`
}

// PaymentTransaction 결제 시도 내역
type PaymentTransaction struct {
	// 결제 시도 아이디
	TxId string `json:"txId"`
	// PG사 거래 아이디
	PgTxId *string `json:"pgTxId,omitempty"`
	// 결제 상태
	Status PaymentStatus `json:"status"`
	// 결제 채널
	Channel *common.SelectedChannel `json:"channel,omitempty"`
	// 결제 시도 시점 (RFC 3339)
	RequestedAt string `json:"requestedAt"`
	// 결제 완료 시점 (RFC 3339)
	PaidAt *string `json:"paidAt,omitempty"`
	// 결제 실패 시점 (RFC 3339)
	FailedAt *string `json:"failedAt,omitempty"`
}

// PaymentEvent 결제 이벤트
type PaymentEvent struct {
	// 결제 이벤트 타입
	Type string `json:"type"`
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 이벤트 발생 시점 (RFC 3339)
	OccurredAt string `json:"occurredAt"`
}

// ConfirmedPaymentSummary 인증 결제 수동 승인 응답
type ConfirmedPaymentSummary struct {
	// 결제 건 아이디
	Id string `json:"id"`
	// PG사
	PgProvider common.PgProvider `json:"pgProvider"`
	// PG사 거래 아이디
	PgTxId string `json:"pgTxId"`
}

// InstantPaymentMethodInput 수기 결제 수단 입력
type InstantPaymentMethodInput struct {
	// 결제수단 타입
	Type string `json:"type"`
	// 카드 정보 (type=card)
	Card *InstantPaymentMethodInputCard `json:"card,omitempty"`
	// 가상계좌 정보 (type=virtualAccount)
	VirtualAccount *InstantPaymentMethodInputVirtualAccount `json:"virtualAccount,omitempty"`
}

// InstantPaymentMethodInputCard 수기 결제 카드 정보 입력
type InstantPaymentMethodInputCard struct {
	// 카드 인증 관련 정보
	Credential CardCredential `json:"credential"`
	// 할부 개월 수
	InstallmentMonth *int `json:"installmentMonth,omitempty"`
	// 무이자 할부 이자를 고객사가 부담할지 여부
	UseFreeInterestFromMerchant *bool `json:"useFreeInterestFromMerchant,omitempty"`
	// 카드 포인트 사용 여부
	UseCardPoint *bool `json:"useCardPoint,omitempty"`
}

// CardCredential 카드 인증 정보
type CardCredential struct {
	// 카드 번호
	Number string `json:"number"`
	// 카드 유효기간 만료 년도 (2자리)
	ExpiryYear string `json:"expiryYear"`
	// 카드 유효기간 만료 월
	ExpiryMonth string `json:"expiryMonth"`
	// 카드 비밀번호 앞 두자리
	PasswordTwoDigits *string `json:"passwordTwoDigits,omitempty"`
	// 생년월일 (YYMMDD) 또는 사업자등록번호
	BirthOrBusinessRegistrationNumber *string `json:"birthOrBusinessRegistrationNumber,omitempty"`
}

// InstantPaymentMethodInputVirtualAccount 수기 결제 가상계좌 정보 입력
type InstantPaymentMethodInputVirtualAccount struct {
	// 은행
	Bank common.Bank `json:"bank"`
	// 입금 만료 기한 (RFC 3339)
	ExpiredAt string `json:"expiredAt"`
	// 가상계좌 유형
	AccountType *string `json:"accountType,omitempty"`
	// 계좌주 명
	AccountHolderName *string `json:"accountHolderName,omitempty"`
	// 고정식 가상계좌 설정
	FixedOption *VirtualAccountFixedOption `json:"fixedOption,omitempty"`
	// 현금영수증 발급 정보
	CashReceipt *VirtualAccountCashReceiptInput `json:"cashReceipt,omitempty"`
	// 송금자명
	RemitterName *string `json:"remitterName,omitempty"`
}

// VirtualAccountFixedOption 고정식 가상계좌 설정
type VirtualAccountFixedOption struct {
	// PG사 가맹점 구분값
	PgAccountId string `json:"pgAccountId"`
}

// VirtualAccountCashReceiptInput 가상계좌 현금영수증 발급 입력
type VirtualAccountCashReceiptInput struct {
	// 현금영수증 발급 타입
	Type common.CashReceiptType `json:"type"`
	// 현금영수증 식별 번호
	CustomerIdentityNumber string `json:"customerIdentityNumber"`
}

// PaymentLogistics 에스크로 물류 정보
type PaymentLogistics struct {
	// 배송사
	Company string `json:"company"`
	// 송장 번호
	InvoiceNumber string `json:"invoiceNumber"`
	// 배송 시작 시점 (RFC 3339)
	SentAt string `json:"sentAt"`
	// 수령 시점 (RFC 3339)
	ReceivedAt *string `json:"receivedAt,omitempty"`
	// 주소
	Address *common.SeparatedAddressInput `json:"address,omitempty"`
}

// PaymentEscrowSenderInput 에스크로 발송자 정보 입력
type PaymentEscrowSenderInput struct {
	// 이름
	Name *string `json:"name,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// 우편번호
	Zipcode *string `json:"zipcode,omitempty"`
	// 주소
	Address *common.SeparatedAddressInput `json:"address,omitempty"`
	// 관계
	Relationship *string `json:"relationship,omitempty"`
}

// PaymentEscrowReceiverInput 에스크로 수취인 정보 입력
type PaymentEscrowReceiverInput struct {
	// 이름
	Name *string `json:"name,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// 우편번호
	Zipcode *string `json:"zipcode,omitempty"`
	// 주소
	Address *common.SeparatedAddressInput `json:"address,omitempty"`
}

// RegisterStoreReceiptBodyItem 영수증 내 하위 상점 거래 아이템
type RegisterStoreReceiptBodyItem struct {
	// 하위 상점 사업자등록번호
	StoreBusinessRegistrationNumber string `json:"storeBusinessRegistrationNumber"`
	// 하위 상점명
	StoreName string `json:"storeName"`
	// 결제 총 금액
	TotalAmount int64 `json:"totalAmount"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 부가세 금액
	VatAmount *int64 `json:"vatAmount,omitempty"`
	// 봉사료 금액
	ServiceFeeAmount *int64 `json:"serviceFeeAmount,omitempty"`
	// 통화
	Currency common.Currency `json:"currency"`
}

// ConfirmPaymentOptions 인증 결제 수동 승인 옵션
type ConfirmPaymentOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 결제 토큰
	PaymentToken string `json:"paymentToken"`
	// 결제 시도 아이디
	TxId *string `json:"txId,omitempty"`
	// 통화
	Currency *common.Currency `json:"currency,omitempty"`
	// 결제 금액
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 테스트 결제 여부
	IsTest *bool `json:"isTest,omitempty"`
}

// PayInstantlyOptions 수기 결제 옵션
type PayInstantlyOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 채널 키
	ChannelKey *string `json:"channelKey,omitempty"`
	// 채널 그룹 아이디
	ChannelGroupId *string `json:"channelGroupId,omitempty"`
	// 결제수단 정보
	Method InstantPaymentMethodInput `json:"method"`
	// 주문명
	OrderName string `json:"orderName"`
	// 문화비 지출 여부
	IsCulturalExpense *bool `json:"isCulturalExpense,omitempty"`
	// 에스크로 결제 여부
	IsEscrow *bool `json:"isEscrow,omitempty"`
	// 고객 정보
	Customer *common.CustomerInput `json:"customer,omitempty"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 결제 금액 세부 입력 정보
	Amount common.PaymentAmountInput `json:"amount"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 결제 국가
	Country *common.Country `json:"country,omitempty"`
	// 웹훅 주소
	NoticeUrls []string `json:"noticeUrls,omitempty"`
	// 상품 정보
	Products []common.PaymentProduct `json:"products,omitempty"`
	// 상품 개수
	ProductCount *int `json:"productCount,omitempty"`
	// 상품 유형
	ProductType *common.PaymentProductType `json:"productType,omitempty"`
	// 배송지 주소
	ShippingAddress *common.SeparatedAddressInput `json:"shippingAddress,omitempty"`
	// 해당 결제에 적용할 프로모션 아이디
	PromotionId *string `json:"promotionId,omitempty"`
}

// ApplyEscrowLogisticsOptions 에스크로 배송 정보 등록 옵션
type ApplyEscrowLogisticsOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 에스크로 발송자 정보
	Sender *PaymentEscrowSenderInput `json:"sender,omitempty"`
	// 에스크로 수취인 정보
	Receiver *PaymentEscrowReceiverInput `json:"receiver,omitempty"`
	// 에스크로 물류 정보
	Logistics PaymentLogistics `json:"logistics"`
	// 이메일 알림 전송 여부
	SendEmail *bool `json:"sendEmail,omitempty"`
	// 상품 정보
	Products []common.PaymentProduct `json:"products,omitempty"`
}

// ModifyEscrowLogisticsOptions 에스크로 배송 정보 수정 옵션
type ModifyEscrowLogisticsOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 에스크로 발송자 정보
	Sender *PaymentEscrowSenderInput `json:"sender,omitempty"`
	// 에스크로 수취인 정보
	Receiver *PaymentEscrowReceiverInput `json:"receiver,omitempty"`
	// 에스크로 물류 정보
	Logistics PaymentLogistics `json:"logistics"`
	// 이메일 알림 전송 여부
	SendEmail *bool `json:"sendEmail,omitempty"`
	// 상품 정보
	Products []common.PaymentProduct `json:"products,omitempty"`
}
