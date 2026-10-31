package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Người dùng / nhân sự (team) */
public class Users extends BaseResource {
    public Users(HttpClient http) { super(http); }

    /** GET /api/v1/users */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/users", null, params, null);
    }

    /** GET /api/v1/users/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/users/" + id, null, null, null);
    }

    /** POST /api/v1/users */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/users", data, null, null);
    }

    /** PATCH /api/v1/users/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/users/" + id, data, null, null);
    }

    /** DELETE /api/v1/users/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/users/" + id, null, null, null);
    }

}
