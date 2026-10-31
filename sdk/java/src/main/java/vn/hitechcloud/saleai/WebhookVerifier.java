package vn.hitechcloud.saleai;

import javax.crypto.Mac;
import javax.crypto.spec.SecretKeySpec;
import java.nio.charset.StandardCharsets;

public final class WebhookVerifier {
    public static final String SIGNATURE_HEADER = "X-SaleAI-Signature";

    private WebhookVerifier() {}

    public static String computeSignature(byte[] rawBody, String secret) {
        try {
            Mac mac = Mac.getInstance("HmacSHA256");
            mac.init(new SecretKeySpec(secret.getBytes(StandardCharsets.UTF_8), "HmacSHA256"));
            byte[] digest = mac.doFinal(rawBody);
            StringBuilder sb = new StringBuilder("sha256=");
            for (byte b : digest) sb.append(String.format("%02x", b));
            return sb.toString();
        } catch (Exception e) {
            throw new RuntimeException("Unable to compute signature", e);
        }
    }

    public static boolean verifySignature(byte[] rawBody, String signature, String secret) {
        if (signature == null || signature.isEmpty()) return false;
        String expected = computeSignature(rawBody, secret);
        return constantTimeEquals(expected, signature);
    }

    private static boolean constantTimeEquals(String a, String b) {
        if (a.length() != b.length()) return false;
        int result = 0;
        for (int i = 0; i < a.length(); i++) result |= a.charAt(i) ^ b.charAt(i);
        return result == 0;
    }
}
