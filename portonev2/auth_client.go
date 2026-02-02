package portonev2

import (
	"context"

	"github.com/Gliese436/portone-go-client/portonev2/auth"
)

// AuthClient 인증 API 클라이언트
type AuthClient struct {
	client *Client
}

// NewAuthClient 새 인증 클라이언트를 생성합니다.
func NewAuthClient(client *Client) *AuthClient {
	return &AuthClient{client: client}
}

// LoginViaApiSecret API 시크릿으로 로그인합니다.
func (c *AuthClient) LoginViaApiSecret(ctx context.Context, apiSecret string) (*auth.LoginViaApiSecretResponse, error) {
	body := auth.LoginViaApiSecretBody{
		ApiSecret: apiSecret,
	}

	var result auth.LoginViaApiSecretResponse
	if err := c.client.Post(ctx, "/login/api-secret", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RefreshToken 토큰을 갱신합니다.
func (c *AuthClient) RefreshToken(ctx context.Context, refreshToken string) (*auth.RefreshTokenResponse, error) {
	body := auth.RefreshTokenBody{
		RefreshToken: refreshToken,
	}

	var result auth.RefreshTokenResponse
	if err := c.client.Post(ctx, "/token/refresh", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
