import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Trợ lý AI bán hàng (config / train / suggest / handover) */
export class AiResource extends BaseResource {
  /** PATCH /api/v1/ai/config */
  updateConfig(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", "/api/v1/ai/config", data);
  }
  /** POST /api/v1/ai/train */
  train(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/ai/train", data);
  }
  /** POST /api/v1/ai/suggest-reply */
  suggestReply(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/ai/suggest-reply", data);
  }
  /** POST /api/v1/ai/handover */
  handover(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/ai/handover", data);
  }
}
