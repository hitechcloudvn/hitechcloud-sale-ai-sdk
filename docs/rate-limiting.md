# Pagination & rate limits

> **English** | [Tiếng Việt](rate-limiting.vi.md)

---

## Pagination

Every list endpoint uses `page` and `per_page` and returns a `meta.pagination`
block.

| Parameter | Type | Default | Description |
|-----------|------|---------|-------------|
| `page` | integer | 1 | Page number |
| `per_page` | integer | 20 | Records per page (max 100) |

```json
{
  "meta": {
    "pagination": {
      "total": 128, "per_page": 20, "current_page": 1,
      "last_page": 7, "has_more": true
    }
  }
}
```

### Iterating through all results (Python)

```python
page = 1
while True:
    res = client.subscribers.list(params={"bot_id": "bot_7c1a9f2e", "page": page, "per_page": 100})
    for sub in res["data"]:
        print(sub["id"])
    if not res["meta"]["pagination"]["has_more"]:
        break
    page += 1
```

## Rate limits

Each access token has a per-minute API call allowance. The information is
returned in headers:

| Header | Meaning |
|--------|---------|
| `X-RateLimit-Limit` | Maximum requests in the window |
| `X-RateLimit-Remaining` | Requests remaining |
| `X-RateLimit-Reset` | Window reset time (epoch) |
| `Retry-After` | Seconds to wait after a `429` |

When the allowance is exceeded, the API returns `429 RATE_LIMITED`. All SDKs in
this suite **retry automatically** honoring `Retry-After` (and exponential
backoff for 5xx errors). Send large broadcasts through the broadcast endpoint
instead of repeating message-sending calls.
