package main

import (
	"fmt"

	saleai "github.com/hitechcloud-vietnam/hitechcloud-sale-ai-sdk/go"
)

func main() {
	client := saleai.NewClient(
		saleai.WithToken("YOUR_TOKEN"),
		saleai.WithPageAccessToken("PAGE_ACCESS_TOKEN"),
	)

	bots, err := client.Bots.List(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(bots)

	_, _ = client.Messages.SendText("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!", "")

	report, _ := client.Analytics.Overview(map[string]string{"bot_id": "bot_7c1a9f2e"})
	fmt.Println(report)

	raw := []byte(`{"event":"order.created"}`)
	sig := saleai.ComputeSignature(raw, "whsec_demo")
	fmt.Println("signature ok:", saleai.VerifyWebhookSignature(raw, sig, "whsec_demo"))
}
