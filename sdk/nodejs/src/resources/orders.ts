import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Đơn hàng (orders) */
export class OrdersResource extends BaseResource {
  /** GET /api/v1/orders */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/orders", undefined, params);
  }
  /** GET /api/v1/orders/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/orders/${id}`);
  }
  /** POST /api/v1/orders */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/orders", data);
  }
  /** PATCH /api/v1/orders/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/orders/${id}`, data);
  }
  /** DELETE /api/v1/orders/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/orders/${id}`);
  }
  /** POST /api/v1/orders/from-conversation */
  createFromConversation(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/orders/from-conversation", data);
  }
  /** POST /api/v1/orders/{id}/status */
  updateStatus(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", `/api/v1/orders/${id}/status`, data);
  }
}
