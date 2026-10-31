<?php

namespace HiTechCloud\SaleAI\Resources;

/** Gửi tin nhắn (text/image/file/quick-reply/button/carousel). */
class Messages extends BaseResource
{
    private function sendRaw(string $botId, array $message, ?string $pat = null): ?array
    {
        $headers = $pat ? ['X-Page-Access-Token' => $pat] : [];
        return $this->http->request('POST', "/api/v1/bots/{$botId}/messages", $message, null, $headers);
    }

    public function send(string $botId, array $message, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, $message, $pageAccessToken);
    }

    public function sendText(string $botId, string $subscriberId, string $text, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'text', 'text' => $text], $pageAccessToken);
    }

    public function sendImage(string $botId, string $subscriberId, string $url, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'image', 'url' => $url], $pageAccessToken);
    }

    public function sendFile(string $botId, string $subscriberId, string $url, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'file', 'url' => $url], $pageAccessToken);
    }

    public function sendQuickReply(string $botId, string $subscriberId, string $text, array $quickReplies, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'quick_reply', 'text' => $text, 'quick_replies' => $quickReplies], $pageAccessToken);
    }

    public function sendButton(string $botId, string $subscriberId, string $text, array $buttons, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'button', 'text' => $text, 'buttons' => $buttons], $pageAccessToken);
    }

    public function sendCarousel(string $botId, string $subscriberId, array $elements, ?string $pageAccessToken = null): ?array
    {
        return $this->sendRaw($botId, ['subscriber_id' => $subscriberId, 'type' => 'carousel', 'elements' => $elements], $pageAccessToken);
    }
}
