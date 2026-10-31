package vn.hitechcloud.saleai;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.lang.reflect.Type;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.Map;

public class HttpClient {
    private final String baseUrl;
    private final int timeout;
    private final int maxRetries;
    private String token;
    private String pageAccessToken;
    private final Gson gson = new Gson();

    public HttpClient(String baseUrl, String token, String pageAccessToken, int timeout, int maxRetries) {
        this.baseUrl = baseUrl.replaceAll("/$", "");
        this.token = token;
        this.pageAccessToken = pageAccessToken;
        this.timeout = timeout;
        this.maxRetries = maxRetries;
    }

    public void setToken(String token) { this.token = token; }
    public void clearToken() { this.token = null; }
    public void setPageAccessToken(String token) { this.pageAccessToken = token; }

    @SuppressWarnings("unchecked")
    public Map<String, Object> request(String method, String path, Object body,
                                       Map<String, String> params, Map<String, String> headers) throws SaleAIException {
        Exception lastError = null;

        for (int attempt = 0; attempt <= maxRetries; attempt++) {
            try {
                StringBuilder urlStr = new StringBuilder(baseUrl + path);
                if (params != null && !params.isEmpty()) {
                    boolean first = true;
                    for (Map.Entry<String, String> e : params.entrySet()) {
                        if (e.getValue() == null) continue;
                        urlStr.append(first ? "?" : "&");
                        urlStr.append(URLEncoder.encode(e.getKey(), StandardCharsets.UTF_8));
                        urlStr.append("=");
                        urlStr.append(URLEncoder.encode(e.getValue(), StandardCharsets.UTF_8));
                        first = false;
                    }
                }

                HttpURLConnection conn = (HttpURLConnection) new URL(urlStr.toString()).openConnection();
                conn.setRequestMethod(method);
                conn.setConnectTimeout(timeout * 1000);
                conn.setReadTimeout(timeout * 1000);
                conn.setRequestProperty("Content-Type", "application/json");
                conn.setRequestProperty("Accept", "application/json");
                conn.setRequestProperty("User-Agent", "HiTechCloud-SaleAI-Java/" + SaleAI.VERSION);
                if (token != null) conn.setRequestProperty("Authorization", "Bearer " + token);
                if (pageAccessToken != null) conn.setRequestProperty("X-Page-Access-Token", pageAccessToken);
                if (headers != null) {
                    for (Map.Entry<String, String> e : headers.entrySet()) {
                        if (e.getValue() != null) conn.setRequestProperty(e.getKey(), e.getValue());
                    }
                }

                if (body != null && ("POST".equals(method) || "PUT".equals(method) || "PATCH".equals(method))) {
                    conn.setDoOutput(true);
                    try (OutputStream os = conn.getOutputStream()) {
                        os.write(gson.toJson(body).getBytes(StandardCharsets.UTF_8));
                    }
                }

                int status = conn.getResponseCode();

                if (status == 429) {
                    int retryAfter = parseRetryAfter(conn);
                    if (attempt < maxRetries) { Thread.sleep(retryAfter * 1000L); continue; }
                }
                if (status >= 500 && attempt < maxRetries) {
                    Thread.sleep((long) Math.pow(2, attempt) * 1000);
                    continue;
                }

                InputStream is = status >= 400 ? conn.getErrorStream() : conn.getInputStream();
                String responseBody = readStream(is);
                Type mapType = new TypeToken<Map<String, Object>>() {}.getType();
                Map<String, Object> result = responseBody.isEmpty() ? null : gson.fromJson(responseBody, mapType);

                if (status >= 400) {
                    throw mapError(status, result, parseRetryAfter(conn));
                }
                return result;
            } catch (SaleAIException e) {
                throw e;
            } catch (Exception e) {
                lastError = e;
                if (attempt < maxRetries) {
                    try { Thread.sleep((long) Math.pow(2, attempt) * 1000); }
                    catch (InterruptedException ie) { Thread.currentThread().interrupt(); }
                    continue;
                }
            }
        }
        throw new SaleAIException("Request failed: " + (lastError != null ? lastError.getMessage() : "unknown"), 0, null, null);
    }

    public Map<String, Object> get(String path, Map<String, String> params, Map<String, String> headers) throws SaleAIException {
        return request("GET", path, null, params, headers);
    }

    public Map<String, Object> post(String path, Object body, Map<String, String> headers) throws SaleAIException {
        return request("POST", path, body, null, headers);
    }

    public Map<String, Object> patch(String path, Object body, Map<String, String> headers) throws SaleAIException {
        return request("PATCH", path, body, null, headers);
    }

    public Map<String, Object> delete(String path, Map<String, String> headers) throws SaleAIException {
        return request("DELETE", path, null, null, headers);
    }

    private int parseRetryAfter(HttpURLConnection conn) {
        String v = conn.getHeaderField("Retry-After");
        try { return v != null ? Integer.parseInt(v) : 2; } catch (NumberFormatException e) { return 2; }
    }

    private String readStream(InputStream is) throws IOException {
        if (is == null) return "";
        StringBuilder sb = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(is, StandardCharsets.UTF_8))) {
            String line;
            while ((line = reader.readLine()) != null) sb.append(line);
        }
        return sb.toString();
    }

    @SuppressWarnings("unchecked")
    private SaleAIException mapError(int status, Map<String, Object> body, int retryAfter) {
        String message = "API error";
        String code = null;
        Object details = null;
        if (body != null && body.get("error") instanceof Map) {
            Map<String, Object> err = (Map<String, Object>) body.get("error");
            if (err.get("message") != null) message = err.get("message").toString();
            if (err.get("code") != null) code = err.get("code").toString();
            details = err.get("details");
        }
        switch (status) {
            case 401: return new AuthenticationException(message, status, code, details);
            case 403: return new AuthorizationException(message, status, code, details);
            case 404: return new NotFoundException(message, status, code, details);
            case 409: return new ConflictException(message, status, code, details);
            case 422: return new ValidationException(message, status, code, details);
            case 429: return new RateLimitException(message, retryAfter, code, details);
            case 502: return new ChannelException(message, status, code, details);
            default:
                if (status >= 500) return new ServerException(message, status, code, details);
                return new SaleAIException(message, status, code, details);
        }
    }
}
