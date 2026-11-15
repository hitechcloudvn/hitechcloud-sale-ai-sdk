# HiTechCloud Sale AI - MCP server

> **English** | [Tiếng Việt](README.vi.md)

---

MCP (Model Context Protocol) server for HiTechCloud Sale AI. It lets AI
assistants invoke core operations directly: list bots, publish flows, send
messages, view reports, suggest replies, and more.

## Install & run

```bash
cd mcp
npm install
export SALE_AI_TOKEN=YOUR_TOKEN
export SALE_AI_PAGE_ACCESS_TOKEN=PAGE_ACCESS_TOKEN   # optional, for sending messages
node server.js
```

## MCP client configuration

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

## Tools

`list_bots`, `get_bot`, `list_flows`, `publish_flow`, `list_subscribers`,
`send_text`, `list_broadcasts`, `send_broadcast`, `list_orders`,
`list_conversations`, `analytics_overview`, `suggest_reply`.
