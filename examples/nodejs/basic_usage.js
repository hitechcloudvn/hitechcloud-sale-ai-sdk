/**
 * Basic example - HiTechCloud Sale AI Node.js SDK. / Ví dụ cơ bản - HiTechCloud Sale AI Node.js SDK.
 *
 * Run: / Cách chạy:
 *   export SALE_AI_TOKEN=your-token
 *   export SALE_AI_PAGE_ACCESS_TOKEN=page-token
 */
const { SaleAI, verifySignature, computeSignature } = require("@hitechcloudvn/sale-ai-sdk");

async function main() {
  const client = new SaleAI({
    token: process.env.SALE_AI_TOKEN,
    pageAccessToken: process.env.SALE_AI_PAGE_ACCESS_TOKEN,
  });

  // List bots / Liệt kê bot
  const bots = await client.bots.list();
  console.log(bots.data);

  // Send a quick-reply message / Gửi tin quick-reply
  await client.messages.sendQuickReply(
    "bot_7c1a9f2e",
    "sub_2a9b8c7d",
    "Ban muon xem gi?",
    ["Ao", "Quan", "Phu kien"],
  );

  // Overview analytics report / Báo cáo tổng quan
  const report = await client.analytics.overview(undefined, { bot_id: "bot_7c1a9f2e" });
  console.log(report);

  // Verify the webhook signature / Xác minh chữ ký webhook
  const raw = '{"event":"order.created"}';
  const sig = computeSignature(raw, "whsec_demo");
  console.log("signature ok:", verifySignature(raw, sig, "whsec_demo"));
}

main().catch(console.error);
