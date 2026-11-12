package portonev2_test

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Gliese436/portone-go-client/portonev2"
	"github.com/Gliese436/portone-go-client/portonev2/b2b/counterparty"
	"github.com/Gliese436/portone-go-client/portonev2/b2b/taxinvoice"
	"github.com/Gliese436/portone-go-client/portonev2/common"
	"github.com/Gliese436/portone-go-client/portonev2/payment"
	"github.com/Gliese436/portone-go-client/portonev2/payment/paymentschedule"
	"github.com/Gliese436/portone-go-client/portonev2/webhook"
)

type captured struct {
	Method      string
	Path        string
	Query       map[string][]string
	Body        string
	ContentType string
	Auth        string
	UserAgent   string
}

// newServer 는 요청을 기록하고 status/response 로 응답하는 테스트 서버와 클라이언트를 만듭니다.
func newServer(t *testing.T, status int, response string) (*portonev2.Client, *captured) {
	t.Helper()
	got := &captured{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		*got = captured{
			Method:      r.Method,
			Path:        r.URL.EscapedPath(),
			Query:       r.URL.Query(),
			Body:        string(body),
			ContentType: r.Header.Get("Content-Type"),
			Auth:        r.Header.Get("Authorization"),
			UserAgent:   r.Header.Get("User-Agent"),
		}
		w.WriteHeader(status)
		_, _ = io.WriteString(w, response)
	}))
	t.Cleanup(srv.Close)
	client := portonev2.NewClient("test-secret",
		portonev2.WithBaseURL(srv.URL),
		portonev2.WithStoreID("store-default"),
	)
	return client, got
}

func ptr[T any](v T) *T { return &v }

func TestGetPaymentUsesDefaultStoreIDAndDecodesUnion(t *testing.T) {
	client, got := newServer(t, 200, `{
		"status": "PAID",
		"id": "pay-1",
		"transactionId": "tx-1",
		"storeId": "store-default",
		"orderName": "주문",
		"amount": {"total": 1000, "taxFree": 0, "vat": 91, "discount": 0, "paid": 1000, "cancelled": 0, "cancelledTaxFree": 0},
		"method": {"type": "PaymentMethodCard", "approvalNumber": "1234"},
		"origin": {"platformType": "PC"}
	}`)

	p, err := client.Payment().GetPayment(context.Background(), &payment.GetPaymentOptions{PaymentId: "pay/1"})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodGet || got.Path != "/payments/pay%2F1" {
		t.Fatalf("unexpected request %s %s", got.Method, got.Path)
	}
	if q := got.Query["storeId"]; len(q) != 1 || q[0] != "store-default" {
		t.Fatalf("storeId query = %v", q)
	}
	if got.Auth != "PortOne test-secret" || got.UserAgent != portonev2.UserAgent {
		t.Fatalf("headers: auth=%q ua=%q", got.Auth, got.UserAgent)
	}
	if got.Body != "" {
		t.Fatalf("GET must not send a body, got %q", got.Body)
	}
	if p.Status != "PAID" || p.PaidPayment == nil {
		t.Fatalf("expected PaidPayment variant, got %+v", p)
	}
	paid, ok := p.Value().(*payment.PaidPayment)
	if !ok || paid.Id != "pay-1" || paid.Amount.Total != 1000 {
		t.Fatalf("unexpected paid payment %+v", paid)
	}
	if paid.Method == nil || paid.Method.PaymentMethodCard == nil || *paid.Method.PaymentMethodCard.ApprovalNumber != "1234" {
		t.Fatalf("payment method not decoded: %+v", paid.Method)
	}
	if paid.Origin.PlatformType != "PC" {
		t.Fatalf("origin not decoded: %+v", paid.Origin)
	}
}

func TestGetSendsBodyAsRequestBodyQuery(t *testing.T) {
	client, got := newServer(t, 200, `{"items": [], "page": {"number": 0, "size": 10, "totalCount": 0}}`)
	_, err := client.Payment().GetPayments(context.Background(), &payment.GetPaymentsOptions{
		Page: &common.PageInput{Number: ptr[int64](0), Size: ptr[int64](10)},
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodGet || got.Path != "/payments" || got.Body != "" {
		t.Fatalf("unexpected request %+v", got)
	}
	if rb := got.Query["requestBody"]; len(rb) != 1 || rb[0] != `{"page":{"number":0,"size":10}}` {
		t.Fatalf("requestBody query = %v", rb)
	}
}

func TestGetWithoutOptionsStillSendsEmptyRequestBody(t *testing.T) {
	client, got := newServer(t, 200, `{"items": [], "page": {"number": 0, "size": 10, "totalCount": 0}}`)
	if _, err := client.Payment().GetPayments(context.Background(), nil); err != nil {
		t.Fatal(err)
	}
	if rb := got.Query["requestBody"]; len(rb) != 1 || rb[0] != `{}` {
		t.Fatalf("requestBody query = %v", rb)
	}
}

func TestDeleteSendsBodyAsRequestBodyQuery(t *testing.T) {
	client, got := newServer(t, 200, `{"revokedScheduleIds": ["s1"]}`)
	res, err := client.Payment().PaymentSchedule().RevokePaymentSchedules(context.Background(),
		&paymentschedule.RevokePaymentSchedulesOptions{ScheduleIds: []string{"s1"}})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodDelete || got.Path != "/payment-schedules" || got.Body != "" {
		t.Fatalf("unexpected request %+v", got)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(got.Query["requestBody"][0]), &body); err != nil {
		t.Fatal(err)
	}
	if body["storeId"] != "store-default" {
		t.Fatalf("storeId default not applied: %v", body)
	}
	if len(res.RevokedScheduleIds) != 1 {
		t.Fatalf("unexpected response %+v", res)
	}
}

func TestEmptyBodyIsNotSent(t *testing.T) {
	// DeleteB2bCounterpartyBody 는 프로퍼티가 없으므로 JS 와 동일하게 바디를 생략해야 합니다.
	client, got := newServer(t, 200, `{}`)
	_, err := client.B2b().Counterparty().DeleteB2bCounterparty(context.Background(),
		&counterparty.DeleteB2bCounterpartyOptions{CounterpartyId: "cp-1", Test: ptr(true)})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodDelete || got.Path != "/b2b/counterparties/cp-1" {
		t.Fatalf("unexpected request %s %s", got.Method, got.Path)
	}
	if _, ok := got.Query["requestBody"]; ok || got.Body != "" {
		t.Fatalf("empty body must not be sent: query=%v body=%q", got.Query, got.Body)
	}
	if got.Query["test"][0] != "true" {
		t.Fatalf("test query = %v", got.Query["test"])
	}

	// POST 에 바디가 없는 경우에도 바디와 Content-Type 을 보내지 않습니다.
	client2, got2 := newServer(t, 200, `{"isValid": true}`)
	_, err = client2.B2b().Counterparty().ValidateB2bCounterpartyCertificate(context.Background(),
		&counterparty.ValidateB2bCounterpartyCertificateOptions{Brn: "1234567890"})
	if err != nil {
		t.Fatal(err)
	}
	if got2.Method != http.MethodPost || got2.Path != "/b2b/counterparties/1234567890/certificate/validate" || got2.Body != "" || got2.ContentType != "" {
		t.Fatalf("unexpected request %+v", got2)
	}
}

func TestPostBodyAppliesStoreIDAndNewFields(t *testing.T) {
	client, got := newServer(t, 200, `{"cancellation": {"status": "SUCCEEDED", "id": "c1", "totalAmount": 1000, "taxFreeAmount": 0, "vatAmount": 0, "reason": "r", "requestedAt": "2026-01-01T00:00:00Z"}}`)
	res, err := client.Payment().CancelPayment(context.Background(), &payment.CancelPaymentOptions{
		PaymentId:   "pay-1",
		Reason:      "고객 요청",
		SkipWebhook: ptr(true),
		RefundEmail: ptr("buyer@example.com"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodPost || got.Path != "/payments/pay-1/cancel" || got.ContentType != "application/json" {
		t.Fatalf("unexpected request %+v", got)
	}
	var body map[string]any
	if err := json.Unmarshal([]byte(got.Body), &body); err != nil {
		t.Fatal(err)
	}
	want := map[string]any{"storeId": "store-default", "reason": "고객 요청", "skipWebhook": true, "refundEmail": "buyer@example.com"}
	for k, v := range want {
		if body[k] != v {
			t.Fatalf("body[%s] = %v, want %v (body=%s)", k, body[k], v, got.Body)
		}
	}
	if _, ok := body["amount"]; ok {
		t.Fatalf("unset optional field must be omitted: %s", got.Body)
	}
	if res.Cancellation.SucceededPaymentCancellation == nil {
		t.Fatalf("cancellation union not decoded: %+v", res.Cancellation)
	}
}

func TestExplicitStoreIDOverridesDefault(t *testing.T) {
	client, got := newServer(t, 200, `{"status": "READY", "id": "p"}`)
	_, err := client.Payment().GetPayment(context.Background(), &payment.GetPaymentOptions{PaymentId: "p", StoreId: ptr("store-other")})
	if err != nil {
		t.Fatal(err)
	}
	if got.Query["storeId"][0] != "store-other" {
		t.Fatalf("storeId = %v", got.Query["storeId"])
	}
}

func TestCreatePaymentScheduleUsesNewPathAndBody(t *testing.T) {
	client, got := newServer(t, 200, `{"schedule": {"id": "sch-1"}}`)
	_, err := client.Payment().PaymentSchedule().CreatePaymentSchedule(context.Background(), &paymentschedule.CreatePaymentScheduleOptions{
		PaymentId: "pay-1",
		Payment: paymentschedule.BillingKeyPaymentScheduleInput{
			BillingKey: "bk",
			OrderName:  "구독",
			Amount:     common.PaymentAmountInput{Total: 1000},
			Currency:   common.CurrencyKrw,
		},
		TimeToPay: "2026-12-01T00:00:00Z",
	})
	if err != nil {
		t.Fatal(err)
	}
	if got.Method != http.MethodPost || got.Path != "/payments/pay-1/schedule" {
		t.Fatalf("unexpected request %s %s", got.Method, got.Path)
	}
	var body struct {
		Payment   map[string]any `json:"payment"`
		TimeToPay string         `json:"timeToPay"`
	}
	if err := json.Unmarshal([]byte(got.Body), &body); err != nil {
		t.Fatal(err)
	}
	if body.TimeToPay != "2026-12-01T00:00:00Z" || body.Payment["billingKey"] != "bk" || body.Payment["currency"] != "KRW" {
		t.Fatalf("unexpected body %s", got.Body)
	}
}

func TestAPIErrorDecodesIntoOperationErrorUnion(t *testing.T) {
	client, _ := newServer(t, 404, `{"type": "PAYMENT_NOT_FOUND", "message": "결제 건이 존재하지 않습니다."}`)
	_, err := client.Payment().GetPayment(context.Background(), &payment.GetPaymentOptions{PaymentId: "missing"})
	apiErr, ok := portonev2.AsAPIError(err)
	if !ok {
		t.Fatalf("expected *APIError, got %T %v", err, err)
	}
	if apiErr.StatusCode != 404 || apiErr.Type != "PAYMENT_NOT_FOUND" {
		t.Fatalf("unexpected api error %+v", apiErr)
	}
	var opErr payment.GetPaymentError
	if err := apiErr.Decode(&opErr); err != nil {
		t.Fatal(err)
	}
	if opErr.PaymentNotFoundError == nil || *opErr.PaymentNotFoundError.Message != "결제 건이 존재하지 않습니다." {
		t.Fatalf("error union not decoded: %+v", opErr)
	}
	var target *portonev2.APIError
	if !errors.As(err, &target) {
		t.Fatal("errors.As failed")
	}
}

func TestCSVResponseIsReturnedAsString(t *testing.T) {
	client, got := newServer(t, 200, "a,b\n1,2\n")
	csv, err := client.B2b().TaxInvoice().DownloadB2bTaxInvoicesSheet(context.Background(), &taxinvoice.DownloadB2bTaxInvoicesSheetOptions{
		Fields: []taxinvoice.TaxInvoicesSheetField{taxinvoice.TaxInvoicesSheetFieldSupplierCounterpartyId},
	})
	if err != nil {
		t.Fatal(err)
	}
	if csv != "a,b\n1,2\n" || got.Path != "/b2b/tax-invoices-sheet" {
		t.Fatalf("csv=%q path=%s", csv, got.Path)
	}
	if !strings.Contains(got.Query["requestBody"][0], "SUPPLIER_COUNTERPARTY_ID") {
		t.Fatalf("requestBody = %v", got.Query["requestBody"])
	}
}

func TestUnionMarshalAddsDiscriminatorAndKeepsUnknown(t *testing.T) {
	method := payment.PaymentMethod{PaymentMethodCard: &payment.PaymentMethodCard{ApprovalNumber: ptr("1")}}
	encoded, err := json.Marshal(method)
	if err != nil {
		t.Fatal(err)
	}
	if string(encoded) != `{"type":"PaymentMethodCard","approvalNumber":"1"}` {
		t.Fatalf("encoded = %s", encoded)
	}

	var unknown payment.PaymentMethod
	raw := `{"type":"SOMETHING_NEW","x":1}`
	if err := json.Unmarshal([]byte(raw), &unknown); err != nil {
		t.Fatal(err)
	}
	if unknown.Type != "SOMETHING_NEW" || unknown.Value() != nil || string(unknown.Unrecognized) != raw {
		t.Fatalf("unknown variant = %+v", unknown)
	}
	reencoded, _ := json.Marshal(unknown)
	if string(reencoded) != raw {
		t.Fatalf("reencoded = %s", reencoded)
	}
}

func signWebhook(t *testing.T, secret []byte, id string, ts int64, payload string) http.Header {
	t.Helper()
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(id + "." + strconv.FormatInt(ts, 10) + "." + payload))
	sig := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	h := http.Header{}
	h.Set("webhook-id", id)
	h.Set("webhook-timestamp", strconv.FormatInt(ts, 10))
	h.Set("webhook-signature", "v1,invalid v1,"+sig)
	return h
}

func TestWebhookVerifyDecodesTypedPayload(t *testing.T) {
	secret := []byte("super-secret-key")
	secretString := "whsec_" + base64.StdEncoding.EncodeToString(secret)
	now := time.Now().Unix()

	payload := `{"type":"Transaction.Paid","timestamp":"2026-11-12T00:00:00Z","data":{"paymentId":"pay-1","storeId":"store-1","transactionId":"tx-1"}}`
	wh, err := webhook.Verify(secretString, []byte(payload), signWebhook(t, secret, "msg-1", now, payload))
	if err != nil {
		t.Fatal(err)
	}
	if !wh.IsTransaction() || wh.WebhookTransactionPaid == nil || wh.WebhookTransactionPaid.Data.PaymentId != "pay-1" {
		t.Fatalf("unexpected webhook %+v", wh)
	}

	payload2 := `{"type":"BillingKey.Issued","timestamp":"2026-11-12T00:00:00Z","data":{"billingKey":"bk-1","storeId":"store-1"}}`
	wh2, err := webhook.Verify(secretString, []byte(payload2), signWebhook(t, secret, "msg-2", now, payload2))
	if err != nil {
		t.Fatal(err)
	}
	if !wh2.IsBillingKey() || wh2.WebhookBillingKeyIssued == nil || wh2.WebhookBillingKeyIssued.Data.BillingKey != "bk-1" {
		t.Fatalf("unexpected webhook %+v", wh2)
	}

	_, err = webhook.Verify(secretString, []byte(payload), signWebhook(t, []byte("wrong"), "msg-1", now, payload))
	var verr *webhook.WebhookVerificationError
	if !errors.As(err, &verr) || verr.Reason != webhook.WebhookVerificationFailureReasonNO_MATCHING_SIGNATURE {
		t.Fatalf("expected NO_MATCHING_SIGNATURE, got %v", err)
	}
}
