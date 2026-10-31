<?php

namespace HiTechCloud\SaleAI\Resources;

/** Vai trò & quyền (roles) */
class Roles extends BaseResource
{
    /** GET /api/v1/roles */
    public function list(): ?array
    {
        return $this->http->request('GET', "/api/v1/roles");
    }

    /** GET /api/v1/roles/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/roles/{$id}");
    }

    /** POST /api/v1/roles */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/roles", $data);
    }

    /** PATCH /api/v1/roles/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/roles/{$id}", $data);
    }

    /** DELETE /api/v1/roles/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/roles/{$id}");
    }

}
