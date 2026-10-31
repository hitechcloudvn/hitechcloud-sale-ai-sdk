<?php

namespace HiTechCloud\SaleAI\Resources;

/** Gửi hàng loạt (broadcasts) */
class Broadcasts extends BaseResource
{
    /** GET /api/v1/broadcasts */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/broadcasts", null, $params);
    }

    /** GET /api/v1/broadcasts/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/broadcasts/{$id}");
    }

    /** POST /api/v1/broadcasts */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/broadcasts", $data);
    }

    /** PATCH /api/v1/broadcasts/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/broadcasts/{$id}", $data);
    }

    /** DELETE /api/v1/broadcasts/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/broadcasts/{$id}");
    }

    /** POST /api/v1/broadcasts/{id}/schedule */
    public function schedule(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/broadcasts/{$id}/schedule", $data);
    }

    /** POST /api/v1/broadcasts/{id}/send */
    public function send(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/broadcasts/{$id}/send");
    }

    /** GET /api/v1/broadcasts/{id}/status */
    public function status(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/broadcasts/{$id}/status");
    }

    /** POST /api/v1/broadcasts/{id}/audience/preview */
    public function previewAudience(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/broadcasts/{$id}/audience/preview", $data);
    }

    /** POST /api/v1/broadcasts/{id}/cancel */
    public function cancel(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/broadcasts/{$id}/cancel");
    }

}
