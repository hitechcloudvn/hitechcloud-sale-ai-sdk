package vn.hitechcloud.saleai;

public class ServerException extends SaleAIException {
    public ServerException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
