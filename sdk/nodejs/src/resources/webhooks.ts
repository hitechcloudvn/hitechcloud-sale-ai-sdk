import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Đăng ký webhook (webhooks) */
export class WebhooksResource extends BaseResource {
  /** GET /api/v1/webhooks */
  list(): Promise<unknown> {
    return this.http.request("GET", "/api/v1/webhooks");
  }
  /** GET /api/v1/webhooks/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/webhooks/${id}`);
  }
  /** POST /api/v1/webhooks */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/webhooks", data);
  }
  /** PATCH /api/v1/webhooks/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/webhooks/${id}`, data);
  }
  /** DELETE /api/v1/webhooks/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/webhooks/${id}`);
  }
}
