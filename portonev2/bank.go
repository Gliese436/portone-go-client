package portonev2

import (
	"context"
	"net/http"

	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// getBankInfos 는 상류 codegen 이 생성 대상에서 제외하므로 수작업으로 유지합니다.

// BankInfoName 은행 명칭
type BankInfoName struct {
	// Ko 한국어 명칭
	Ko string `json:"ko"`
}

// BankInfo 은행 정보
type BankInfo struct {
	// Bank 은행
	Bank common.Bank `json:"bank"`
	// Name 언어별 명칭
	Name BankInfoName `json:"name"`
}

// GetBankInfosResponse 은행 정보 조회 성공 응답 정보
type GetBankInfosResponse struct {
	// Items 조회된 은행 정보 리스트
	Items []BankInfo `json:"items"`
}

// GetBankInfos 은행 정보 조회
//
// 포트원에서 지원하는 은행 정보를 조회합니다.
//
// GET /banks
func (c *Client) GetBankInfos(ctx context.Context) (*GetBankInfosResponse, error) {
	var result GetBankInfosResponse
	if err := c.do(ctx, http.MethodGet, "/banks", nil, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
