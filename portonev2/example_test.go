package portonev2_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/Gliese436/portone-go-client/portonev2"
	"github.com/Gliese436/portone-go-client/portonev2/payment"
	"github.com/Gliese436/portone-go-client/portonev2/webhook"
)

func Example() {
	ctx := context.Background()
	client := portonev2.NewClient(
		os.Getenv("PORTONE_API_SECRET"),
		portonev2.WithStoreID(os.Getenv("PORTONE_STORE_ID")),
	)

	p, err := client.Payment().GetPayment(ctx, &payment.GetPaymentOptions{PaymentId: "order-1"})
	if err != nil {
		if apiErr, ok := portonev2.AsAPIError(err); ok {
			var e payment.GetPaymentError
			_ = apiErr.Decode(&e)
			if e.PaymentNotFoundError != nil {
				fmt.Println("결제 건 없음")
			}
		}
		return
	}
	switch v := p.Value().(type) {
	case *payment.PaidPayment:
		fmt.Println("결제 완료", v.Amount.Total)
	case *payment.CancelledPayment:
		fmt.Println("취소됨")
	}
}

func Example_webhook() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		wh, err := webhook.Verify(os.Getenv("PORTONE_WEBHOOK_SECRET"), body, r.Header)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if wh.WebhookTransactionPaid != nil {
			fmt.Println(wh.WebhookTransactionPaid.Data.PaymentId)
		}
	})
}
