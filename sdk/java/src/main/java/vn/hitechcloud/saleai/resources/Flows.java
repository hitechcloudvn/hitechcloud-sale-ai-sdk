package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Kịch bản hội thoại (flows) và khối (blocks) */
public class Flows extends BaseResource {
    public Flows(HttpClient http) { super(http); }

    /** GET /api/v1/flows */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/flows", null, params, null);
    }

    /** GET /api/v1/flows/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/flows/" + id, null, null, null);
    }

    /** POST /api/v1/flows */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/flows", data, null, null);
    }

    /** PATCH /api/v1/flows/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/flows/" + id, data, null, null);
    }

    /** DELETE /api/v1/flows/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/flows/" + id, null, null, null);
    }

    /** POST /api/v1/flows/{id}/publish */
    public Map<String, Object> publish(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/flows/" + id + "/publish", null, null, null);
    }

    /** POST /api/v1/flows/{id}/duplicate */
    public Map<String, Object> duplicate(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/flows/" + id + "/duplicate", null, null, null);
    }

    /** POST /api/v1/flows/{id}/trigger */
    public Map<String, Object> trigger(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/flows/" + id + "/trigger", data, null, null);
    }

    /** GET /api/v1/flows/{id}/blocks */
    public Map<String, Object> listBlocks(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/flows/" + id + "/blocks", null, null, null);
    }

    /** POST /api/v1/flows/{id}/blocks */
    public Map<String, Object> addBlock(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/flows/" + id + "/blocks", data, null, null);
    }

    /** PATCH /api/v1/flows/{id}/blocks/{block_id} */
    public Map<String, Object> updateBlock(String id, String blockId, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/flows/" + id + "/blocks/" + blockId, data, null, null);
    }

    /** DELETE /api/v1/flows/{id}/blocks/{block_id} */
    public Map<String, Object> deleteBlock(String id, String blockId) throws SaleAIException {
        return http.request("DELETE", "/api/v1/flows/" + id + "/blocks/" + blockId, null, null, null);
    }

}
