# Đóng góp

> [English](CONTRIBUTING.md) | **Tiếng Việt**

---

Cảm ơn bạn quan tâm đóng góp cho HiTechCloud Sale AI SDK!

## Cấu trúc kho mã

```
sdk/nodejs   - SDK TypeScript (@hitechcloudvn/sale-ai-sdk)
sdk/python   - SDK Python (hitechcloud-sale-ai)
sdk/php      - SDK PHP (hitechcloudvn/sale-ai-sdk)
sdk/go       - SDK Go
sdk/java     - SDK Java (vn.hitechcloud:sale-ai-sdk)
mcp          - MCP server (Node)
docs         - Tài liệu tiếng Việt
examples     - Ví dụ theo từng ngôn ngữ
openapi.json - Đặc tả OpenAPI 3.0
```

## Nguyên tắc

- `openapi.json` và các SDK phải bám sát endpoint thật của Sale AI.
- Giữ chữ ký phương thức nhất quán giữa các ngôn ngữ.
- Mọi thao tác kênh (gửi tin, ẩn/thích bình luận) phải hỗ trợ `X-Page-Access-Token`.
- Thêm/sửa endpoint: cập nhật đồng thời cả 5 SDK, `openapi.json` và `docs/`.

## Kiểm tra cú pháp nhanh

```bash
python -m py_compile sdk/python/hitechcloud_sale_ai/*.py sdk/python/hitechcloud_sale_ai/resources/*.py
node --check mcp/server.js
php -l sdk/php/src/HttpClient.php
```

## Quy trình

1. Fork và tạo nhánh tính năng.
2. Commit rõ ràng, có kiểm thử khi cần.
3. Mở Pull Request mô tả thay đổi.
