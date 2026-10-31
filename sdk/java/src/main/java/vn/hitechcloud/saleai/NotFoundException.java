package vn.hitechcloud.saleai;

public class NotFoundException extends SaleAIException {
    public NotFoundException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
