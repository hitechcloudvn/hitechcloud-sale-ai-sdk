<?php

namespace HiTechCloud\SaleAI\Resources;

/** Bot & kênh kết nối (Facebook/Instagram/Zalo/Website) */
class Bots extends BaseResource
{
    /** GET /api/v1/bots */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/bots", null, $params);
    }

    /** GET /api/v1/bots/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/bots/{$id}");
    }

    /** POST /api/v1/bots */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/bots", $data);
    }

    /** PATCH /api/v1/bots/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/bots/{$id}", $data);
    }

    /** DELETE /api/v1/bots/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/bots/{$id}");
    }

    /** POST /api/v1/bots/{bot_id}/connect */
    public function connect(string $botId, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/bots/{$botId}/connect", $data);
    }

    /** POST /api/v1/bots/{bot_id}/disconnect */
    public function disconnect(string $botId): ?array
    {
        return $this->http->request('POST', "/api/v1/bots/{$botId}/disconnect");
    }

    /** PATCH /api/v1/bots/{bot_id}/settings */
    public function updateSettings(string $botId, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/bots/{$botId}/settings", $data);
    }

    /** POST /api/v1/bots/{bot_id}/access-token/refresh */
    public function refreshAccessToken(string $botId, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/bots/{$botId}/access-token/refresh", $data);
    }

}
