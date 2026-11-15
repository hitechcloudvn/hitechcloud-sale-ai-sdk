// Basic example - HiTechCloud Sale AI Java SDK. / Ví dụ cơ bản - HiTechCloud Sale AI Java SDK.
//
// Run: / Cách chạy:
//   export SALE_AI_TOKEN=your-token
//   export SALE_AI_PAGE_ACCESS_TOKEN=page-token
import java.util.List;
import java.util.Map;

import vn.hitechcloud.saleai.SaleAI;
import vn.hitechcloud.saleai.WebhookVerifier;

public class BasicUsage {
    public static void main(String[] args) throws Exception {
        SaleAI client = new SaleAI(
            System.getenv().getOrDefault("SALE_AI_TOKEN", "YOUR_TOKEN"),
            System.getenv("SALE_AI_PAGE_ACCESS_TOKEN"),
            SaleAI.PRODUCTION_URL, 30, 3
        );

        // List bots / Liệt kê bot
        Map<String, Object> bots = client.bots.list(null);
        System.out.println(bots.get("data"));

        // Send a quick-reply message / Gửi tin quick-reply
        client.messages.sendQuickReply(
            "bot_7c1a9f2e", "sub_2a9b8c7d", "Ban muon xem gi?",
            List.of("Ao", "Quan"), null
        );

        // Verify the webhook signature / Xác minh chữ ký webhook
        byte[] raw = "{\"event\":\"order.created\"}".getBytes();
        String sig = WebhookVerifier.computeSignature(raw, "whsec_demo");
        System.out.println("signature ok: " + WebhookVerifier.verifySignature(raw, sig, "whsec_demo"));
    }
}
