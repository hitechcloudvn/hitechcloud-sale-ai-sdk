package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Ý định AI (intents) */
public class Intents extends BaseResource {
    public Intents(HttpClient http) { super(http); }

    /** GET /api/v1/ai/intents */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/ai/intents", null, params, null);
    }

    /** GET /api/v1/ai/intents/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/ai/intents/" + id, null, null, null);
    }

    /** POST /api/v1/ai/intents */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/ai/intents", data, null, null);
    }

    /** PATCH /api/v1/ai/intents/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/ai/intents/" + id, data, null, null);
    }

    /** DELETE /api/v1/ai/intents/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/ai/intents/" + id, null, null, null);
    }

}
