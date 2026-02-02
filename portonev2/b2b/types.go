package b2b

import (
	"github.com/gliese436/portone-go-client/portonev2/common"
)

// B2bTaxInvoiceStatus 세금계산서 상태
type B2bTaxInvoiceStatus string

const (
	B2bTaxInvoiceStatusDRAFT       B2bTaxInvoiceStatus = "DRAFT"
	B2bTaxInvoiceStatusREGISTERED  B2bTaxInvoiceStatus = "REGISTERED"
	B2bTaxInvoiceStatusISSUED      B2bTaxInvoiceStatus = "ISSUED"
	B2bTaxInvoiceStatusSENT        B2bTaxInvoiceStatus = "SENT"
	B2bTaxInvoiceStatusAPPROVED    B2bTaxInvoiceStatus = "APPROVED"
	B2bTaxInvoiceStatusCANCELLED   B2bTaxInvoiceStatus = "CANCELLED"
	B2bTaxInvoiceStatusREFUSED     B2bTaxInvoiceStatus = "REFUSED"
	B2bTaxInvoiceStatusISSUE_FAILED B2bTaxInvoiceStatus = "ISSUE_FAILED"
)

// B2bTaxInvoiceType 세금계산서 유형
type B2bTaxInvoiceType string

const (
	B2bTaxInvoiceTypeTAX           B2bTaxInvoiceType = "TAX"
	B2bTaxInvoiceTypeTAX_FREE      B2bTaxInvoiceType = "TAX_FREE"
	B2bTaxInvoiceTypeZERO_RATE     B2bTaxInvoiceType = "ZERO_RATE"
)

// B2bTaxInvoicePurposeType 세금계산서 용도
type B2bTaxInvoicePurposeType string

const (
	B2bTaxInvoicePurposeTypeFORWARD  B2bTaxInvoicePurposeType = "FORWARD"
	B2bTaxInvoicePurposeTypeREVERSE  B2bTaxInvoicePurposeType = "REVERSE"
)

// B2bCompanyContact 사업자 연락처
type B2bCompanyContact struct {
	// 담당자명
	Name *string `json:"name,omitempty"`
	// 담당자 이메일
	Email *string `json:"email,omitempty"`
	// 담당자 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// B2bCompanyInfo 사업자 정보
type B2bCompanyInfo struct {
	// 사업자등록번호
	BrNo string `json:"brNo"`
	// 상호명
	Name *string `json:"name,omitempty"`
	// 대표자명
	RepresentativeName *string `json:"representativeName,omitempty"`
	// 사업장 주소
	Address *string `json:"address,omitempty"`
	// 업태
	BusinessType *string `json:"businessType,omitempty"`
	// 종목
	BusinessCategory *string `json:"businessCategory,omitempty"`
	// 연락처
	Contact *B2bCompanyContact `json:"contact,omitempty"`
}

// B2bTaxInvoiceItem 세금계산서 품목
type B2bTaxInvoiceItem struct {
	// 품목명
	Name string `json:"name"`
	// 규격
	Specification *string `json:"specification,omitempty"`
	// 수량
	Quantity *int64 `json:"quantity,omitempty"`
	// 단가
	UnitPrice *int64 `json:"unitPrice,omitempty"`
	// 공급가액
	SupplyPrice int64 `json:"supplyPrice"`
	// 세액
	Tax int64 `json:"tax"`
	// 비고
	Remark *string `json:"remark,omitempty"`
	// 일자 (yyyyMMdd)
	Date *string `json:"date,omitempty"`
}

// B2bTaxInvoice 세금계산서
type B2bTaxInvoice struct {
	// 세금계산서 상태
	Status B2bTaxInvoiceStatus `json:"status"`
	// 세금계산서 아이디
	ID string `json:"id"`
	// 세금계산서 유형
	Type B2bTaxInvoiceType `json:"type"`
	// 세금계산서 용도
	PurposeType B2bTaxInvoicePurposeType `json:"purposeType"`
	// 공급자 정보
	Supplier B2bCompanyInfo `json:"supplier"`
	// 공급받는자 정보
	Recipient B2bCompanyInfo `json:"recipient"`
	// 품목 목록
	Items []B2bTaxInvoiceItem `json:"items"`
	// 총 공급가액
	TotalSupplyPrice int64 `json:"totalSupplyPrice"`
	// 총 세액
	TotalTax int64 `json:"totalTax"`
	// 총 금액
	TotalAmount int64 `json:"totalAmount"`
	// 비고
	Remark *string `json:"remark,omitempty"`
	// 작성 일자 (yyyyMMdd)
	WriteDate *string `json:"writeDate,omitempty"`
	// 발급 일시 (RFC 3339)
	IssuedAt *string `json:"issuedAt,omitempty"`
	// 생성 일시 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 수정 일시 (RFC 3339)
	ModifiedAt string `json:"modifiedAt"`
	// 국세청 승인번호
	NtsApprovalNo *string `json:"ntsApprovalNo,omitempty"`
}

// B2bTaxInvoiceSummary 세금계산서 요약
type B2bTaxInvoiceSummary struct {
	// 세금계산서 아이디
	ID string `json:"id"`
	// 세금계산서 상태
	Status B2bTaxInvoiceStatus `json:"status"`
	// 총 금액
	TotalAmount int64 `json:"totalAmount"`
}

// GetB2bTaxInvoiceResponse 세금계산서 조회 응답
type GetB2bTaxInvoiceResponse = B2bTaxInvoice

// CreateB2bTaxInvoiceBody 세금계산서 생성 요청
type CreateB2bTaxInvoiceBody struct {
	// 세금계산서 유형
	Type B2bTaxInvoiceType `json:"type"`
	// 세금계산서 용도
	PurposeType B2bTaxInvoicePurposeType `json:"purposeType"`
	// 공급자 정보
	Supplier B2bCompanyInfo `json:"supplier"`
	// 공급받는자 정보
	Recipient B2bCompanyInfo `json:"recipient"`
	// 품목 목록
	Items []B2bTaxInvoiceItem `json:"items"`
	// 비고
	Remark *string `json:"remark,omitempty"`
	// 작성 일자 (yyyyMMdd)
	WriteDate *string `json:"writeDate,omitempty"`
}

// CreateB2bTaxInvoiceResponse 세금계산서 생성 응답
type CreateB2bTaxInvoiceResponse struct {
	// 세금계산서 아이디
	TaxInvoiceId string `json:"taxInvoiceId"`
}

// IssueB2bTaxInvoiceBody 세금계산서 발급 요청
type IssueB2bTaxInvoiceBody struct {
	// 세금계산서 아이디
	TaxInvoiceId string `json:"taxInvoiceId"`
	// 메모
	Memo *string `json:"memo,omitempty"`
	// 이메일 발송 여부
	SendEmail *bool `json:"sendEmail,omitempty"`
}

// IssueB2bTaxInvoiceResponse 세금계산서 발급 응답
type IssueB2bTaxInvoiceResponse struct{}

// CancelB2bTaxInvoiceIssuanceBody 세금계산서 발급 취소 요청
type CancelB2bTaxInvoiceIssuanceBody struct {
	// 세금계산서 아이디
	TaxInvoiceId string `json:"taxInvoiceId"`
	// 메모
	Memo *string `json:"memo,omitempty"`
}

// CancelB2bTaxInvoiceIssuanceResponse 세금계산서 발급 취소 응답
type CancelB2bTaxInvoiceIssuanceResponse struct{}

// B2bCompanyStateStatus 사업자 상태
type B2bCompanyStateStatus string

const (
	B2bCompanyStateStatusACTIVE          B2bCompanyStateStatus = "ACTIVE"
	B2bCompanyStateStatusSUSPENDED       B2bCompanyStateStatus = "SUSPENDED"
	B2bCompanyStateStatusCLOSED          B2bCompanyStateStatus = "CLOSED"
	B2bCompanyStateStatusNOT_REGISTERED  B2bCompanyStateStatus = "NOT_REGISTERED"
)

// B2bCompanyState 사업자 상태 정보
type B2bCompanyState struct {
	// 사업자등록번호
	BrNo string `json:"brNo"`
	// 사업자 상태
	Status B2bCompanyStateStatus `json:"status"`
	// 상태 확인 일시 (RFC 3339)
	CheckedAt string `json:"checkedAt"`
	// 폐업일 (yyyyMMdd)
	ClosedAt *string `json:"closedAt,omitempty"`
}

// GetB2bCompanyStateBody 사업자 상태 조회 요청
type GetB2bCompanyStateBody struct {
	// 사업자등록번호 목록
	BrNos []string `json:"brNos"`
}

// GetB2bCompanyStateResponse 사업자 상태 조회 응답
type GetB2bCompanyStateResponse struct {
	// 사업자 상태 목록
	Items []B2bCompanyState `json:"items"`
}

// B2bCertificate 공인인증서 정보
type B2bCertificate struct {
	// 인증서 DN
	Dn string `json:"dn"`
	// 유효 시작 일시 (RFC 3339)
	ValidFrom string `json:"validFrom"`
	// 유효 종료 일시 (RFC 3339)
	ValidTo string `json:"validTo"`
	// 발급자
	Issuer string `json:"issuer"`
	// 인증서 유형
	Type *string `json:"type,omitempty"`
}

// B2bMemberCompany 회원사 정보
type B2bMemberCompany struct {
	// 사업자등록번호
	BrNo string `json:"brNo"`
	// 상호명
	Name string `json:"name"`
	// 대표자명
	RepresentativeName string `json:"representativeName"`
	// 통화
	Currency common.Currency `json:"currency"`
	// 등록 일시 (RFC 3339)
	RegisteredAt string `json:"registeredAt"`
	// 공인인증서 정보
	Certificate *B2bCertificate `json:"certificate,omitempty"`
}

// RegisterB2bMemberCompanyBody 회원사 등록 요청
type RegisterB2bMemberCompanyBody struct {
	// 사업자등록번호
	BrNo string `json:"brNo"`
	// 상호명
	Name string `json:"name"`
	// 대표자명
	RepresentativeName string `json:"representativeName"`
	// 사업장 주소
	Address *string `json:"address,omitempty"`
	// 업태
	BusinessType *string `json:"businessType,omitempty"`
	// 종목
	BusinessCategory *string `json:"businessCategory,omitempty"`
	// 담당자 정보
	Contact *B2bCompanyContact `json:"contact,omitempty"`
}

// RegisterB2bMemberCompanyResponse 회원사 등록 응답
type RegisterB2bMemberCompanyResponse struct {
	// 회원사 정보
	Company B2bMemberCompany `json:"company"`
}

// ---- Additional B2B Tax Invoice Types ----

// B2bTaxInvoiceKeyType 세금계산서 문서 번호 유형
type B2bTaxInvoiceKeyType string

const (
	B2bTaxInvoiceKeyTypeTAX_INVOICE_ID B2bTaxInvoiceKeyType = "TAX_INVOICE_ID"
	B2bTaxInvoiceKeyTypeSUPPLIER       B2bTaxInvoiceKeyType = "SUPPLIER"
	B2bTaxInvoiceKeyTypeRECIPIENT      B2bTaxInvoiceKeyType = "RECIPIENT"
)

// B2bBulkTaxInvoiceStatus 일괄 세금계산서 상태
type B2bBulkTaxInvoiceStatus string

const (
	B2bBulkTaxInvoiceStatusPENDING    B2bBulkTaxInvoiceStatus = "PENDING"
	B2bBulkTaxInvoiceStatusPROCESSING B2bBulkTaxInvoiceStatus = "PROCESSING"
	B2bBulkTaxInvoiceStatusCOMPLETED  B2bBulkTaxInvoiceStatus = "COMPLETED"
	B2bBulkTaxInvoiceStatusFAILED     B2bBulkTaxInvoiceStatus = "FAILED"
)

// B2bBulkTaxInvoice 일괄 세금계산서
type B2bBulkTaxInvoice struct {
	// 일괄 세금계산서 아이디
	ID string `json:"id"`
	// 상태
	Status B2bBulkTaxInvoiceStatus `json:"status"`
	// 전체 건수
	TotalCount int32 `json:"totalCount"`
	// 성공 건수
	SuccessCount int32 `json:"successCount"`
	// 실패 건수
	FailedCount int32 `json:"failedCount"`
	// 생성 일시 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 수정 일시 (RFC 3339)
	ModifiedAt string `json:"modifiedAt"`
}

// CreateB2bFileUploadUrlPayload 파일 업로드 URL 생성 응답
type CreateB2bFileUploadUrlPayload struct {
	// 파일 아이디
	FileId string `json:"fileId"`
	// 업로드 URL
	UploadUrl string `json:"uploadUrl"`
}

// B2bTaxInvoiceInput 세금계산서 생성 요청 정보
type B2bTaxInvoiceInput struct {
	// 세금계산서 유형
	Type B2bTaxInvoiceType `json:"type"`
	// 세금계산서 용도
	PurposeType B2bTaxInvoicePurposeType `json:"purposeType"`
	// 공급자 정보
	Supplier B2bCompanyInfo `json:"supplier"`
	// 공급받는자 정보
	Recipient B2bCompanyInfo `json:"recipient"`
	// 품목 목록
	Items []B2bTaxInvoiceItem `json:"items"`
	// 비고
	Remark *string `json:"remark,omitempty"`
	// 작성 일자 (yyyyMMdd)
	WriteDate *string `json:"writeDate,omitempty"`
	// 공급자 문서번호
	SupplierDocumentKey *string `json:"supplierDocumentKey,omitempty"`
	// 공급받는자 문서번호
	RecipientDocumentKey *string `json:"recipientDocumentKey,omitempty"`
}

// B2bTaxInvoiceModificationCreateBody 수정 세금계산서 입력 정보
type B2bTaxInvoiceModificationCreateBody struct {
	// 원본 세금계산서 아이디
	OriginalTaxInvoiceId *string `json:"originalTaxInvoiceId,omitempty"`
	// 수정 사유 코드
	ModificationCode *string `json:"modificationCode,omitempty"`
}

// DraftB2bTaxInvoiceBody 세금계산서 임시저장 요청
type DraftB2bTaxInvoiceBody struct {
	// 세금계산서 생성 요청 정보
	TaxInvoice B2bTaxInvoiceInput `json:"taxInvoice"`
	// 수정 세금계산서 입력 정보
	Modification *B2bTaxInvoiceModificationCreateBody `json:"modification,omitempty"`
	// 메모
	Memo *string `json:"memo,omitempty"`
}

// DraftB2bTaxInvoiceResponse 세금계산서 임시저장 응답
type DraftB2bTaxInvoiceResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// UpdateB2bTaxInvoiceDraftBody 세금계산서 임시저장 수정 요청
type UpdateB2bTaxInvoiceDraftBody struct {
	// 사업자등록번호
	Brn *string `json:"brn,omitempty"`
	// 세금계산서 문서 번호
	TaxInvoiceKey string `json:"taxInvoiceKey"`
	// 문서 번호 유형
	TaxInvoiceKeyType *B2bTaxInvoiceKeyType `json:"taxInvoiceKeyType,omitempty"`
	// 세금계산서 생성 요청 정보
	TaxInvoice B2bTaxInvoiceInput `json:"taxInvoice"`
	// 메모
	Memo *string `json:"memo,omitempty"`
}

// UpdateB2bTaxInvoiceDraftResponse 세금계산서 임시저장 수정 응답
type UpdateB2bTaxInvoiceDraftResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// IssueB2bTaxInvoiceImmediatelyBody 세금계산서 즉시 정발행 요청
type IssueB2bTaxInvoiceImmediatelyBody struct {
	// 세금계산서 생성 요청 정보
	TaxInvoice B2bTaxInvoiceInput `json:"taxInvoice"`
	// 메모
	Memo *string `json:"memo,omitempty"`
	// 수정 세금계산서 입력 정보
	Modification *B2bTaxInvoiceModificationCreateBody `json:"modification,omitempty"`
}

// IssueB2bTaxInvoiceImmediatelyResponse 세금계산서 즉시 정발행 응답
type IssueB2bTaxInvoiceImmediatelyResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// RequestB2bTaxInvoiceReverseIssuanceBody 세금계산서 역발행 즉시 요청
type RequestB2bTaxInvoiceReverseIssuanceBody struct {
	// 세금계산서 생성 요청 정보
	TaxInvoice B2bTaxInvoiceInput `json:"taxInvoice"`
	// 메모
	Memo *string `json:"memo,omitempty"`
	// 수정 세금계산서 입력 정보
	Modification *B2bTaxInvoiceModificationCreateBody `json:"modification,omitempty"`
}

// RequestB2bTaxInvoiceReverseIssuanceResponse 세금계산서 역발행 즉시 요청 응답
type RequestB2bTaxInvoiceReverseIssuanceResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// B2bTaxInvoiceAttachment 세금계산서 첨부파일
type B2bTaxInvoiceAttachment struct {
	// 첨부파일 아이디
	ID string `json:"id"`
	// 파일명
	FileName string `json:"fileName"`
	// 파일 URL
	Url *string `json:"url,omitempty"`
}

// GetB2bTaxInvoiceAttachmentsResponse 세금계산서 첨부파일 목록 조회 응답
type GetB2bTaxInvoiceAttachmentsResponse struct {
	// 첨부파일 목록
	Attachments []B2bTaxInvoiceAttachment `json:"attachments"`
}

// CancelB2bTaxInvoiceRequestResponse 세금계산서 역발행 요청 취소 응답
type CancelB2bTaxInvoiceRequestResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// GetB2bTaxInvoicePdfDownloadUrlResponse PDF 다운로드 URL 조회 응답
type GetB2bTaxInvoicePdfDownloadUrlResponse struct {
	// 다운로드 URL
	Url string `json:"url"`
}

// GetB2bTaxInvoicePopupUrlResponse 팝업 URL 조회 응답
type GetB2bTaxInvoicePopupUrlResponse struct {
	// 팝업 URL
	Url string `json:"url"`
}

// GetB2bTaxInvoicePrintUrlResponse 프린트 URL 조회 응답
type GetB2bTaxInvoicePrintUrlResponse struct {
	// 프린트 URL
	Url string `json:"url"`
}

// RefuseB2bTaxInvoiceRequestResponse 세금계산서 역발행 요청 거부 응답
type RefuseB2bTaxInvoiceRequestResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// RequestB2bTaxInvoiceResponse 세금계산서 역발행 요청 응답
type RequestB2bTaxInvoiceResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// SendToNtsB2bTaxInvoiceResponse 세금계산서 국세청 즉시 전송 응답
type SendToNtsB2bTaxInvoiceResponse struct {
	// 세금계산서
	TaxInvoice B2bTaxInvoice `json:"taxInvoice"`
}

// DeleteB2bTaxInvoiceResponse 세금계산서 삭제 응답
type DeleteB2bTaxInvoiceResponse struct{}

// GetB2bTaxInvoicesBodyFilter 세금계산서 목록 조회 필터
type GetB2bTaxInvoicesBodyFilter struct {
	// 세금계산서 상태 목록
	Statuses []B2bTaxInvoiceStatus `json:"statuses,omitempty"`
	// 세금계산서 유형 목록
	Types []B2bTaxInvoiceType `json:"types,omitempty"`
	// 세금계산서 용도 목록
	PurposeTypes []B2bTaxInvoicePurposeType `json:"purposeTypes,omitempty"`
	// 조회 시작 일시 (RFC 3339)
	From *string `json:"from,omitempty"`
	// 조회 종료 일시 (RFC 3339)
	Until *string `json:"until,omitempty"`
}

// GetB2bTaxInvoicesBody 세금계산서 목록 조회 요청
type GetB2bTaxInvoicesBody struct {
	// 테스트 모드 여부
	Test *bool `json:"test,omitempty"`
	// 페이지 번호
	PageNumber *int32 `json:"pageNumber,omitempty"`
	// 페이지 크기
	PageSize *int32 `json:"pageSize,omitempty"`
	// 필터
	Filter *GetB2bTaxInvoicesBodyFilter `json:"filter,omitempty"`
}

// GetB2bTaxInvoicesResponse 세금계산서 목록 조회 응답
type GetB2bTaxInvoicesResponse struct {
	// 세금계산서 목록
	Items []B2bTaxInvoice `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}

// TaxInvoicesSheetField 다운로드 시트 필드
type TaxInvoicesSheetField string

const (
	TaxInvoicesSheetFieldTAX_INVOICE_ID  TaxInvoicesSheetField = "TAX_INVOICE_ID"
	TaxInvoicesSheetFieldSTATUS          TaxInvoicesSheetField = "STATUS"
	TaxInvoicesSheetFieldTYPE            TaxInvoicesSheetField = "TYPE"
	TaxInvoicesSheetFieldPURPOSE_TYPE    TaxInvoicesSheetField = "PURPOSE_TYPE"
	TaxInvoicesSheetFieldSUPPLIER_BRN    TaxInvoicesSheetField = "SUPPLIER_BRN"
	TaxInvoicesSheetFieldRECIPIENT_BRN   TaxInvoicesSheetField = "RECIPIENT_BRN"
	TaxInvoicesSheetFieldTOTAL_AMOUNT    TaxInvoicesSheetField = "TOTAL_AMOUNT"
	TaxInvoicesSheetFieldWRITE_DATE      TaxInvoicesSheetField = "WRITE_DATE"
	TaxInvoicesSheetFieldISSUED_AT       TaxInvoicesSheetField = "ISSUED_AT"
	TaxInvoicesSheetFieldNTS_APPROVAL_NO TaxInvoicesSheetField = "NTS_APPROVAL_NO"
)

// ---- B2B Business Info Types ----

// B2bBusinessInfo 사업자등록 정보
type B2bBusinessInfo struct {
	// 사업자등록번호
	Brn string `json:"brn"`
	// 상호
	Name string `json:"name"`
	// 대표자명
	CeoName string `json:"ceoName"`
	// 우편번호
	ZipCode string `json:"zipCode"`
	// 주소
	Address string `json:"address"`
	// 사업자 유형
	BusinessEntityType string `json:"businessEntityType"`
	// 사업 상태
	BusinessStatus string `json:"businessStatus"`
	// 과세 유형
	TaxationType string `json:"taxationType"`
	// 간이과세-일반과세 전환일
	SimplifiedTaxationTypeDate *string `json:"simplifiedTaxationTypeDate,omitempty"`
	// 폐업일
	ClosingDate *string `json:"closingDate,omitempty"`
	// 개업일
	OpeningDate string `json:"openingDate"`
	// 업태
	BusinessType string `json:"businessType"`
	// 종목
	BusinessClass string `json:"businessClass"`
	// 업종코드
	BusinessCategoryCode string `json:"businessCategoryCode"`
	// 법인등록번호
	CorpRegNo *string `json:"corpRegNo,omitempty"`
	// 전화번호
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	// 관할세무서코드
	TaxOfficeCode *string `json:"taxOfficeCode,omitempty"`
	// 관할세무서명
	TaxOfficeName *string `json:"taxOfficeName,omitempty"`
}

// B2bBusinessInfoResult 사업자등록 정보조회 결과
type B2bBusinessInfoResult struct {
	// 사업자등록번호
	Brn string `json:"brn"`
	// 사업자등록 정보
	BusinessInfo *B2bBusinessInfo `json:"businessInfo,omitempty"`
	// 조회 실패 시 에러 메시지
	Error *string `json:"error,omitempty"`
}

// GetB2bBusinessInfosBody 사업자등록 정보 조회 요청
type GetB2bBusinessInfosBody struct {
	// 조회할 사업자등록번호 리스트
	BrnList []string `json:"brnList"`
}

// GetB2bBusinessInfosResponse 사업자등록 정보 조회 응답
type GetB2bBusinessInfosResponse struct {
	// 사업자등록 정보 리스트
	Result []B2bBusinessInfoResult `json:"result"`
}
