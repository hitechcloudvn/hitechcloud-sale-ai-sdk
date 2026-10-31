package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Đăng ký webhook (webhooks) */
public class Webhooks extends BaseResource {
    public Webhooks(HttpClient http) { super(http); }

    /** GET /api/v1/webhooks */
    public Map<String, Object> list() throws SaleAIException {
        return http.request("GET", "/api/v1/webhooks", null, null, null);
    }

    /** GET /api/v1/webhooks/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/webhooks/" + id, null, null, null);
    }

    /** POST /api/v1/webhooks */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/webhooks", data, null, null);
    }

    /** PATCH /api/v1/webhooks/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/webhooks/" + id, data, null, null);
    }

    /** DELETE /api/v1/webhooks/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/webhooks/" + id, null, null, null);
    }

}
