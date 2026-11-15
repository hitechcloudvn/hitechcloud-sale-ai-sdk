# Contributing

> **English** | [Tiếng Việt](CONTRIBUTING.vi.md)

---

Thank you for your interest in contributing to the HiTechCloud Sale AI SDK!

## Repository layout

```
sdk/nodejs   - TypeScript SDK (@hitechcloudvn/sale-ai-sdk)
sdk/python   - Python SDK (hitechcloud-sale-ai)
sdk/php      - PHP SDK (hitechcloudvn/sale-ai-sdk)
sdk/go       - Go SDK
sdk/java     - Java SDK (vn.hitechcloud:sale-ai-sdk)
mcp          - MCP server (Node)
docs         - Documentation
examples     - Per-language examples
openapi.json - OpenAPI 3.0 specification
```

## Principles

- `openapi.json` and the SDKs must track the real HiTechCloud Sale AI endpoints.
- Keep method signatures consistent across languages.
- Every channel operation (send messages, hide/like comments) must support
  `X-Page-Access-Token`.
- When adding or changing an endpoint: update all 5 SDKs, `openapi.json` and
  `docs/` together.

## Quick syntax checks

```bash
python -m py_compile sdk/python/hitechcloud_sale_ai/*.py sdk/python/hitechcloud_sale_ai/resources/*.py
node --check mcp/server.js
php -l sdk/php/src/HttpClient.php
```

## Workflow

1. Fork and create a feature branch.
2. Commit clear, well-tested changes.
3. Open a Pull Request describing the changes.
