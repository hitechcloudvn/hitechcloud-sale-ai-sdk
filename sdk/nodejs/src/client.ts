import { HttpClient, HttpClientOptions } from "./http-client";
import { BotsResource } from "./resources/bots";
import { FlowsResource } from "./resources/flows";
import { KeywordsResource } from "./resources/keywords";
import { SequencesResource } from "./resources/sequences";
import { BroadcastsResource } from "./resources/broadcasts";
import { SubscribersResource } from "./resources/subscribers";
import { TagsResource } from "./resources/tags";
import { SegmentsResource } from "./resources/segments";
import { ProductsResource } from "./resources/products";
import { OrdersResource } from "./resources/orders";
import { CommentRulesResource } from "./resources/comment-rules";
import { WidgetsResource } from "./resources/widgets";
import { IntentsResource } from "./resources/intents";
import { KnowledgeBaseResource } from "./resources/knowledge-base";
import { UsersResource } from "./resources/users";
import { RolesResource } from "./resources/roles";
import { WebhooksResource } from "./resources/webhooks";
import { CommentsResource } from "./resources/comments";
import { AiResource } from "./resources/ai";
import { ConversationsResource } from "./resources/conversations";
import { AnalyticsResource } from "./resources/analytics";
import { SettingsResource } from "./resources/settings";
import { MessagesResource } from "./resources/messages";

export interface SaleAIOptions {
  token?: string;
  pageAccessToken?: string;
  baseUrl?: string;
  timeout?: number;
  maxRetries?: number;
}

export const PRODUCTION_URL = "https://sale-ai.hitechcloud.vn";
export const SANDBOX_URL = "https://sale-ai-sandbox.hitechcloud.vn";

export class SaleAI {
  private http: HttpClient;
  readonly bots: BotsResource;
  readonly flows: FlowsResource;
  readonly keywords: KeywordsResource;
  readonly sequences: SequencesResource;
  readonly broadcasts: BroadcastsResource;
  readonly subscribers: SubscribersResource;
  readonly tags: TagsResource;
  readonly segments: SegmentsResource;
  readonly products: ProductsResource;
  readonly orders: OrdersResource;
  readonly commentRules: CommentRulesResource;
  readonly widgets: WidgetsResource;
  readonly intents: IntentsResource;
  readonly knowledgeBase: KnowledgeBaseResource;
  readonly users: UsersResource;
  readonly roles: RolesResource;
  readonly webhooks: WebhooksResource;
  readonly comments: CommentsResource;
  readonly ai: AiResource;
  readonly conversations: ConversationsResource;
  readonly analytics: AnalyticsResource;
  readonly settings: SettingsResource;
  readonly messages: MessagesResource;

  constructor(options: SaleAIOptions = {}) {
    const opts: HttpClientOptions = {
      baseUrl: options.baseUrl || PRODUCTION_URL,
      token: options.token,
      pageAccessToken: options.pageAccessToken,
      timeout: options.timeout,
      maxRetries: options.maxRetries,
    };
    this.http = new HttpClient(opts);
    this.bots = new BotsResource(this.http);
    this.flows = new FlowsResource(this.http);
    this.keywords = new KeywordsResource(this.http);
    this.sequences = new SequencesResource(this.http);
    this.broadcasts = new BroadcastsResource(this.http);
    this.subscribers = new SubscribersResource(this.http);
    this.tags = new TagsResource(this.http);
    this.segments = new SegmentsResource(this.http);
    this.products = new ProductsResource(this.http);
    this.orders = new OrdersResource(this.http);
    this.commentRules = new CommentRulesResource(this.http);
    this.widgets = new WidgetsResource(this.http);
    this.intents = new IntentsResource(this.http);
    this.knowledgeBase = new KnowledgeBaseResource(this.http);
    this.users = new UsersResource(this.http);
    this.roles = new RolesResource(this.http);
    this.webhooks = new WebhooksResource(this.http);
    this.comments = new CommentsResource(this.http);
    this.ai = new AiResource(this.http);
    this.conversations = new ConversationsResource(this.http);
    this.analytics = new AnalyticsResource(this.http);
    this.settings = new SettingsResource(this.http);
    this.messages = new MessagesResource(this.http);
  }

  setToken(token: string): void { this.http.setToken(token); }
  setPageAccessToken(token: string): void { this.http.setPageAccessToken(token); }
}
