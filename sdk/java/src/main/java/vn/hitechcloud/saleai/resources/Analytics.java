package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Báo cáo & phân tích (analytics) */
public class Analytics extends BaseResource {
    public Analytics(HttpClient http) { super(http); }

    /** GET /api/v1/analytics/overview */
    public Map<String, Object> overview(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/analytics/overview", null, params, null);
    }

    /** GET /api/v1/analytics/bots/{bot_id} */
    public Map<String, Object> bot(String botId, Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/analytics/bots/" + botId, null, params, null);
    }

    /** GET /api/v1/analytics/flows/{id} */
    public Map<String, Object> flow(String id, Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/analytics/flows/" + id, null, params, null);
    }

    /** GET /api/v1/analytics/broadcasts/{id} */
    public Map<String, Object> broadcast(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/analytics/broadcasts/" + id, null, null, null);
    }

    /** GET /api/v1/analytics/subscriber-growth */
    public Map<String, Object> subscriberGrowth(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/analytics/subscriber-growth", null, params, null);
    }

}
