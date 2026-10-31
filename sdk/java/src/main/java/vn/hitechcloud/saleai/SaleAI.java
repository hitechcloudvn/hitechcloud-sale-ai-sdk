package vn.hitechcloud.saleai;

import vn.hitechcloud.saleai.resources.*;

public class SaleAI {
    public static final String PRODUCTION_URL = "https://sale-ai.hitechcloud.vn";
    public static final String SANDBOX_URL = "https://sale-ai-sandbox.hitechcloud.vn";
    public static final String VERSION = "1.0.0";

    private final HttpClient http;

    public final Bots bots;
    public final Flows flows;
    public final Keywords keywords;
    public final Sequences sequences;
    public final Broadcasts broadcasts;
    public final Subscribers subscribers;
    public final Tags tags;
    public final Segments segments;
    public final Products products;
    public final Orders orders;
    public final CommentRules commentRules;
    public final Widgets widgets;
    public final Intents intents;
    public final KnowledgeBase knowledgeBase;
    public final Users users;
    public final Roles roles;
    public final Webhooks webhooks;
    public final Comments comments;
    public final Ai ai;
    public final Conversations conversations;
    public final Analytics analytics;
    public final Settings settings;
    public final Messages messages;

    public SaleAI(String token) {
        this(token, null, PRODUCTION_URL, 30, 3);
    }

    public SaleAI(String token, String pageAccessToken, String baseUrl, int timeout, int maxRetries) {
        this.http = new HttpClient(baseUrl, token, pageAccessToken, timeout, maxRetries);
        this.bots = new Bots(http);
        this.flows = new Flows(http);
        this.keywords = new Keywords(http);
        this.sequences = new Sequences(http);
        this.broadcasts = new Broadcasts(http);
        this.subscribers = new Subscribers(http);
        this.tags = new Tags(http);
        this.segments = new Segments(http);
        this.products = new Products(http);
        this.orders = new Orders(http);
        this.commentRules = new CommentRules(http);
        this.widgets = new Widgets(http);
        this.intents = new Intents(http);
        this.knowledgeBase = new KnowledgeBase(http);
        this.users = new Users(http);
        this.roles = new Roles(http);
        this.webhooks = new Webhooks(http);
        this.comments = new Comments(http);
        this.ai = new Ai(http);
        this.conversations = new Conversations(http);
        this.analytics = new Analytics(http);
        this.settings = new Settings(http);
        this.messages = new Messages(http);
    }

    public void setToken(String token) { http.setToken(token); }
    public void setPageAccessToken(String token) { http.setPageAccessToken(token); }
}
