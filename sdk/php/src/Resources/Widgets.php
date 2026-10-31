<?php

namespace HiTechCloud\SaleAI\Resources;

/** Công cụ tăng trưởng (widgets) */
class Widgets extends BaseResource
{
    /** GET /api/v1/widgets */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/widgets", null, $params);
    }

    /** GET /api/v1/widgets/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/widgets/{$id}");
    }

    /** POST /api/v1/widgets */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/widgets", $data);
    }

    /** PATCH /api/v1/widgets/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/widgets/{$id}", $data);
    }

    /** DELETE /api/v1/widgets/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/widgets/{$id}");
    }

}
