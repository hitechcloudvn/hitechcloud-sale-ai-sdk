<?php

namespace HiTechCloud\SaleAI\Resources;

/** Người dùng / nhân sự (team) */
class Users extends BaseResource
{
    /** GET /api/v1/users */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/users", null, $params);
    }

    /** GET /api/v1/users/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/users/{$id}");
    }

    /** POST /api/v1/users */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/users", $data);
    }

    /** PATCH /api/v1/users/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/users/{$id}", $data);
    }

    /** DELETE /api/v1/users/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/users/{$id}");
    }

}
