import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú */
export class SubscribersResource extends BaseResource {
  /** GET /api/v1/subscribers */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/subscribers", undefined, params);
  }
  /** GET /api/v1/subscribers/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/subscribers/${id}`);
  }
  /** POST /api/v1/subscribers */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/subscribers", data);
  }
  /** PATCH /api/v1/subscribers/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/subscribers/${id}`, data);
  }
  /** DELETE /api/v1/subscribers/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/subscribers/${id}`);
  }
  /** POST /api/v1/subscribers/{id}/tags */
  addTags(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/tags`, data);
  }
  /** DELETE /api/v1/subscribers/{id}/tags/{tag} */
  removeTag(id: string, tag: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/subscribers/${id}/tags/${tag}`);
  }
  /** PATCH /api/v1/subscribers/{id}/custom-fields */
  setCustomFields(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/subscribers/${id}/custom-fields`, data);
  }
  /** POST /api/v1/subscribers/{id}/notes */
  addNote(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/notes`, data);
  }
  /** POST /api/v1/subscribers/{id}/block */
  block(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/block`);
  }
  /** POST /api/v1/subscribers/{id}/unblock */
  unblock(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/unblock`);
  }
  /** GET /api/v1/subscribers/{id}/messages */
  listMessages(id: string, params?: Query): Promise<unknown> {
    return this.http.request("GET", `/api/v1/subscribers/${id}/messages`, undefined, params);
  }
  /** POST /api/v1/subscribers/{id}/messages */
  sendMessage(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/messages`, data);
  }
  /** POST /api/v1/subscribers/{id}/mark-seen */
  markSeen(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/mark-seen`);
  }
  /** POST /api/v1/subscribers/{id}/typing */
  typing(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/subscribers/${id}/typing`, data);
  }
}
