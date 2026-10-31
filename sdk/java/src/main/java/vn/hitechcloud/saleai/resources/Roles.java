package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Vai trò & quyền (roles) */
public class Roles extends BaseResource {
    public Roles(HttpClient http) { super(http); }

    /** GET /api/v1/roles */
    public Map<String, Object> list() throws SaleAIException {
        return http.request("GET", "/api/v1/roles", null, null, null);
    }

    /** GET /api/v1/roles/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/roles/" + id, null, null, null);
    }

    /** POST /api/v1/roles */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/roles", data, null, null);
    }

    /** PATCH /api/v1/roles/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/roles/" + id, data, null, null);
    }

    /** DELETE /api/v1/roles/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/roles/" + id, null, null, null);
    }

}
