package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Nhãn (tags) */
public class Tags extends BaseResource {
    public Tags(HttpClient http) { super(http); }

    /** GET /api/v1/tags */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/tags", null, params, null);
    }

    /** GET /api/v1/tags/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/tags/" + id, null, null, null);
    }

    /** POST /api/v1/tags */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/tags", data, null, null);
    }

    /** PATCH /api/v1/tags/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/tags/" + id, data, null, null);
    }

    /** DELETE /api/v1/tags/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/tags/" + id, null, null, null);
    }

}
