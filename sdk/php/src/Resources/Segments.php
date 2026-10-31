<?php

namespace HiTechCloud\SaleAI\Resources;

/** Phân khúc (segments) */
class Segments extends BaseResource
{
    /** GET /api/v1/segments */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/segments", null, $params);
    }

    /** GET /api/v1/segments/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/segments/{$id}");
    }

    /** POST /api/v1/segments */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/segments", $data);
    }

    /** PATCH /api/v1/segments/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/segments/{$id}", $data);
    }

    /** DELETE /api/v1/segments/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/segments/{$id}");
    }

}
