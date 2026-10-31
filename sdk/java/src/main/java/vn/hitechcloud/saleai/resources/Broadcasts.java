package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Gửi hàng loạt (broadcasts) */
public class Broadcasts extends BaseResource {
    public Broadcasts(HttpClient http) { super(http); }

    /** GET /api/v1/broadcasts */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/broadcasts", null, params, null);
    }

    /** GET /api/v1/broadcasts/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/broadcasts/" + id, null, null, null);
    }

    /** POST /api/v1/broadcasts */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/broadcasts", data, null, null);
    }

    /** PATCH /api/v1/broadcasts/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/broadcasts/" + id, data, null, null);
    }

    /** DELETE /api/v1/broadcasts/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/broadcasts/" + id, null, null, null);
    }

    /** POST /api/v1/broadcasts/{id}/schedule */
    public Map<String, Object> schedule(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/broadcasts/" + id + "/schedule", data, null, null);
    }

    /** POST /api/v1/broadcasts/{id}/send */
    public Map<String, Object> send(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/broadcasts/" + id + "/send", null, null, null);
    }

    /** GET /api/v1/broadcasts/{id}/status */
    public Map<String, Object> status(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/broadcasts/" + id + "/status", null, null, null);
    }

    /** POST /api/v1/broadcasts/{id}/audience/preview */
    public Map<String, Object> previewAudience(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/broadcasts/" + id + "/audience/preview", data, null, null);
    }

    /** POST /api/v1/broadcasts/{id}/cancel */
    public Map<String, Object> cancel(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/broadcasts/" + id + "/cancel", null, null, null);
    }

}
