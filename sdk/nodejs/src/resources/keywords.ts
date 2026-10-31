import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Từ khoá tự trả lời (keywords) */
export class KeywordsResource extends BaseResource {
  /** GET /api/v1/keywords */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/keywords", undefined, params);
  }
  /** GET /api/v1/keywords/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/keywords/${id}`);
  }
  /** POST /api/v1/keywords */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/keywords", data);
  }
  /** PATCH /api/v1/keywords/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/keywords/${id}`, data);
  }
  /** DELETE /api/v1/keywords/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/keywords/${id}`);
  }
  /** POST /api/v1/keywords/test */
  test(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/keywords/test", data);
  }
}
