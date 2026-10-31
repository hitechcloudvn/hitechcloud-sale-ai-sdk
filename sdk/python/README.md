# HiTechCloud Sale AI - Python SDK

SDK Python chính thức cho **HiTechCloud Sale AI** — nền tảng chatbot bán hàng & tự
động hoá hội thoại đa kênh (Facebook Messenger, Instagram, Zalo, Website).

- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`
- Xác thực: `Authorization: Bearer <token>` + `X-Page-Access-Token` cho thao tác kênh

## Cài đặt

```bash
pip install hitechcloud-sale-ai
```

## Bắt đầu nhanh

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")

# Liệt kê bot
bots = client.bots.list()

# Gửi tin nhắn văn bản (cần page access token)
client.messages.send_text(
    "bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
    page_access_token="PAGE_ACCESS_TOKEN",
)
```

Có thể đặt page access token mặc định cho client:

```python
client = SaleAI(token="YOUR_TOKEN", page_access_token="PAGE_ACCESS_TOKEN")
client.messages.send_text("bot_7c1a9f2e", "sub_2a9b8c7d", "Xin chao!")
```

## Xác minh chữ ký webhook

```python
from hitechcloud_sale_ai import verify_signature, SIGNATURE_HEADER

ok = verify_signature(raw_body, request.headers.get(SIGNATURE_HEADER), "whsec_...")
```

## Tài nguyên

`bots`, `flows`, `keywords`, `sequences`, `broadcasts`, `subscribers`, `tags`,
`segments`, `products`, `orders`, `comment_rules`, `widgets`, `intents`,
`knowledge_base`, `users`, `roles`, `webhooks`, `comments`, `ai`,
`conversations`, `analytics`, `settings`, `messages`.

## Lỗi

Tất cả lỗi kế thừa `SaleAIError`: `AuthenticationError`, `AuthorizationError`,
`NotFoundError`, `ConflictError`, `ValidationError`, `RateLimitError`,
`ChannelError`, `ServerError`.

MIT License.
