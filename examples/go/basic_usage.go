// Basic example - HiTechCloud Sale AI Go SDK. / Ví dụ cơ bản - HiTechCloud Sale AI Go SDK.
//
// Run: / Cách chạy:
//
//	export SALE_AI_TOKEN=your-token
//	export SALE_AI_PAGE_ACCESS_TOKEN=page-token
package main

import (
	"fmt"

	saleai "github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go"
)

func main() {
	client := saleai.NewClient(
		saleai.WithToken("YOUR_TOKEN"),
		saleai.WithPageAccessToken("PAGE_ACCESS_TOKEN"),
	)

	// List bots / Liệt kê bot
	bots, err := client.Bots.List(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bots)

	// Send a text message / Gửi tin nhắn văn bản
	_, _ = client.Messages.SendText("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!", "")

	// Overview analytics report / Báo cáo tổng quan
	report, _ := client.Analytics.Overview(map[string]string{"bot_id": "bot_7c1a9f2e"})
	fmt.Println(report)

	// Verify the webhook signature / Xác minh chữ ký webhook
	raw := []byte(`{"event":"order.created"}`)
	sig := saleai.ComputeSignature(raw, "whsec_demo")
	fmt.Println("signature ok:", saleai.VerifyWebhookSignature(raw, sig, "whsec_demo"))
}
