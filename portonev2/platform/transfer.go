package platform

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// PlatformTransfer 정산건 (discriminated union)
type PlatformTransfer struct {
	Type PlatformTransferType `json:"type"`
	// Common fields
	ID                    string                                `json:"id"`
	GraphqlID             string                                `json:"graphqlId"`
	PartnerID             string                                `json:"partnerId"`
	Status                PlatformTransferStatus                `json:"status"`
	Memo                  *string                               `json:"memo,omitempty"`
	SettlementDate        string                                `json:"settlementDate"`
	SettlementCurrency    common.Currency                       `json:"settlementCurrency"`
	IsForTest             bool                                  `json:"isForTest"`
	UserDefinedProperties []PlatformUserDefinedPropertyKeyValue `json:"userDefinedProperties,omitempty"`
	// Order fields
	PaymentID               *string `json:"paymentId,omitempty"`
	ContractID              *string `json:"contractId,omitempty"`
	OrderAmount             *int64  `json:"orderAmount,omitempty"`
	OrderTaxFreeAmount      *int64  `json:"orderTaxFreeAmount,omitempty"`
	SettlementStartDate     *string `json:"settlementStartDate,omitempty"`
	DiscountAmount          *int64  `json:"discountAmount,omitempty"`
	DiscountTaxFreeAmount   *int64  `json:"discountTaxFreeAmount,omitempty"`
	DiscountShareAmount     *int64  `json:"discountShareAmount,omitempty"`
	AdditionalFeeAmount     *int64  `json:"additionalFeeAmount,omitempty"`
	PlatformFeeAmount       *int64  `json:"platformFeeAmount,omitempty"`
	SettlementAmount        *int64  `json:"settlementAmount,omitempty"`
	SettlementTaxFreeAmount *int64  `json:"settlementTaxFreeAmount,omitempty"`
	// Manual fields
	// (uses SettlementAmount, SettlementTaxFreeAmount)
	// OrderCancel fields
	CancellationID *string `json:"cancellationId,omitempty"`
}

// IsManual 수기 정산건인지 확인
func (t *PlatformTransfer) IsManual() bool {
	return t.Type == PlatformTransferTypeManual
}

// IsOrder 주문 정산건인지 확인
func (t *PlatformTransfer) IsOrder() bool {
	return t.Type == PlatformTransferTypeOrder
}

// IsOrderCancel 주문 취소 정산건인지 확인
func (t *PlatformTransfer) IsOrderCancel() bool {
	return t.Type == PlatformTransferTypeOrderCancel
}

// PlatformTransferSummary 정산건 요약
type PlatformTransferSummary struct {
	ID                 string                 `json:"id"`
	GraphqlID          string                 `json:"graphqlId"`
	Type               PlatformTransferType   `json:"type"`
	PartnerID          string                 `json:"partnerId"`
	Status             PlatformTransferStatus `json:"status"`
	Memo               *string                `json:"memo,omitempty"`
	SettlementDate     string                 `json:"settlementDate"`
	SettlementCurrency common.Currency        `json:"settlementCurrency"`
	SettlementAmount   int64                  `json:"settlementAmount"`
	PaymentID          *string                `json:"paymentId,omitempty"`
	IsForTest          bool                   `json:"isForTest"`
}

// GetPlatformTransferSummariesResponse 정산건 다건 조회 응답
type GetPlatformTransferSummariesResponse struct {
	Items []PlatformTransferSummary `json:"items"`
	Page  PageInfo                  `json:"page"`
}

// DeletePlatformTransferResponse 정산건 삭제 응답
type DeletePlatformTransferResponse struct {
	// empty
}

// CreateManualTransferResponse 수기 정산건 생성 응답
type CreateManualTransferResponse struct {
	Transfer PlatformTransfer `json:"transfer"`
}

// CreateOrderTransferResponse 주문 정산건 생성 응답
type CreateOrderTransferResponse struct {
	Transfer PlatformTransfer `json:"transfer"`
}

// CreateOrderCancelTransferResponse 주문 취소 정산건 생성 응답
type CreateOrderCancelTransferResponse struct {
	Transfer PlatformTransfer `json:"transfer"`
}

// ---- Transfer Create Body Types ----

// CreatePlatformManualTransferBody 수기 정산건 생성 요청 바디
type CreatePlatformManualTransferBody struct {
	PartnerID               string                                `json:"partnerId"`
	Memo                    *string                               `json:"memo,omitempty"`
	SettlementAmount        int64                                 `json:"settlementAmount"`
	SettlementTaxFreeAmount *int64                                `json:"settlementTaxFreeAmount,omitempty"`
	SettlementDate          string                                `json:"settlementDate"`
	IsForTest               *bool                                 `json:"isForTest,omitempty"`
	UserDefinedProperties   []PlatformUserDefinedPropertyKeyValue `json:"userDefinedProperties,omitempty"`
}

// CreatePlatformOrderTransferBodyOrderDetail 주문 정보
type CreatePlatformOrderTransferBodyOrderDetail struct {
	OrderName  string                                     `json:"orderName"`
	OrderLines []CreatePlatformOrderTransferBodyOrderLine `json:"orderLines,omitempty"`
}

// CreatePlatformOrderTransferBodyOrderLine 주문 항목
type CreatePlatformOrderTransferBodyOrderLine struct {
	ProductID     string `json:"productId"`
	ProductName   string `json:"productName"`
	Quantity      int32  `json:"quantity"`
	UnitPrice     int64  `json:"unitPrice"`
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
}

// CreatePlatformOrderTransferBodyDiscount 할인 정보
type CreatePlatformOrderTransferBodyDiscount struct {
	SharePolicyID string `json:"sharePolicyId"`
	Amount        int64  `json:"amount"`
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
}

// CreatePlatformOrderTransferBodyAdditionalFee 추가 수수료 정보
type CreatePlatformOrderTransferBodyAdditionalFee struct {
	PolicyID string `json:"policyId"`
}

// CreatePlatformOrderTransferBodyExternalPaymentDetail 외부 결제 상세 정보
type CreatePlatformOrderTransferBodyExternalPaymentDetail struct {
	Currency      common.Currency             `json:"currency"`
	OrderName     *string                     `json:"orderName,omitempty"`
	PaidAt        string                      `json:"paidAt"`
	Method        *PlatformPaymentMethodInput `json:"method,omitempty"`
	Amount        int64                       `json:"amount"`
	TaxFreeAmount *int64                      `json:"taxFreeAmount,omitempty"`
}

// PlatformPaymentMethodInput 결제 수단 입력
type PlatformPaymentMethodInput struct {
	Type string `json:"type"`
	// Card fields
	CardBrand *string `json:"cardBrand,omitempty"`
	// EasyPay fields
	EasyPayProvider   *string `json:"easyPayProvider,omitempty"`
	EasyPayMethodType *string `json:"easyPayMethodType,omitempty"`
}

// TransferParameters 정산 파라미터 (실험기능)
type TransferParameters struct {
	PaymentAmount        *int64 `json:"paymentAmount,omitempty"`
	PaymentTaxFreeAmount *int64 `json:"paymentTaxFreeAmount,omitempty"`
	PaymentSupplyAmount  *int64 `json:"paymentSupplyAmount,omitempty"`
	PaymentVatAmount     *int64 `json:"paymentVatAmount,omitempty"`
}

// CreatePlatformOrderTransferBody 주문 정산건 생성 요청 바디
type CreatePlatformOrderTransferBody struct {
	PartnerID             string                                                `json:"partnerId"`
	ContractID            *string                                               `json:"contractId,omitempty"`
	Memo                  *string                                               `json:"memo,omitempty"`
	PaymentID             string                                                `json:"paymentId"`
	OrderDetail           CreatePlatformOrderTransferBodyOrderDetail            `json:"orderDetail"`
	TaxFreeAmount         *int64                                                `json:"taxFreeAmount,omitempty"`
	SettlementStartDate   *string                                               `json:"settlementStartDate,omitempty"`
	SettlementDate        *string                                               `json:"settlementDate,omitempty"`
	Discounts             []CreatePlatformOrderTransferBodyDiscount             `json:"discounts"`
	AdditionalFees        []CreatePlatformOrderTransferBodyAdditionalFee        `json:"additionalFees"`
	ExternalPaymentDetail *CreatePlatformOrderTransferBodyExternalPaymentDetail `json:"externalPaymentDetail,omitempty"`
	IsForTest             *bool                                                 `json:"isForTest,omitempty"`
	Parameters            *TransferParameters                                   `json:"parameters,omitempty"`
	UserDefinedProperties []PlatformUserDefinedPropertyKeyValue                 `json:"userDefinedProperties,omitempty"`
}

// CreatePlatformOrderCancelTransferBodyOrderDetail 주문 취소 정보
type CreatePlatformOrderCancelTransferBodyOrderDetail struct {
	OrderLines []CreatePlatformOrderCancelTransferBodyOrderLine `json:"orderLines,omitempty"`
}

// CreatePlatformOrderCancelTransferBodyOrderLine 주문 취소 항목
type CreatePlatformOrderCancelTransferBodyOrderLine struct {
	ProductID string `json:"productId"`
	Quantity  int32  `json:"quantity"`
}

// CreatePlatformOrderCancelTransferBodyDiscount 취소 할인 정보
type CreatePlatformOrderCancelTransferBodyDiscount struct {
	SharePolicyID string `json:"sharePolicyId"`
	Amount        int64  `json:"amount"`
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
}

// CreatePlatformOrderCancelTransferBodyExternalCancellationDetail 외부 취소 상세 정보
type CreatePlatformOrderCancelTransferBodyExternalCancellationDetail struct {
	CancelledAt   string `json:"cancelledAt"`
	Amount        int64  `json:"amount"`
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
}

// CreatePlatformOrderCancelTransferBody 주문 취소 정산건 생성 요청 바디
type CreatePlatformOrderCancelTransferBody struct {
	PartnerID                  *string                                                          `json:"partnerId,omitempty"`
	PaymentID                  *string                                                          `json:"paymentId,omitempty"`
	TransferID                 *string                                                          `json:"transferId,omitempty"`
	CancellationID             string                                                           `json:"cancellationId"`
	Memo                       *string                                                          `json:"memo,omitempty"`
	OrderDetail                *CreatePlatformOrderCancelTransferBodyOrderDetail                `json:"orderDetail,omitempty"`
	TaxFreeAmount              *int64                                                           `json:"taxFreeAmount,omitempty"`
	Discounts                  []CreatePlatformOrderCancelTransferBodyDiscount                  `json:"discounts"`
	SettlementStartDate        *string                                                          `json:"settlementStartDate,omitempty"`
	SettlementDate             *string                                                          `json:"settlementDate,omitempty"`
	ExternalCancellationDetail *CreatePlatformOrderCancelTransferBodyExternalCancellationDetail `json:"externalCancellationDetail,omitempty"`
	IsForTest                  *bool                                                            `json:"isForTest,omitempty"`
	UserDefinedProperties      []PlatformUserDefinedPropertyKeyValue                            `json:"userDefinedProperties,omitempty"`
}

// ---- Payout Types ----

// PlatformPayoutStatus 지급 상태
type PlatformPayoutStatus string

const (
	PlatformPayoutStatusScheduled PlatformPayoutStatus = "SCHEDULED"
	PlatformPayoutStatusInProcess PlatformPayoutStatus = "IN_PROCESS"
	PlatformPayoutStatusCompleted PlatformPayoutStatus = "COMPLETED"
	PlatformPayoutStatusFailed    PlatformPayoutStatus = "FAILED"
	PlatformPayoutStatusCancelled PlatformPayoutStatus = "CANCELLED"
)

// PlatformPayout 지급
type PlatformPayout struct {
	ID            string               `json:"id"`
	GraphqlID     string               `json:"graphqlId"`
	Status        PlatformPayoutStatus `json:"status"`
	PartnerID     string               `json:"partnerId"`
	Amount        int64                `json:"amount"`
	Currency      common.Currency      `json:"currency"`
	ScheduledDate string               `json:"scheduledDate"`
	CompletedAt   *string              `json:"completedAt,omitempty"`
	IsForTest     bool                 `json:"isForTest"`
}

// PlatformPayoutFilterInput 지급 필터 입력
type PlatformPayoutFilterInput struct {
	PartnerIDs []string               `json:"partnerIds,omitempty"`
	Statuses   []PlatformPayoutStatus `json:"statuses,omitempty"`
}

// GetPlatformPayoutsResponse 지급 다건 조회 응답
type GetPlatformPayoutsResponse struct {
	Items []PlatformPayout `json:"items"`
	Page  PageInfo         `json:"page"`
}

// ---- BulkPayout Types ----

// PlatformBulkPayoutStatus 일괄 지급 상태
type PlatformBulkPayoutStatus string

const (
	PlatformBulkPayoutStatusScheduled  PlatformBulkPayoutStatus = "SCHEDULED"
	PlatformBulkPayoutStatusInProgress PlatformBulkPayoutStatus = "IN_PROGRESS"
	PlatformBulkPayoutStatusCompleted  PlatformBulkPayoutStatus = "COMPLETED"
	PlatformBulkPayoutStatusFailed     PlatformBulkPayoutStatus = "FAILED"
)

// PlatformBulkPayout 일괄 지급
type PlatformBulkPayout struct {
	ID            string                   `json:"id"`
	GraphqlID     string                   `json:"graphqlId"`
	Status        PlatformBulkPayoutStatus `json:"status"`
	ScheduledDate string                   `json:"scheduledDate"`
	CompletedAt   *string                  `json:"completedAt,omitempty"`
	TotalAmount   int64                    `json:"totalAmount"`
	TotalCount    int32                    `json:"totalCount"`
	IsForTest     bool                     `json:"isForTest"`
}

// PlatformBulkPayoutFilterInput 일괄 지급 필터 입력
type PlatformBulkPayoutFilterInput struct {
	Statuses []PlatformBulkPayoutStatus `json:"statuses,omitempty"`
}

// GetPlatformBulkPayoutsResponse 일괄 지급 다건 조회 응답
type GetPlatformBulkPayoutsResponse struct {
	Items []PlatformBulkPayout `json:"items"`
	Page  PageInfo             `json:"page"`
}

// ---- AccountTransfer Types ----

// PlatformAccountTransferType 계좌 이체 유형
type PlatformAccountTransferType string

const (
	PlatformAccountTransferTypeDeposit    PlatformAccountTransferType = "DEPOSIT"
	PlatformAccountTransferTypeWithdrawal PlatformAccountTransferType = "WITHDRAWAL"
)

// PlatformAccountTransferStatus 계좌 이체 상태
type PlatformAccountTransferStatus string

const (
	PlatformAccountTransferStatusScheduled PlatformAccountTransferStatus = "SCHEDULED"
	PlatformAccountTransferStatusInProcess PlatformAccountTransferStatus = "IN_PROCESS"
	PlatformAccountTransferStatusCompleted PlatformAccountTransferStatus = "COMPLETED"
	PlatformAccountTransferStatusFailed    PlatformAccountTransferStatus = "FAILED"
	PlatformAccountTransferStatusCancelled PlatformAccountTransferStatus = "CANCELLED"
)

// PlatformAccountTransfer 계좌 이체
type PlatformAccountTransfer struct {
	ID          string                        `json:"id"`
	GraphqlID   string                        `json:"graphqlId"`
	Type        PlatformAccountTransferType   `json:"type"`
	Status      PlatformAccountTransferStatus `json:"status"`
	Amount      int64                         `json:"amount"`
	Currency    common.Currency               `json:"currency"`
	Memo        *string                       `json:"memo,omitempty"`
	CompletedAt *string                       `json:"completedAt,omitempty"`
	IsForTest   bool                          `json:"isForTest"`
}

// PlatformAccountTransferFilterInput 계좌 이체 필터 입력
type PlatformAccountTransferFilterInput struct {
	Types    []PlatformAccountTransferType   `json:"types,omitempty"`
	Statuses []PlatformAccountTransferStatus `json:"statuses,omitempty"`
}

// GetPlatformAccountTransfersResponse 계좌 이체 다건 조회 응답
type GetPlatformAccountTransfersResponse struct {
	Items []PlatformAccountTransfer `json:"items"`
	Page  PageInfo                  `json:"page"`
}

// ---- Account Holder Types ----

// GetPlatformAccountHolderResponse 예금주 조회 응답
type GetPlatformAccountHolderResponse struct {
	Holder    string  `json:"holder"`
	BirthDate *string `json:"birthDate,omitempty"`
}

// ---- Company State Types ----

// GetPlatformCompanyStateResponse 사업자 상태 조회 응답
type GetPlatformCompanyStateResponse struct {
	BusinessRegistrationNumber string                        `json:"businessRegistrationNumber"`
	CompanyName                *string                       `json:"companyName,omitempty"`
	RepresentativeName         *string                       `json:"representativeName,omitempty"`
	Status                     PlatformPartnerBusinessStatus `json:"status"`
}

// ---- Partner Filter Options ----

// GetPlatformPartnerFilterOptionsResponse 파트너 필터 옵션 응답
type GetPlatformPartnerFilterOptionsResponse struct {
	Tags []string `json:"tags"`
}

// GetPlatformDiscountSharePolicyFilterOptionsResponse 할인 분담 정책 필터 옵션 응답
type GetPlatformDiscountSharePolicyFilterOptionsResponse struct {
	// empty for now
}

// ---- Partner Settlements ----

// PlatformPartnerSettlement 파트너 정산 내역
type PlatformPartnerSettlement struct {
	PartnerID      string          `json:"partnerId"`
	SettlementDate string          `json:"settlementDate"`
	Amount         int64           `json:"amount"`
	Currency       common.Currency `json:"currency"`
}

// PlatformPartnerSettlementFilterInput 파트너 정산 내역 필터 입력
type PlatformPartnerSettlementFilterInput struct {
	PartnerIDs []string `json:"partnerIds,omitempty"`
}

// GetPlatformPartnerSettlementsResponse 파트너 정산 내역 다건 조회 응답
type GetPlatformPartnerSettlementsResponse struct {
	Items []PlatformPartnerSettlement `json:"items"`
	Page  PageInfo                    `json:"page"`
}
