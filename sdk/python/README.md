# HiTechCloud Sale AI - Python SDK

> **English** | [Tiếng Việt](README.vi.md)

---

Official Python SDK for **HiTechCloud Sale AI** — the multichannel sales
chatbot and conversation automation platform (Facebook Messenger, Instagram,
Zalo, Website).

- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`
- Authentication: `Authorization: Bearer <token>` + `X-Page-Access-Token` for channel operations

## Installation

```bash
pip install hitechcloud-sale-ai
```

## Quick start

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")

# List bots
bots = client.bots.list()

# Send a text message (page access token required)
client.messages.send_text(
    "bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
    page_access_token="PAGE_ACCESS_TOKEN",
)
```

You can set a default page access token on the client:

```python
client = SaleAI(token="YOUR_TOKEN", page_access_token="PAGE_ACCESS_TOKEN")
client.messages.send_text("bot_7c1a9f2e", "sub_2a9b8c7d", "Xin chao!")
```

## Webhook signature verification

```python
from hitechcloud_sale_ai import verify_signature, SIGNATURE_HEADER

ok = verify_signature(raw_body, request.headers.get(SIGNATURE_HEADER), "whsec_...")
```

## Resources

`bots`, `flows`, `keywords`, `sequences`, `broadcasts`, `subscribers`, `tags`,
`segments`, `products`, `orders`, `comment_rules`, `widgets`, `intents`,
`knowledge_base`, `users`, `roles`, `webhooks`, `comments`, `ai`,
`conversations`, `analytics`, `settings`, `messages`.

## Errors

All errors extend `SaleAIError`: `AuthenticationError`, `AuthorizationError`,
`NotFoundError`, `ConflictError`, `ValidationError`, `RateLimitError`,
`ChannelError`, `ServerError`.

MIT License.
