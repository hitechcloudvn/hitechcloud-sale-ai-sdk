# HiTechCloud Sale AI - Go SDK

SDK Go chính thức cho **HiTechCloud Sale AI**.

- Module: `github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go`
- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`

## Cài đặt

```bash
go get github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go
```

## Bắt đầu nhanh

```go
package main

import (
    "fmt"
    saleai "github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go"
)

func main() {
    client := saleai.NewClient(saleai.WithToken("YOUR_TOKEN"))

    bots, err := client.Bots.List(nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(bots)

    _, _ = client.Messages.SendText("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!", "PAGE_ACCESS_TOKEN")
}
```

## Webhook

```go
ok := saleai.VerifyWebhookSignature(rawBody, signatureHeader, "whsec_...")
```

MIT License.
