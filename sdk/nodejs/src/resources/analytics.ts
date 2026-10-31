import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Báo cáo & phân tích (analytics) */
export class AnalyticsResource extends BaseResource {
  /** GET /api/v1/analytics/overview */
  overview(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/analytics/overview", undefined, params);
  }
  /** GET /api/v1/analytics/bots/{bot_id} */
  bot(botId: string, params?: Query): Promise<unknown> {
    return this.http.request("GET", `/api/v1/analytics/bots/${botId}`, undefined, params);
  }
  /** GET /api/v1/analytics/flows/{id} */
  flow(id: string, params?: Query): Promise<unknown> {
    return this.http.request("GET", `/api/v1/analytics/flows/${id}`, undefined, params);
  }
  /** GET /api/v1/analytics/broadcasts/{id} */
  broadcast(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/analytics/broadcasts/${id}`);
  }
  /** GET /api/v1/analytics/subscriber-growth */
  subscriberGrowth(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/analytics/subscriber-growth", undefined, params);
  }
}
