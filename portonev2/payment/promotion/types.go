package promotion

import (
	"github.com/Gliese436/portone-go-client/portonev2/common"
)

// PromotionStatus 프로모션 상태
type PromotionStatus string

const (
	PromotionStatusSCHEDULED  PromotionStatus = "SCHEDULED"
	PromotionStatusINPROGRESS PromotionStatus = "IN_PROGRESS"
	PromotionStatusPAUSED     PromotionStatus = "PAUSED"
	PromotionStatusSTOPPED    PromotionStatus = "STOPPED"
	PromotionStatusCOMPLETED  PromotionStatus = "COMPLETED"
)

// DiscountType 할인 타입
type DiscountType string

const (
	DiscountTypeFIXED_AMOUNT DiscountType = "FIXED_AMOUNT"
	DiscountTypeFIXED_RATE   DiscountType = "FIXED_RATE"
)

// Promotion 프로모션
type Promotion struct {
	// 프로모션 아이디
	ID string `json:"id"`
	// 상점 아이디
	StoreId string `json:"storeId"`
	// 프로모션 이름
	Name string `json:"name"`
	// 할인 타입
	DiscountType DiscountType `json:"discountType"`
	// 총 예산
	TotalBudget *int64 `json:"totalBudget,omitempty"`
	// 최대 할인 금액 (건당)
	MaxDiscountAmount *int64 `json:"maxDiscountAmount,omitempty"`
	// 최소 결제 금액
	MinPaymentAmount *int64 `json:"minPaymentAmount,omitempty"`
	// 고정 금액 할인 (FIXED_AMOUNT 타입)
	DiscountAmount *int64 `json:"discountAmount,omitempty"`
	// 고정 비율 할인 (FIXED_RATE 타입, 0~100)
	DiscountRate *int `json:"discountRate,omitempty"`
	// 프로모션 상태
	Status PromotionStatus `json:"status"`
	// 시작 시점 (RFC 3339)
	StartedAt *string `json:"startedAt,omitempty"`
	// 종료 시점 (RFC 3339)
	EndedAt *string `json:"endedAt,omitempty"`
	// 생성 시점 (RFC 3339)
	CreatedAt string `json:"createdAt"`
	// 사용된 예산
	SpentBudget *int64 `json:"spentBudget,omitempty"`
	// 카드사 목록
	CardCompanies []string `json:"cardCompanies,omitempty"`
	// 통화
	Currency common.Currency `json:"currency"`
}

// GetPromotionResponse 프로모션 조회 응답
type GetPromotionResponse = Promotion

// GetPromotionsResponse 프로모션 목록 조회 응답
type GetPromotionsResponse struct {
	// 프로모션 목록
	Items []Promotion `json:"items"`
	// 페이지 정보
	Page common.PageInfo `json:"page"`
}
