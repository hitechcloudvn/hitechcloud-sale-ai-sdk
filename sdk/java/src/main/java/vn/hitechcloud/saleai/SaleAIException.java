package vn.hitechcloud.saleai;

public class SaleAIException extends Exception {
    private final int statusCode;
    private final String code;
    private final Object details;

    public SaleAIException(String message, int statusCode, String code, Object details) {
        super(message);
        this.statusCode = statusCode;
        this.code = code;
        this.details = details;
    }

    public int getStatusCode() { return statusCode; }
    public String getCode() { return code; }
    public Object getDetails() { return details; }
}
