package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Công cụ tăng trưởng (widgets) */
public class Widgets extends BaseResource {
    public Widgets(HttpClient http) { super(http); }

    /** GET /api/v1/widgets */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/widgets", null, params, null);
    }

    /** GET /api/v1/widgets/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/widgets/" + id, null, null, null);
    }

    /** POST /api/v1/widgets */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/widgets", data, null, null);
    }

    /** PATCH /api/v1/widgets/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/widgets/" + id, data, null, null);
    }

    /** DELETE /api/v1/widgets/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/widgets/" + id, null, null, null);
    }

}
