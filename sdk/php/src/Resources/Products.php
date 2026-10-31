<?php

namespace HiTechCloud\SaleAI\Resources;

/** Sản phẩm (products) */
class Products extends BaseResource
{
    /** GET /api/v1/products */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/products", null, $params);
    }

    /** GET /api/v1/products/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/products/{$id}");
    }

    /** POST /api/v1/products */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/products", $data);
    }

    /** PATCH /api/v1/products/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/products/{$id}", $data);
    }

    /** DELETE /api/v1/products/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/products/{$id}");
    }

}
