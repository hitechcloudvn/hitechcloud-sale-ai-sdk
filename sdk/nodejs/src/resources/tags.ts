import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Nhãn (tags) */
export class TagsResource extends BaseResource {
  /** GET /api/v1/tags */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/tags", undefined, params);
  }
  /** GET /api/v1/tags/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/tags/${id}`);
  }
  /** POST /api/v1/tags */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/tags", data);
  }
  /** PATCH /api/v1/tags/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/tags/${id}`, data);
  }
  /** DELETE /api/v1/tags/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/tags/${id}`);
  }
}
