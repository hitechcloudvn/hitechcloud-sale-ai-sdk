import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Kịch bản hội thoại (flows) và khối (blocks) */
export class FlowsResource extends BaseResource {
  /** GET /api/v1/flows */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/flows", undefined, params);
  }
  /** GET /api/v1/flows/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/flows/${id}`);
  }
  /** POST /api/v1/flows */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/flows", data);
  }
  /** PATCH /api/v1/flows/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/flows/${id}`, data);
  }
  /** DELETE /api/v1/flows/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/flows/${id}`);
  }
  /** POST /api/v1/flows/{id}/publish */
  publish(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/flows/${id}/publish`);
  }
  /** POST /api/v1/flows/{id}/duplicate */
  duplicate(id: string): Promise<unknown> {
    return this.http.request("POST", `/api/v1/flows/${id}/duplicate`);
  }
  /** POST /api/v1/flows/{id}/trigger */
  trigger(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/flows/${id}/trigger`, data);
  }
  /** GET /api/v1/flows/{id}/blocks */
  listBlocks(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/flows/${id}/blocks`);
  }
  /** POST /api/v1/flows/{id}/blocks */
  addBlock(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/flows/${id}/blocks`, data);
  }
  /** PATCH /api/v1/flows/{id}/blocks/{block_id} */
  updateBlock(id: string, blockId: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/flows/${id}/blocks/${blockId}`, data);
  }
  /** DELETE /api/v1/flows/{id}/blocks/{block_id} */
  deleteBlock(id: string, blockId: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/flows/${id}/blocks/${blockId}`);
  }
}
