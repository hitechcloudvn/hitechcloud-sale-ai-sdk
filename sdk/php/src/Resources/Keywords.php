<?php

namespace HiTechCloud\SaleAI\Resources;

/** Từ khoá tự trả lời (keywords) */
class Keywords extends BaseResource
{
    /** GET /api/v1/keywords */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/keywords", null, $params);
    }

    /** GET /api/v1/keywords/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/keywords/{$id}");
    }

    /** POST /api/v1/keywords */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/keywords", $data);
    }

    /** PATCH /api/v1/keywords/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/keywords/{$id}", $data);
    }

    /** DELETE /api/v1/keywords/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/keywords/{$id}");
    }

    /** POST /api/v1/keywords/test */
    public function test(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/keywords/test", $data);
    }

}
