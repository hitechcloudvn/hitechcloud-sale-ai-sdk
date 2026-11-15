# HiTechCloud Sale AI SDK

> **English** | [Tiếng Việt](README.vi.md)

---

[![CI](https://github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

Official multi-language SDK for **HiTechCloud Sale AI** — the multichannel
sales chatbot and conversation automation platform (Facebook Messenger,
Instagram, Zalo, Website): flows, keywords, message sequences, broadcasts,
an AI assistant, in-conversation order closing and real-time webhooks.

- Production host: `https://sale-ai.hitechcloud.vn`
- Sandbox host: `https://sale-ai-sandbox.hitechcloud.vn`
- Base path: `/api/v1`
- Authentication: `Authorization: Bearer <token>` + `X-Page-Access-Token` (channel operations)
- Webhook signature: `X-SaleAI-Signature` (HMAC-SHA256)

## Supported languages

| Language | Package | Directory |
|----------|---------|-----------|
| Node.js / TypeScript | `@hitechcloudvn/sale-ai-sdk` | [`sdk/nodejs`](sdk/nodejs) |
| Python | `hitechcloud-sale-ai` | [`sdk/python`](sdk/python) |
| PHP | `hitechcloudvn/sale-ai-sdk` (namespace `HiTechCloud\SaleAI`) | [`sdk/php`](sdk/php) |
| Go | `github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go` | [`sdk/go`](sdk/go) |
| Java | `vn.hitechcloud:sale-ai-sdk` | [`sdk/java`](sdk/java) |

## Features

- Full CRUD: bots, flows (+blocks), keywords, sequences (+steps), broadcasts,
  subscribers (+tags/custom-fields/notes), tags, segments, products, orders,
  comment-rules, widgets, AI intents, knowledge base, users, roles, webhooks.
- Rich message sending: text, image, file, quick-reply, button, carousel.
- AI assistant: configuration, training, suggested replies, human handover.
- Unified inbox (conversations/inbox), analytics, settings.
- HMAC-SHA256 webhook signature verification helpers in all 5 languages.
- Page access token handling and automatic retry for `429`/`5xx`.

## Quick start (Python)

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")
bots = client.bots.list()
client.messages.send_text("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
                          page_access_token="PAGE_ACCESS_TOKEN")
```

## Documentation

See the [`docs/`](docs/README.md) directory: quick start, authentication, API
reference, error handling, pagination and rate limits. OpenAPI specification:
[`openapi.json`](openapi.json).

## MCP server

[`mcp/`](mcp/README.md) — MCP server (Node) that lets AI assistants call
HiTechCloud Sale AI directly.

## License

[MIT](LICENSE) © HiTechCloud Vietnam
