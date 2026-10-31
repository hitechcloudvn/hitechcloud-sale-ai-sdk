/**
 * Ví dụ cơ bản - HiTechCloud Sale AI Node.js SDK.
 *
 *   export SALE_AI_TOKEN=your-token
 *   export SALE_AI_PAGE_ACCESS_TOKEN=page-token
 */
const { SaleAI, verifySignature, computeSignature } = require("@hitechcloud-vietnam/sale-ai-sdk");

async function main() {
  const client = new SaleAI({
    token: process.env.SALE_AI_TOKEN,
    pageAccessToken: process.env.SALE_AI_PAGE_ACCESS_TOKEN,
  });

  const bots = await client.bots.list();
  console.log(bots.data);

  await client.messages.sendQuickReply(
    "bot_7c1a9f2e",
    "sub_2a9b8c7d",
    "Ban muon xem gi?",
    ["Ao", "Quan", "Phu kien"],
  );

  const report = await client.analytics.overview(undefined, { bot_id: "bot_7c1a9f2e" });
  console.log(report);

  const raw = '{"event":"order.created"}';
  const sig = computeSignature(raw, "whsec_demo");
  console.log("signature ok:", verifySignature(raw, sig, "whsec_demo"));
}

main().catch(console.error);
