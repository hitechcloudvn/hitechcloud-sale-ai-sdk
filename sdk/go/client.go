package saleai

import "time"

const (
	ProductionURL = "https://sale-ai.hitechcloud.vn"
	SandboxURL    = "https://sale-ai-sandbox.hitechcloud.vn"
	Version       = "1.0.0"
)

// Client is the main HiTechCloud Sale AI client.
type Client struct {
	http *HTTPClient
	Bots *BotsResource
	Flows *FlowsResource
	Keywords *KeywordsResource
	Sequences *SequencesResource
	Broadcasts *BroadcastsResource
	Subscribers *SubscribersResource
	Tags *TagsResource
	Segments *SegmentsResource
	Products *ProductsResource
	Orders *OrdersResource
	CommentRules *CommentRulesResource
	Widgets *WidgetsResource
	Intents *IntentsResource
	KnowledgeBase *KnowledgeBaseResource
	Users *UsersResource
	Roles *RolesResource
	Webhooks *WebhooksResource
	Comments *CommentsResource
	Ai *AiResource
	Conversations *ConversationsResource
	Analytics *AnalyticsResource
	Settings *SettingsResource
	Messages *MessagesResource
}

// Option configures the client.
type Option func(*config)

type config struct {
	baseURL         string
	token           string
	pageAccessToken string
	timeout         time.Duration
	maxRetries      int
}

// WithToken sets the account access token.
func WithToken(token string) Option { return func(c *config) { c.token = token } }

// WithPageAccessToken sets a default page access token.
func WithPageAccessToken(token string) Option { return func(c *config) { c.pageAccessToken = token } }

// WithBaseURL overrides the base URL.
func WithBaseURL(u string) Option { return func(c *config) { c.baseURL = u } }

// WithSandbox targets the sandbox environment.
func WithSandbox() Option { return func(c *config) { c.baseURL = SandboxURL } }

// WithTimeout sets the request timeout.
func WithTimeout(d time.Duration) Option { return func(c *config) { c.timeout = d } }

// WithMaxRetries sets retry attempts.
func WithMaxRetries(n int) Option { return func(c *config) { c.maxRetries = n } }

// NewClient creates a new Sale AI client.
func NewClient(opts ...Option) *Client {
	cfg := &config{baseURL: ProductionURL, timeout: 30 * time.Second, maxRetries: 3}
	for _, o := range opts {
		o(cfg)
	}
	http := newHTTPClient(cfg.baseURL, cfg.token, cfg.pageAccessToken, cfg.timeout, cfg.maxRetries)
	c := &Client{http: http}
	c.Bots = &BotsResource{client: http}
	c.Flows = &FlowsResource{client: http}
	c.Keywords = &KeywordsResource{client: http}
	c.Sequences = &SequencesResource{client: http}
	c.Broadcasts = &BroadcastsResource{client: http}
	c.Subscribers = &SubscribersResource{client: http}
	c.Tags = &TagsResource{client: http}
	c.Segments = &SegmentsResource{client: http}
	c.Products = &ProductsResource{client: http}
	c.Orders = &OrdersResource{client: http}
	c.CommentRules = &CommentRulesResource{client: http}
	c.Widgets = &WidgetsResource{client: http}
	c.Intents = &IntentsResource{client: http}
	c.KnowledgeBase = &KnowledgeBaseResource{client: http}
	c.Users = &UsersResource{client: http}
	c.Roles = &RolesResource{client: http}
	c.Webhooks = &WebhooksResource{client: http}
	c.Comments = &CommentsResource{client: http}
	c.Ai = &AiResource{client: http}
	c.Conversations = &ConversationsResource{client: http}
	c.Analytics = &AnalyticsResource{client: http}
	c.Settings = &SettingsResource{client: http}
	c.Messages = &MessagesResource{client: http}
	return c
}

// SetToken updates the account access token.
func (c *Client) SetToken(token string) { c.http.SetToken(token) }

// SetPageAccessToken updates the default page access token.
func (c *Client) SetPageAccessToken(token string) { c.http.SetPageAccessToken(token) }
