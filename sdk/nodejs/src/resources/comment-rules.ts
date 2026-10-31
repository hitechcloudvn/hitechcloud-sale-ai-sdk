import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Quy tắc trả lời bình luận (comment rules) */
export class CommentRulesResource extends BaseResource {
  /** GET /api/v1/comment-rules */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/comment-rules", undefined, params);
  }
  /** GET /api/v1/comment-rules/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/comment-rules/${id}`);
  }
  /** POST /api/v1/comment-rules */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/comment-rules", data);
  }
  /** PATCH /api/v1/comment-rules/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/comment-rules/${id}`, data);
  }
  /** DELETE /api/v1/comment-rules/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/comment-rules/${id}`);
  }
}
