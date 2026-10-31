/** HiTechCloud Sale AI SDK - HTTP client (fetch based) */
import {
  SaleAIError, AuthenticationError, AuthorizationError, NotFoundError,
  ConflictError, ValidationError, RateLimitError, ChannelError, ServerError,
} from "./errors";

export interface HttpClientOptions {
  baseUrl: string;
  token?: string;
  pageAccessToken?: string;
  timeout?: number;
  maxRetries?: number;
}

export type Query = Record<string, string | number | boolean | undefined>;
export type Headers = Record<string, string | undefined>;

const USER_AGENT = "HiTechCloud-SaleAI-NodeJS/1.0.0";

function sleep(ms: number): Promise<void> {
  return new Promise((r) => setTimeout(r, ms));
}

export class HttpClient {
  baseUrl: string;
  private token?: string;
  private pageAccessToken?: string;
  private timeout: number;
  private maxRetries: number;

  constructor(opts: HttpClientOptions) {
    this.baseUrl = opts.baseUrl.replace(/\/$/, "");
    this.token = opts.token;
    this.pageAccessToken = opts.pageAccessToken;
    this.timeout = opts.timeout ?? 30000;
    this.maxRetries = opts.maxRetries ?? 3;
  }

  setToken(token: string): void { this.token = token; }
  clearToken(): void { this.token = undefined; }
  setPageAccessToken(token: string): void { this.pageAccessToken = token; }

  private buildHeaders(extra?: Headers): Record<string, string> {
    const h: Record<string, string> = {
      "Content-Type": "application/json",
      Accept: "application/json",
      "User-Agent": USER_AGENT,
    };
    if (this.token) h["Authorization"] = `Bearer ${this.token}`;
    if (this.pageAccessToken) h["X-Page-Access-Token"] = this.pageAccessToken;
    if (extra) for (const [k, v] of Object.entries(extra)) if (v !== undefined) h[k] = v;
    return h;
  }

  async request<T = unknown>(
    method: string,
    path: string,
    body?: unknown,
    params?: Query,
    headers?: Headers,
  ): Promise<T> {
    let url = this.baseUrl + path;
    if (params) {
      const qs = Object.entries(params)
        .filter(([, v]) => v !== undefined)
        .map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(String(v))}`)
        .join("&");
      if (qs) url += (url.includes("?") ? "&" : "?") + qs;
    }

    let lastError: unknown;
    for (let attempt = 0; attempt <= this.maxRetries; attempt++) {
      const controller = new AbortController();
      const timer = setTimeout(() => controller.abort(), this.timeout);
      try {
        const res = await fetch(url, {
          method,
          headers: this.buildHeaders(headers),
          body: body !== undefined && body !== null ? JSON.stringify(body) : undefined,
          signal: controller.signal,
        });
        clearTimeout(timer);

        if (res.status === 429) {
          const retryAfter = parseInt(res.headers.get("Retry-After") || "2", 10);
          if (attempt < this.maxRetries) { await sleep(retryAfter * 1000); continue; }
          throw new RateLimitError("Rate limit exceeded", retryAfter, "RATE_LIMITED");
        }
        if (res.status >= 500 && attempt < this.maxRetries) {
          await sleep(Math.pow(2, attempt) * 1000); continue;
        }
        if (res.status >= 400) { await this.raise(res); }
        if (res.status === 204) return undefined as T;
        const text = await res.text();
        return (text ? JSON.parse(text) : undefined) as T;
      } catch (err) {
        clearTimeout(timer);
        if (err instanceof SaleAIError) throw err;
        lastError = err;
        if (attempt < this.maxRetries) { await sleep(Math.pow(2, attempt) * 1000); continue; }
        throw new SaleAIError(`Request failed: ${(err as Error).message}`);
      }
    }
    throw new SaleAIError(`Request failed: ${String(lastError)}`);
  }

  private async raise(res: Response): Promise<never> {
    let message = `HTTP ${res.status}`;
    let code: string | undefined;
    let details: never[] | undefined;
    try {
      const body = await res.json();
      const e = (body && body.error) || {};
      message = e.message || message;
      code = e.code;
      details = e.details;
    } catch {
      /* non-json error body */
    }
    switch (res.status) {
      case 401: throw new AuthenticationError(message, code, details);
      case 403: throw new AuthorizationError(message, code, details);
      case 404: throw new NotFoundError(message, code, details);
      case 409: throw new ConflictError(message, code, details);
      case 422: throw new ValidationError(message, code, details);
      case 429: {
        const retryAfter = parseInt(res.headers.get("Retry-After") || "2", 10);
        throw new RateLimitError(message, retryAfter, code, details);
      }
      case 502: throw new ChannelError(message, code, details);
      default:
        if (res.status >= 500) throw new ServerError(message, res.status, code, details);
        throw new SaleAIError(message, res.status, code, details);
    }
  }
}
