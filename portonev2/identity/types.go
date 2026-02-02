package identity

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// IdentityVerificationStatus 본인인증 상태
type IdentityVerificationStatus string

const (
	IdentityVerificationStatusREADY    IdentityVerificationStatus = "READY"
	IdentityVerificationStatusVERIFIED IdentityVerificationStatus = "VERIFIED"
	IdentityVerificationStatusFAILED   IdentityVerificationStatus = "FAILED"
)

// IdentityVerificationMethod 본인인증 수단
type IdentityVerificationMethod string

const (
	IdentityVerificationMethodSMS  IdentityVerificationMethod = "SMS"
	IdentityVerificationMethodAPP  IdentityVerificationMethod = "APP"
	IdentityVerificationMethodBANK IdentityVerificationMethod = "BANK"
)

// IdentityVerificationOperator 본인인증 사업자
type IdentityVerificationOperator string

const (
	IdentityVerificationOperatorSKT      IdentityVerificationOperator = "SKT"
	IdentityVerificationOperatorKT       IdentityVerificationOperator = "KT"
	IdentityVerificationOperatorLGU      IdentityVerificationOperator = "LGU"
	IdentityVerificationOperatorSKT_MVNO IdentityVerificationOperator = "SKT_MVNO"
	IdentityVerificationOperatorKT_MVNO  IdentityVerificationOperator = "KT_MVNO"
	IdentityVerificationOperatorLGU_MVNO IdentityVerificationOperator = "LGU_MVNO"
)

// IdentityVerificationFailure 본인인증 실패 정보
type IdentityVerificationFailure struct {
	// 실패 사유
	Reason *string `json:"reason,omitempty"`
	// PG사 실패 코드
	PgCode *string `json:"pgCode,omitempty"`
	// PG사 실패 메시지
	PgMessage *string `json:"pgMessage,omitempty"`
}

// VerifiedCustomer 본인인증된 고객 정보
type VerifiedCustomer struct {
	// 고객 아이디
	ID *string `json:"id,omitempty"`
	// 이름
	Name string `json:"name"`
	// 성별
	Gender *common.Gender `json:"gender,omitempty"`
	// 생년월일 (yyyy-MM-dd)
	BirthDate *string `json:"birthDate,omitempty"`
	// 국적
	IsForeigner *bool `json:"isForeigner,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// CI (Connecting Information)
	Ci *string `json:"ci,omitempty"`
	// DI (Duplication Information)
	Di *string `json:"di,omitempty"`
}

// IdentityVerification 본인인증 내역
type IdentityVerification struct {
	// 본인인증 상태
	Status IdentityVerificationStatus `json:"status"`
	// 본인인증 아이디
	ID string `json:"id"`
	// 채널
	Channel *common.SelectedChannel `json:"channel,omitempty"`
	// 본인인증된 고객 정보
	VerifiedCustomer *VerifiedCustomer `json:"verifiedCustomer,omitempty"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// 본인인증 요청 시점 (RFC 3339)
	RequestedAt string `json:"requestedAt"`
	// 업데이트 시점 (RFC 3339)
	UpdatedAt string `json:"updatedAt"`
	// 상태 업데이트 시점 (RFC 3339)
	StatusChangedAt string `json:"statusChangedAt"`
	// 본인인증 완료 시점 (RFC 3339)
	VerifiedAt *string `json:"verifiedAt,omitempty"`
	// PG사 거래 아이디
	PgTxId *string `json:"pgTxId,omitempty"`
	// PG사 거래 응답 본문
	PgRawResponse *string `json:"pgRawResponse,omitempty"`
	// 본인인증 실패 정보
	Failure *IdentityVerificationFailure `json:"failure,omitempty"`
}

// IsReady 준비 상태인지 확인
func (i *IdentityVerification) IsReady() bool {
	return i.Status == IdentityVerificationStatusREADY
}

// IsVerified 인증 완료 상태인지 확인
func (i *IdentityVerification) IsVerified() bool {
	return i.Status == IdentityVerificationStatusVERIFIED
}

// IsFailed 실패 상태인지 확인
func (i *IdentityVerification) IsFailed() bool {
	return i.Status == IdentityVerificationStatusFAILED
}

// SendIdentityVerificationBodyCustomer 본인인증 요청 고객 정보
type SendIdentityVerificationBodyCustomer struct {
	// 고객 아이디
	ID *string `json:"id,omitempty"`
	// 이름
	Name string `json:"name"`
	// 전화번호
	PhoneNumber string `json:"phoneNumber"`
	// 주민등록번호 앞 7자리
	IdentityNumber *string `json:"identityNumber,omitempty"`
	// 통신사
	Operator *IdentityVerificationOperator `json:"operator,omitempty"`
}

// SendIdentityVerificationBody 본인인증 요청
type SendIdentityVerificationBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 채널 키
	ChannelKey string `json:"channelKey"`
	// 고객 정보
	Customer SendIdentityVerificationBodyCustomer `json:"customer"`
	// 사용자 지정 데이터
	CustomData *string `json:"customData,omitempty"`
	// PG사 추가 파라미터
	Bypass interface{} `json:"bypass,omitempty"`
	// 본인인증 수단
	Method *IdentityVerificationMethod `json:"method,omitempty"`
}

// SendIdentityVerificationResponse 본인인증 요청 응답
type SendIdentityVerificationResponse struct{}

// ConfirmIdentityVerificationBody 본인인증 확인 요청
type ConfirmIdentityVerificationBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// OTP
	Otp *string `json:"otp,omitempty"`
}

// ConfirmIdentityVerificationResponse 본인인증 확인 응답
type ConfirmIdentityVerificationResponse struct {
	// 본인인증 내역
	IdentityVerification IdentityVerification `json:"identityVerification"`
}

// ResendIdentityVerificationBody 본인인증 재요청
type ResendIdentityVerificationBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
}

// ResendIdentityVerificationResponse 본인인증 재요청 응답
type ResendIdentityVerificationResponse struct{}

// GetIdentityVerificationResponse 본인인증 조회 응답
type GetIdentityVerificationResponse = IdentityVerification

// GetIdentityVerificationsResponse 본인인증 목록 조회 응답
type GetIdentityVerificationsResponse struct {
	// 본인인증 내역 목록
	Items []IdentityVerification `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// IdentityVerificationSortBy 본인인증 정렬 기준
type IdentityVerificationSortBy string

const (
	IdentityVerificationSortByREQUESTED_AT      IdentityVerificationSortBy = "REQUESTED_AT"
	IdentityVerificationSortByVERIFIED_AT       IdentityVerificationSortBy = "VERIFIED_AT"
	IdentityVerificationSortBySTATUS_CHANGED_AT IdentityVerificationSortBy = "STATUS_CHANGED_AT"
)

// IdentityVerificationSortInput 본인인증 정렬 입력
type IdentityVerificationSortInput struct {
	// 정렬 기준
	By *IdentityVerificationSortBy `json:"by,omitempty"`
	// 정렬 순서
	Order *common.SortOrder `json:"order,omitempty"`
}

// IdentityVerificationTimeRangeField 본인인증 시간 범위 필드
type IdentityVerificationTimeRangeField string

const (
	IdentityVerificationTimeRangeFieldREQUESTED_AT      IdentityVerificationTimeRangeField = "REQUESTED_AT"
	IdentityVerificationTimeRangeFieldVERIFIED_AT       IdentityVerificationTimeRangeField = "VERIFIED_AT"
	IdentityVerificationTimeRangeFieldSTATUS_CHANGED_AT IdentityVerificationTimeRangeField = "STATUS_CHANGED_AT"
)

// IdentityVerificationFilterInput 본인인증 필터 입력
type IdentityVerificationFilterInput struct {
	// 상점 아이디 목록
	StoreIds []string `json:"storeIds,omitempty"`
	// 시간 범위 필드
	TimeRangeField *IdentityVerificationTimeRangeField `json:"timeRangeField,omitempty"`
	// 시작 시간 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 종료 시간 (RFC 3339)
	Until *string `json:"until,omitempty"`
	// 본인인증 상태 목록
	Statuses []IdentityVerificationStatus `json:"statuses,omitempty"`
	// PG사 목록
	PgProviders []common.PgProvider `json:"pgProviders,omitempty"`
}

// GetIdentityVerificationsBody 본인인증 목록 조회 요청
type GetIdentityVerificationsBody struct {
	// 페이지 정보
	Page *common.PageInput `json:"page,omitempty"`
	// 정렬 정보
	Sort *IdentityVerificationSortInput `json:"sort,omitempty"`
	// 필터 정보
	Filter *IdentityVerificationFilterInput `json:"filter,omitempty"`
}
