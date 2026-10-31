import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Chuỗi tin hẹn giờ (sequences) và các bước */
export class SequencesResource extends BaseResource {
  /** GET /api/v1/sequences */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/sequences", undefined, params);
  }
  /** GET /api/v1/sequences/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/sequences/${id}`);
  }
  /** POST /api/v1/sequences */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/sequences", data);
  }
  /** PATCH /api/v1/sequences/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/sequences/${id}`, data);
  }
  /** DELETE /api/v1/sequences/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/sequences/${id}`);
  }
  /** GET /api/v1/sequences/{id}/steps */
  listSteps(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/sequences/${id}/steps`);
  }
  /** POST /api/v1/sequences/{id}/steps */
  addStep(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/sequences/${id}/steps`, data);
  }
  /** PATCH /api/v1/sequences/{id}/steps/{step_id} */
  updateStep(id: string, stepId: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/sequences/${id}/steps/${stepId}`, data);
  }
  /** DELETE /api/v1/sequences/{id}/steps/{step_id} */
  deleteStep(id: string, stepId: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/sequences/${id}/steps/${stepId}`);
  }
  /** POST /api/v1/sequences/{id}/subscribe */
  subscribe(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/sequences/${id}/subscribe`, data);
  }
  /** POST /api/v1/sequences/{id}/unsubscribe */
  unsubscribe(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/sequences/${id}/unsubscribe`, data);
  }
}
