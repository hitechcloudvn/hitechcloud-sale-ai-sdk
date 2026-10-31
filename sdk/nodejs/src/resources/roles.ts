import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Vai trò & quyền (roles) */
export class RolesResource extends BaseResource {
  /** GET /api/v1/roles */
  list(): Promise<unknown> {
    return this.http.request("GET", "/api/v1/roles");
  }
  /** GET /api/v1/roles/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/roles/${id}`);
  }
  /** POST /api/v1/roles */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/roles", data);
  }
  /** PATCH /api/v1/roles/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/roles/${id}`, data);
  }
  /** DELETE /api/v1/roles/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/roles/${id}`);
  }
}
