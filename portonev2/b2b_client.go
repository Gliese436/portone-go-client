package portonev2

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/Gliese436/portone-go-client/portonev2/b2b"
)

// B2BClient B2B API 클라이언트
type B2BClient struct {
	client *Client
}

// NewB2BClient 새 B2B 클라이언트를 생성합니다.
func NewB2BClient(client *Client) *B2BClient {
	return &B2BClient{client: client}
}

// GetB2bTaxInvoice 세금계산서를 조회합니다.
func (c *B2BClient) GetB2bTaxInvoice(ctx context.Context, taxInvoiceID string) (*b2b.B2bTaxInvoice, error) {
	var result b2b.B2bTaxInvoice
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateB2bTaxInvoice 세금계산서를 생성합니다.
func (c *B2BClient) CreateB2bTaxInvoice(ctx context.Context, opts *b2b.CreateB2bTaxInvoiceBody) (*b2b.CreateB2bTaxInvoiceResponse, error) {
	var result b2b.CreateB2bTaxInvoiceResponse
	if err := c.client.Post(ctx, "/b2b/tax-invoices", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IssueB2bTaxInvoice 세금계산서를 발급합니다.
func (c *B2BClient) IssueB2bTaxInvoice(ctx context.Context, opts *b2b.IssueB2bTaxInvoiceBody) error {
	return c.client.Post(ctx, "/b2b/tax-invoices/"+url.PathEscape(opts.TaxInvoiceId)+"/issue", opts, nil)
}

// CancelB2bTaxInvoiceIssuance 세금계산서 발급을 취소합니다.
func (c *B2BClient) CancelB2bTaxInvoiceIssuance(ctx context.Context, opts *b2b.CancelB2bTaxInvoiceIssuanceBody) error {
	return c.client.Post(ctx, "/b2b/tax-invoices/"+url.PathEscape(opts.TaxInvoiceId)+"/cancel-issuance", opts, nil)
}

// GetB2bCompanyState 사업자 상태를 조회합니다.
func (c *B2BClient) GetB2bCompanyState(ctx context.Context, opts *b2b.GetB2bCompanyStateBody) (*b2b.GetB2bCompanyStateResponse, error) {
	var result b2b.GetB2bCompanyStateResponse
	if err := c.client.Post(ctx, "/b2b/company-state", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RegisterB2bMemberCompany 회원사를 등록합니다.
func (c *B2BClient) RegisterB2bMemberCompany(ctx context.Context, opts *b2b.RegisterB2bMemberCompanyBody) (*b2b.RegisterB2bMemberCompanyResponse, error) {
	var result b2b.RegisterB2bMemberCompanyResponse
	if err := c.client.Post(ctx, "/b2b/member-companies", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bBulkTaxInvoice 일괄 세금계산서를 조회합니다.
func (c *B2BClient) GetB2bBulkTaxInvoice(ctx context.Context, bulkTaxInvoiceID string, test *bool) (*b2b.B2bBulkTaxInvoice, error) {
	query := url.Values{}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.B2bBulkTaxInvoice
	if err := c.client.Get(ctx, "/b2b/bulk-tax-invoices/"+url.PathEscape(bulkTaxInvoiceID), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreateB2bFileUploadUrl 파일 업로드 URL을 생성합니다.
func (c *B2BClient) CreateB2bFileUploadUrl(ctx context.Context, test *bool, fileName string) (*b2b.CreateB2bFileUploadUrlPayload, error) {
	body := map[string]interface{}{
		"fileName": fileName,
	}
	path := "/b2b/file-upload-url"
	if test != nil && *test {
		path += "?test=true"
	}
	var result b2b.CreateB2bFileUploadUrlPayload
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DraftB2bTaxInvoice 세금계산서를 임시 저장합니다.
func (c *B2BClient) DraftB2bTaxInvoice(ctx context.Context, test *bool, opts *b2b.DraftB2bTaxInvoiceBody) (*b2b.DraftB2bTaxInvoiceResponse, error) {
	path := "/b2b/tax-invoices/draft"
	if test != nil && *test {
		path += "?test=true"
	}
	var result b2b.DraftB2bTaxInvoiceResponse
	if err := c.client.Post(ctx, path, opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdateB2bTaxInvoiceDraft 임시 저장된 세금계산서를 수정합니다.
func (c *B2BClient) UpdateB2bTaxInvoiceDraft(ctx context.Context, test *bool, opts *b2b.UpdateB2bTaxInvoiceDraftBody) (*b2b.UpdateB2bTaxInvoiceDraftResponse, error) {
	path := "/b2b/tax-invoices/draft"
	if test != nil && *test {
		path += "?test=true"
	}
	var result b2b.UpdateB2bTaxInvoiceDraftResponse
	if err := c.client.Put(ctx, path, opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IssueB2bTaxInvoiceImmediately 세금계산서를 즉시 정발행합니다.
func (c *B2BClient) IssueB2bTaxInvoiceImmediately(ctx context.Context, test *bool, opts *b2b.IssueB2bTaxInvoiceImmediatelyBody) (*b2b.IssueB2bTaxInvoiceImmediatelyResponse, error) {
	path := "/b2b/tax-invoices/issue-immediately"
	if test != nil && *test {
		path += "?test=true"
	}
	var result b2b.IssueB2bTaxInvoiceImmediatelyResponse
	if err := c.client.Post(ctx, path, opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestB2bTaxInvoiceReverseIssuance 세금계산서 역발행을 즉시 요청합니다.
func (c *B2BClient) RequestB2bTaxInvoiceReverseIssuance(ctx context.Context, test *bool, opts *b2b.RequestB2bTaxInvoiceReverseIssuanceBody) (*b2b.RequestB2bTaxInvoiceReverseIssuanceResponse, error) {
	path := "/b2b/tax-invoices/request-reverse-issuance"
	if test != nil && *test {
		path += "?test=true"
	}
	var result b2b.RequestB2bTaxInvoiceReverseIssuanceResponse
	if err := c.client.Post(ctx, path, opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bTaxInvoiceWithKey 세금계산서를 문서번호로 조회합니다.
func (c *B2BClient) GetB2bTaxInvoiceWithKey(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.B2bTaxInvoice, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.B2bTaxInvoice
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteB2bTaxInvoice 세금계산서를 삭제합니다.
func (c *B2BClient) DeleteB2bTaxInvoice(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.DeleteB2bTaxInvoiceResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.DeleteB2bTaxInvoiceResponse
	if err := c.client.Delete(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// AttachB2bTaxInvoiceFile 세금계산서에 파일을 첨부합니다.
func (c *B2BClient) AttachB2bTaxInvoiceFile(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool, fileID string) error {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	body := map[string]interface{}{
		"fileId": fileID,
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/attach-file"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	return c.client.Post(ctx, path, body, nil)
}

// GetB2bTaxInvoiceAttachments 세금계산서 첨부파일 목록을 조회합니다.
func (c *B2BClient) GetB2bTaxInvoiceAttachments(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.GetB2bTaxInvoiceAttachmentsResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.GetB2bTaxInvoiceAttachmentsResponse
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey)+"/attachments", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteB2bTaxInvoiceAttachment 세금계산서 첨부파일을 삭제합니다.
func (c *B2BClient) DeleteB2bTaxInvoiceAttachment(ctx context.Context, taxInvoiceKey string, attachmentID string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) error {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	return c.client.Delete(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey)+"/attachments/"+url.PathEscape(attachmentID), query, nil)
}

// IssueB2bTaxInvoiceWithKey 세금계산서 발급을 승인합니다.
func (c *B2BClient) IssueB2bTaxInvoiceWithKey(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool, memo *string, emailSubject *string) (*b2b.B2bTaxInvoice, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	body := map[string]interface{}{}
	if memo != nil {
		body["memo"] = *memo
	}
	if emailSubject != nil {
		body["emailSubject"] = *emailSubject
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/issue"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.B2bTaxInvoice
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelB2bTaxInvoiceIssuanceWithKey 세금계산서 발급을 취소합니다.
func (c *B2BClient) CancelB2bTaxInvoiceIssuanceWithKey(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool, memo *string) (*b2b.B2bTaxInvoice, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	body := map[string]interface{}{}
	if memo != nil {
		body["memo"] = *memo
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/cancel-issuance"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.B2bTaxInvoice
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelB2bTaxInvoiceRequest 세금계산서 역발행 요청을 취소합니다.
func (c *B2BClient) CancelB2bTaxInvoiceRequest(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool, memo *string) (*b2b.CancelB2bTaxInvoiceRequestResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	body := map[string]interface{}{}
	if memo != nil {
		body["memo"] = *memo
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/cancel-request"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.CancelB2bTaxInvoiceRequestResponse
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bTaxInvoicePdfDownloadUrl 세금계산서 PDF 다운로드 URL을 조회합니다.
func (c *B2BClient) GetB2bTaxInvoicePdfDownloadUrl(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.GetB2bTaxInvoicePdfDownloadUrlResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.GetB2bTaxInvoicePdfDownloadUrlResponse
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey)+"/pdf-download-url", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bTaxInvoicePopupUrl 세금계산서 팝업 URL을 조회합니다.
func (c *B2BClient) GetB2bTaxInvoicePopupUrl(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, includeMenu *bool, test *bool) (*b2b.GetB2bTaxInvoicePopupUrlResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if includeMenu != nil {
		if *includeMenu {
			query.Set("includeMenu", "true")
		} else {
			query.Set("includeMenu", "false")
		}
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.GetB2bTaxInvoicePopupUrlResponse
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey)+"/popup-url", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bTaxInvoicePrintUrl 세금계산서 프린트 URL을 조회합니다.
func (c *B2BClient) GetB2bTaxInvoicePrintUrl(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.GetB2bTaxInvoicePrintUrlResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	var result b2b.GetB2bTaxInvoicePrintUrlResponse
	if err := c.client.Get(ctx, "/b2b/tax-invoices/"+url.PathEscape(taxInvoiceKey)+"/print-url", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RefuseB2bTaxInvoiceRequest 세금계산서 역발행 요청을 거부합니다.
func (c *B2BClient) RefuseB2bTaxInvoiceRequest(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool, memo *string) (*b2b.RefuseB2bTaxInvoiceRequestResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	body := map[string]interface{}{}
	if memo != nil {
		body["memo"] = *memo
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/refuse-request"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.RefuseB2bTaxInvoiceRequestResponse
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RequestB2bTaxInvoice 세금계산서 역발행을 요청합니다.
func (c *B2BClient) RequestB2bTaxInvoice(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.RequestB2bTaxInvoiceResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/request"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.RequestB2bTaxInvoiceResponse
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SendToNtsB2bTaxInvoice 세금계산서를 국세청에 즉시 전송합니다.
func (c *B2BClient) SendToNtsB2bTaxInvoice(ctx context.Context, taxInvoiceKey string, brn *string, taxInvoiceKeyType *b2b.B2bTaxInvoiceKeyType, test *bool) (*b2b.SendToNtsB2bTaxInvoiceResponse, error) {
	query := url.Values{}
	if brn != nil {
		query.Set("brn", *brn)
	}
	if taxInvoiceKeyType != nil {
		query.Set("taxInvoiceKeyType", string(*taxInvoiceKeyType))
	}
	if test != nil && *test {
		query.Set("test", "true")
	}
	path := "/b2b/tax-invoices/" + url.PathEscape(taxInvoiceKey) + "/send-to-nts"
	if len(query) > 0 {
		path += "?" + query.Encode()
	}
	var result b2b.SendToNtsB2bTaxInvoiceResponse
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetB2bTaxInvoices 세금계산서 목록을 조회합니다.
func (c *B2BClient) GetB2bTaxInvoices(ctx context.Context, opts *b2b.GetB2bTaxInvoicesBody) (*b2b.GetB2bTaxInvoicesResponse, error) {
	query := url.Values{}
	if opts != nil {
		jsonBytes, _ := json.Marshal(opts)
		query.Set("requestBody", string(jsonBytes))
	}
	var result b2b.GetB2bTaxInvoicesResponse
	if err := c.client.Get(ctx, "/b2b/tax-invoices", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DownloadB2bTaxInvoicesSheet 세금계산서 엑셀 파일(csv)을 다운로드합니다.
func (c *B2BClient) DownloadB2bTaxInvoicesSheet(ctx context.Context, test *bool, filter *b2b.GetB2bTaxInvoicesBodyFilter, fields []b2b.TaxInvoicesSheetField) (string, error) {
	requestBody := map[string]interface{}{}
	if test != nil {
		requestBody["test"] = *test
	}
	if filter != nil {
		requestBody["filter"] = filter
	}
	if len(fields) > 0 {
		requestBody["fields"] = fields
	}
	query := url.Values{}
	if len(requestBody) > 0 {
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}
	var result string
	if err := c.client.Get(ctx, "/b2b/tax-invoices-sheet", query, &result); err != nil {
		return "", err
	}
	return result, nil
}

// GetB2bBusinessInfos 사업자등록 정보를 조회합니다.
func (c *B2BClient) GetB2bBusinessInfos(ctx context.Context, brnList []string) (*b2b.GetB2bBusinessInfosResponse, error) {
	body := b2b.GetB2bBusinessInfosBody{
		BrnList: brnList,
	}
	var result b2b.GetB2bBusinessInfosResponse
	if err := c.client.Post(ctx, "/b2b/companies/business-info", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
