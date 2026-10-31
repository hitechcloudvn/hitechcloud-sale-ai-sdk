<?php

namespace HiTechCloud\SaleAI\Resources;

/** Đăng ký webhook (webhooks) */
class Webhooks extends BaseResource
{
    /** GET /api/v1/webhooks */
    public function list(): ?array
    {
        return $this->http->request('GET', "/api/v1/webhooks");
    }

    /** GET /api/v1/webhooks/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/webhooks/{$id}");
    }

    /** POST /api/v1/webhooks */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/webhooks", $data);
    }

    /** PATCH /api/v1/webhooks/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/webhooks/{$id}", $data);
    }

    /** DELETE /api/v1/webhooks/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/webhooks/{$id}");
    }

}
