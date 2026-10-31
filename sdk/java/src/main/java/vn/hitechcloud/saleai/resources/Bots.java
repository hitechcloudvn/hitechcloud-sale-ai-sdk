package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Bot & kênh kết nối (Facebook/Instagram/Zalo/Website) */
public class Bots extends BaseResource {
    public Bots(HttpClient http) { super(http); }

    /** GET /api/v1/bots */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/bots", null, params, null);
    }

    /** GET /api/v1/bots/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/bots/" + id, null, null, null);
    }

    /** POST /api/v1/bots */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/bots", data, null, null);
    }

    /** PATCH /api/v1/bots/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/bots/" + id, data, null, null);
    }

    /** DELETE /api/v1/bots/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/bots/" + id, null, null, null);
    }

    /** POST /api/v1/bots/{bot_id}/connect */
    public Map<String, Object> connect(String botId, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/bots/" + botId + "/connect", data, null, null);
    }

    /** POST /api/v1/bots/{bot_id}/disconnect */
    public Map<String, Object> disconnect(String botId) throws SaleAIException {
        return http.request("POST", "/api/v1/bots/" + botId + "/disconnect", null, null, null);
    }

    /** PATCH /api/v1/bots/{bot_id}/settings */
    public Map<String, Object> updateSettings(String botId, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/bots/" + botId + "/settings", data, null, null);
    }

    /** POST /api/v1/bots/{bot_id}/access-token/refresh */
    public Map<String, Object> refreshAccessToken(String botId, Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/bots/" + botId + "/access-token/refresh", data, null, null);
    }

}
