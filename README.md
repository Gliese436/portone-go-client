# 포트원 Golang v2 클라이언트 라이브러리

## 상세
- 포트원 측에서 제공하는 REST API v2.0.0의 OpenAPI 3.0.3 JSON 파일에 담긴 사양을 바탕으로 함
- oapi-codegen을 사용하여 Go struct 및 HTTP 요청 / 응답을 포함한 종합 클라이언트를 생성
- 통합 자동화 및 업데이트 시 생성되도록 할 수도 있을 것이나, 이는 현재로서는 TODO
- 서브모듈로 포트원 저장소에서 API 사양 연동



## 코드 생성
```console
$ chmod +x generate-client.sh
$ ./generate-client.sh
```
**generate-type.sh 파일은 실험용으로, 필요 시 요청 / 응답 JSON에 대응되는 구조체만 생성하기 위함임**

**사용 시에는 generate-client.sh 파일만 사용하더라도 구조체 역시 생성됨**

## 사용법
### 설치
```
go get github.com/Gliese436/portone-go-client
```
### 예시
```go
yourStoreId := os.Getenv("PORTONE_STORE_ID")

client, err := portonev2client.NewClient(
  os.Getenv("PORTONE_URL"),
  portonev2client.WithRequestEditorFn(func(ctx context.Context, req *http.Request) error {
			req.Header.Add("Authorization", "Portone "+os.Getenv("PORTONE_SECRET_KEY"))
			return nil
  }),
)
if err != nil {
  log.Panic(err)
}
response, err := client.GetPayment(
  context.Background(),
  "YOUR_PAYMENT_ID",
  &portonev2client.GetPaymentParams{
    StoreId: &yourStoreId,
  },
)
// ... //
```
