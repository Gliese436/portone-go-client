package portonev2

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Gliese436/portone-go-client/portonev2/common"
	"github.com/Gliese436/portone-go-client/portonev2/payment"
	"github.com/Gliese436/portone-go-client/portonev2/payment/billingkey"
	"github.com/Gliese436/portone-go-client/portonev2/payment/cashreceipt"
	"github.com/Gliese436/portone-go-client/portonev2/payment/promotion"
	"github.com/Gliese436/portone-go-client/portonev2/payment/schedule"
)

// PaymentClient Payment API 클라이언트
type PaymentClient struct {
	client *Client

	billingKey  *BillingKeyClient
	cashReceipt *CashReceiptClient
	schedule    *PaymentScheduleClient
}

// NewPaymentClient 새 Payment 클라이언트를 생성합니다.
func NewPaymentClient(client *Client) *PaymentClient {
	return &PaymentClient{
		client: client,
	}
}

// BillingKey BillingKey API 클라이언트를 반환합니다.
func (c *PaymentClient) BillingKey() *BillingKeyClient {
	if c.billingKey == nil {
		c.billingKey = NewBillingKeyClient(c.client)
	}
	return c.billingKey
}

// CashReceipt CashReceipt API 클라이언트를 반환합니다.
func (c *PaymentClient) CashReceipt() *CashReceiptClient {
	if c.cashReceipt == nil {
		c.cashReceipt = NewCashReceiptClient(c.client)
	}
	return c.cashReceipt
}

// Schedule PaymentSchedule API 클라이언트를 반환합니다.
func (c *PaymentClient) Schedule() *PaymentScheduleClient {
	if c.schedule == nil {
		c.schedule = NewPaymentScheduleClient(c.client)
	}
	return c.schedule
}

// GetPayment 결제 건을 조회합니다.
func (c *PaymentClient) GetPayment(ctx context.Context, paymentID string, storeID *string) (*payment.Payment, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result payment.Payment
	if err := c.client.Get(ctx, "/payments/"+url.PathEscape(paymentID), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPaymentsOptions 결제 목록 조회 옵션
type GetPaymentsOptions struct {
	Page   *common.PageInput           `json:"page,omitempty"`
	Filter *payment.PaymentFilterInput `json:"filter,omitempty"`
}

// GetPayments 결제 목록을 조회합니다.
func (c *PaymentClient) GetPayments(ctx context.Context, opts *GetPaymentsOptions) (*payment.GetPaymentsResponse, error) {
	var result payment.GetPaymentsResponse
	if err := c.client.Post(ctx, "/payments", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllPaymentsByCursorOptions 커서 기반 결제 목록 조회 옵션
type GetAllPaymentsByCursorOptions struct {
	StoreID *string `json:"storeId,omitempty"`
	From    *string `json:"from,omitempty"`
	Until   *string `json:"until,omitempty"`
	Cursor  *string `json:"cursor,omitempty"`
	Size    *int    `json:"size,omitempty"`
}

// GetAllPaymentsByCursor 커서 기반으로 결제 목록을 조회합니다.
func (c *PaymentClient) GetAllPaymentsByCursor(ctx context.Context, opts *GetAllPaymentsByCursorOptions) (*payment.GetAllPaymentsByCursorResponse, error) {
	query := url.Values{}
	if opts != nil {
		if opts.StoreID != nil {
			query.Set("storeId", *opts.StoreID)
		} else if c.client.storeID != "" {
			query.Set("storeId", c.client.storeID)
		}
		if opts.From != nil {
			query.Set("from", *opts.From)
		}
		if opts.Until != nil {
			query.Set("until", *opts.Until)
		}
		if opts.Cursor != nil {
			query.Set("cursor", *opts.Cursor)
		}
		if opts.Size != nil {
			query.Set("size", fmt.Sprintf("%d", *opts.Size))
		}
	}

	var result payment.GetAllPaymentsByCursorResponse
	if err := c.client.Get(ctx, "/payments-by-cursor", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PayWithBillingKeyOptions 빌링키 결제 옵션
type PayWithBillingKeyOptions struct {
	StoreID                     *string                       `json:"storeId,omitempty"`
	BillingKey                  string                        `json:"billingKey"`
	ChannelKey                  *string                       `json:"channelKey,omitempty"`
	OrderName                   string                        `json:"orderName"`
	Customer                    *common.CustomerInput         `json:"customer,omitempty"`
	CustomData                  *string                       `json:"customData,omitempty"`
	Amount                      common.PaymentAmountInput     `json:"amount"`
	Currency                    common.Currency               `json:"currency"`
	InstallmentMonth            *int                          `json:"installmentMonth,omitempty"`
	UseFreeInterestFromMerchant *bool                         `json:"useFreeInterestFromMerchant,omitempty"`
	UseCardPoint                *bool                         `json:"useCardPoint,omitempty"`
	CashReceipt                 *common.CashReceiptInput      `json:"cashReceipt,omitempty"`
	Country                     *common.Country               `json:"country,omitempty"`
	NoticeUrls                  []string                      `json:"noticeUrls,omitempty"`
	Products                    []common.PaymentProduct       `json:"products,omitempty"`
	ProductCount                *int                          `json:"productCount,omitempty"`
	ProductType                 *common.PaymentProductType    `json:"productType,omitempty"`
	ShippingAddress             *common.SeparatedAddressInput `json:"shippingAddress,omitempty"`
	PromotionId                 *string                       `json:"promotionId,omitempty"`
	Bypass                      interface{}                   `json:"bypass,omitempty"`
}

// PayWithBillingKey 빌링키로 결제합니다.
func (c *PaymentClient) PayWithBillingKey(ctx context.Context, paymentID string, opts *PayWithBillingKeyOptions) (*payment.PayWithBillingKeyResponse, error) {
	body := map[string]interface{}{
		"storeId":    opts.StoreID,
		"billingKey": opts.BillingKey,
		"orderName":  opts.OrderName,
		"amount":     opts.Amount,
		"currency":   opts.Currency,
	}
	if opts.StoreID == nil && c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.ChannelKey != nil {
		body["channelKey"] = opts.ChannelKey
	}
	if opts.Customer != nil {
		body["customer"] = opts.Customer
	}
	if opts.CustomData != nil {
		body["customData"] = opts.CustomData
	}
	if opts.InstallmentMonth != nil {
		body["installmentMonth"] = opts.InstallmentMonth
	}
	if opts.UseFreeInterestFromMerchant != nil {
		body["useFreeInterestFromMerchant"] = opts.UseFreeInterestFromMerchant
	}
	if opts.UseCardPoint != nil {
		body["useCardPoint"] = opts.UseCardPoint
	}
	if opts.CashReceipt != nil {
		body["cashReceipt"] = opts.CashReceipt
	}
	if opts.Country != nil {
		body["country"] = opts.Country
	}
	if len(opts.NoticeUrls) > 0 {
		body["noticeUrls"] = opts.NoticeUrls
	}
	if len(opts.Products) > 0 {
		body["products"] = opts.Products
	}
	if opts.ProductCount != nil {
		body["productCount"] = opts.ProductCount
	}
	if opts.ProductType != nil {
		body["productType"] = opts.ProductType
	}
	if opts.ShippingAddress != nil {
		body["shippingAddress"] = opts.ShippingAddress
	}
	if opts.PromotionId != nil {
		body["promotionId"] = opts.PromotionId
	}
	if opts.Bypass != nil {
		body["bypass"] = opts.Bypass
	}

	var result payment.PayWithBillingKeyResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/billing-key", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelPayment 결제를 취소합니다.
func (c *PaymentClient) CancelPayment(ctx context.Context, paymentID string, opts *payment.CancelPaymentOptions) (*payment.CancelPaymentResponse, error) {
	body := map[string]interface{}{
		"reason": opts.Reason,
	}
	if opts.StoreId != nil {
		body["storeId"] = opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.Amount != nil {
		body["amount"] = opts.Amount
	}
	if opts.TaxFreeAmount != nil {
		body["taxFreeAmount"] = opts.TaxFreeAmount
	}
	if opts.VatAmount != nil {
		body["vatAmount"] = opts.VatAmount
	}
	if opts.Requester != nil {
		body["requester"] = opts.Requester
	}
	if opts.RefundAccount != nil {
		body["refundAccount"] = opts.RefundAccount
	}

	var result payment.CancelPaymentResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/cancel", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PreRegisterPayment 결제를 사전 등록합니다.
func (c *PaymentClient) PreRegisterPayment(ctx context.Context, paymentID string, storeID *string, totalAmount *int64, taxFreeAmount *int64, currency *common.Currency) error {
	body := map[string]interface{}{}
	if storeID != nil {
		body["storeId"] = storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if totalAmount != nil {
		body["totalAmount"] = totalAmount
	}
	if taxFreeAmount != nil {
		body["taxFreeAmount"] = taxFreeAmount
	}
	if currency != nil {
		body["currency"] = currency
	}

	return c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/pre-register", body, nil)
}

// ConfirmPayment 인증 결제를 수동 승인합니다.
func (c *PaymentClient) ConfirmPayment(ctx context.Context, paymentID string, opts *payment.ConfirmPaymentOptions) (*payment.ConfirmedPaymentSummary, error) {
	body := map[string]interface{}{
		"paymentToken": opts.PaymentToken,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.TxId != nil {
		body["txId"] = *opts.TxId
	}
	if opts.Currency != nil {
		body["currency"] = *opts.Currency
	}
	if opts.TotalAmount != nil {
		body["totalAmount"] = *opts.TotalAmount
	}
	if opts.TaxFreeAmount != nil {
		body["taxFreeAmount"] = *opts.TaxFreeAmount
	}
	if opts.IsTest != nil {
		body["isTest"] = *opts.IsTest
	}

	var result payment.ConfirmedPaymentSummary
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/confirm", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ResendWebhook 웹훅을 재전송합니다.
func (c *PaymentClient) ResendWebhook(ctx context.Context, paymentID string, webhookID *string, storeID *string) error {
	body := map[string]interface{}{}
	if webhookID != nil {
		body["webhookId"] = *webhookID
	}
	if storeID != nil {
		body["storeId"] = *storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}

	return c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/resend-webhook", body, nil)
}

// PayInstantly 수기 결제를 요청합니다.
func (c *PaymentClient) PayInstantly(ctx context.Context, paymentID string, opts *payment.PayInstantlyOptions) (*payment.PayInstantlyResponse, error) {
	body := map[string]interface{}{
		"method":    opts.Method,
		"orderName": opts.OrderName,
		"amount":    opts.Amount,
		"currency":  opts.Currency,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.ChannelKey != nil {
		body["channelKey"] = *opts.ChannelKey
	}
	if opts.ChannelGroupId != nil {
		body["channelGroupId"] = *opts.ChannelGroupId
	}
	if opts.IsCulturalExpense != nil {
		body["isCulturalExpense"] = *opts.IsCulturalExpense
	}
	if opts.IsEscrow != nil {
		body["isEscrow"] = *opts.IsEscrow
	}
	if opts.Customer != nil {
		body["customer"] = opts.Customer
	}
	if opts.CustomData != nil {
		body["customData"] = *opts.CustomData
	}
	if opts.Country != nil {
		body["country"] = *opts.Country
	}
	if len(opts.NoticeUrls) > 0 {
		body["noticeUrls"] = opts.NoticeUrls
	}
	if len(opts.Products) > 0 {
		body["products"] = opts.Products
	}
	if opts.ProductCount != nil {
		body["productCount"] = *opts.ProductCount
	}
	if opts.ProductType != nil {
		body["productType"] = *opts.ProductType
	}
	if opts.ShippingAddress != nil {
		body["shippingAddress"] = opts.ShippingAddress
	}
	if opts.PromotionId != nil {
		body["promotionId"] = *opts.PromotionId
	}

	var result payment.PayInstantlyResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/instant", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CapturePayment 수동 매입을 요청합니다.
func (c *PaymentClient) CapturePayment(ctx context.Context, paymentID string, storeID *string) (*payment.CapturePaymentResponse, error) {
	body := map[string]interface{}{}
	if storeID != nil {
		body["storeId"] = *storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}

	var result payment.CapturePaymentResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/capture", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CloseVirtualAccount 가상계좌를 말소합니다.
func (c *PaymentClient) CloseVirtualAccount(ctx context.Context, paymentID string, storeID *string) (*payment.CloseVirtualAccountResponse, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result payment.CloseVirtualAccountResponse
	if err := c.client.Request(ctx, "POST", "/payments/"+url.PathEscape(paymentID)+"/virtual-account/close", query, nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ApplyEscrowLogistics 에스크로 배송 정보를 등록합니다.
func (c *PaymentClient) ApplyEscrowLogistics(ctx context.Context, paymentID string, opts *payment.ApplyEscrowLogisticsOptions) (*payment.ApplyEscrowLogisticsResponse, error) {
	body := map[string]interface{}{
		"logistics": opts.Logistics,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.Sender != nil {
		body["sender"] = opts.Sender
	}
	if opts.Receiver != nil {
		body["receiver"] = opts.Receiver
	}
	if opts.SendEmail != nil {
		body["sendEmail"] = *opts.SendEmail
	}
	if len(opts.Products) > 0 {
		body["products"] = opts.Products
	}

	var result payment.ApplyEscrowLogisticsResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/escrow/logistics", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ModifyEscrowLogistics 에스크로 배송 정보를 수정합니다.
func (c *PaymentClient) ModifyEscrowLogistics(ctx context.Context, paymentID string, opts *payment.ModifyEscrowLogisticsOptions) (*payment.ModifyEscrowLogisticsResponse, error) {
	body := map[string]interface{}{
		"logistics": opts.Logistics,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.Sender != nil {
		body["sender"] = opts.Sender
	}
	if opts.Receiver != nil {
		body["receiver"] = opts.Receiver
	}
	if opts.SendEmail != nil {
		body["sendEmail"] = *opts.SendEmail
	}
	if len(opts.Products) > 0 {
		body["products"] = opts.Products
	}

	var result payment.ModifyEscrowLogisticsResponse
	if err := c.client.Patch(ctx, "/payments/"+url.PathEscape(paymentID)+"/escrow/logistics", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmEscrow 에스크로 결제를 구매 확정 처리합니다.
func (c *PaymentClient) ConfirmEscrow(ctx context.Context, paymentID string, storeID *string, fromStore *bool) (*payment.ConfirmEscrowResponse, error) {
	body := map[string]interface{}{}
	if storeID != nil {
		body["storeId"] = *storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if fromStore != nil {
		body["fromStore"] = *fromStore
	}

	var result payment.ConfirmEscrowResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/escrow/complete", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RegisterStoreReceipt 영수증 내 하위 상점 거래를 등록합니다.
func (c *PaymentClient) RegisterStoreReceipt(ctx context.Context, paymentID string, items []payment.RegisterStoreReceiptBodyItem, storeID *string) (*payment.RegisterStoreReceiptResponse, error) {
	body := map[string]interface{}{
		"items": items,
	}
	if storeID != nil {
		body["storeId"] = *storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}

	var result payment.RegisterStoreReceiptResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/register-store-receipt", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPaymentTransactions 결제 시도 내역을 조회합니다.
func (c *PaymentClient) GetPaymentTransactions(ctx context.Context, paymentID string, storeID *string) (*payment.GetPaymentTransactionsResponse, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result payment.GetPaymentTransactionsResponse
	if err := c.client.Get(ctx, "/payments/"+url.PathEscape(paymentID)+"/transactions", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetAllPaymentEventsByCursorOptions 커서 기반 결제 이벤트 목록 조회 옵션
type GetAllPaymentEventsByCursorOptions struct {
	StoreID *string `json:"storeId,omitempty"`
	From    *string `json:"from,omitempty"`
	Until   *string `json:"until,omitempty"`
	Cursor  *string `json:"cursor,omitempty"`
	Size    *int    `json:"size,omitempty"`
}

// GetAllPaymentEventsByCursor 커서 기반으로 결제 이벤트 목록을 조회합니다.
func (c *PaymentClient) GetAllPaymentEventsByCursor(ctx context.Context, opts *GetAllPaymentEventsByCursorOptions) (*payment.GetAllPaymentEventsByCursorResponse, error) {
	query := url.Values{}
	if opts != nil {
		if opts.StoreID != nil {
			query.Set("storeId", *opts.StoreID)
		} else if c.client.storeID != "" {
			query.Set("storeId", c.client.storeID)
		}
		if opts.From != nil {
			query.Set("from", *opts.From)
		}
		if opts.Until != nil {
			query.Set("until", *opts.Until)
		}
		if opts.Cursor != nil {
			query.Set("cursor", *opts.Cursor)
		}
		if opts.Size != nil {
			query.Set("size", fmt.Sprintf("%d", *opts.Size))
		}
	}

	var result payment.GetAllPaymentEventsByCursorResponse
	if err := c.client.Get(ctx, "/payment-events-by-cursor", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// BillingKeyClient BillingKey API 클라이언트
type BillingKeyClient struct {
	client *Client
}

// NewBillingKeyClient 새 BillingKey 클라이언트를 생성합니다.
func NewBillingKeyClient(client *Client) *BillingKeyClient {
	return &BillingKeyClient{client: client}
}

// GetBillingKeyInfo 빌링키 정보를 조회합니다.
func (c *BillingKeyClient) GetBillingKeyInfo(ctx context.Context, billingKeyValue string, storeID *string) (*billingkey.BillingKeyInfo, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result billingkey.BillingKeyInfo
	if err := c.client.Get(ctx, "/billing-keys/"+url.PathEscape(billingKeyValue), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// DeleteBillingKey 빌링키를 삭제합니다.
func (c *BillingKeyClient) DeleteBillingKey(ctx context.Context, billingKeyValue string, storeID *string) (*billingkey.DeleteBillingKeyResponse, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result billingkey.DeleteBillingKeyResponse
	if err := c.client.Delete(ctx, "/billing-keys/"+url.PathEscape(billingKeyValue), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetBillingKeyInfos 빌링키 목록을 조회합니다.
func (c *BillingKeyClient) GetBillingKeyInfos(ctx context.Context, opts *billingkey.GetBillingKeyInfosBody) (*billingkey.GetBillingKeyInfosResponse, error) {
	var result billingkey.GetBillingKeyInfosResponse
	if err := c.client.Post(ctx, "/billing-keys", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IssueBillingKey 빌링키를 발급합니다.
func (c *BillingKeyClient) IssueBillingKey(ctx context.Context, opts *billingkey.IssueBillingKeyBody) (*billingkey.IssueBillingKeyResponse, error) {
	var result billingkey.IssueBillingKeyResponse
	if err := c.client.Post(ctx, "/billing-keys", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmBillingKey 빌링키 발급을 수동 승인합니다.
func (c *BillingKeyClient) ConfirmBillingKey(ctx context.Context, opts *billingkey.ConfirmBillingKeyOptions) (*billingkey.ConfirmedBillingKeySummary, error) {
	body := map[string]interface{}{
		"billingIssueToken": opts.BillingIssueToken,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.IsTest != nil {
		body["isTest"] = *opts.IsTest
	}

	var result billingkey.ConfirmedBillingKeySummary
	if err := c.client.Post(ctx, "/billing-keys/confirm", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// ConfirmBillingKeyIssueAndPay 빌링키 발급 및 초회 결제를 수동 승인합니다.
func (c *BillingKeyClient) ConfirmBillingKeyIssueAndPay(ctx context.Context, opts *billingkey.ConfirmBillingKeyIssueAndPayOptions) (*billingkey.ConfirmedBillingKeyIssueAndPaySummary, error) {
	body := map[string]interface{}{
		"billingIssueToken": opts.BillingIssueToken,
	}
	if opts.StoreId != nil {
		body["storeId"] = *opts.StoreId
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}
	if opts.PaymentId != nil {
		body["paymentId"] = *opts.PaymentId
	}
	if opts.Currency != nil {
		body["currency"] = *opts.Currency
	}
	if opts.TotalAmount != nil {
		body["totalAmount"] = *opts.TotalAmount
	}
	if opts.TaxFreeAmount != nil {
		body["taxFreeAmount"] = *opts.TaxFreeAmount
	}
	if opts.IsTest != nil {
		body["isTest"] = *opts.IsTest
	}

	var result billingkey.ConfirmedBillingKeyIssueAndPaySummary
	if err := c.client.Post(ctx, "/billing-keys/confirm-issue-and-pay", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CashReceiptClient CashReceipt API 클라이언트
type CashReceiptClient struct {
	client *Client
}

// NewCashReceiptClient 새 CashReceipt 클라이언트를 생성합니다.
func NewCashReceiptClient(client *Client) *CashReceiptClient {
	return &CashReceiptClient{client: client}
}

// GetCashReceiptByPaymentId 결제 건의 현금영수증을 조회합니다.
func (c *CashReceiptClient) GetCashReceiptByPaymentId(ctx context.Context, paymentID string, storeID *string) (*cashreceipt.CashReceipt, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result cashreceipt.CashReceipt
	if err := c.client.Get(ctx, "/payments/"+url.PathEscape(paymentID)+"/cash-receipt", query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetCashReceipts 현금영수증 목록을 조회합니다.
func (c *CashReceiptClient) GetCashReceipts(ctx context.Context, opts *cashreceipt.GetCashReceiptsBody) (*cashreceipt.GetCashReceiptsResponse, error) {
	var result cashreceipt.GetCashReceiptsResponse
	if err := c.client.Post(ctx, "/cash-receipts", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// IssueCashReceipt 현금영수증을 발급합니다.
func (c *CashReceiptClient) IssueCashReceipt(ctx context.Context, opts *cashreceipt.IssueCashReceiptBody) (*cashreceipt.IssueCashReceiptResponse, error) {
	var result cashreceipt.IssueCashReceiptResponse
	if err := c.client.Post(ctx, "/cash-receipts", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CancelCashReceiptByPaymentId 결제 건의 현금영수증을 취소합니다.
func (c *CashReceiptClient) CancelCashReceiptByPaymentId(ctx context.Context, paymentID string, storeID *string) (*cashreceipt.CancelCashReceiptByPaymentIdResponse, error) {
	body := map[string]interface{}{}
	if storeID != nil {
		body["storeId"] = storeID
	} else if c.client.storeID != "" {
		body["storeId"] = c.client.storeID
	}

	var result cashreceipt.CancelCashReceiptByPaymentIdResponse
	if err := c.client.Post(ctx, "/payments/"+url.PathEscape(paymentID)+"/cash-receipt/cancel", body, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// PaymentScheduleClient PaymentSchedule API 클라이언트
type PaymentScheduleClient struct {
	client *Client
}

// NewPaymentScheduleClient 새 PaymentSchedule 클라이언트를 생성합니다.
func NewPaymentScheduleClient(client *Client) *PaymentScheduleClient {
	return &PaymentScheduleClient{client: client}
}

// GetPaymentSchedule 결제 예약 건을 조회합니다.
func (c *PaymentScheduleClient) GetPaymentSchedule(ctx context.Context, paymentScheduleID string, storeID *string) (*schedule.PaymentSchedule, error) {
	query := url.Values{}
	if storeID != nil {
		query.Set("storeId", *storeID)
	} else if c.client.storeID != "" {
		query.Set("storeId", c.client.storeID)
	}

	var result schedule.PaymentSchedule
	if err := c.client.Get(ctx, "/payment-schedules/"+url.PathEscape(paymentScheduleID), query, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetPaymentSchedules 결제 예약 목록을 조회합니다.
func (c *PaymentScheduleClient) GetPaymentSchedules(ctx context.Context, opts *schedule.GetPaymentSchedulesBody) (*schedule.GetPaymentSchedulesResponse, error) {
	var result schedule.GetPaymentSchedulesResponse
	if err := c.client.Post(ctx, "/payment-schedules", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// CreatePaymentSchedule 결제 예약을 생성합니다.
func (c *PaymentScheduleClient) CreatePaymentSchedule(ctx context.Context, opts *schedule.CreatePaymentScheduleBody) (*schedule.CreatePaymentScheduleResponse, error) {
	var result schedule.CreatePaymentScheduleResponse
	if err := c.client.Post(ctx, "/payment-schedules/create", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// RevokePaymentSchedules 결제 예약을 취소합니다.
func (c *PaymentScheduleClient) RevokePaymentSchedules(ctx context.Context, opts *schedule.RevokePaymentSchedulesBody) (*schedule.RevokePaymentSchedulesResponse, error) {
	var result schedule.RevokePaymentSchedulesResponse
	if err := c.client.Post(ctx, "/payment-schedules/revoke", opts, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// =============== Promotion ===============

// GetPromotion 프로모션을 조회합니다.
func (c *PaymentClient) GetPromotion(ctx context.Context, promotionID string) (*promotion.Promotion, error) {
	var result promotion.Promotion
	if err := c.client.Get(ctx, "/promotions/"+url.PathEscape(promotionID), nil, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
