package platform

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// PlatformPartnerStatus 플랫폼 파트너 상태
type PlatformPartnerStatus string

const (
	PlatformPartnerStatusPending  PlatformPartnerStatus = "PENDING"
	PlatformPartnerStatusApproved PlatformPartnerStatus = "APPROVED"
	PlatformPartnerStatusRejected PlatformPartnerStatus = "REJECTED"
)

// PlatformAccountStatus 플랫폼 계좌 상태
type PlatformAccountStatus string

const (
	PlatformAccountStatusVerified     PlatformAccountStatus = "VERIFIED"
	PlatformAccountStatusVerifyFailed PlatformAccountStatus = "VERIFY_FAILED"
	PlatformAccountStatusVerifyError  PlatformAccountStatus = "VERIFY_ERROR"
	PlatformAccountStatusNotVerified  PlatformAccountStatus = "NOT_VERIFIED"
	PlatformAccountStatusUnknown      PlatformAccountStatus = "UNKNOWN"
)

// PlatformPayer 금액 부담 주체
type PlatformPayer string

const (
	PlatformPayerPartner  PlatformPayer = "PARTNER"
	PlatformPayerMerchant PlatformPayer = "MERCHANT"
)

// PlatformSettlementCycleDatePolicy 플랫폼 정산 기준일
type PlatformSettlementCycleDatePolicy string

const (
	PlatformSettlementCycleDatePolicyHolidayBefore PlatformSettlementCycleDatePolicy = "HOLIDAY_BEFORE"
	PlatformSettlementCycleDatePolicyHolidayAfter  PlatformSettlementCycleDatePolicy = "HOLIDAY_AFTER"
	PlatformSettlementCycleDatePolicyCalendarDay   PlatformSettlementCycleDatePolicy = "CALENDAR_DAY"
)

// DayOfWeek 요일
type DayOfWeek string

const (
	DayOfWeekSun DayOfWeek = "SUN"
	DayOfWeekMon DayOfWeek = "MON"
	DayOfWeekTue DayOfWeek = "TUE"
	DayOfWeekWed DayOfWeek = "WED"
	DayOfWeekThu DayOfWeek = "THU"
	DayOfWeekFri DayOfWeek = "FRI"
	DayOfWeekSat DayOfWeek = "SAT"
)

// SettlementAmountType 정산 금액 취급 기준
type SettlementAmountType string

const (
	SettlementAmountTypeNet   SettlementAmountType = "NET"
	SettlementAmountTypeGross SettlementAmountType = "GROSS"
)

// PlatformPartnerTaxationType 플랫폼 파트너 과세 유형
type PlatformPartnerTaxationType string

const (
	PlatformPartnerTaxationTypeNormal                 PlatformPartnerTaxationType = "NORMAL"
	PlatformPartnerTaxationTypeSimpleTaxInvoiceIssuer PlatformPartnerTaxationType = "SIMPLE_TAX_INVOICE_ISSUER"
	PlatformPartnerTaxationTypeSimple                 PlatformPartnerTaxationType = "SIMPLE"
	PlatformPartnerTaxationTypeTaxFree                PlatformPartnerTaxationType = "TAX_FREE"
)

// PlatformPartnerBusinessStatus 플랫폼 파트너 사업자 상태
type PlatformPartnerBusinessStatus string

const (
	PlatformPartnerBusinessStatusNotVerified PlatformPartnerBusinessStatus = "NOT_VERIFIED"
	PlatformPartnerBusinessStatusVerifyError PlatformPartnerBusinessStatus = "VERIFY_ERROR"
	PlatformPartnerBusinessStatusNotFound    PlatformPartnerBusinessStatus = "NOT_FOUND"
	PlatformPartnerBusinessStatusInBusiness  PlatformPartnerBusinessStatus = "IN_BUSINESS"
	PlatformPartnerBusinessStatusClosed      PlatformPartnerBusinessStatus = "CLOSED"
	PlatformPartnerBusinessStatusSuspended   PlatformPartnerBusinessStatus = "SUSPENDED"
)

// PlatformPartnerMemberCompanyConnectionStatus 파트너 연동 사업자 연결 상태
type PlatformPartnerMemberCompanyConnectionStatus string

const (
	PlatformPartnerMemberCompanyConnectionStatusNotConnected      PlatformPartnerMemberCompanyConnectionStatus = "NOT_CONNECTED"
	PlatformPartnerMemberCompanyConnectionStatusConnectPending    PlatformPartnerMemberCompanyConnectionStatus = "CONNECT_PENDING"
	PlatformPartnerMemberCompanyConnectionStatusConnected         PlatformPartnerMemberCompanyConnectionStatus = "CONNECTED"
	PlatformPartnerMemberCompanyConnectionStatusConnectFailed     PlatformPartnerMemberCompanyConnectionStatus = "CONNECT_FAILED"
	PlatformPartnerMemberCompanyConnectionStatusDisconnectPending PlatformPartnerMemberCompanyConnectionStatus = "DISCONNECT_PENDING"
)

// PlatformFeeType 수수료 타입
type PlatformFeeType string

const (
	PlatformFeeTypeFixedAmount PlatformFeeType = "FIXED_AMOUNT"
	PlatformFeeTypeFixedRate   PlatformFeeType = "FIXED_RATE"
)

// PlatformSettlementCycleMethodType 정산 주기 계산 방식 타입
type PlatformSettlementCycleMethodType string

const (
	PlatformSettlementCycleMethodTypeDaily       PlatformSettlementCycleMethodType = "DAILY"
	PlatformSettlementCycleMethodTypeWeekly      PlatformSettlementCycleMethodType = "WEEKLY"
	PlatformSettlementCycleMethodTypeMonthly     PlatformSettlementCycleMethodType = "MONTHLY"
	PlatformSettlementCycleMethodTypeManualDates PlatformSettlementCycleMethodType = "MANUAL_DATES"
)

// PlatformPartnerTypeValue 파트너 유형
type PlatformPartnerTypeValue string

const (
	PlatformPartnerTypeValueBusiness    PlatformPartnerTypeValue = "BUSINESS"
	PlatformPartnerTypeValueWhtPayer    PlatformPartnerTypeValue = "WHT_PAYER"
	PlatformPartnerTypeValueNonWhtPayer PlatformPartnerTypeValue = "NON_WHT_PAYER"
)

// PlatformTransferType 정산건 유형
type PlatformTransferType string

const (
	PlatformTransferTypeManual      PlatformTransferType = "MANUAL"
	PlatformTransferTypeOrder       PlatformTransferType = "ORDER"
	PlatformTransferTypeOrderCancel PlatformTransferType = "ORDER_CANCEL"
)

// PlatformTransferStatus 정산건 상태
type PlatformTransferStatus string

const (
	PlatformTransferStatusScheduled PlatformTransferStatus = "SCHEDULED"
	PlatformTransferStatusInProcess PlatformTransferStatus = "IN_PROCESS"
	PlatformTransferStatusWithdrawn PlatformTransferStatus = "WITHDRAWN"
	PlatformTransferStatusPaidOut   PlatformTransferStatus = "PAID_OUT"
	PlatformTransferStatusCancelled PlatformTransferStatus = "CANCELLED"
)

// PlatformContact 플랫폼 파트너 담당자 연락 정보
type PlatformContact struct {
	Name        string  `json:"name"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Email       string  `json:"email"`
}

// PlatformAccount 플랫폼 정산 계좌
type PlatformAccount struct {
	Bank     common.Bank           `json:"bank"`
	Currency common.Currency       `json:"currency"`
	Number   string                `json:"number"`
	Holder   string                `json:"holder"`
	Status   PlatformAccountStatus `json:"status"`
}

// PlatformFee 플랫폼 중개수수료 정보
type PlatformFee struct {
	Type   PlatformFeeType `json:"type"`
	Amount *int64          `json:"amount,omitempty"` // FIXED_AMOUNT인 경우
	Rate   *int32          `json:"rate,omitempty"`   // FIXED_RATE인 경우 (밀리 퍼센트)
}

// PlatformSettlementCycleMethod 정산 주기 계산 방식
type PlatformSettlementCycleMethod struct {
	Type        PlatformSettlementCycleMethodType `json:"type"`
	DaysOfWeek  []DayOfWeek                       `json:"daysOfWeek,omitempty"`  // WEEKLY
	DaysOfMonth []int32                           `json:"daysOfMonth,omitempty"` // MONTHLY
	Dates       []MonthDay                        `json:"dates,omitempty"`       // MANUAL_DATES
}

// MonthDay 월/일 조합
type MonthDay struct {
	Month int32 `json:"month"`
	Day   int32 `json:"day"`
}

// PlatformSettlementCycle 정산 주기
type PlatformSettlementCycle struct {
	LagDays    int32                             `json:"lagDays"`
	DatePolicy PlatformSettlementCycleDatePolicy `json:"datePolicy"`
	Method     PlatformSettlementCycleMethod     `json:"method"`
}

// PlatformProperties 사용자 정의 속성 (map)
type PlatformProperties map[string]string

// PlatformPartnerType 파트너 유형별 추가 정보 (discriminated union)
type PlatformPartnerType struct {
	Type PlatformPartnerTypeValue `json:"type"`
	// Business fields
	CompanyName                         string                                       `json:"companyName,omitempty"`
	TaxationType                        PlatformPartnerTaxationType                  `json:"taxationType,omitempty"`
	BusinessStatus                      PlatformPartnerBusinessStatus                `json:"businessStatus,omitempty"`
	BusinessRegistrationNumber          string                                       `json:"businessRegistrationNumber,omitempty"`
	RepresentativeName                  string                                       `json:"representativeName,omitempty"`
	CompanyAddress                      *string                                      `json:"companyAddress,omitempty"`
	BusinessType                        *string                                      `json:"businessType,omitempty"`
	BusinessClass                       *string                                      `json:"businessClass,omitempty"`
	MemberCompanyConnectionStatus       PlatformPartnerMemberCompanyConnectionStatus `json:"memberCompanyConnectionStatus,omitempty"`
	MemberCompanyConnectionFailedReason *string                                      `json:"memberCompanyConnectionFailedReason,omitempty"`
	// WhtPayer / NonWhtPayer fields
	Birthdate *string `json:"birthdate,omitempty"`
}

// IsBusiness 사업자 파트너인지 확인
func (p *PlatformPartnerType) IsBusiness() bool {
	return p.Type == PlatformPartnerTypeValueBusiness
}

// IsWhtPayer 원천징수 대상자인지 확인
func (p *PlatformPartnerType) IsWhtPayer() bool {
	return p.Type == PlatformPartnerTypeValueWhtPayer
}

// IsNonWhtPayer 원천징수 비대상자인지 확인
func (p *PlatformPartnerType) IsNonWhtPayer() bool {
	return p.Type == PlatformPartnerTypeValueNonWhtPayer
}

// PlatformPartner 파트너
type PlatformPartner struct {
	ID                    string                `json:"id"`
	GraphqlID             string                `json:"graphqlId"`
	Name                  string                `json:"name"`
	Contact               PlatformContact       `json:"contact"`
	Account               PlatformAccount       `json:"account"`
	Status                PlatformPartnerStatus `json:"status"`
	DefaultContractID     string                `json:"defaultContractId"`
	Memo                  *string               `json:"memo,omitempty"`
	Tags                  []string              `json:"tags"`
	Type                  PlatformPartnerType   `json:"type"`
	IsArchived            bool                  `json:"isArchived"`
	AppliedAt             string                `json:"appliedAt"`
	UserDefinedProperties PlatformProperties    `json:"userDefinedProperties"`
	IsForTest             bool                  `json:"isForTest"`
}

// PlatformContract 계약
type PlatformContract struct {
	ID                       string                  `json:"id"`
	GraphqlID                string                  `json:"graphqlId"`
	Name                     string                  `json:"name"`
	Memo                     *string                 `json:"memo,omitempty"`
	PlatformFee              PlatformFee             `json:"platformFee"`
	SettlementCycle          PlatformSettlementCycle `json:"settlementCycle"`
	PlatformFeeVatPayer      PlatformPayer           `json:"platformFeeVatPayer"`
	SubtractPaymentVatAmount bool                    `json:"subtractPaymentVatAmount"`
	IsArchived               bool                    `json:"isArchived"`
	AppliedAt                string                  `json:"appliedAt"`
	IsForTest                bool                    `json:"isForTest"`
}

// PlatformAdditionalFeePolicy 추가 수수료 정책
type PlatformAdditionalFeePolicy struct {
	ID         string        `json:"id"`
	GraphqlID  string        `json:"graphqlId"`
	Name       string        `json:"name"`
	Fee        PlatformFee   `json:"fee"`
	Memo       *string       `json:"memo,omitempty"`
	VatPayer   PlatformPayer `json:"vatPayer"`
	IsArchived bool          `json:"isArchived"`
	AppliedAt  string        `json:"appliedAt"`
	IsForTest  bool          `json:"isForTest"`
}

// PlatformDiscountSharePolicy 할인 분담 정책
type PlatformDiscountSharePolicy struct {
	ID               string  `json:"id"`
	GraphqlID        string  `json:"graphqlId"`
	Name             string  `json:"name"`
	PartnerShareRate int32   `json:"partnerShareRate"`
	Memo             *string `json:"memo,omitempty"`
	IsArchived       bool    `json:"isArchived"`
	AppliedAt        string  `json:"appliedAt"`
	IsForTest        bool    `json:"isForTest"`
}

// PlatformSetting 플랫폼 설정
type PlatformSetting struct {
	DefaultWithdrawalMemo                     *string              `json:"defaultWithdrawalMemo,omitempty"`
	DefaultDepositMemo                        *string              `json:"defaultDepositMemo,omitempty"`
	SupportsMultipleOrderTransfersPerPartner  bool                 `json:"supportsMultipleOrderTransfersPerPartner"`
	AdjustSettlementDateAfterHolidayIfEarlier bool                 `json:"adjustSettlementDateAfterHolidayIfEarlier"`
	DeductWht                                 bool                 `json:"deductWht"`
	SettlementAmountType                      SettlementAmountType `json:"settlementAmountType"`
	IsForTest                                 bool                 `json:"isForTest"`
}

// PageInput 페이지 입력
type PageInput struct {
	Number *int32 `json:"number,omitempty"`
	Size   *int32 `json:"size,omitempty"`
}

// PageInfo 페이지 정보
type PageInfo struct {
	Number     int32 `json:"number"`
	Size       int32 `json:"size"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int32 `json:"totalPages"`
}

// PlatformPartnerFilterInput 파트너 필터 입력
type PlatformPartnerFilterInput struct {
	IDs          []string                   `json:"ids,omitempty"`
	Keyword      *string                    `json:"keyword,omitempty"`
	Statuses     []PlatformPartnerStatus    `json:"statuses,omitempty"`
	ContractIDs  []string                   `json:"contractIds,omitempty"`
	Tags         []string                   `json:"tags,omitempty"`
	IsArchived   *bool                      `json:"isArchived,omitempty"`
	PartnerTypes []PlatformPartnerTypeValue `json:"partnerTypes,omitempty"`
}

// PlatformContractFilterInput 계약 필터 입력
type PlatformContractFilterInput struct {
	IDs        []string `json:"ids,omitempty"`
	Keyword    *string  `json:"keyword,omitempty"`
	IsArchived *bool    `json:"isArchived,omitempty"`
}

// PlatformAdditionalFeePolicyFilterInput 추가 수수료 정책 필터 입력
type PlatformAdditionalFeePolicyFilterInput struct {
	IDs        []string `json:"ids,omitempty"`
	Keyword    *string  `json:"keyword,omitempty"`
	IsArchived *bool    `json:"isArchived,omitempty"`
}

// PlatformDiscountSharePolicyFilterInput 할인 분담 정책 필터 입력
type PlatformDiscountSharePolicyFilterInput struct {
	IDs        []string `json:"ids,omitempty"`
	Keyword    *string  `json:"keyword,omitempty"`
	IsArchived *bool    `json:"isArchived,omitempty"`
}

// PlatformTransferFilterInput 정산건 필터 입력
type PlatformTransferFilterInput struct {
	PartnerIDs       []string                 `json:"partnerIds,omitempty"`
	ContractIDs      []string                 `json:"contractIds,omitempty"`
	TransferTypes    []PlatformTransferType   `json:"transferTypes,omitempty"`
	TransferStatuses []PlatformTransferStatus `json:"transferStatuses,omitempty"`
	PaymentIDs       []string                 `json:"paymentIds,omitempty"`
}

// PlatformUserDefinedPropertyKeyValue 사용자 정의 속성 키-값
type PlatformUserDefinedPropertyKeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// ---- Response Types ----

// GetPlatformPartnersResponse 파트너 다건 조회 응답
type GetPlatformPartnersResponse struct {
	Items []PlatformPartner `json:"items"`
	Page  PageInfo          `json:"page"`
}

// CreatePlatformPartnerResponse 파트너 생성 응답
type CreatePlatformPartnerResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// CreatePlatformPartnersResponse 파트너 다건 생성 응답
type CreatePlatformPartnersResponse struct {
	Partners []PlatformPartner `json:"partners"`
}

// UpdatePlatformPartnerResponse 파트너 수정 응답
type UpdatePlatformPartnerResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// ArchivePlatformPartnerResponse 파트너 보관 응답
type ArchivePlatformPartnerResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// RecoverPlatformPartnerResponse 파트너 복원 응답
type RecoverPlatformPartnerResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// ConnectPartnerMemberCompanyResponse 파트너 국세청 연동 응답
type ConnectPartnerMemberCompanyResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// ConnectBulkPartnerMemberCompanyResponse 파트너 일괄 국세청 연동 응답
type ConnectBulkPartnerMemberCompanyResponse struct {
	ConnectedPartnerCount int32 `json:"connectedPartnerCount"`
}

// DisconnectPartnerMemberCompanyResponse 파트너 국세청 연동 해제 응답
type DisconnectPartnerMemberCompanyResponse struct {
	Partner PlatformPartner `json:"partner"`
}

// DisconnectBulkPartnerMemberCompanyResponse 파트너 일괄 국세청 연동 해제 응답
type DisconnectBulkPartnerMemberCompanyResponse struct {
	DisconnectedPartnerCount int32 `json:"disconnectedPartnerCount"`
}

// GetPlatformContractsResponse 계약 다건 조회 응답
type GetPlatformContractsResponse struct {
	Items []PlatformContract `json:"items"`
	Page  PageInfo           `json:"page"`
}

// CreatePlatformContractResponse 계약 생성 응답
type CreatePlatformContractResponse struct {
	Contract PlatformContract `json:"contract"`
}

// UpdatePlatformContractResponse 계약 수정 응답
type UpdatePlatformContractResponse struct {
	Contract PlatformContract `json:"contract"`
}

// ArchivePlatformContractResponse 계약 보관 응답
type ArchivePlatformContractResponse struct {
	Contract PlatformContract `json:"contract"`
}

// RecoverPlatformContractResponse 계약 복원 응답
type RecoverPlatformContractResponse struct {
	Contract PlatformContract `json:"contract"`
}

// GetPlatformAdditionalFeePoliciesResponse 추가 수수료 정책 다건 조회 응답
type GetPlatformAdditionalFeePoliciesResponse struct {
	Items []PlatformAdditionalFeePolicy `json:"items"`
	Page  PageInfo                      `json:"page"`
}

// CreatePlatformAdditionalFeePolicyResponse 추가 수수료 정책 생성 응답
type CreatePlatformAdditionalFeePolicyResponse struct {
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// UpdatePlatformAdditionalFeePolicyResponse 추가 수수료 정책 수정 응답
type UpdatePlatformAdditionalFeePolicyResponse struct {
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// ArchivePlatformAdditionalFeePolicyResponse 추가 수수료 정책 보관 응답
type ArchivePlatformAdditionalFeePolicyResponse struct {
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// RecoverPlatformAdditionalFeePolicyResponse 추가 수수료 정책 복원 응답
type RecoverPlatformAdditionalFeePolicyResponse struct {
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// GetPlatformDiscountSharePoliciesResponse 할인 분담 정책 다건 조회 응답
type GetPlatformDiscountSharePoliciesResponse struct {
	Items []PlatformDiscountSharePolicy `json:"items"`
	Page  PageInfo                      `json:"page"`
}

// CreatePlatformDiscountSharePolicyResponse 할인 분담 정책 생성 응답
type CreatePlatformDiscountSharePolicyResponse struct {
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// UpdatePlatformDiscountSharePolicyResponse 할인 분담 정책 수정 응답
type UpdatePlatformDiscountSharePolicyResponse struct {
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// ArchivePlatformDiscountSharePolicyResponse 할인 분담 정책 보관 응답
type ArchivePlatformDiscountSharePolicyResponse struct {
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// RecoverPlatformDiscountSharePolicyResponse 할인 분담 정책 복원 응답
type RecoverPlatformDiscountSharePolicyResponse struct {
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// UpdatePlatformSettingResponse 플랫폼 설정 수정 응답
type UpdatePlatformSettingResponse struct {
	Setting PlatformSetting `json:"setting"`
}

// ---- Body/Input Types ----

// CreatePlatformPartnerBodyContact 파트너 생성 연락처 정보
type CreatePlatformPartnerBodyContact struct {
	Name        string  `json:"name"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Email       string  `json:"email"`
}

// CreatePlatformPartnerBodyAccount 파트너 생성 계좌 정보
type CreatePlatformPartnerBodyAccount struct {
	Bank                  common.Bank     `json:"bank"`
	Currency              common.Currency `json:"currency"`
	Number                string          `json:"number"`
	Holder                string          `json:"holder"`
	AccountVerificationID *string         `json:"accountVerificationId,omitempty"`
}

// CreatePlatformPartnerBodyTypeBusiness 사업자 파트너 생성 정보
type CreatePlatformPartnerBodyTypeBusiness struct {
	Type                       PlatformPartnerTypeValue    `json:"type"` // "BUSINESS"
	CompanyName                string                      `json:"companyName"`
	TaxationType               PlatformPartnerTaxationType `json:"taxationType"`
	BusinessRegistrationNumber string                      `json:"businessRegistrationNumber"`
	RepresentativeName         string                      `json:"representativeName"`
	CompanyAddress             *string                     `json:"companyAddress,omitempty"`
	BusinessType               *string                     `json:"businessType,omitempty"`
	BusinessClass              *string                     `json:"businessClass,omitempty"`
	CompanyVerificationID      *string                     `json:"companyVerificationId,omitempty"`
}

// CreatePlatformPartnerBodyTypeWhtPayer 원천징수 대상자 파트너 생성 정보
type CreatePlatformPartnerBodyTypeWhtPayer struct {
	Type      PlatformPartnerTypeValue `json:"type"` // "WHT_PAYER"
	Birthdate *string                  `json:"birthdate,omitempty"`
}

// CreatePlatformPartnerBodyTypeNonWhtPayer 원천징수 비대상자 파트너 생성 정보
type CreatePlatformPartnerBodyTypeNonWhtPayer struct {
	Type      PlatformPartnerTypeValue `json:"type"` // "NON_WHT_PAYER"
	Birthdate *string                  `json:"birthdate,omitempty"`
}

// CreatePlatformPartnerBodyType 파트너 생성 유형 정보
type CreatePlatformPartnerBodyType struct {
	Type                       PlatformPartnerTypeValue    `json:"type"`
	CompanyName                string                      `json:"companyName,omitempty"`
	TaxationType               PlatformPartnerTaxationType `json:"taxationType,omitempty"`
	BusinessRegistrationNumber string                      `json:"businessRegistrationNumber,omitempty"`
	RepresentativeName         string                      `json:"representativeName,omitempty"`
	CompanyAddress             *string                     `json:"companyAddress,omitempty"`
	BusinessType               *string                     `json:"businessType,omitempty"`
	BusinessClass              *string                     `json:"businessClass,omitempty"`
	CompanyVerificationID      *string                     `json:"companyVerificationId,omitempty"`
	Birthdate                  *string                     `json:"birthdate,omitempty"`
}

// CreatePlatformPartnerBody 파트너 생성 요청 바디
type CreatePlatformPartnerBody struct {
	ID                    *string                          `json:"id,omitempty"`
	Name                  string                           `json:"name"`
	Contact               CreatePlatformPartnerBodyContact `json:"contact"`
	Account               CreatePlatformPartnerBodyAccount `json:"account"`
	DefaultContractID     string                           `json:"defaultContractId"`
	Memo                  *string                          `json:"memo,omitempty"`
	Tags                  []string                         `json:"tags"`
	Type                  CreatePlatformPartnerBodyType    `json:"type"`
	UserDefinedProperties PlatformProperties               `json:"userDefinedProperties,omitempty"`
}

// UpdatePlatformPartnerBodyContact 파트너 수정 연락처 정보
type UpdatePlatformPartnerBodyContact struct {
	Name        *string `json:"name,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Email       *string `json:"email,omitempty"`
}

// UpdatePlatformPartnerBodyAccount 파트너 수정 계좌 정보
type UpdatePlatformPartnerBodyAccount struct {
	Bank                  *common.Bank     `json:"bank,omitempty"`
	Currency              *common.Currency `json:"currency,omitempty"`
	Number                *string          `json:"number,omitempty"`
	Holder                *string          `json:"holder,omitempty"`
	AccountVerificationID *string          `json:"accountVerificationId,omitempty"`
}

// UpdatePlatformPartnerBodyType 파트너 수정 유형 정보
type UpdatePlatformPartnerBodyType struct {
	Type                       *PlatformPartnerTypeValue    `json:"type,omitempty"`
	CompanyName                *string                      `json:"companyName,omitempty"`
	TaxationType               *PlatformPartnerTaxationType `json:"taxationType,omitempty"`
	BusinessRegistrationNumber *string                      `json:"businessRegistrationNumber,omitempty"`
	RepresentativeName         *string                      `json:"representativeName,omitempty"`
	CompanyAddress             *string                      `json:"companyAddress,omitempty"`
	BusinessType               *string                      `json:"businessType,omitempty"`
	BusinessClass              *string                      `json:"businessClass,omitempty"`
	CompanyVerificationID      *string                      `json:"companyVerificationId,omitempty"`
	Birthdate                  *string                      `json:"birthdate,omitempty"`
}

// UpdatePlatformPartnerBody 파트너 수정 요청 바디
type UpdatePlatformPartnerBody struct {
	Name                  *string                           `json:"name,omitempty"`
	Contact               *UpdatePlatformPartnerBodyContact `json:"contact,omitempty"`
	Account               *UpdatePlatformPartnerBodyAccount `json:"account,omitempty"`
	DefaultContractID     *string                           `json:"defaultContractId,omitempty"`
	Memo                  *string                           `json:"memo,omitempty"`
	Tags                  []string                          `json:"tags,omitempty"`
	Type                  *UpdatePlatformPartnerBodyType    `json:"type,omitempty"`
	UserDefinedProperties PlatformProperties                `json:"userDefinedProperties,omitempty"`
}

// PlatformFeeInput 수수료 입력
type PlatformFeeInput struct {
	Type   PlatformFeeType `json:"type"`
	Amount *int64          `json:"amount,omitempty"`
	Rate   *int32          `json:"rate,omitempty"`
}

// PlatformSettlementCycleMethodInput 정산 주기 계산 방식 입력
type PlatformSettlementCycleMethodInput struct {
	Type        PlatformSettlementCycleMethodType `json:"type"`
	DaysOfWeek  []DayOfWeek                       `json:"daysOfWeek,omitempty"`
	DaysOfMonth []int32                           `json:"daysOfMonth,omitempty"`
	Dates       []MonthDay                        `json:"dates,omitempty"`
}

// PlatformSettlementCycleInput 정산 주기 입력
type PlatformSettlementCycleInput struct {
	LagDays    int32                              `json:"lagDays"`
	DatePolicy PlatformSettlementCycleDatePolicy  `json:"datePolicy"`
	Method     PlatformSettlementCycleMethodInput `json:"method"`
}

// CreatePlatformContractBody 계약 생성 요청 바디
type CreatePlatformContractBody struct {
	ID                       *string                      `json:"id,omitempty"`
	Name                     string                       `json:"name"`
	Memo                     *string                      `json:"memo,omitempty"`
	PlatformFee              PlatformFeeInput             `json:"platformFee"`
	SettlementCycle          PlatformSettlementCycleInput `json:"settlementCycle"`
	PlatformFeeVatPayer      PlatformPayer                `json:"platformFeeVatPayer"`
	SubtractPaymentVatAmount bool                         `json:"subtractPaymentVatAmount"`
}

// UpdatePlatformContractBody 계약 수정 요청 바디
type UpdatePlatformContractBody struct {
	Name                     *string                       `json:"name,omitempty"`
	Memo                     *string                       `json:"memo,omitempty"`
	PlatformFee              *PlatformFeeInput             `json:"platformFee,omitempty"`
	SettlementCycle          *PlatformSettlementCycleInput `json:"settlementCycle,omitempty"`
	PlatformFeeVatPayer      *PlatformPayer                `json:"platformFeeVatPayer,omitempty"`
	SubtractPaymentVatAmount *bool                         `json:"subtractPaymentVatAmount,omitempty"`
}

// CreatePlatformAdditionalFeePolicyBody 추가 수수료 정책 생성 요청 바디
type CreatePlatformAdditionalFeePolicyBody struct {
	ID       *string          `json:"id,omitempty"`
	Name     string           `json:"name"`
	Fee      PlatformFeeInput `json:"fee"`
	Memo     *string          `json:"memo,omitempty"`
	VatPayer PlatformPayer    `json:"vatPayer"`
}

// UpdatePlatformAdditionalFeePolicyBody 추가 수수료 정책 수정 요청 바디
type UpdatePlatformAdditionalFeePolicyBody struct {
	Name     *string           `json:"name,omitempty"`
	Fee      *PlatformFeeInput `json:"fee,omitempty"`
	Memo     *string           `json:"memo,omitempty"`
	VatPayer *PlatformPayer    `json:"vatPayer,omitempty"`
}

// CreatePlatformDiscountSharePolicyBody 할인 분담 정책 생성 요청 바디
type CreatePlatformDiscountSharePolicyBody struct {
	ID               *string `json:"id,omitempty"`
	Name             string  `json:"name"`
	PartnerShareRate int32   `json:"partnerShareRate"`
	Memo             *string `json:"memo,omitempty"`
}

// UpdatePlatformDiscountSharePolicyBody 할인 분담 정책 수정 요청 바디
type UpdatePlatformDiscountSharePolicyBody struct {
	Name             *string `json:"name,omitempty"`
	PartnerShareRate *int32  `json:"partnerShareRate,omitempty"`
	Memo             *string `json:"memo,omitempty"`
}

// UpdatePlatformSettingBody 플랫폼 설정 수정 요청 바디
type UpdatePlatformSettingBody struct {
	DefaultWithdrawalMemo                     *string               `json:"defaultWithdrawalMemo,omitempty"`
	DefaultDepositMemo                        *string               `json:"defaultDepositMemo,omitempty"`
	SupportsMultipleOrderTransfersPerPartner  *bool                 `json:"supportsMultipleOrderTransfersPerPartner,omitempty"`
	AdjustSettlementDateAfterHolidayIfEarlier *bool                 `json:"adjustSettlementDateAfterHolidayIfEarlier,omitempty"`
	DeductWht                                 *bool                 `json:"deductWht,omitempty"`
	SettlementAmountType                      *SettlementAmountType `json:"settlementAmountType,omitempty"`
}

// ---- Schedule Types ----

// SchedulePlatformPartnerBody 파트너 예약 요청 바디
type SchedulePlatformPartnerBody struct {
	// 반영할 업데이트 내용
	Update UpdatePlatformPartnerBody `json:"update"`
	// 업데이트 적용 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// SchedulePlatformPartnerResponse 파트너 예약 응답
type SchedulePlatformPartnerResponse struct {
	// 예약된 파트너 정보
	Partner PlatformPartner `json:"partner"`
}

// ReschedulePlatformPartnerResponse 파트너 예약 재설정 응답
type ReschedulePlatformPartnerResponse struct {
	// 예약된 파트너 정보
	Partner PlatformPartner `json:"partner"`
}

// CancelPlatformPartnerScheduleResponse 파트너 예약 취소 응답
type CancelPlatformPartnerScheduleResponse struct{}

// SchedulePlatformPartnersBodyUpdate 파트너 일괄 예약 업데이트 내용
type SchedulePlatformPartnersBodyUpdate struct {
	// 이름
	Name *string `json:"name,omitempty"`
	// 연락처
	Contact *UpdatePlatformPartnerBodyContact `json:"contact,omitempty"`
	// 정산 계좌
	Account *UpdatePlatformPartnerBodyAccount `json:"account,omitempty"`
	// 기본 계약 아이디
	DefaultContractID *string `json:"defaultContractId,omitempty"`
	// 메모
	Memo *string `json:"memo,omitempty"`
	// 태그
	Tags []string `json:"tags,omitempty"`
	// 사용자 정의 속성
	UserDefinedProperties PlatformProperties `json:"userDefinedProperties,omitempty"`
}

// SchedulePlatformPartnersBody 파트너 일괄 예약 요청 바디
type SchedulePlatformPartnersBody struct {
	// 필터
	Filter *PlatformPartnerFilterInput `json:"filter,omitempty"`
	// 반영할 업데이트 내용
	Update SchedulePlatformPartnersBodyUpdate `json:"update"`
	// 업데이트 적용 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// SchedulePlatformPartnersResponse 파트너 일괄 예약 응답
type SchedulePlatformPartnersResponse struct{}

// SchedulePlatformContractBody 계약 예약 요청 바디
type SchedulePlatformContractBody struct {
	// 반영할 업데이트 내용
	Update UpdatePlatformContractBody `json:"update"`
	// 업데이트 적용 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// SchedulePlatformContractResponse 계약 예약 응답
type SchedulePlatformContractResponse struct {
	// 예약된 계약 정보
	Contract PlatformContract `json:"contract"`
}

// ReschedulePlatformContractResponse 계약 예약 재설정 응답
type ReschedulePlatformContractResponse struct {
	// 예약된 계약 정보
	Contract PlatformContract `json:"contract"`
}

// CancelPlatformContractScheduleResponse 계약 예약 취소 응답
type CancelPlatformContractScheduleResponse struct{}

// SchedulePlatformAdditionalFeePolicyBody 추가 수수료 정책 예약 요청 바디
type SchedulePlatformAdditionalFeePolicyBody struct {
	// 반영할 업데이트 내용
	Update UpdatePlatformAdditionalFeePolicyBody `json:"update"`
	// 업데이트 적용 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// SchedulePlatformAdditionalFeePolicyResponse 추가 수수료 정책 예약 응답
type SchedulePlatformAdditionalFeePolicyResponse struct {
	// 예약된 추가 수수료 정책 정보
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// ReschedulePlatformAdditionalFeePolicyResponse 추가 수수료 정책 예약 재설정 응답
type ReschedulePlatformAdditionalFeePolicyResponse struct {
	// 예약된 추가 수수료 정책 정보
	AdditionalFeePolicy PlatformAdditionalFeePolicy `json:"additionalFeePolicy"`
}

// CancelPlatformAdditionalFeePolicyScheduleResponse 추가 수수료 정책 예약 취소 응답
type CancelPlatformAdditionalFeePolicyScheduleResponse struct{}

// SchedulePlatformDiscountSharePolicyBody 할인 분담 정책 예약 요청 바디
type SchedulePlatformDiscountSharePolicyBody struct {
	// 반영할 업데이트 내용
	Update UpdatePlatformDiscountSharePolicyBody `json:"update"`
	// 업데이트 적용 시점 (RFC 3339)
	AppliedAt string `json:"appliedAt"`
}

// SchedulePlatformDiscountSharePolicyResponse 할인 분담 정책 예약 응답
type SchedulePlatformDiscountSharePolicyResponse struct {
	// 예약된 할인 분담 정책 정보
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// ReschedulePlatformDiscountSharePolicyResponse 할인 분담 정책 예약 재설정 응답
type ReschedulePlatformDiscountSharePolicyResponse struct {
	// 예약된 할인 분담 정책 정보
	DiscountSharePolicy PlatformDiscountSharePolicy `json:"discountSharePolicy"`
}

// CancelPlatformDiscountSharePolicyScheduleResponse 할인 분담 정책 예약 취소 응답
type CancelPlatformDiscountSharePolicyScheduleResponse struct{}

// ---- Platform Bulk Account Transfer Types ----

// PlatformAccountTransferStatusStats 계좌 이체 상태별 통계
type PlatformAccountTransferStatusStats struct {
	// 준비됨
	Prepared int64 `json:"prepared"`
	// 예약됨
	Scheduled int64 `json:"scheduled"`
	// 취소됨
	Cancelled int64 `json:"cancelled"`
	// 중단됨
	Stopped int64 `json:"stopped"`
	// 처리 중
	Processing int64 `json:"processing"`
	// 비동기 처리 중
	AsyncProcessing int64 `json:"asyncProcessing"`
	// 성공
	Succeeded int64 `json:"succeeded"`
	// 실패
	Failed int64 `json:"failed"`
}

// PlatformBulkAccountTransferStatus 일괄 이체 상태
type PlatformBulkAccountTransferStatus string

const (
	PlatformBulkAccountTransferStatusPREPARED  PlatformBulkAccountTransferStatus = "PREPARED"
	PlatformBulkAccountTransferStatusSCHEDULED PlatformBulkAccountTransferStatus = "SCHEDULED"
	PlatformBulkAccountTransferStatusONGOING   PlatformBulkAccountTransferStatus = "ONGOING"
	PlatformBulkAccountTransferStatusCOMPLETED PlatformBulkAccountTransferStatus = "COMPLETED"
)

// PlatformBulkAccountTransferStats 일괄 이체 통계
type PlatformBulkAccountTransferStats struct {
	// 금액 통계
	Amount PlatformAccountTransferStatusStats `json:"amount"`
	// 건수 통계
	Count PlatformAccountTransferStatusStats `json:"count"`
}

// PlatformBulkAccountTransfer 일괄 이체 내역
type PlatformBulkAccountTransfer struct {
	// 일괄 이체 고유 아이디
	ID string `json:"id"`
	// GraphQL ID
	GraphqlId string `json:"graphqlId"`
	// 생성자 ID
	CreatorId string `json:"creatorId"`
	// 출금 계좌 아이디
	BankAccountId string `json:"bankAccountId"`
	// 출금 계좌 GraphQL ID
	BankAccountGraphqlId string `json:"bankAccountGraphqlId"`
	// 총 금액
	TotalAmount int64 `json:"totalAmount"`
	// 상태
	Status PlatformBulkAccountTransferStatus `json:"status"`
	// 통계
	Stats PlatformBulkAccountTransferStats `json:"stats"`
	// 상태 업데이트 일시 (RFC 3339)
	StatusUpdatedAt string `json:"statusUpdatedAt"`
	// 생성 일시 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 수정 일시 (RFC 3339)
	UpdatedAt string `json:"updatedAt"`
	// 예약 일시 (RFC 3339)
	ScheduledAt *string `json:"scheduledAt,omitempty"`
}

// PlatformBulkAccountTransferStatusStats 일괄 이체 상태별 통계
type PlatformBulkAccountTransferStatusStats struct {
	// 준비됨
	Prepared int64 `json:"prepared"`
	// 예약됨
	Scheduled int64 `json:"scheduled"`
	// 진행 중
	Ongoing int64 `json:"ongoing"`
	// 완료됨
	Completed int64 `json:"completed"`
}

// PlatformBulkAccountTransferFilterInputCriteria 일괄 이체 필터 기준
type PlatformBulkAccountTransferFilterInputCriteria struct {
	// 생성 일시 범위
	TimestampRange *common.DateTimeRange `json:"timestampRange,omitempty"`
	// 상태 업데이트 일시 범위
	StatusUpdatedTimestampRange *common.DateTimeRange `json:"statusUpdatedTimestampRange,omitempty"`
	// 일괄 이체 아이디
	BulkAccountTransferId *string `json:"bulkAccountTransferId,omitempty"`
}

// PlatformBulkAccountTransferFilterInput 일괄 이체 필터 입력
type PlatformBulkAccountTransferFilterInput struct {
	// 상태 목록
	Statuses []PlatformBulkAccountTransferStatus `json:"statuses,omitempty"`
	// 필터 기준
	Criteria *PlatformBulkAccountTransferFilterInputCriteria `json:"criteria,omitempty"`
}

// GetPlatformBulkAccountTransfersBody 일괄 이체 내역 조회 요청
type GetPlatformBulkAccountTransfersBody struct {
	// 테스트 모드 여부
	IsForTest *bool `json:"isForTest,omitempty"`
	// 페이지 정보
	Page *common.PageInput `json:"page,omitempty"`
	// 필터
	Filter *PlatformBulkAccountTransferFilterInput `json:"filter,omitempty"`
}

// GetPlatformBulkAccountTransfersResponse 일괄 이체 내역 조회 응답
type GetPlatformBulkAccountTransfersResponse struct {
	// 일괄 이체 내역 목록
	Items []PlatformBulkAccountTransfer `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
	// 상태별 통계
	Counts PlatformBulkAccountTransferStatusStats `json:"counts"`
}
