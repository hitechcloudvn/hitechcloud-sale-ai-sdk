#!/usr/bin/env node
/**
 * HiTechCloud Sale AI - MCP server
 *
 * Expose core Sale AI operations as MCP tools over stdio.
 *
 * Env:
 *   SALE_AI_TOKEN             - account access token (Bearer)
 *   SALE_AI_PAGE_ACCESS_TOKEN - default page access token (for sending messages)
 *   SALE_AI_BASE_URL          - override base host (default production)
 */
import { Server } from "@modelcontextprotocol/sdk/server/index.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";

const BASE_URL = process.env.SALE_AI_BASE_URL || "https://sale-ai.hitechcloud.vn";
const TOKEN = process.env.SALE_AI_TOKEN;
const PAGE_ACCESS_TOKEN = process.env.SALE_AI_PAGE_ACCESS_TOKEN;
const API = "/api/v1";

async function api(method, path, { body, params, pageAccessToken } = {}) {
  let url = BASE_URL + API + path;
  if (params) {
    const qs = Object.entries(params)
      .filter(([, v]) => v !== undefined && v !== null)
      .map(([k, v]) => `${encodeURIComponent(k)}=${encodeURIComponent(String(v))}`)
      .join("&");
    if (qs) url += (url.includes("?") ? "&" : "?") + qs;
  }
  const headers = {
    "Content-Type": "application/json",
    Accept: "application/json",
    Authorization: `Bearer ${TOKEN}`,
    "User-Agent": "HiTechCloud-SaleAI-MCP/1.0.0",
  };
  const pat = pageAccessToken || PAGE_ACCESS_TOKEN;
  if (pat) headers["X-Page-Access-Token"] = pat;
  const res = await fetch(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  });
  const text = await res.text();
  return text ? JSON.parse(text) : { status: res.status };
}

const TOOLS = [
  { name: "list_bots", description: "Liệt kê bot", inputSchema: { type: "object", properties: {} } },
  { name: "get_bot", description: "Chi tiết bot theo ID", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },
  { name: "list_flows", description: "Liệt kê kịch bản", inputSchema: { type: "object", properties: { bot_id: { type: "string" } } } },
  { name: "publish_flow", description: "Xuất bản kịch bản", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },
  { name: "list_subscribers", description: "Liệt kê người theo dõi", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, tag: { type: "string" } } } },
  { name: "send_text", description: "Gửi tin nhắn văn bản (cần page access token)", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, subscriber_id: { type: "string" }, text: { type: "string" }, page_access_token: { type: "string" } }, required: ["bot_id", "subscriber_id", "text"] } },
  { name: "list_broadcasts", description: "Liệt kê chiến dịch broadcast", inputSchema: { type: "object", properties: { bot_id: { type: "string" } } } },
  { name: "send_broadcast", description: "Gửi broadcast ngay", inputSchema: { type: "object", properties: { id: { type: "string" } }, required: ["id"] } },
  { name: "list_orders", description: "Liệt kê đơn hàng", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, status: { type: "string" } } } },
  { name: "list_conversations", description: "Liệt kê hội thoại", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, status: { type: "string" } } } },
  { name: "analytics_overview", description: "Báo cáo tổng quan", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, from: { type: "string" }, to: { type: "string" } } } },
  { name: "suggest_reply", description: "Gợi ý câu trả lời cho nhân viên", inputSchema: { type: "object", properties: { bot_id: { type: "string" }, subscriber_id: { type: "string" }, message: { type: "string" } }, required: ["bot_id", "subscriber_id", "message"] } },
];

const HANDLERS = {
  list_bots: () => api("GET", "/bots"),
  get_bot: (a) => api("GET", `/bots/${a.id}`),
  list_flows: (a) => api("GET", "/flows", { params: { bot_id: a.bot_id } }),
  publish_flow: (a) => api("POST", `/flows/${a.id}/publish`),
  list_subscribers: (a) => api("GET", "/subscribers", { params: { bot_id: a.bot_id, tag: a.tag } }),
  send_text: (a) => api("POST", `/bots/${a.bot_id}/messages`, { body: { subscriber_id: a.subscriber_id, type: "text", text: a.text }, pageAccessToken: a.page_access_token }),
  list_broadcasts: (a) => api("GET", "/broadcasts", { params: { bot_id: a.bot_id } }),
  send_broadcast: (a) => api("POST", `/broadcasts/${a.id}/send`),
  list_orders: (a) => api("GET", "/orders", { params: { bot_id: a.bot_id, status: a.status } }),
  list_conversations: (a) => api("GET", "/conversations", { params: { bot_id: a.bot_id, status: a.status } }),
  analytics_overview: (a) => api("GET", "/analytics/overview", { params: { bot_id: a.bot_id, from: a.from, to: a.to } }),
  suggest_reply: (a) => api("POST", "/ai/suggest-reply", { body: { bot_id: a.bot_id, subscriber_id: a.subscriber_id, message: a.message } }),
};

async function main() {
  const server = new Server(
    { name: "hitechcloud-sale-ai", version: "1.0.0" },
    { capabilities: { tools: {} } },
  );

  server.setRequestHandler("tools/list", async () => ({ tools: TOOLS }));

  server.setRequestHandler("tools/call", async (request) => {
    const { name, arguments: args } = request.params;
    const handler = HANDLERS[name];
    if (!handler) throw new Error(`Unknown tool: ${name}`);
    const result = await handler(args || {});
    return { content: [{ type: "text", text: JSON.stringify(result, null, 2) }] };
  });

  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("HiTechCloud Sale AI MCP server running on stdio");
}

main().catch((err) => {
  console.error(err);
  process.exit(1);
});
