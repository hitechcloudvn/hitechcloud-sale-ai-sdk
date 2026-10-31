package vn.hitechcloud.saleai.resources;

import vn.hitechcloud.saleai.HttpClient;

public abstract class BaseResource {
    protected final HttpClient http;

    protected BaseResource(HttpClient http) {
        this.http = http;
    }
}
