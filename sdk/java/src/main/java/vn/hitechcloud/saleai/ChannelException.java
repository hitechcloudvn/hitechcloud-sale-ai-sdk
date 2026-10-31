package vn.hitechcloud.saleai;

public class ChannelException extends SaleAIException {
    public ChannelException(String message, int statusCode, String code, Object details) {
        super(message, statusCode, code, details);
    }
}
