# HiTechCloud Sale AI SDK

[![CI](https://github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

SDK đa ngôn ngữ chính thức cho **HiTechCloud Sale AI** — nền tảng chatbot bán
hàng & tự động hoá hội thoại đa kênh (Facebook Messenger, Instagram, Zalo,
Website): kịch bản, từ khoá, chuỗi tin, broadcast, trợ lý AI, chốt đơn ngay
trong hội thoại và webhook thời gian thực.

- Host production: `https://sale-ai.hitechcloud.vn`
- Host sandbox: `https://sale-ai-sandbox.hitechcloud.vn`
- Base path: `/api/v1`
- Xác thực: `Authorization: Bearer <token>` + `X-Page-Access-Token` (thao tác kênh)
- Webhook signature: `X-SaleAI-Signature` (HMAC-SHA256)

## Ngôn ngữ hỗ trợ

| Ngôn ngữ | Gói | Thư mục |
|----------|-----|---------|
| Node.js / TypeScript | `@hitechcloudvn/sale-ai-sdk` | [`sdk/nodejs`](sdk/nodejs) |
| Python | `hitechcloud-sale-ai` | [`sdk/python`](sdk/python) |
| PHP | `hitechcloudvn/sale-ai-sdk` (ns `HiTechCloud\SaleAI`) | [`sdk/php`](sdk/php) |
| Go | `github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go` | [`sdk/go`](sdk/go) |
| Java | `vn.hitechcloud:sale-ai-sdk` | [`sdk/java`](sdk/java) |

## Tính năng

- CRUD đầy đủ: bots, flows (+blocks), keywords, sequences (+steps), broadcasts,
  subscribers (+tags/custom-fields/notes), tags, segments, products, orders,
  comment-rules, widgets, AI intents, knowledge base, users, roles, webhooks.
- Gửi tin đa định dạng: text, image, file, quick-reply, button, carousel.
- Trợ lý AI: cấu hình, huấn luyện, gợi ý trả lời, bàn giao người thật.
- Hộp thư đa kênh (conversations/inbox), analytics, settings.
- Helper xác minh chữ ký webhook HMAC-SHA256 ở cả 5 ngôn ngữ.
- Xử lý page access token, tự động retry `429`/`5xx`.

## Bắt đầu nhanh (Python)

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")
bots = client.bots.list()
client.messages.send_text("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
                          page_access_token="PAGE_ACCESS_TOKEN")
```

## Tài liệu

Xem thư mục [`docs/`](docs/README.md): bắt đầu nhanh, xác thực, tham chiếu API,
xử lý lỗi, phân trang & giới hạn tốc độ. Đặc tả OpenAPI: [`openapi.json`](openapi.json).

## MCP server

[`mcp/`](mcp/README.md) — MCP server (Node) để trợ lý AI gọi trực tiếp Sale AI.

## Giấy phép

[MIT](LICENSE) © HiTechCloud Vietnam
