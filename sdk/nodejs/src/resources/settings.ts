import { BaseResource } from "./base";
import type { Query } from "../http-client";

/** Cài đặt tài khoản (settings) */
export class SettingsResource extends BaseResource {
  /** GET /api/v1/settings */
  get(): Promise<unknown> {
    return this.http.request("GET", "/api/v1/settings");
  }
  /** PATCH /api/v1/settings */
  update(data?: Record<string, unknown>): Promise<unknown> {
    return this.http.request("PATCH", "/api/v1/settings", data);
  }
}
