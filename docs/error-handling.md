# Error handling

> **English** | [Tiếng Việt](error-handling.vi.md)

---

When an error occurs, the API returns `status: "error"` with `error.code`,
`error.message` and (where applicable) `error.details`.

```json
{
  "status": "error",
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Field 'subscriber_id' is required.",
    "details": [{ "field": "subscriber_id", "issue": "required" }]
  }
}
```

## Error code table

| HTTP | Error code | SDK exception class |
|------|------------|---------------------|
| 400 | `BAD_REQUEST` | `SaleAIError` |
| 401 | `UNAUTHORIZED` / `TOKEN_EXPIRED` | `AuthenticationError` |
| 403 | `FORBIDDEN` / `PAGE_TOKEN_REQUIRED` | `AuthorizationError` |
| 404 | `NOT_FOUND` | `NotFoundError` |
| 409 | `CONFLICT` | `ConflictError` |
| 422 | `VALIDATION_ERROR` / `OUTSIDE_MESSAGING_WINDOW` | `ValidationError` |
| 429 | `RATE_LIMITED` | `RateLimitError` |
| 500 | `INTERNAL_ERROR` | `ServerError` |
| 502 | `CHANNEL_ERROR` | `ChannelError` |
| 503 | `SERVICE_UNAVAILABLE` | `ServerError` |

## Example (Python)

```python
from hitechcloud_sale_ai import SaleAI, ValidationError, RateLimitError

client = SaleAI(token="YOUR_TOKEN")
try:
    client.orders.create(data={"bot_id": "bot_7c1a9f2e"})
except ValidationError as e:
    print(e.code, e.message, e.details)
except RateLimitError as e:
    print("Retry after", e.retry_after, "seconds")
```

> Always check `error.code` rather than the `message` string (which may change).
> For `CHANNEL_ERROR`, see `error.details` for the underlying channel-platform
> error.
