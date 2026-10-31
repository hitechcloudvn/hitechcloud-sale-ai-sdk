import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Sản phẩm (products) */
export class ProductsResource extends BaseResource {
  /** GET /api/v1/products */
  list(params?: Query): Promise<unknown> {
    return this.http.request("GET", "/api/v1/products", undefined, params);
  }
  /** GET /api/v1/products/{id} */
  get(id: string): Promise<unknown> {
    return this.http.request("GET", `/api/v1/products/${id}`);
  }
  /** POST /api/v1/products */
  create(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("POST", "/api/v1/products", data);
  }
  /** PATCH /api/v1/products/{id} */
  update(id: string, data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", `/api/v1/products/${id}`, data);
  }
  /** DELETE /api/v1/products/{id} */
  delete(id: string): Promise<unknown> {
    return this.http.request("DELETE", `/api/v1/products/${id}`);
  }
}
