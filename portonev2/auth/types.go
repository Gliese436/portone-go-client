package auth

// LoginViaApiSecretBody API 시크릿으로 로그인 요청
type LoginViaApiSecretBody struct {
	// API 시크릿
	ApiSecret string `json:"apiSecret"`
}

// LoginViaApiSecretResponse API 시크릿 로그인 응답
type LoginViaApiSecretResponse struct {
	// 액세스 토큰
	AccessToken string `json:"accessToken"`
	// 리프레시 토큰
	RefreshToken string `json:"refreshToken"`
}

// RefreshTokenBody 토큰 갱신 요청
type RefreshTokenBody struct {
	// 리프레시 토큰
	RefreshToken string `json:"refreshToken"`
}

// RefreshTokenResponse 토큰 갱신 응답
type RefreshTokenResponse struct {
	// 새 액세스 토큰
	AccessToken string `json:"accessToken"`
	// 새 리프레시 토큰
	RefreshToken string `json:"refreshToken"`
}
