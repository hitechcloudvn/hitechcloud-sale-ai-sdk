/** HiTechCloud Sale AI SDK - error types */

export interface ErrorDetail {
  field?: string;
  issue?: string;
  [key: string]: unknown;
}

export class SaleAIError extends Error {
  statusCode?: number;
  code?: string;
  details?: ErrorDetail[];
  constructor(message: string, statusCode?: number, code?: string, details?: ErrorDetail[]) {
    super(message);
    this.name = "SaleAIError";
    this.statusCode = statusCode;
    this.code = code;
    this.details = details;
  }
}

export class AuthenticationError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 401, c, d); this.name = "AuthenticationError"; }
}
export class AuthorizationError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 403, c, d); this.name = "AuthorizationError"; }
}
export class NotFoundError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 404, c, d); this.name = "NotFoundError"; }
}
export class ConflictError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 409, c, d); this.name = "ConflictError"; }
}
export class ValidationError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 422, c, d); this.name = "ValidationError"; }
}
export class RateLimitError extends SaleAIError {
  retryAfter?: number;
  constructor(m: string, retryAfter?: number, c?: string, d?: ErrorDetail[]) {
    super(m, 429, c, d); this.name = "RateLimitError"; this.retryAfter = retryAfter;
  }
}
export class ChannelError extends SaleAIError {
  constructor(m: string, c?: string, d?: ErrorDetail[]) { super(m, 502, c, d); this.name = "ChannelError"; }
}
export class ServerError extends SaleAIError {
  constructor(m: string, statusCode = 500, c?: string, d?: ErrorDetail[]) { super(m, statusCode, c, d); this.name = "ServerError"; }
}
