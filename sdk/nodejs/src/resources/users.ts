import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Người dùng / nhân sự (team) */
export class UsersResource extends BaseResource {
  /** GET /api/v1/users */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/users", undefined, params);
  }
  /** GET /api/v1/users/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/users/${id}`);
  }
  /** POST /api/v1/users */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/users", data);
  }
  /** PATCH /api/v1/users/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/users/${id}`, data);
  }
  /** DELETE /api/v1/users/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/users/${id}`);
  }
}
