package vn.hitechcloud.saleai;

public class AuthorizationException extends SaleAIException {
    public AuthorizationException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
