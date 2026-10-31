package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Cài đặt tài khoản (settings) */
public class Settings extends BaseResource {
    public Settings(HttpClient http) { super(http); }

    /** GET /api/v1/settings */
    public Map<String, Object> get() throws SaleAIException {
        return http.request("GET", "/api/v1/settings", null, null, null);
    }

    /** PATCH /api/v1/settings */
    public Map<String, Object> update(Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/settings", data, null, null);
    }

}
