package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Trợ lý AI bán hàng (config / train / suggest / handover) */
public class Ai extends BaseResource {
    public Ai(HttpClient http) { super(http); }

    /** PATCH /api/v1/ai/config */
    public Map<String, Object> updateConfig(Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/ai/config", data, null, null);
    }

    /** POST /api/v1/ai/train */
    public Map<String, Object> train(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/ai/train", data, null, null);
    }

    /** POST /api/v1/ai/suggest-reply */
    public Map<String, Object> suggestReply(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/ai/suggest-reply", data, null, null);
    }

    /** POST /api/v1/ai/handover */
    public Map<String, Object> handover(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/ai/handover", data, null, null);
    }

}
