import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Ý định AI (intents) */
export class IntentsResource extends BaseResource {
  /** GET /api/v1/ai/intents */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/ai/intents", undefined, params);
  }
  /** GET /api/v1/ai/intents/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/ai/intents/${id}`);
  }
  /** POST /api/v1/ai/intents */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/ai/intents", data);
  }
  /** PATCH /api/v1/ai/intents/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/ai/intents/${id}`, data);
  }
  /** DELETE /api/v1/ai/intents/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/ai/intents/${id}`);
  }
}
