import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Bot & kênh kết nối (Facebook/Instagram/Zalo/Website) */
export class BotsResource extends BaseResource {
  /** GET /api/v1/bots */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/bots", undefined, params);
  }
  /** GET /api/v1/bots/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/bots/${id}`);
  }
  /** POST /api/v1/bots */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/bots", data);
  }
  /** PATCH /api/v1/bots/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/bots/${id}`, data);
  }
  /** DELETE /api/v1/bots/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/bots/${id}`);
  }
  /** POST /api/v1/bots/{bot_id}/connect */
  connect(botId: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/bots/${botId}/connect`, data);
  }
  /** POST /api/v1/bots/{bot_id}/disconnect */
  disconnect(botId: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/bots/${botId}/disconnect`);
  }
  /** PATCH /api/v1/bots/{bot_id}/settings */
  updateSettings(botId: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/bots/${botId}/settings`, data);
  }
  /** POST /api/v1/bots/{bot_id}/access-token/refresh */
  refreshAccessToken(botId: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/bots/${botId}/access-token/refresh`, data);
  }
}
