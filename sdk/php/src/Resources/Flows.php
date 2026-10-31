<?php

namespace HiTechCloud\SaleAI\Resources;

/** Kịch bản hội thoại (flows) và khối (blocks) */
class Flows extends BaseResource
{
    /** GET /api/v1/flows */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/flows", null, $params);
    }

    /** GET /api/v1/flows/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/flows/{$id}");
    }

    /** POST /api/v1/flows */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/flows", $data);
    }

    /** PATCH /api/v1/flows/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/flows/{$id}", $data);
    }

    /** DELETE /api/v1/flows/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/flows/{$id}");
    }

    /** POST /api/v1/flows/{id}/publish */
    public function publish(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/flows/{$id}/publish");
    }

    /** POST /api/v1/flows/{id}/duplicate */
    public function duplicate(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/flows/{$id}/duplicate");
    }

    /** POST /api/v1/flows/{id}/trigger */
    public function trigger(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/flows/{$id}/trigger", $data);
    }

    /** GET /api/v1/flows/{id}/blocks */
    public function listBlocks(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/flows/{$id}/blocks");
    }

    /** POST /api/v1/flows/{id}/blocks */
    public function addBlock(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/flows/{$id}/blocks", $data);
    }

    /** PATCH /api/v1/flows/{id}/blocks/{block_id} */
    public function updateBlock(string $id, string $blockId, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/flows/{$id}/blocks/{$blockId}", $data);
    }

    /** DELETE /api/v1/flows/{id}/blocks/{block_id} */
    public function deleteBlock(string $id, string $blockId): ?array
    {
        return $this->http->request('DELETE', "/api/v1/flows/{$id}/blocks/{$blockId}");
    }

}
