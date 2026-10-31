package vn.hitechcloud.saleai;

public class AuthenticationException extends SaleAIException {
    public AuthenticationException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
