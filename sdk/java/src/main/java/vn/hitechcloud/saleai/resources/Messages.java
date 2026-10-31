package vn.hitechcloud.saleai.resources;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import vn.hitechcloud.saleai.HttpClient;
import vn.hitechcloud.saleai.SaleAIException;

/** Gửi tin nhắn (text/image/file/quick-reply/button/carousel). */
public class Messages extends BaseResource {
    public Messages(HttpClient http) { super(http); }

    private Map<String, Object> sendRaw(String botId, Map<String, Object> message, String pat) throws SaleAIException {
        Map<String, String> headers = new HashMap<>();
        if (pat != null) headers.put("X-Page-Access-Token", pat);
        return http.request("POST", "/api/v1/bots/" + botId + "/messages", message, null, headers);
    }

    public Map<String, Object> send(String botId, Map<String, Object> message, String pageAccessToken) throws SaleAIException {
        return sendRaw(botId, message, pageAccessToken);
    }

    public Map<String, Object> sendText(String botId, String subscriberId, String text, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "text");
        m.put("text", text);
        return sendRaw(botId, m, pageAccessToken);
    }

    public Map<String, Object> sendImage(String botId, String subscriberId, String url, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "image");
        m.put("url", url);
        return sendRaw(botId, m, pageAccessToken);
    }

    public Map<String, Object> sendFile(String botId, String subscriberId, String url, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "file");
        m.put("url", url);
        return sendRaw(botId, m, pageAccessToken);
    }

    public Map<String, Object> sendQuickReply(String botId, String subscriberId, String text, List<String> quickReplies, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "quick_reply");
        m.put("text", text);
        m.put("quick_replies", quickReplies);
        return sendRaw(botId, m, pageAccessToken);
    }

    public Map<String, Object> sendButton(String botId, String subscriberId, String text, List<Map<String, Object>> buttons, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "button");
        m.put("text", text);
        m.put("buttons", buttons);
        return sendRaw(botId, m, pageAccessToken);
    }

    public Map<String, Object> sendCarousel(String botId, String subscriberId, List<Map<String, Object>> elements, String pageAccessToken) throws SaleAIException {
        Map<String, Object> m = new HashMap<>();
        m.put("subscriber_id", subscriberId);
        m.put("type", "carousel");
        m.put("elements", elements);
        return sendRaw(botId, m, pageAccessToken);
    }
}
