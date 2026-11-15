<?php
/**
 * Basic example - HiTechCloud Sale AI PHP SDK. / Ví dụ cơ bản - HiTechCloud Sale AI PHP SDK.
 *
 * Run: / Cách chạy:
 *   export SALE_AI_TOKEN=your-token
 *   export SALE_AI_PAGE_ACCESS_TOKEN=page-token
 */
require __DIR__ . '/../../sdk/php/vendor/autoload.php';

use HiTechCloud\SaleAI\SaleAI;
use HiTechCloud\SaleAI\Webhook;

$client = new SaleAI(
    token: getenv('SALE_AI_TOKEN') ?: 'YOUR_TOKEN',
    pageAccessToken: getenv('SALE_AI_PAGE_ACCESS_TOKEN') ?: null,
);

// List bots / Liệt kê bot
$bots = $client->bots->list();
foreach ($bots['data'] ?? [] as $bot) {
    echo $bot['id'] . ' ' . $bot['name'] . PHP_EOL;
}

// Send a text message / Gửi tin nhắn văn bản
$client->messages->sendText('bot_7c1a9f2e', 'sub_2a9b8c7d', 'Chao ban!');

// Verify the webhook signature / Xác minh chữ ký webhook
$raw = '{"event":"message.received"}';
$sig = Webhook::computeSignature($raw, 'whsec_demo');
var_dump(Webhook::verifySignature($raw, $sig, 'whsec_demo'));
