# HiTechCloud Sale AI - Java SDK

SDK Java chính thức cho **HiTechCloud Sale AI**.

- Maven: `vn.hitechcloud:sale-ai-sdk`
- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`

## Cài đặt (Maven)

```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>sale-ai-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Bắt đầu nhanh

```java
import vn.hitechcloud.saleai.SaleAI;

SaleAI client = new SaleAI("YOUR_TOKEN");

var bots = client.bots.list(null);

client.messages.sendText("bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!", "PAGE_ACCESS_TOKEN");
```

## Webhook

```java
import vn.hitechcloud.saleai.WebhookVerifier;

boolean ok = WebhookVerifier.verifySignature(rawBody, signatureHeader, "whsec_...");
```

MIT License.
