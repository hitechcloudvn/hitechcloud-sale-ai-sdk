<?php
/**
 * Ví dụ cơ bản - HiTechCloud Sale AI PHP SDK.
 */
require __DIR__ . '/../../sdk/php/vendor/autoload.php';

use HiTechCloud\SaleAI\SaleAI;
use HiTechCloud\SaleAI\Webhook;

$client = new SaleAI(
    token: getenv('SALE_AI_TOKEN') ?: 'YOUR_TOKEN',
    pageAccessToken: getenv('SALE_AI_PAGE_ACCESS_TOKEN') ?: null,
);

$bots = $client->bots->list();
foreach ($bots['data'] ?? [] as $bot) {
    echo $bot['id'] . ' ' . $bot['name'] . PHP_EOL;
}

$client->messages->sendText('bot_7c1a9f2e', 'sub_2a9b8c7d', 'Chao ban!');

$raw = '{"event":"message.received"}';
$sig = Webhook::computeSignature($raw, 'whsec_demo');
var_dump(Webhook::verifySignature($raw, $sig, 'whsec_demo'));
