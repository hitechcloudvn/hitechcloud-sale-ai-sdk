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

        Map<String, Object> bots = client.bots.list(null);
        System.out.println(bots.get("data"));

        client.messages.sendQuickReply(
            "bot_7c1a9f2e", "sub_2a9b8c7d", "Ban muon xem gi?",
            List.of("Ao", "Quan"), null
        );

        byte[] raw = "{\"event\":\"order.created\"}".getBytes();
        String sig = WebhookVerifier.computeSignature(raw, "whsec_demo");
        System.out.println("signature ok: " + WebhookVerifier.verifySignature(raw, sig, "whsec_demo"));
    }
}
