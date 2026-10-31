<?php

namespace HiTechCloud\SaleAI\Resources;

/** Chuỗi tin hẹn giờ (sequences) và các bước */
class Sequences extends BaseResource
{
    /** GET /api/v1/sequences */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/sequences", null, $params);
    }

    /** GET /api/v1/sequences/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/sequences/{$id}");
    }

    /** POST /api/v1/sequences */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/sequences", $data);
    }

    /** PATCH /api/v1/sequences/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/sequences/{$id}", $data);
    }

    /** DELETE /api/v1/sequences/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/sequences/{$id}");
    }

    /** GET /api/v1/sequences/{id}/steps */
    public function listSteps(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/sequences/{$id}/steps");
    }

    /** POST /api/v1/sequences/{id}/steps */
    public function addStep(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/sequences/{$id}/steps", $data);
    }

    /** PATCH /api/v1/sequences/{id}/steps/{step_id} */
    public function updateStep(string $id, string $stepId, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/sequences/{$id}/steps/{$stepId}", $data);
    }

    /** DELETE /api/v1/sequences/{id}/steps/{step_id} */
    public function deleteStep(string $id, string $stepId): ?array
    {
        return $this->http->request('DELETE', "/api/v1/sequences/{$id}/steps/{$stepId}");
    }

    /** POST /api/v1/sequences/{id}/subscribe */
    public function subscribe(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/sequences/{$id}/subscribe", $data);
    }

    /** POST /api/v1/sequences/{id}/unsubscribe */
    public function unsubscribe(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/sequences/{$id}/unsubscribe", $data);
    }

}
