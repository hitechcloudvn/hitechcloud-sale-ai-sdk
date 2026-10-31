import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Công cụ tăng trưởng (widgets) */
export class WidgetsResource extends BaseResource {
  /** GET /api/v1/widgets */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/widgets", undefined, params);
  }
  /** GET /api/v1/widgets/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/widgets/${id}`);
  }
  /** POST /api/v1/widgets */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/widgets", data);
  }
  /** PATCH /api/v1/widgets/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/widgets/${id}`, data);
  }
  /** DELETE /api/v1/widgets/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/widgets/${id}`);
  }
}
