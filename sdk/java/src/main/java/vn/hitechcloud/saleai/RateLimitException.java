package vn.hitechcloud.saleai;

public class RateLimitException extends SaleAIException {
    private final int retryAfter;

    public RateLimitException(String message, int retryAfter, String code, Object details) {
        super(message, 429, code, details);
        this.retryAfter = retryAfter;
    }

    public int getRetryAfter() { return retryAfter; }
}
