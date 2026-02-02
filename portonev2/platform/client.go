package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// ClientInterface Platform API 클라이언트 인터페이스
type ClientInterface interface {
	Request(ctx context.Context, method, path string, query url.Values, body interface{}, result interface{}) error
	Get(ctx context.Context, path string, query url.Values, result interface{}) error
	Post(ctx context.Context, path string, body interface{}, result interface{}) error
	Delete(ctx context.Context, path string, query url.Values, result interface{}) error
	Patch(ctx context.Context, path string, body interface{}, result interface{}) error
	Put(ctx context.Context, path string, body interface{}, result interface{}) error
}

// Client Platform API 클라이언트
type Client struct {
	client ClientInterface
}

// NewClient 새 Platform 클라이언트를 생성합니다.
func NewClient(client ClientInterface) *Client {
	return &Client{client: client}
}

// buildQuery query parameter를 생성합니다.
func buildQuery(test *bool) url.Values {
	query := url.Values{}
	if test != nil && *test {
		query.Set("test", "true")
	}
	return query
}

// =============== Setting ===============

// GetPlatformSetting 플랫폼 설정을 조회합니다.
func (c *Client) GetPlatformSetting(ctx context.Context, test *bool) (*PlatformSetting, error) {
	var result PlatformSetting
	if err := c.client.Get(ctx, "/platform/setting", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePlatformSetting 플랫폼 설정을 수정합니다.
func (c *Client) UpdatePlatformSetting(ctx context.Context, test *bool, body *UpdatePlatformSettingBody) (*UpdatePlatformSettingResponse, error) {
	var result UpdatePlatformSettingResponse
	path := "/platform/setting"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Patch(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Partner ===============

// GetPlatformPartners 파트너 다건 조회
func (c *Client) GetPlatformPartners(ctx context.Context, test *bool, page *PageInput, filter *PlatformPartnerFilterInput) (*GetPlatformPartnersResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformPartnersResponse
	if err := c.client.Get(ctx, "/platform/partners", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformPartner 파트너 조회
func (c *Client) GetPlatformPartner(ctx context.Context, id string, test *bool) (*PlatformPartner, error) {
	var result PlatformPartner
	if err := c.client.Get(ctx, "/platform/partners/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformPartner 파트너 생성
func (c *Client) CreatePlatformPartner(ctx context.Context, test *bool, body *CreatePlatformPartnerBody) (*CreatePlatformPartnerResponse, error) {
	var result CreatePlatformPartnerResponse
	path := "/platform/partners"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformPartners 파트너 다건 생성
func (c *Client) CreatePlatformPartners(ctx context.Context, test *bool, partners []CreatePlatformPartnerBody) (*CreatePlatformPartnersResponse, error) {
	body := map[string]interface{}{
		"partners": partners,
	}
	var result CreatePlatformPartnersResponse
	path := "/platform/partners/batch"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePlatformPartner 파트너 수정
func (c *Client) UpdatePlatformPartner(ctx context.Context, id string, test *bool, body *UpdatePlatformPartnerBody) (*UpdatePlatformPartnerResponse, error) {
	var result UpdatePlatformPartnerResponse
	path := "/platform/partners/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Patch(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchivePlatformPartner 파트너 보관
func (c *Client) ArchivePlatformPartner(ctx context.Context, id string, test *bool) (*ArchivePlatformPartnerResponse, error) {
	var result ArchivePlatformPartnerResponse
	path := "/platform/partners/" + url.PathEscape(id) + "/archive"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecoverPlatformPartner 파트너 복원
func (c *Client) RecoverPlatformPartner(ctx context.Context, id string, test *bool) (*RecoverPlatformPartnerResponse, error) {
	var result RecoverPlatformPartnerResponse
	path := "/platform/partners/" + url.PathEscape(id) + "/recover"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConnectPartnerMemberCompany 파트너 국세청 연동
func (c *Client) ConnectPartnerMemberCompany(ctx context.Context, id string, test *bool) (*ConnectPartnerMemberCompanyResponse, error) {
	var result ConnectPartnerMemberCompanyResponse
	path := "/platform/partners/member-company-connect/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConnectBulkPartnerMemberCompany 파트너 일괄 국세청 연동
func (c *Client) ConnectBulkPartnerMemberCompany(ctx context.Context, test *bool, filter *PlatformPartnerFilterInput) (*ConnectBulkPartnerMemberCompanyResponse, error) {
	body := map[string]interface{}{}
	if filter != nil {
		body["filter"] = filter
	}
	var result ConnectBulkPartnerMemberCompanyResponse
	path := "/platform/partners/member-company-connect"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DisconnectPartnerMemberCompany 파트너 국세청 연동 해제
func (c *Client) DisconnectPartnerMemberCompany(ctx context.Context, id string, test *bool) (*DisconnectPartnerMemberCompanyResponse, error) {
	var result DisconnectPartnerMemberCompanyResponse
	path := "/platform/partners/member-company-disconnect/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DisconnectBulkPartnerMemberCompany 파트너 일괄 국세청 연동 해제
func (c *Client) DisconnectBulkPartnerMemberCompany(ctx context.Context, test *bool, filter *PlatformPartnerFilterInput) (*DisconnectBulkPartnerMemberCompanyResponse, error) {
	body := map[string]interface{}{}
	if filter != nil {
		body["filter"] = filter
	}
	var result DisconnectBulkPartnerMemberCompanyResponse
	path := "/platform/partners/member-company-disconnect"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Partner Schedule ===============

// GetPlatformPartnerSchedule 파트너 예약 업데이트 조회
func (c *Client) GetPlatformPartnerSchedule(ctx context.Context, id string, test *bool) (*PlatformPartner, error) {
	var result PlatformPartner
	if err := c.client.Get(ctx, "/platform/partners/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SchedulePlatformPartner 파트너 업데이트 예약
func (c *Client) SchedulePlatformPartner(ctx context.Context, id string, test *bool, body *SchedulePlatformPartnerBody) (*SchedulePlatformPartnerResponse, error) {
	var result SchedulePlatformPartnerResponse
	path := "/platform/partners/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReschedulePlatformPartner 파트너 예약 업데이트 재설정
func (c *Client) ReschedulePlatformPartner(ctx context.Context, id string, test *bool, body *SchedulePlatformPartnerBody) (*ReschedulePlatformPartnerResponse, error) {
	var result ReschedulePlatformPartnerResponse
	path := "/platform/partners/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Put(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPlatformPartnerSchedule 파트너 예약 업데이트 취소
func (c *Client) CancelPlatformPartnerSchedule(ctx context.Context, id string, test *bool) (*CancelPlatformPartnerScheduleResponse, error) {
	var result CancelPlatformPartnerScheduleResponse
	if err := c.client.Delete(ctx, "/platform/partners/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SchedulePlatformPartners 파트너 일괄 업데이트 예약
func (c *Client) SchedulePlatformPartners(ctx context.Context, test *bool, body *SchedulePlatformPartnersBody) (*SchedulePlatformPartnersResponse, error) {
	var result SchedulePlatformPartnersResponse
	path := "/platform/partners/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Contract ===============

// GetPlatformContracts 계약 다건 조회
func (c *Client) GetPlatformContracts(ctx context.Context, test *bool, page *PageInput, filter *PlatformContractFilterInput) (*GetPlatformContractsResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformContractsResponse
	if err := c.client.Get(ctx, "/platform/contracts", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformContract 계약 조회
func (c *Client) GetPlatformContract(ctx context.Context, id string, test *bool) (*PlatformContract, error) {
	var result PlatformContract
	if err := c.client.Get(ctx, "/platform/contracts/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformContract 계약 생성
func (c *Client) CreatePlatformContract(ctx context.Context, test *bool, body *CreatePlatformContractBody) (*CreatePlatformContractResponse, error) {
	var result CreatePlatformContractResponse
	path := "/platform/contracts"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePlatformContract 계약 수정
func (c *Client) UpdatePlatformContract(ctx context.Context, id string, test *bool, body *UpdatePlatformContractBody) (*UpdatePlatformContractResponse, error) {
	var result UpdatePlatformContractResponse
	path := "/platform/contracts/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Patch(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchivePlatformContract 계약 보관
func (c *Client) ArchivePlatformContract(ctx context.Context, id string, test *bool) (*ArchivePlatformContractResponse, error) {
	var result ArchivePlatformContractResponse
	path := "/platform/contracts/" + url.PathEscape(id) + "/archive"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecoverPlatformContract 계약 복원
func (c *Client) RecoverPlatformContract(ctx context.Context, id string, test *bool) (*RecoverPlatformContractResponse, error) {
	var result RecoverPlatformContractResponse
	path := "/platform/contracts/" + url.PathEscape(id) + "/recover"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Contract Schedule ===============

// GetPlatformContractSchedule 계약 예약 업데이트 조회
func (c *Client) GetPlatformContractSchedule(ctx context.Context, id string, test *bool) (*PlatformContract, error) {
	var result PlatformContract
	if err := c.client.Get(ctx, "/platform/contracts/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SchedulePlatformContract 계약 업데이트 예약
func (c *Client) SchedulePlatformContract(ctx context.Context, id string, test *bool, body *SchedulePlatformContractBody) (*SchedulePlatformContractResponse, error) {
	var result SchedulePlatformContractResponse
	path := "/platform/contracts/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReschedulePlatformContract 계약 예약 업데이트 재설정
func (c *Client) ReschedulePlatformContract(ctx context.Context, id string, test *bool, body *SchedulePlatformContractBody) (*ReschedulePlatformContractResponse, error) {
	var result ReschedulePlatformContractResponse
	path := "/platform/contracts/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Put(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPlatformContractSchedule 계약 예약 업데이트 취소
func (c *Client) CancelPlatformContractSchedule(ctx context.Context, id string, test *bool) (*CancelPlatformContractScheduleResponse, error) {
	var result CancelPlatformContractScheduleResponse
	if err := c.client.Delete(ctx, "/platform/contracts/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Additional Fee Policy ===============

// GetPlatformAdditionalFeePolicies 추가 수수료 정책 다건 조회
func (c *Client) GetPlatformAdditionalFeePolicies(ctx context.Context, test *bool, page *PageInput, filter *PlatformAdditionalFeePolicyFilterInput) (*GetPlatformAdditionalFeePoliciesResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformAdditionalFeePoliciesResponse
	if err := c.client.Get(ctx, "/platform/additional-fee-policies", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformAdditionalFeePolicy 추가 수수료 정책 조회
func (c *Client) GetPlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool) (*PlatformAdditionalFeePolicy, error) {
	var result PlatformAdditionalFeePolicy
	if err := c.client.Get(ctx, "/platform/additional-fee-policies/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformAdditionalFeePolicy 추가 수수료 정책 생성
func (c *Client) CreatePlatformAdditionalFeePolicy(ctx context.Context, test *bool, body *CreatePlatformAdditionalFeePolicyBody) (*CreatePlatformAdditionalFeePolicyResponse, error) {
	var result CreatePlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePlatformAdditionalFeePolicy 추가 수수료 정책 수정
func (c *Client) UpdatePlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool, body *UpdatePlatformAdditionalFeePolicyBody) (*UpdatePlatformAdditionalFeePolicyResponse, error) {
	var result UpdatePlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Patch(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchivePlatformAdditionalFeePolicy 추가 수수료 정책 보관
func (c *Client) ArchivePlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool) (*ArchivePlatformAdditionalFeePolicyResponse, error) {
	var result ArchivePlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies/" + url.PathEscape(id) + "/archive"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecoverPlatformAdditionalFeePolicy 추가 수수료 정책 복원
func (c *Client) RecoverPlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool) (*RecoverPlatformAdditionalFeePolicyResponse, error) {
	var result RecoverPlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies/" + url.PathEscape(id) + "/recover"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Additional Fee Policy Schedule ===============

// GetPlatformAdditionalFeePolicySchedule 추가 수수료 정책 예약 업데이트 조회
func (c *Client) GetPlatformAdditionalFeePolicySchedule(ctx context.Context, id string, test *bool) (*PlatformAdditionalFeePolicy, error) {
	var result PlatformAdditionalFeePolicy
	if err := c.client.Get(ctx, "/platform/additional-fee-policies/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SchedulePlatformAdditionalFeePolicy 추가 수수료 정책 업데이트 예약
func (c *Client) SchedulePlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool, body *SchedulePlatformAdditionalFeePolicyBody) (*SchedulePlatformAdditionalFeePolicyResponse, error) {
	var result SchedulePlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReschedulePlatformAdditionalFeePolicy 추가 수수료 정책 예약 업데이트 재설정
func (c *Client) ReschedulePlatformAdditionalFeePolicy(ctx context.Context, id string, test *bool, body *SchedulePlatformAdditionalFeePolicyBody) (*ReschedulePlatformAdditionalFeePolicyResponse, error) {
	var result ReschedulePlatformAdditionalFeePolicyResponse
	path := "/platform/additional-fee-policies/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Put(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPlatformAdditionalFeePolicySchedule 추가 수수료 정책 예약 업데이트 취소
func (c *Client) CancelPlatformAdditionalFeePolicySchedule(ctx context.Context, id string, test *bool) (*CancelPlatformAdditionalFeePolicyScheduleResponse, error) {
	var result CancelPlatformAdditionalFeePolicyScheduleResponse
	if err := c.client.Delete(ctx, "/platform/additional-fee-policies/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Discount Share Policy ===============

// GetPlatformDiscountSharePolicies 할인 분담 정책 다건 조회
func (c *Client) GetPlatformDiscountSharePolicies(ctx context.Context, test *bool, page *PageInput, filter *PlatformDiscountSharePolicyFilterInput) (*GetPlatformDiscountSharePoliciesResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformDiscountSharePoliciesResponse
	if err := c.client.Get(ctx, "/platform/discount-share-policies", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformDiscountSharePolicy 할인 분담 정책 조회
func (c *Client) GetPlatformDiscountSharePolicy(ctx context.Context, id string, test *bool) (*PlatformDiscountSharePolicy, error) {
	var result PlatformDiscountSharePolicy
	if err := c.client.Get(ctx, "/platform/discount-share-policies/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformDiscountSharePolicy 할인 분담 정책 생성
func (c *Client) CreatePlatformDiscountSharePolicy(ctx context.Context, test *bool, body *CreatePlatformDiscountSharePolicyBody) (*CreatePlatformDiscountSharePolicyResponse, error) {
	var result CreatePlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// UpdatePlatformDiscountSharePolicy 할인 분담 정책 수정
func (c *Client) UpdatePlatformDiscountSharePolicy(ctx context.Context, id string, test *bool, body *UpdatePlatformDiscountSharePolicyBody) (*UpdatePlatformDiscountSharePolicyResponse, error) {
	var result UpdatePlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies/" + url.PathEscape(id)
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Patch(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ArchivePlatformDiscountSharePolicy 할인 분담 정책 보관
func (c *Client) ArchivePlatformDiscountSharePolicy(ctx context.Context, id string, test *bool) (*ArchivePlatformDiscountSharePolicyResponse, error) {
	var result ArchivePlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies/" + url.PathEscape(id) + "/archive"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RecoverPlatformDiscountSharePolicy 할인 분담 정책 복원
func (c *Client) RecoverPlatformDiscountSharePolicy(ctx context.Context, id string, test *bool) (*RecoverPlatformDiscountSharePolicyResponse, error) {
	var result RecoverPlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies/" + url.PathEscape(id) + "/recover"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformDiscountSharePolicyFilterOptions 할인 분담 정책 필터 옵션 조회
func (c *Client) GetPlatformDiscountSharePolicyFilterOptions(ctx context.Context, test *bool) (*GetPlatformDiscountSharePolicyFilterOptionsResponse, error) {
	var result GetPlatformDiscountSharePolicyFilterOptionsResponse
	if err := c.client.Get(ctx, "/platform/discount-share-policy-filter-options", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Discount Share Policy Schedule ===============

// GetPlatformDiscountSharePolicySchedule 할인 분담 정책 예약 업데이트 조회
func (c *Client) GetPlatformDiscountSharePolicySchedule(ctx context.Context, id string, test *bool) (*PlatformDiscountSharePolicy, error) {
	var result PlatformDiscountSharePolicy
	if err := c.client.Get(ctx, "/platform/discount-share-policies/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SchedulePlatformDiscountSharePolicy 할인 분담 정책 업데이트 예약
func (c *Client) SchedulePlatformDiscountSharePolicy(ctx context.Context, id string, test *bool, body *SchedulePlatformDiscountSharePolicyBody) (*SchedulePlatformDiscountSharePolicyResponse, error) {
	var result SchedulePlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ReschedulePlatformDiscountSharePolicy 할인 분담 정책 예약 업데이트 재설정
func (c *Client) ReschedulePlatformDiscountSharePolicy(ctx context.Context, id string, test *bool, body *SchedulePlatformDiscountSharePolicyBody) (*ReschedulePlatformDiscountSharePolicyResponse, error) {
	var result ReschedulePlatformDiscountSharePolicyResponse
	path := "/platform/discount-share-policies/" + url.PathEscape(id) + "/schedule"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Put(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPlatformDiscountSharePolicySchedule 할인 분담 정책 예약 업데이트 취소
func (c *Client) CancelPlatformDiscountSharePolicySchedule(ctx context.Context, id string, test *bool) (*CancelPlatformDiscountSharePolicyScheduleResponse, error) {
	var result CancelPlatformDiscountSharePolicyScheduleResponse
	if err := c.client.Delete(ctx, "/platform/discount-share-policies/"+url.PathEscape(id)+"/schedule", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Transfer ===============

// GetPlatformTransferSummaries 정산건 다건 조회
func (c *Client) GetPlatformTransferSummaries(ctx context.Context, test *bool, page *PageInput, filter *PlatformTransferFilterInput) (*GetPlatformTransferSummariesResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformTransferSummariesResponse
	if err := c.client.Get(ctx, "/platform/transfer-summaries", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPlatformTransfer 정산건 조회
func (c *Client) GetPlatformTransfer(ctx context.Context, id string, test *bool) (*PlatformTransfer, error) {
	var result PlatformTransfer
	if err := c.client.Get(ctx, "/platform/transfers/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeletePlatformTransfer 정산건 삭제
func (c *Client) DeletePlatformTransfer(ctx context.Context, id string, test *bool) (*DeletePlatformTransferResponse, error) {
	var result DeletePlatformTransferResponse
	if err := c.client.Delete(ctx, "/platform/transfers/"+url.PathEscape(id), buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformManualTransfer 수기 정산건 생성
func (c *Client) CreatePlatformManualTransfer(ctx context.Context, test *bool, body *CreatePlatformManualTransferBody) (*CreateManualTransferResponse, error) {
	var result CreateManualTransferResponse
	path := "/platform/transfers/manual"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformOrderTransfer 주문 정산건 생성
func (c *Client) CreatePlatformOrderTransfer(ctx context.Context, test *bool, body *CreatePlatformOrderTransferBody) (*CreateOrderTransferResponse, error) {
	var result CreateOrderTransferResponse
	path := "/platform/transfers/order"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePlatformOrderCancelTransfer 주문 취소 정산건 생성
func (c *Client) CreatePlatformOrderCancelTransfer(ctx context.Context, test *bool, body *CreatePlatformOrderCancelTransferBody) (*CreateOrderCancelTransferResponse, error) {
	var result CreateOrderCancelTransferResponse
	path := "/platform/transfers/order-cancel"
	if test != nil && *test {
		path += "?test=true"
	}
	if err := c.client.Post(ctx, path, body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DownloadPlatformTransferSheet 정산 상세 내역 다운로드
func (c *Client) DownloadPlatformTransferSheet(ctx context.Context, test *bool, filter *PlatformTransferFilterInput, fields []string) (string, error) {
	query := buildQuery(test)
	if filter != nil || len(fields) > 0 {
		requestBody := map[string]interface{}{}
		if filter != nil {
			requestBody["filter"] = filter
		}
		if len(fields) > 0 {
			requestBody["fields"] = fields
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result string
	if err := c.client.Get(ctx, "/platform/transfer-summaries/sheet-file", query, &result); err != nil {
		return "", err
	}
	return result, nil
}

// =============== Payout ===============

// GetPlatformPayouts 지급 다건 조회
func (c *Client) GetPlatformPayouts(ctx context.Context, test *bool, page *PageInput, filter *PlatformPayoutFilterInput) (*GetPlatformPayoutsResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformPayoutsResponse
	if err := c.client.Get(ctx, "/platform/payouts", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Bulk Payout ===============

// GetPlatformBulkPayouts 일괄 지급 다건 조회
func (c *Client) GetPlatformBulkPayouts(ctx context.Context, test *bool, page *PageInput, filter *PlatformBulkPayoutFilterInput) (*GetPlatformBulkPayoutsResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformBulkPayoutsResponse
	if err := c.client.Get(ctx, "/platform/bulk-payouts", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Account Transfer ===============

// GetPlatformAccountTransfers 계좌 이체 다건 조회
func (c *Client) GetPlatformAccountTransfers(ctx context.Context, test *bool, page *PageInput, filter *PlatformAccountTransferFilterInput) (*GetPlatformAccountTransfersResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformAccountTransfersResponse
	if err := c.client.Get(ctx, "/platform/account-transfers", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Account Holder ===============

// GetPlatformAccountHolder 예금주 조회
func (c *Client) GetPlatformAccountHolder(ctx context.Context, bank common.Bank, accountNumber string, test *bool) (*GetPlatformAccountHolderResponse, error) {
	query := buildQuery(test)
	path := fmt.Sprintf("/platform/account-holders/%s/%s", url.PathEscape(string(bank)), url.PathEscape(accountNumber))
	var result GetPlatformAccountHolderResponse
	if err := c.client.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Company State ===============

// GetPlatformCompanyState 사업자 상태 조회
func (c *Client) GetPlatformCompanyState(ctx context.Context, businessRegistrationNumber string, test *bool) (*GetPlatformCompanyStateResponse, error) {
	query := buildQuery(test)
	path := "/platform/company-state/" + url.PathEscape(businessRegistrationNumber)
	var result GetPlatformCompanyStateResponse
	if err := c.client.Get(ctx, path, query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Partner Filter Options ===============

// GetPlatformPartnerFilterOptions 파트너 필터 옵션 조회
func (c *Client) GetPlatformPartnerFilterOptions(ctx context.Context, test *bool) (*GetPlatformPartnerFilterOptionsResponse, error) {
	var result GetPlatformPartnerFilterOptionsResponse
	if err := c.client.Get(ctx, "/platform/partner-filter-options", buildQuery(test), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Partner Settlements ===============

// GetPlatformPartnerSettlements 파트너 정산 내역 다건 조회
func (c *Client) GetPlatformPartnerSettlements(ctx context.Context, test *bool, page *PageInput, filter *PlatformPartnerSettlementFilterInput) (*GetPlatformPartnerSettlementsResponse, error) {
	query := buildQuery(test)
	if page != nil || filter != nil {
		requestBody := map[string]interface{}{}
		if page != nil {
			requestBody["page"] = page
		}
		if filter != nil {
			requestBody["filter"] = filter
		}
		jsonBytes, _ := json.Marshal(requestBody)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformPartnerSettlementsResponse
	if err := c.client.Get(ctx, "/platform/partner-settlements", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Bulk Account Transfers ===============

// GetPlatformBulkAccountTransfers 일괄 이체 내역 다건 조회
func (c *Client) GetPlatformBulkAccountTransfers(ctx context.Context, test *bool, body *GetPlatformBulkAccountTransfersBody) (*GetPlatformBulkAccountTransfersResponse, error) {
	query := buildQuery(test)
	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		query.Set("requestBody", string(jsonBytes))
	}

	var result GetPlatformBulkAccountTransfersResponse
	if err := c.client.Get(ctx, "/platform/bulk-account-transfers", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// Ensure unused imports are used
var _ = http.StatusOK
