package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Quy tắc trả lời bình luận (comment rules) */
public class CommentRules extends BaseResource {
    public CommentRules(HttpClient http) { super(http); }

    /** GET /api/v1/comment-rules */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/comment-rules", null, params, null);
    }

    /** GET /api/v1/comment-rules/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/comment-rules/" + id, null, null, null);
    }

    /** POST /api/v1/comment-rules */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/comment-rules", data, null, null);
    }

    /** PATCH /api/v1/comment-rules/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/comment-rules/" + id, data, null, null);
    }

    /** DELETE /api/v1/comment-rules/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/comment-rules/" + id, null, null, null);
    }

}
