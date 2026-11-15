# Authentication

> **English** | [Tiếng Việt](authentication.vi.md)

---

HiTechCloud Sale AI uses two layers of authentication.

## 1. Access token (Bearer)

Required on every request, sent in the `Authorization` header:

```
Authorization: Bearer YOUR_TOKEN
```

In the SDK: pass `token` when constructing the client, or call `setToken()`.

## 2. Per-bot page access token

Endpoints that act directly on a channel (send messages, hide/like comments)
also require the `X-Page-Access-Token` header carrying that bot's page token.
This token is stored encrypted server-side and never returned in full.

```
X-Page-Access-Token: PAGE_ACCESS_TOKEN
```

In the SDK, you can:

- Set a default on the client: `SaleAI(token=..., page_access_token=...)`.
- Or pass it per call: `client.messages.send_text(..., page_access_token=...)`.

## Authentication error codes

| HTTP | Error code | Meaning |
|------|------------|---------|
| 401 | `UNAUTHORIZED` | Missing or invalid access token |
| 401 | `TOKEN_EXPIRED` | Token has expired |
| 403 | `FORBIDDEN` | No permission for the resource/bot |
| 403 | `PAGE_TOKEN_REQUIRED` | Missing `X-Page-Access-Token` for a channel operation |

> Always call the API from your server. Never expose the access token or page
> access token in client-side applications.
