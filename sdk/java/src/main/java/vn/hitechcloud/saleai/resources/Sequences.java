package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Chuỗi tin hẹn giờ (sequences) và các bước */
public class Sequences extends BaseResource {
    public Sequences(HttpClient http) { super(http); }

    /** GET /api/v1/sequences */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/sequences", null, params, null);
    }

    /** GET /api/v1/sequences/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/sequences/" + id, null, null, null);
    }

    /** POST /api/v1/sequences */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/sequences", data, null, null);
    }

    /** PATCH /api/v1/sequences/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/sequences/" + id, data, null, null);
    }

    /** DELETE /api/v1/sequences/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/sequences/" + id, null, null, null);
    }

    /** GET /api/v1/sequences/{id}/steps */
    public Map<String, Object> listSteps(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/sequences/" + id + "/steps", null, null, null);
    }

    /** POST /api/v1/sequences/{id}/steps */
    public Map<String, Object> addStep(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/sequences/" + id + "/steps", data, null, null);
    }

    /** PATCH /api/v1/sequences/{id}/steps/{step_id} */
    public Map<String, Object> updateStep(String id, String stepId, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/sequences/" + id + "/steps/" + stepId, data, null, null);
    }

    /** DELETE /api/v1/sequences/{id}/steps/{step_id} */
    public Map<String, Object> deleteStep(String id, String stepId) throws SaleAIException {
        return http.request("DELETE", "/api/v1/sequences/" + id + "/steps/" + stepId, null, null, null);
    }

    /** POST /api/v1/sequences/{id}/subscribe */
    public Map<String, Object> subscribe(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/sequences/" + id + "/subscribe", data, null, null);
    }

    /** POST /api/v1/sequences/{id}/unsubscribe */
    public Map<String, Object> unsubscribe(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/sequences/" + id + "/unsubscribe", data, null, null);
    }

}
