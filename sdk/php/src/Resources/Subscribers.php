<?php

namespace HiTechCloud\SaleAI\Resources;

/** Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú */
class Subscribers extends BaseResource
{
    /** GET /api/v1/subscribers */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/subscribers", null, $params);
    }

    /** GET /api/v1/subscribers/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/subscribers/{$id}");
    }

    /** POST /api/v1/subscribers */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers", $data);
    }

    /** PATCH /api/v1/subscribers/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/subscribers/{$id}", $data);
    }

    /** DELETE /api/v1/subscribers/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/subscribers/{$id}");
    }

    /** POST /api/v1/subscribers/{id}/tags */
    public function addTags(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/tags", $data);
    }

    /** DELETE /api/v1/subscribers/{id}/tags/{tag} */
    public function removeTag(string $id, string $tag): ?array
    {
        return $this->http->request('DELETE', "/api/v1/subscribers/{$id}/tags/{$tag}");
    }

    /** PATCH /api/v1/subscribers/{id}/custom-fields */
    public function setCustomFields(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/subscribers/{$id}/custom-fields", $data);
    }

    /** POST /api/v1/subscribers/{id}/notes */
    public function addNote(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/notes", $data);
    }

    /** POST /api/v1/subscribers/{id}/block */
    public function block(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/block");
    }

    /** POST /api/v1/subscribers/{id}/unblock */
    public function unblock(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/unblock");
    }

    /** GET /api/v1/subscribers/{id}/messages */
    public function listMessages(string $id, array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/subscribers/{$id}/messages", null, $params);
    }

    /** POST /api/v1/subscribers/{id}/messages */
    public function sendMessage(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/messages", $data);
    }

    /** POST /api/v1/subscribers/{id}/mark-seen */
    public function markSeen(string $id): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/mark-seen");
    }

    /** POST /api/v1/subscribers/{id}/typing */
    public function typing(string $id, array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/subscribers/{$id}/typing", $data);
    }

}
