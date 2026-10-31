import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Gửi hàng loạt (broadcasts) */
export class BroadcastsResource extends BaseResource {
  /** GET /api/v1/broadcasts */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/broadcasts", undefined, params);
  }
  /** GET /api/v1/broadcasts/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/broadcasts/${id}`);
  }
  /** POST /api/v1/broadcasts */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/broadcasts", data);
  }
  /** PATCH /api/v1/broadcasts/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/broadcasts/${id}`, data);
  }
  /** DELETE /api/v1/broadcasts/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/broadcasts/${id}`);
  }
  /** POST /api/v1/broadcasts/{id}/schedule */
  schedule(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/broadcasts/${id}/schedule`, data);
  }
  /** POST /api/v1/broadcasts/{id}/send */
  send(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/broadcasts/${id}/send`);
  }
  /** GET /api/v1/broadcasts/{id}/status */
  status(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/broadcasts/${id}/status`);
  }
  /** POST /api/v1/broadcasts/{id}/audience/preview */
  previewAudience(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/broadcasts/${id}/audience/preview`, data);
  }
  /** POST /api/v1/broadcasts/{id}/cancel */
  cancel(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/broadcasts/${id}/cancel`);
  }
}
