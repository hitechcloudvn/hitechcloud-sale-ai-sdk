<?php

namespace HiTechCloud\SaleAI\Resources;

/** Hộp thư đa kênh (conversations / inbox) */
class Conversations extends BaseResource
{
    /** GET /api/v1/conversations */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/conversations", null, $params);
    }

    /** GET /api/v1/conversations/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/conversations/{$id}");
    }

    /** POST /api/v1/conversations/{id}/assign */
    public function assign(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/conversations/{$id}/assign", $data);
    }

    /** POST /api/v1/conversations/{id}/notes */
    public function addNote(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/conversations/{$id}/notes", $data);
    }

    /** POST /api/v1/conversations/{id}/labels */
    public function addLabels(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/conversations/{$id}/labels", $data);
    }

    /** POST /api/v1/conversations/{id}/close */
    public function close(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/conversations/{$id}/close");
    }

}
