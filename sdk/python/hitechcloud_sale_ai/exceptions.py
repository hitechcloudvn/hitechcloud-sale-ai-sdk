"""HiTechCloud Sale AI SDK - Exceptions"""


class SaleAIError(Exception):
    """Base exception for the Sale AI SDK."""

    def __init__(self, message, status_code=None, code=None, details=None):
        self.message = message
        self.status_code = status_code
        self.code = code
        self.details = details or []
        super().__init__(self.message)


class AuthenticationError(SaleAIError):
    """401 UNAUTHORIZED / TOKEN_EXPIRED."""


class AuthorizationError(SaleAIError):
    """403 FORBIDDEN / PAGE_TOKEN_REQUIRED."""


class NotFoundError(SaleAIError):
    """404 NOT_FOUND."""


class ConflictError(SaleAIError):
    """409 CONFLICT."""


class ValidationError(SaleAIError):
    """422 VALIDATION_ERROR / OUTSIDE_MESSAGING_WINDOW."""


class RateLimitError(SaleAIError):
    """429 RATE_LIMITED."""

    def __init__(self, message, retry_after=None, **kwargs):
        super().__init__(message, **kwargs)
        self.retry_after = retry_after


class ChannelError(SaleAIError):
    """502 CHANNEL_ERROR - the messaging channel returned an error."""


class ServerError(SaleAIError):
    """500 / 503 server-side errors."""


ERROR_MAP = {
    401: AuthenticationError,
    403: AuthorizationError,
    404: NotFoundError,
    409: ConflictError,
    422: ValidationError,
    429: RateLimitError,
    500: ServerError,
    502: ChannelError,
    503: ServerError,
}
