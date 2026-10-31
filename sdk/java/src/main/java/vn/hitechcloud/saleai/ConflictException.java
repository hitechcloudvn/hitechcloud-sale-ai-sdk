package vn.hitechcloud.saleai;

public class ConflictException extends SaleAIException {
    public ConflictException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
