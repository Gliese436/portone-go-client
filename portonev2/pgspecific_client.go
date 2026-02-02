package portonev2

import (
	"context"
	"net/url"

	"github.com/gliese436/portone-go-client/portonev2/pgspecific"
)

// PgSpecificClient PG사별 API 클라이언트
type PgSpecificClient struct {
	client *Client
}

// NewPgSpecificClient 새 PG사별 클라이언트를 생성합니다.
func NewPgSpecificClient(client *Client) *PgSpecificClient {
	return &PgSpecificClient{client: client}
}

// GetKakaopayPaymentOrder 카카오페이 주문을 조회합니다.
func (c *PgSpecificClient) GetKakaopayPaymentOrder(ctx context.Context, pgTxID string, channelKey string) (*pgspecific.GetKakaopayPaymentOrderResponse, error) {
	query := url.Values{}
	query.Set("pgTxId", pgTxID)
	query.Set("channelKey", channelKey)

	var result pgspecific.GetKakaopayPaymentOrderResponse
	if err := c.client.Get(ctx, "/kakaopay/payment/order", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
