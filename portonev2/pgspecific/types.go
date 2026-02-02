package pgspecific

import (
	"github.com/gliese436/portone-go-client/portonev2/common"
)

// KakaopayPaymentOrder 카카오페이 주문 정보
type KakaopayPaymentOrder struct {
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 결제 금액
	Amount int64 `json:"amount"`
	// 주문명
	OrderName string `json:"orderName"`
}

// NaverpayPaymentOrder 네이버페이 주문 정보
type NaverpayPaymentOrder struct {
	// 결제 건 아이디
	PaymentId string `json:"paymentId"`
	// 결제 금액
	Amount int64 `json:"amount"`
	// 주문명
	OrderName string `json:"orderName"`
}

// GetKakaopayPaymentOrderBody 카카오페이 주문 조회 요청
type GetKakaopayPaymentOrderBody struct {
	// 채널 키
	ChannelKey string `json:"channelKey"`
	// PG사 거래 아이디
	PgTxId string `json:"pgTxId"`
}

// GetKakaopayPaymentOrderResponse 카카오페이 주문 조회 응답
type GetKakaopayPaymentOrderResponse struct {
	// 주문 상태
	StatusCode string `json:"statusCode"`
	// 주문 정보
	Orders []KakaopayPaymentOrder `json:"orders"`
}

// GetNaverpayPaymentOrderBody 네이버페이 주문 조회 요청
type GetNaverpayPaymentOrderBody struct {
	// 채널 키
	ChannelKey string `json:"channelKey"`
	// 네이버페이 결제 아이디
	NaverPayPaymentId string `json:"naverPayPaymentId"`
}

// GetNaverpayPaymentOrderResponse 네이버페이 주문 조회 응답
type GetNaverpayPaymentOrderResponse struct {
	// 주문 상태
	StatusCode string `json:"statusCode"`
	// 주문 정보
	Orders []NaverpayPaymentOrder `json:"orders"`
}

// PgSpecificOptions PG사별 옵션
type PgSpecificOptions struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 채널 키
	ChannelKey string `json:"channelKey"`
}

// TosspaymentsRegisterEscrowLogisticsBody 토스페이먼츠 에스크로 배송 등록 요청
type TosspaymentsRegisterEscrowLogisticsBody struct {
	// 상점 아이디
	StoreId *string `json:"storeId,omitempty"`
	// 택배사 코드
	Company string `json:"company"`
	// 송장 번호
	InvoiceNumber string `json:"invoiceNumber"`
	// 발송 일시 (yyyyMMdd)
	SentAt string `json:"sentAt"`
}

// TosspaymentsRegisterEscrowLogisticsResponse 토스페이먼츠 에스크로 배송 등록 응답
type TosspaymentsRegisterEscrowLogisticsResponse struct{}

// NaverpayPointInfo 네이버페이 포인트 정보
type NaverpayPointInfo struct {
	// 사용 가능 포인트
	Available int64 `json:"available"`
	// 총 포인트
	Total int64 `json:"total"`
}

// PgInfo PG사 정보
type PgInfo struct {
	// PG사
	PgProvider common.PgProvider `json:"pgProvider"`
	// PG사 상점 아이디
	PgMerchantId string `json:"pgMerchantId"`
}
