<?php

namespace HiTechCloud\SaleAI\Resources;

/** Ý định AI (intents) */
class Intents extends BaseResource
{
    /** GET /api/v1/ai/intents */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/ai/intents", null, $params);
    }

    /** GET /api/v1/ai/intents/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/ai/intents/{$id}");
    }

    /** POST /api/v1/ai/intents */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/ai/intents", $data);
    }

    /** PATCH /api/v1/ai/intents/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/ai/intents/{$id}", $data);
    }

    /** DELETE /api/v1/ai/intents/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/ai/intents/{$id}");
    }

}
