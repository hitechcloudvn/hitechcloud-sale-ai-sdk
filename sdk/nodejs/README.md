# HiTechCloud Sale AI - Node.js / TypeScript SDK

SDK TypeScript chính thức cho **HiTechCloud Sale AI**.

- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`

## Cài đặt

```bash
npm install @hitechcloud-vietnam/sale-ai-sdk
```

## Bắt đầu nhanh

```ts
import { SaleAI } from "@hitechcloud-vietnam/sale-ai-sdk";

const client = new SaleAI({ token: "YOUR_TOKEN" });

const bots = await client.bots.list();

await client.messages.sendText(
  "bot_7c1a9f2e",
  "sub_2a9b8c7d",
  "Chao ban!",
  "PAGE_ACCESS_TOKEN",
);
```

## Webhook

```ts
import { verifySignature, SIGNATURE_HEADER } from "@hitechcloud-vietnam/sale-ai-sdk";

const ok = verifySignature(rawBody, req.headers[SIGNATURE_HEADER.toLowerCase()], "whsec_...");
```

## Build

```bash
npm run build   # tsc -> dist/
```

MIT License.
