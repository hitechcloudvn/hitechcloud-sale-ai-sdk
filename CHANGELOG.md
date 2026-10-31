# Changelog

Tất cả thay đổi đáng chú ý sẽ được ghi tại đây. Định dạng theo
[Keep a Changelog](https://keepachangelog.com/) và
[Semantic Versioning](https://semver.org/).

## [1.0.0] - 2026-11-01

### Added
- Phát hành đầu tiên bộ SDK đa ngôn ngữ cho HiTechCloud Sale AI.
- Hỗ trợ 5 ngôn ngữ: Node.js/TypeScript, Python, PHP, Go, Java.
- Bao phủ đầy đủ tài nguyên: bots, flows (+blocks), keywords, sequences
  (+steps), broadcasts, subscribers (+tags/custom-fields/notes), tags,
  segments, products, orders, comment-rules, widgets, AI intents, knowledge
  base, users, roles, webhooks, comments, trợ lý AI, conversations/inbox,
  analytics, settings và gửi tin (text/image/file/quick-reply/button/carousel).
- Helper xác minh chữ ký webhook `X-SaleAI-Signature` (HMAC-SHA256) ở cả 5 ngôn ngữ.
- Xử lý page access token (`X-Page-Access-Token`): mặc định ở client hoặc theo lời gọi.
- Tự động retry cho `429` (theo `Retry-After`) và `5xx` (exponential backoff).
- Đặc tả `openapi.json`, tài liệu tiếng Việt và MCP server (Node).
