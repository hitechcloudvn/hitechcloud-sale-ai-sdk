import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Phân khúc (segments) */
export class SegmentsResource extends BaseResource {
  /** GET /api/v1/segments */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/segments", undefined, params);
  }
  /** GET /api/v1/segments/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/segments/${id}`);
  }
  /** POST /api/v1/segments */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/segments", data);
  }
  /** PATCH /api/v1/segments/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/segments/${id}`, data);
  }
  /** DELETE /api/v1/segments/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/segments/${id}`);
  }
}
