package vn.hitechcloud.saleai;

public class ValidationException extends SaleAIException {
    public ValidationException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
