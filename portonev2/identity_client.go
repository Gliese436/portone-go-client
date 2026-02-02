package portonev2

import (
	"context"
	"net/url"

	"github.com/gliese436/portone-go-client/portonev2/identity"
)

// IdentityVerificationClient 본인인증 API 클라이언트
type IdentityVerificationClient struct {
	client *Client
}

// NewIdentityVerificationClient 새 본인인증 클라이언트를 생성합니다.
func NewIdentityVerificationClient(client *Client) *IdentityVerificationClient {
	return &IdentityVerificationClient{client: client}
}

// GetIdentityVerification 본인인증 정보를 조회합니다.
func (c *IdentityVerificationClient) GetIdentityVerification(ctx context.Context, identityVerificationID string, storeID *string) (*identity.IdentityVerification, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result identity.IdentityVerification
	if err := c.client.Get(ctx, "/identity-verifications/"+url.PathEscape(identityVerificationID), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// SendIdentityVerification 본인인증 요청을 전송합니다.
func (c *IdentityVerificationClient) SendIdentityVerification(ctx context.Context, identityVerificationID string, opts *identity.SendIdentityVerificationBody) error {
	body := map[string]interface{}{
		"channelKey": opts.ChannelKey,
		"customer":   opts.Customer,
	}
	if opts.StoreId != nil {
		body["storeId"] = opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.CustomData != nil {
		body["customData"] = opts.CustomData
	}
	if opts.Bypass != nil {
		body["bypass"] = opts.Bypass
	}
	if opts.Method != nil {
		body["method"] = opts.Method
	}

	return c.client.Post(ctx, "/identity-verifications/"+url.PathEscape(identityVerificationID)+"/send", body, nil)
}

// ConfirmIdentityVerification 본인인증을 확인합니다.
func (c *IdentityVerificationClient) ConfirmIdentityVerification(ctx context.Context, identityVerificationID string, opts *identity.ConfirmIdentityVerificationBody) (*identity.ConfirmIdentityVerificationResponse, error) {
	body := map[string]interface{}{}
	if opts != nil {
		if opts.StoreId != nil {
			body["storeId"] = opts.StoreId
		} else if c.client.storeID != "" {
			body["storeId"] = c.client.storeID
		}
		if opts.Otp != nil {
			body["otp"] = opts.Otp
		}
	}

	var result identity.ConfirmIdentityVerificationResponse
	if err := c.client.Post(ctx, "/identity-verifications/"+url.PathEscape(identityVerificationID)+"/confirm", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ResendIdentityVerification 본인인증을 재전송합니다.
func (c *IdentityVerificationClient) ResendIdentityVerification(ctx context.Context, identityVerificationID string, storeID *string) error {
	body := map[string]interface{}{}
	if storeID != nil {
		body["storeId"] = storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}

	return c.client.Post(ctx, "/identity-verifications/"+url.PathEscape(identityVerificationID)+"/resend", body, nil)
}

// GetIdentityVerifications 본인인증 목록을 조회합니다.
func (c *IdentityVerificationClient) GetIdentityVerifications(ctx context.Context, opts *identity.GetIdentityVerificationsBody) (*identity.GetIdentityVerificationsResponse, error) {
	var result identity.GetIdentityVerificationsResponse
	if err := c.client.Post(ctx, "/identity-verifications", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
