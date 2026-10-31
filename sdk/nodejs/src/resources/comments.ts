import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Thao tác bình luận trên kênh (comments) */
export class CommentsResource extends BaseResource {
  /** POST /api/v1/comments/{comment_id}/hide */
  hide(commentId: string, pageAccessToken?: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/comments/${commentId}/hide`, undefined, undefined, pageAccessToken ? { "X-Page-Access-Token": pageAccessToken } : undefined);
  }
  /** POST /api/v1/comments/{comment_id}/like */
  like(commentId: string, pageAccessToken?: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/comments/${commentId}/like`, undefined, undefined, pageAccessToken ? { "X-Page-Access-Token": pageAccessToken } : undefined);
  }
}
