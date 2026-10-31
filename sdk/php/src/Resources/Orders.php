<?php

namespace HiTechCloud\SaleAI\Resources;

/** Đơn hàng (orders) */
class Orders extends BaseResource
{
    /** GET /api/v1/orders */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/orders", null, $params);
    }

    /** GET /api/v1/orders/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/orders/{$id}");
    }

    /** POST /api/v1/orders */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/orders", $data);
    }

    /** PATCH /api/v1/orders/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/orders/{$id}", $data);
    }

    /** DELETE /api/v1/orders/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/orders/{$id}");
    }

    /** POST /api/v1/orders/from-conversation */
    public function createFromConversation(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/orders/from-conversation", $data);
    }

    /** POST /api/v1/orders/{id}/status */
    public function updateStatus(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/orders/{$id}/status", $data);
    }

}
