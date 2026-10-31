package vn.hitechcloud.saleai.resources;

import java.util.Map;
import java.util.HashMap;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Sản phẩm (products) */
public class Products extends BaseResource {
    public Products(HttpClient http) { super(http); }

    /** GET /api/v1/products */
    public Map<String, Object> list(Map<String, String> params) throws SaleAIException {
        return http.request("GET", "/api/v1/products", null, params, null);
    }

    /** GET /api/v1/products/{id} */
    public Map<String, Object> get(String id) throws SaleAIException {
        return http.request("GET", "/api/v1/products/" + id, null, null, null);
    }

    /** POST /api/v1/products */
    public Map<String, Object> create(Map<String, Object> data) throws SaleAIException {
        return http.request("POST", "/api/v1/products", data, null, null);
    }

    /** PATCH /api/v1/products/{id} */
    public Map<String, Object> update(String id, Map<String, Object> data) throws SaleAIException {
        return http.request("PATCH", "/api/v1/products/" + id, data, null, null);
    }

    /** DELETE /api/v1/products/{id} */
    public Map<String, Object> delete(String id) throws SaleAIException {
        return http.request("DELETE", "/api/v1/products/" + id, null, null, null);
    }

}
