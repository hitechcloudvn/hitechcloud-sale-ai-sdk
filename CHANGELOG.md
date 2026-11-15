# Changelog

> **English** | [Tiếng Việt](CHANGELOG.vi.md)

---

All notable changes to the HiTechCloud Sale AI SDK are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/), and this
project adheres to [Semantic Versioning](https://semver.org/).

## [1.0.1] - 2026-11-16

### Fixed
- PHP SDK: add `phpunit.xml.dist` and smoke tests so `vendor/bin/phpunit`
  runs the suite instead of printing help and failing the test job.

## [1.0.0] - 2026-11-01

### Added
- First release of the multi-language SDK for HiTechCloud Sale AI.
- Five languages supported: Node.js/TypeScript, Python, PHP, Go, Java.
- Full resource coverage: bots, flows (+blocks), keywords, sequences
  (+steps), broadcasts, subscribers (+tags/custom-fields/notes), tags,
  segments, products, orders, comment-rules, widgets, AI intents, knowledge
  base, users, roles, webhooks, comments, AI assistant, conversations/inbox,
  analytics, settings and messaging (text/image/file/quick-reply/button/carousel).
- `X-SaleAI-Signature` (HMAC-SHA256) webhook signature verification helpers in
  all 5 languages.
- Page access token handling (`X-Page-Access-Token`): client default or per call.
- Automatic retry for `429` (honoring `Retry-After`) and `5xx` (exponential
  backoff).
- `openapi.json` specification, Vietnamese documentation and an MCP server (Node).
