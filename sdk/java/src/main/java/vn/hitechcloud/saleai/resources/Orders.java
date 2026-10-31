package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Đơn hàng (orders) */
public class Orders extends BaseResource {
    public Orders(HttpClient http) { super(http); }

    /** GET /api/v1/orders */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/orders", null, params, null);
    }

    /** GET /api/v1/orders/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/orders/" + id, null, null, null);
    }

    /** POST /api/v1/orders */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/orders", data, null, null);
    }

    /** PATCH /api/v1/orders/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/orders/" + id, data, null, null);
    }

    /** DELETE /api/v1/orders/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/orders/" + id, null, null, null);
    }

    /** POST /api/v1/orders/from-conversation */
    public Map<String, Object> createFromConversation(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/orders/from-conversation", data, null, null);
    }

    /** POST /api/v1/orders/{id}/status */
    public Map<String, Object> updateStatus(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/orders/" + id + "/status", data, null, null);
    }

}
