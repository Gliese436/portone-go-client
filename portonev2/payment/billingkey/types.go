package billingkey

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// BillingKeyStatus 빌링키 상태
type BillingKeyStatus string

const (
	BillingKeyStatusISSUED  BillingKeyStatus = "ISSUED"
	BillingKeyStatusDELETED BillingKeyStatus = "DELETED"
)

// BillingKeyPaymentMethodType 빌링키 결제수단 타입
type BillingKeyPaymentMethodType string

const (
	BillingKeyPaymentMethodTypeCARD     BillingKeyPaymentMethodType = "BillingKeyPaymentMethodCard"
	BillingKeyPaymentMethodTypeTRANSFER BillingKeyPaymentMethodType = "BillingKeyPaymentMethodTransfer"
	BillingKeyPaymentMethodTypeMOBILE   BillingKeyPaymentMethodType = "BillingKeyPaymentMethodMobile"
	BillingKeyPaymentMethodTypeEASYPAY  BillingKeyPaymentMethodType = "BillingKeyPaymentMethodEasyPay"
	BillingKeyPaymentMethodTypePAYPAL   BillingKeyPaymentMethodType = "BillingKeyPaymentMethodPaypal"
)

// BillingKeyPaymentMethod 빌링키 결제수단 정보
type BillingKeyPaymentMethod struct {
	Type string `json:"type"`
	// Card fields
	Card *common.Card `json:"card,omitempty"`
	// Transfer fields
	Bank *common.Bank `json:"bank,omitempty"`
	// Mobile fields
	Carrier     *string `json:"carrier,omitempty"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// EasyPay fields
	Provider *common.EasyPayProvider `json:"provider,omitempty"`
}

// BillingKeyInfo 빌링키 정보
type BillingKeyInfo struct {
	// 빌링키 상태
	Status BillingKeyStatus `json:"status"`
	// 빌링키
	BillingKey string `json:"billingKey"`
	// 고객사 아이디
	MerchantId string `json:"merchantId"`
	// 상점 아이디
	StoreId string `json:"storeId"`
	// 빌링키 결제수단 상세 정보
	Methods []BillingKeyPaymentMethod `json:"methods,omitempty"`
	// 빌링키 발급 시 사용된 채널
	Channels []common.SelectedChannel `json:"channels,omitempty"`
	// 고객 정보
	Customer common.Customer `json:"customer"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 고객사가 채번하는 빌링키 발급 건 고유 아이디
	IssueId *string `json:"issueId,omitempty"`
	// 빌링키 발급 건 이름
	IssueName *string `json:"issueName,omitempty"`
	// 발급 요청 시점 (RFC 3339)
	RequestedAt *string `json:"requestedAt,omitempty"`
	// 발급 시점 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
	// 채널 그룹
	ChannelGroup *common.ChannelGroupSummary `json:"channelGroup,omitempty"`
	// 삭제 시점 (RFC 3339) - DELETED 상태에서만 존재
	DeletedAt *string `json:"deletedAt,omitempty"`
}

// IsIssued 발급 완료 상태인지 확인
func (b *BillingKeyInfo) IsIssued() bool {
	return b.Status == BillingKeyStatusISSUED
}

// IsDeleted 삭제 완료 상태인지 확인
func (b *BillingKeyInfo) IsDeleted() bool {
	return b.Status == BillingKeyStatusDELETED
}

// BillingKeyInfoSummary 빌링키 요약 정보
type BillingKeyInfoSummary struct {
	// 빌링키
	BillingKey string `json:"billingKey"`
	// 빌링키 발급 시 사용된 채널
	Channels []common.SelectedChannel `json:"channels,omitempty"`
	// 발급 시점 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
}

// BillingKeyFailure 빌링키 발급 실패 정보
type BillingKeyFailure struct {
	// 실패 사유
	Reason *string `json:"reason,omitempty"`
	// PG사 실패 코드
	PgCode *string `json:"pgCode,omitempty"`
	// PG사 실패 메시지
	PgMessage *string `json:"pgMessage,omitempty"`
}

// BillingKeySortBy 빌링키 정렬 기준
type BillingKeySortBy string

const (
	BillingKeySortByREQUESTED_AT      BillingKeySortBy = "REQUESTED_AT"
	BillingKeySortByISSUED_AT         BillingKeySortBy = "ISSUED_AT"
	BillingKeySortByDELETED_AT        BillingKeySortBy = "DELETED_AT"
	BillingKeySortBySTATUS_CHANGED_AT BillingKeySortBy = "STATUS_CHANGED_AT"
)

// BillingKeySortInput 빌링키 정렬 입력
type BillingKeySortInput struct {
	// 정렬 기준
	By *BillingKeySortBy `json:"by,omitempty"`
	// 정렬 순서
	Order *common.SortOrder `json:"order,omitempty"`
}

// BillingKeyTextSearchField 빌링키 텍스트 검색 필드
type BillingKeyTextSearchField string

const (
	BillingKeyTextSearchFieldCARD_BIN           BillingKeyTextSearchField = "CARD_BIN"
	BillingKeyTextSearchFieldCARD_NUMBER        BillingKeyTextSearchField = "CARD_NUMBER"
	BillingKeyTextSearchFieldPG_MERCHANT_ID     BillingKeyTextSearchField = "PG_MERCHANT_ID"
	BillingKeyTextSearchFieldCUSTOMER_NAME      BillingKeyTextSearchField = "CUSTOMER_NAME"
	BillingKeyTextSearchFieldCUSTOMER_EMAIL     BillingKeyTextSearchField = "CUSTOMER_EMAIL"
	BillingKeyTextSearchFieldCUSTOMER_PHONE     BillingKeyTextSearchField = "CUSTOMER_PHONE"
	BillingKeyTextSearchFieldCUSTOMER_ADDRESS   BillingKeyTextSearchField = "CUSTOMER_ADDRESS"
	BillingKeyTextSearchFieldCUSTOMER_ZIPCODE   BillingKeyTextSearchField = "CUSTOMER_ZIPCODE"
	BillingKeyTextSearchFieldUSER_AGENT         BillingKeyTextSearchField = "USER_AGENT"
	BillingKeyTextSearchFieldBILLING_KEY        BillingKeyTextSearchField = "BILLING_KEY"
	BillingKeyTextSearchFieldCHANNEL_GROUP_NAME BillingKeyTextSearchField = "CHANNEL_GROUP_NAME"
)

// BillingKeyTextSearch 빌링키 텍스트 검색
type BillingKeyTextSearch struct {
	// 검색 필드
	Field BillingKeyTextSearchField `json:"field"`
	// 검색어
	Value string `json:"value"`
}

// BillingKeyTimeRangeField 빌링키 시간 범위 필드
type BillingKeyTimeRangeField string

const (
	BillingKeyTimeRangeFieldREQUESTED_AT      BillingKeyTimeRangeField = "REQUESTED_AT"
	BillingKeyTimeRangeFieldISSUED_AT         BillingKeyTimeRangeField = "ISSUED_AT"
	BillingKeyTimeRangeFieldDELETED_AT        BillingKeyTimeRangeField = "DELETED_AT"
	BillingKeyTimeRangeFieldSTATUS_CHANGED_AT BillingKeyTimeRangeField = "STATUS_CHANGED_AT"
)

// BillingKeyFilterInput 빌링키 필터 입력
type BillingKeyFilterInput struct {
	// 상점 아이디 목록
	StoreIds []string `json:"storeIds,omitempty"`
	// 시간 범위 필드
	TimeRangeField *BillingKeyTimeRangeField `json:"timeRangeField,omitempty"`
	// 시작 시간 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 종료 시간 (RFC 3339)
	Until *string `json:"until,omitempty"`
	// 빌링키 상태 목록
	Statuses []BillingKeyStatus `json:"statuses,omitempty"`
	// 채널 그룹 아이디 목록
	ChannelGroupIds []string `json:"channelGroupIds,omitempty"`
	// 텍스트 검색
	TextSearch *BillingKeyTextSearch `json:"textSearch,omitempty"`
	// PG사 목록
	PgProviders []common.PgProvider `json:"pgProviders,omitempty"`
	// PG사 상점 아이디 목록
	PgMerchantIds []string `json:"pgMerchantIds,omitempty"`
	// 포트원 버전
	Version *common.PortOneVersion `json:"version,omitempty"`
}

// GetBillingKeyInfosBody 빌링키 목록 조회 요청
type GetBillingKeyInfosBody struct {
	// 페이지 정보
	Page *common.PageInput `json:"page,omitempty"`
	// 정렬 정보
	Sort *BillingKeySortInput `json:"sort,omitempty"`
	// 필터 정보
	Filter *BillingKeyFilterInput `json:"filter,omitempty"`
}

// GetBillingKeyInfosResponse 빌링키 목록 조회 응답
type GetBillingKeyInfosResponse struct {
	// 빌링키 목록
	Items []BillingKeyInfo `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// DeleteBillingKeyResponse 빌링키 삭제 응답
type DeleteBillingKeyResponse struct {
	// 삭제 시점 (RFC 3339)
	DeletedAt string `json:"deletedAt"`
}

// IssueBillingKeyBody 빌링키 발급 요청
type IssueBillingKeyBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 빌링키 결제수단 정보
	Method InstantBillingKeyPaymentMethodInput `json:"method"`
	// 채널 키
	ChannelKey *string `json:"channelKey,omitempty"`
	// 채널 그룹 아이디
	ChannelGroupId *string `json:"channelGroupId,omitempty"`
	// 고객 정보
	Customer *common.CustomerInput `json:"customer,omitempty"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 빌링키 발급 건 이름
	IssueName *string `json:"issueName,omitempty"`
	// PG사 추가 파라미터 (bypass)
	Bypass interface{} `json:"bypass,omitempty"`
}

// InstantBillingKeyPaymentMethodInput 빌링키 즉시 발급 결제수단 입력
type InstantBillingKeyPaymentMethodInput struct {
	// 카드 정보
	Card *InstantBillingKeyPaymentMethodInputCard `json:"card,omitempty"`
}

// InstantBillingKeyPaymentMethodInputCard 빌링키 카드 발급 입력
type InstantBillingKeyPaymentMethodInputCard struct {
	// 카드 인증 정보
	Credential common.CardCredential `json:"credential"`
}

// IssueBillingKeyResponse 빌링키 발급 응답
type IssueBillingKeyResponse struct {
	// 빌링키 정보
	BillingKeyInfo BillingKeyInfo `json:"billingKeyInfo"`
}

// BillingKeyDeleteRequester 빌링키 삭제 요청자
type BillingKeyDeleteRequester string

const (
	BillingKeyDeleteRequesterCUSTOMER BillingKeyDeleteRequester = "CUSTOMER"
	BillingKeyDeleteRequesterADMIN    BillingKeyDeleteRequester = "ADMIN"
)

// ConfirmedBillingKeySummary 빌링키 발급 수동 승인 응답
type ConfirmedBillingKeySummary struct {
	// 빌링키
	BillingKey string `json:"billingKey"`
}

// ConfirmedBillingKeyIssueAndPaySummary 빌링키 발급 및 초회 결제 수동 승인 응답
type ConfirmedBillingKeyIssueAndPaySummary struct {
	// 빌링키
	BillingKey string `json:"billingKey"`
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
}

// ConfirmBillingKeyOptions 빌링키 발급 수동 승인 옵션
type ConfirmBillingKeyOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 빌링키 발급 토큰
	BillingIssueToken string `json:"billingIssueToken"`
	// 테스트 결제 여부
	IsTest *bool `json:"isTest,omitempty"`
}

// ConfirmBillingKeyIssueAndPayOptions 빌링키 발급 및 초회 결제 수동 승인 옵션
type ConfirmBillingKeyIssueAndPayOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 빌링키 발급 토큰
	BillingIssueToken string `json:"billingIssueToken"`
	// 결제 건 아이디
	PaymentId *string `json:"paymentId,omitempty"`
	// 통화
	Currency *common.Currency `json:"currency,omitempty"`
	// 결제 금액
	TotalAmount *int64 `json:"totalAmount,omitempty"`
	// 면세 금액
	TaxFreeAmount *int64 `json:"taxFreeAmount,omitempty"`
	// 테스트 결제 여부
	IsTest *bool `json:"isTest,omitempty"`
}
