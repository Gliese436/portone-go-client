package schedule

import (
	"github.com/gliese436/portone-go-client/portonev2/common"
)

// PaymentScheduleStatus 결제 예약 상태
type PaymentScheduleStatus string

const (
	PaymentScheduleStatusSCHEDULED PaymentScheduleStatus = "SCHEDULED"
	PaymentScheduleStatusSTARTED   PaymentScheduleStatus = "STARTED"
	PaymentScheduleStatusSUCCEEDED PaymentScheduleStatus = "SUCCEEDED"
	PaymentScheduleStatusFAILED    PaymentScheduleStatus = "FAILED"
	PaymentScheduleStatusREVOKED   PaymentScheduleStatus = "REVOKED"
	PaymentScheduleStatusPENDING   PaymentScheduleStatus = "PENDING"
)

// PaymentSchedule 결제 예약 건
type PaymentSchedule struct {
	// 결제 예약 상태
	Status PaymentScheduleStatus `json:"status"`
	// 결제 예약 아이디
	ID string `json:"id"`
	// 고객사 아이디
	MerchantId string `json:"merchantId"`
	// 상점 아이디
	StoreId string `json:"storeId"`
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 빌링키
	BillingKey string `json:"billingKey"`
	// 주문명
	OrderName string `json:"orderName"`
	// 문화비 지출 여부
	IsCulturalExpense bool `json:"isCulturalExpense"`
	// 에스크로 결제 여부
	IsEscrow bool `json:"isEscrow"`
	// 고객 정보
	Customer *common.Customer `json:"customer,omitempty"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 총 결제 금액
	TotalAmount int64 `json:"totalAmount"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 부가세
	VatAmount *int64 `json:"vatAmount,omitempty"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 할부 개월 수
	InstallMonth *int `json:"installMonth,omitempty"`
	// 웹훅 알림 URL 목록
	NoticeUrls []string `json:"noticeUrls,omitempty"`
	// 상품 목록
	Products []common.PaymentProduct `json:"products,omitempty"`
	// 결제 예약 시점 (RFC 3339)
	TimeToPay string `json:"timeToPay"`
	// 결제 예약 생성 시점 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 결제 예약 취소 시점 (RFC 3339)
	RevokedAt *string `json:"revokedAt,omitempty"`
	// 결제 시작 시점 (RFC 3339)
	StartedAt *string `json:"startedAt,omitempty"`
	// 결제 완료 시점 (RFC 3339)
	CompletedAt *string `json:"completedAt,omitempty"`
}

// IsScheduled 예약 완료 상태인지 확인
func (p *PaymentSchedule) IsScheduled() bool {
	return p.Status == PaymentScheduleStatusSCHEDULED
}

// IsSucceeded 결제 성공 상태인지 확인
func (p *PaymentSchedule) IsSucceeded() bool {
	return p.Status == PaymentScheduleStatusSUCCEEDED
}

// IsFailed 결제 실패 상태인지 확인
func (p *PaymentSchedule) IsFailed() bool {
	return p.Status == PaymentScheduleStatusFAILED
}

// IsRevoked 취소 상태인지 확인
func (p *PaymentSchedule) IsRevoked() bool {
	return p.Status == PaymentScheduleStatusREVOKED
}

// PaymentScheduleSortBy 결제 예약 정렬 기준
type PaymentScheduleSortBy string

const (
	PaymentScheduleSortByCREATED_AT    PaymentScheduleSortBy = "CREATED_AT"
	PaymentScheduleSortByTIME_TO_PAY   PaymentScheduleSortBy = "TIME_TO_PAY"
	PaymentScheduleSortByCOMPLETED_AT  PaymentScheduleSortBy = "COMPLETED_AT"
)

// PaymentScheduleSortInput 결제 예약 정렬 입력
type PaymentScheduleSortInput struct {
	// 정렬 기준
	By *PaymentScheduleSortBy `json:"by,omitempty"`
	// 정렬 순서
	Order *common.SortOrder `json:"order,omitempty"`
}

// PaymentScheduleFilterInput 결제 예약 필터 입력
type PaymentScheduleFilterInput struct {
	// 상점 아이디 목록
	StoreIds []string `json:"storeIds,omitempty"`
	// 빌링키
	BillingKey *string `json:"billingKey,omitempty"`
	// 결제 예약 시작 시점 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 결제 예약 종료 시점 (RFC 3339)
	Until *string `json:"until,omitempty"`
	// 결제 예약 상태 목록
	Statuses []PaymentScheduleStatus `json:"statuses,omitempty"`
}

// GetPaymentSchedulesBody 결제 예약 목록 조회 요청
type GetPaymentSchedulesBody struct {
	// 페이지 정보
	Page *common.PageInput `json:"page,omitempty"`
	// 정렬 정보
	Sort *PaymentScheduleSortInput `json:"sort,omitempty"`
	// 필터 정보
	Filter *PaymentScheduleFilterInput `json:"filter,omitempty"`
}

// GetPaymentSchedulesResponse 결제 예약 목록 조회 응답
type GetPaymentSchedulesResponse struct {
	// 결제 예약 목록
	Items []PaymentSchedule `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// CreatePaymentScheduleBody 결제 예약 생성 요청
type CreatePaymentScheduleBody struct {
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 결제 예약 정보
	Payment PaymentScheduleInput `json:"payment"`
}

// PaymentScheduleInput 결제 예약 입력
type PaymentScheduleInput struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 빌링키
	BillingKey string `json:"billingKey"`
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
	// 총 결제 금액
	TotalAmount int64 `json:"totalAmount"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 부가세
	VatAmount *int64 `json:"vatAmount,omitempty"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 할부 개월 수
	InstallMonth *int `json:"installMonth,omitempty"`
	// 웹훅 알림 URL 목록
	NoticeUrls []string `json:"noticeUrls,omitempty"`
	// 상품 목록
	Products []common.PaymentProduct `json:"products,omitempty"`
	// 결제 예약 시점 (RFC 3339)
	TimeToPay string `json:"timeToPay"`
}

// CreatePaymentScheduleResponse 결제 예약 생성 응답
type CreatePaymentScheduleResponse struct {
	// 결제 예약 정보
	Schedule PaymentSchedule `json:"schedule"`
}

// RevokePaymentSchedulesBody 결제 예약 취소 요청
type RevokePaymentSchedulesBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 빌링키
	BillingKey *string `json:"billingKey,omitempty"`
	// 결제 예약 아이디 목록
	ScheduleIds []string `json:"scheduleIds,omitempty"`
}

// RevokePaymentSchedulesResponse 결제 예약 취소 응답
type RevokePaymentSchedulesResponse struct {
	// 취소된 결제 예약 아이디 목록
	RevokedScheduleIds []string `json:"revokedScheduleIds"`
	// 취소된 결제 예약 건 수
	RevokedCount int `json:"revokedCount"`
}
