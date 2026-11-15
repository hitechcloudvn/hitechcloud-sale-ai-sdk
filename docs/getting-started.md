# Quick start

> **English** | [Tiếng Việt](getting-started.vi.md)

---

HiTechCloud Sale AI helps you sell and automate conversations directly in
messaging channels (Facebook Messenger, Instagram, Zalo, Website).

## 1. Get an access token

1. Sign in to the admin portal at `https://sale-ai.hitechcloud.vn`.
2. Go to **Settings → API & Integration → Access Token**.
3. Create a new token, choose the scopes and copy it (shown only once).
4. Store the token in an environment variable; never commit it to the repository.

## 2. Install the SDK

| Language | Command |
|----------|---------|
| Node.js / TS | `npm install @hitechcloudvn/sale-ai-sdk` |
| Python | `pip install hitechcloud-sale-ai` |
| PHP | `composer require hitechcloudvn/sale-ai-sdk` |
| Go | `go get github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go` |
| Java | `vn.hitechcloud:sale-ai-sdk` (Maven) |

## 3. Make your first API call

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")
bots = client.bots.list()
print(bots["data"])
```

## 4. Send your first message

Message-sending endpoints also require the `X-Page-Access-Token` header (the
bot's page token):

```python
client.messages.send_text(
    "bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
    page_access_token="PAGE_ACCESS_TOKEN",
)
```

## 5. Sandbox environment

Use `base_url="https://sale-ai-sandbox.hitechcloud.vn"` for testing. On
sandbox, messages are NOT delivered to real channels, but webhooks still emit
real events so you can test your receiver.
