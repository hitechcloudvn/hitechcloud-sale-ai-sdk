import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Hộp thư đa kênh (conversations / inbox) */
export class ConversationsResource extends BaseResource {
  /** GET /api/v1/conversations */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/conversations", undefined, params);
  }
  /** GET /api/v1/conversations/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/conversations/${id}`);
  }
  /** POST /api/v1/conversations/{id}/assign */
  assign(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/conversations/${id}/assign`, data);
  }
  /** POST /api/v1/conversations/{id}/notes */
  addNote(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/conversations/${id}/notes`, data);
  }
  /** POST /api/v1/conversations/{id}/labels */
  addLabels(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/conversations/${id}/labels`, data);
  }
  /** POST /api/v1/conversations/{id}/close */
  close(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/conversations/${id}/close`);
  }
}
