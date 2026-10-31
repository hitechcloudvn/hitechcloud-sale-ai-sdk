<?php

namespace HiTechCloud\SaleAI\Resources;

/** Nhãn (tags) */
class Tags extends BaseResource
{
    /** GET /api/v1/tags */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/tags", null, $params);
    }

    /** GET /api/v1/tags/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/tags/{$id}");
    }

    /** POST /api/v1/tags */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/tags", $data);
    }

    /** PATCH /api/v1/tags/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/tags/{$id}", $data);
    }

    /** DELETE /api/v1/tags/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/tags/{$id}");
    }

}
