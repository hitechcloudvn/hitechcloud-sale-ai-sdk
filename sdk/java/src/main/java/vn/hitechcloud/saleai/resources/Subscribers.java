package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú */
public class Subscribers extends BaseResource {
    public Subscribers(HttpClient http) { super(http); }

    /** GET /api/v1/subscribers */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/subscribers", null, params, null);
    }

    /** GET /api/v1/subscribers/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/subscribers/" + id, null, null, null);
    }

    /** POST /api/v1/subscribers */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers", data, null, null);
    }

    /** PATCH /api/v1/subscribers/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/subscribers/" + id, data, null, null);
    }

    /** DELETE /api/v1/subscribers/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/subscribers/" + id, null, null, null);
    }

    /** POST /api/v1/subscribers/{id}/tags */
    public Map<String, Object> addTags(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/tags", data, null, null);
    }

    /** DELETE /api/v1/subscribers/{id}/tags/{tag} */
    public Map<String, Object> removeTag(String id, String tag) throws SaleAIException {
        return http.request("DELETE", "/api/v1/subscribers/" + id + "/tags/" + tag, null, null, null);
    }

    /** PATCH /api/v1/subscribers/{id}/custom-fields */
    public Map<String, Object> setCustomFields(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/subscribers/" + id + "/custom-fields", data, null, null);
    }

    /** POST /api/v1/subscribers/{id}/notes */
    public Map<String, Object> addNote(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/notes", data, null, null);
    }

    /** POST /api/v1/subscribers/{id}/block */
    public Map<String, Object> block(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/block", null, null, null);
    }

    /** POST /api/v1/subscribers/{id}/unblock */
    public Map<String, Object> unblock(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/unblock", null, null, null);
    }

    /** GET /api/v1/subscribers/{id}/messages */
    public Map<String, Object> listMessages(String id, Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/subscribers/" + id + "/messages", null, params, null);
    }

    /** POST /api/v1/subscribers/{id}/messages */
    public Map<String, Object> sendMessage(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/messages", data, null, null);
    }

    /** POST /api/v1/subscribers/{id}/mark-seen */
    public Map<String, Object> markSeen(String id) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/mark-seen", null, null, null);
    }

    /** POST /api/v1/subscribers/{id}/typing */
    public Map<String, Object> typing(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/subscribers/" + id + "/typing", data, null, null);
    }

}
