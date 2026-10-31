<?php

namespace HiTechCloud\SaleAI\Resources;

/** Kho tri thức AI (knowledge base) */
class KnowledgeBase extends BaseResource
{
    /** GET /api/v1/ai/knowledge-base */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/ai/knowledge-base", null, $params);
    }

    /** GET /api/v1/ai/knowledge-base/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/ai/knowledge-base/{$id}");
    }

    /** POST /api/v1/ai/knowledge-base */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/ai/knowledge-base", $data);
    }

    /** PATCH /api/v1/ai/knowledge-base/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/ai/knowledge-base/{$id}", $data);
    }

    /** DELETE /api/v1/ai/knowledge-base/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/ai/knowledge-base/{$id}");
    }

}
