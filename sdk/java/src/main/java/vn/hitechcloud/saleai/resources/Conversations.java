package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Hộp thư đa kênh (conversations / inbox) */
public class Conversations extends BaseResource {
    public Conversations(HttpClient http) { super(http); }

    /** GET /api/v1/conversations */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/conversations", null, params, null);
    }

    /** GET /api/v1/conversations/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/conversations/" + id, null, null, null);
    }

    /** POST /api/v1/conversations/{id}/assign */
    public Map<String, Object> assign(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/conversations/" + id + "/assign", data, null, null);
    }

    /** POST /api/v1/conversations/{id}/notes */
    public Map<String, Object> addNote(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/conversations/" + id + "/notes", data, null, null);
    }

    /** POST /api/v1/conversations/{id}/labels */
    public Map<String, Object> addLabels(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/conversations/" + id + "/labels", data, null, null);
    }

    /** POST /api/v1/conversations/{id}/close */
    public Map<String, Object> close(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/conversations/" + id + "/close", null, null, null);
    }

}
