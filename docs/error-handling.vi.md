# Xử lý lỗi

> [English](error-handling.md) | **Tiếng Việt**

---

Khi có lỗi, API trả `status: "error"` kèm `error.code`, `error.message` và (tuỳ
trường hợp) `error.details`.

```json
{
  "status": "error",
  "error": {
    "code": "VALIDATION_ERROR",
    "message": "Truong 'subscriber_id' la bat buoc.",
    "details": [{ "field": "subscriber_id", "issue": "required" }]
  }
}
```

## Bảng mã lỗi

| HTTP | Mã lỗi | Lớp ngoại lệ SDK |
|------|--------|------------------|
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

## Ví dụ (Python)

```python
from hitechcloud_sale_ai import SaleAI, ValidationError, RateLimitError

client = SaleAI(token="YOUR_TOKEN")
try:
    client.orders.create(data={"bot_id": "bot_7c1a9f2e"})
except ValidationError as e:
    print(e.code, e.message, e.details)
except RateLimitError as e:
    print("Thử lại sau", e.retry_after, "giây")
```

> Luôn kiểm tra `error.code` thay vì chuỗi `message` (có thể đổi). Với
> `CHANNEL_ERROR`, xem `error.details` để biết lỗi gốc từ nền tảng kênh.
