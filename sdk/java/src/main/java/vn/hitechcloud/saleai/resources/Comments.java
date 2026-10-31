package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Thao tác bình luận trên kênh (comments) */
public class Comments extends BaseResource {
    public Comments(HttpClient http) { super(http); }

    /** POST /api/v1/comments/{comment_id}/hide */
    public Map<String, Object> hide(String commentId, String pageAccessToken) throws SaleAIException {
        Map<String, String> headers = new HashMap<>();
        if (pageAccessToken != null) headers.put("X-Page-Access-Token", pageAccessToken);
        return http.request("POST", "/api/v1/comments/" + commentId + "/hide", null, null, headers);
    }

    /** POST /api/v1/comments/{comment_id}/like */
    public Map<String, Object> like(String commentId, String pageAccessToken) throws SaleAIException {
        Map<String, String> headers = new HashMap<>();
        if (pageAccessToken != null) headers.put("X-Page-Access-Token", pageAccessToken);
        return http.request("POST", "/api/v1/comments/" + commentId + "/like", null, null, headers);
    }

}
