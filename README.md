# 포트원 Golang v2 클라이언트 라이브러리

포트원 REST API V2 Go 클라이언트입니다. API 사양은 서브모듈 [`server-sdk`](https://github.com/portone-io/server-sdk)의 `codegen/openapi.json`을 따릅니다.

두 가지 클라이언트를 제공합니다.

| 패키지 | 생성 방식 | 특징 |
| --- | --- | --- |
| `portonev2` (권장) | `codegen/main.ts` — server-sdk 공식 codegen 파서 + Go 출력 | JS/Python/Kotlin SDK와 동일한 구조·동작, 카테고리별 패키지, oneOf/enum 타입, 웹훅 검증 |
| `portonev2client` | [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) | OpenAPI 그대로의 저수준 클라이언트 (`portonev2client.gen.go`) |

## 설치

```console
go get github.com/Gliese436/portone-go-client
```

## 사용법 (`portonev2`)

```go
import (
	"github.com/Gliese436/portone-go-client/portonev2"
	"github.com/Gliese436/portone-go-client/portonev2/payment"
)

client := portonev2.NewClient(
	os.Getenv("PORTONE_API_SECRET"),
	portonev2.WithStoreID(os.Getenv("PORTONE_STORE_ID")), // 하위 상점 사용 시
)

p, err := client.Payment().GetPayment(ctx, &payment.GetPaymentOptions{PaymentId: "order-1"})
if err != nil {
	return err
}
switch v := p.Value().(type) { // 결제 건(Payment)은 status 로 구분되는 oneOf 타입입니다.
case *payment.PaidPayment:
	fmt.Println("결제 완료", v.Amount.Total)
case *payment.CancelledPayment:
	fmt.Println("취소됨")
}
```

- 클라이언트 구조는 JS SDK와 같습니다: `client.B2b().Counterparty()`, `client.Payment().BillingKey()`, `client.Platform().Partner()`, `client.Reconciliation()` 등.
- 모든 파라미터(경로·쿼리·바디)는 메서드별 `XxxOptions` 구조체 하나로 전달합니다. 선택 값은 포인터이며 `nil`이면 전송하지 않습니다.
- `StoreId`를 지정하지 않으면 `WithStoreID`로 설정한 값을 사용합니다.
- GET/DELETE 요청의 바디는 JS SDK와 같이 `requestBody` 쿼리 파라미터(JSON)로 전송됩니다.
- oneOf 타입은 판별자 값(`Type`/`Status` 등)과 variant 포인터 필드를 가진 구조체입니다. 알 수 없는 variant는 `Unrecognized`에 원본 JSON이 보관됩니다.

### 오류 처리

2xx가 아닌 응답은 `*portonev2.APIError`로 반환됩니다. 메서드 문서에 적힌 오류 타입으로 디코딩하면 오류별 필드를 확인할 수 있습니다.

```go
if apiErr, ok := portonev2.AsAPIError(err); ok {
	var e payment.GetPaymentError
	_ = apiErr.Decode(&e)
	if e.PaymentNotFoundError != nil { /* ... */ }
}
```

### 웹훅 검증

```go
wh, err := webhook.Verify(os.Getenv("PORTONE_WEBHOOK_SECRET"), body, r.Header)
if err != nil {
	return err
}
if wh.WebhookTransactionPaid != nil {
	paymentID := wh.WebhookTransactionPaid.Data.PaymentId
	// ...
}
```

## 코드 생성 / 최신화

```console
$ git submodule update --remote server-sdk   # 서버 SDK 최신화
$ ./generate-sdk.sh                          # portonev2/*.gen.go 재생성 (deno, gofmt 필요)
$ ./generate-client.sh                       # portonev2client.gen.go 재생성 (oapi-codegen v2.5.1)
$ go test ./...
```

- `generate-sdk.sh`는 `server-sdk/codegen/parser`를 그대로 import 하므로 엔티티 배치·파라미터 구성이 공식 SDK와 항상 일치합니다. 서브모듈은 수정하지 않습니다.
- 생성 파일(`*.gen.go`)은 직접 수정하지 마세요. 수작업 코드는 `portonev2/client.go`(HTTP), `errors.go`, `bank.go`(상류에서 제외된 은행 조회 API), `webhook/verify.go`, `webhook/errors.go`, `internal/jsonutil` 입니다.
- `generate-type.sh`는 실험용으로, 요청/응답 JSON에 대응되는 구조체만 생성합니다.

### oapi-codegen 클라이언트 예시 (`portonev2client`)

```go
client, err := portonev2client.NewClient(
	"https://api.portone.io",
	portonev2client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Authorization", "PortOne "+os.Getenv("PORTONE_API_SECRET"))
		return nil
	}),
)
storeID := os.Getenv("PORTONE_STORE_ID")
response, err := client.GetPayment(ctx, "YOUR_PAYMENT_ID", &portonev2client.GetPaymentParams{StoreId: &storeID})
```
