<?php

namespace HiTechCloud\SaleAI\Resources;

/** Báo cáo & phân tích (analytics) */
class Analytics extends BaseResource
{
    /** GET /api/v1/analytics/overview */
    public function overview(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/analytics/overview", null, $params);
    }

    /** GET /api/v1/analytics/bots/{bot_id} */
    public function bot(string $botId, array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/analytics/bots/{$botId}", null, $params);
    }

    /** GET /api/v1/analytics/flows/{id} */
    public function flow(string $id, array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/analytics/flows/{$id}", null, $params);
    }

    /** GET /api/v1/analytics/broadcasts/{id} */
    public function broadcast(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/analytics/broadcasts/{$id}");
    }

    /** GET /api/v1/analytics/subscriber-growth */
    public function subscriberGrowth(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/analytics/subscriber-growth", null, $params);
    }

}
