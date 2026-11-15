# HiTechCloud Sale AI - MCP server

> [English](README.md) | **Tiếng Việt**

---

MCP (Model Context Protocol) server cho HiTechCloud Sale AI. Cho phép trợ lý AI
gọi trực tiếp các thao tác chính: liệt kê bot, xuất bản kịch bản, gửi tin, xem
báo cáo, gợi ý trả lời...

## Cài đặt & chạy

```bash
cd mcp
npm install
export SALE_AI_TOKEN=YOUR_TOKEN
export SALE_AI_PAGE_ACCESS_TOKEN=PAGE_ACCESS_TOKEN   # tuỳ chọn, cho gửi tin
node server.js
```

## Cấu hình client MCP

```json
{
  "mcpServers": {
    "hitechcloud-sale-ai": {
      "command": "node",
      "args": ["/path/to/mcp/server.js"],
      "env": { "SALE_AI_TOKEN": "YOUR_TOKEN" }
    }
  }
}
```

## Tool

`list_bots`, `get_bot`, `list_flows`, `publish_flow`, `list_subscribers`,
`send_text`, `list_broadcasts`, `send_broadcast`, `list_orders`,
`list_conversations`, `analytics_overview`, `suggest_reply`.
