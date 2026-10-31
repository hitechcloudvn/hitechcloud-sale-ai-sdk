import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Kho tri thức AI (knowledge base) */
export class KnowledgeBaseResource extends BaseResource {
  /** GET /api/v1/ai/knowledge-base */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/ai/knowledge-base", undefined, params);
  }
  /** GET /api/v1/ai/knowledge-base/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/ai/knowledge-base/${id}`);
  }
  /** POST /api/v1/ai/knowledge-base */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/ai/knowledge-base", data);
  }
  /** PATCH /api/v1/ai/knowledge-base/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/ai/knowledge-base/${id}`, data);
  }
  /** DELETE /api/v1/ai/knowledge-base/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/ai/knowledge-base/${id}`);
  }
}
