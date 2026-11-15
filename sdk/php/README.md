# HiTechCloud Sale AI - PHP SDK

> **English** | [Tiếng Việt](README.vi.md)

---

Official PHP SDK for **HiTechCloud Sale AI** (PHP 8+, Guzzle).

- Host: `https://sale-ai.hitechcloud.vn`
- Base path: `/api/v1`
- Namespace: `HiTechCloud\SaleAI`

## Installation

```bash
composer require hitechcloudvn/sale-ai-sdk
```

## Quick start

```php
<?php
require 'vendor/autoload.php';

use HiTechCloud\SaleAI\SaleAI;

$client = new SaleAI(token: 'YOUR_TOKEN');

$bots = $client->bots->list();

$client->messages->sendText(
    'bot_7c1a9f2e', 'sub_2a9b8c7d', 'Chao ban!', 'PAGE_ACCESS_TOKEN'
);
```

## Webhook

```php
use HiTechCloud\SaleAI\Webhook;

$ok = Webhook::verifySignature($rawBody, $_SERVER['HTTP_X_SALEAI_SIGNATURE'] ?? null, 'whsec_...');
```

MIT License.
