<?php

namespace HiTechCloud\SaleAI\Resources;

/** Quy tắc trả lời bình luận (comment rules) */
class CommentRules extends BaseResource
{
    /** GET /api/v1/comment-rules */
    public function list(array $params = []): ?array
    {
        return $this->http->request('GET', "/api/v1/comment-rules", null, $params);
    }

    /** GET /api/v1/comment-rules/{id} */
    public function get(string $id): ?array
    {
        return $this->http->request('GET', "/api/v1/comment-rules/{$id}");
    }

    /** POST /api/v1/comment-rules */
    public function create(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/comment-rules", $data);
    }

    /** PATCH /api/v1/comment-rules/{id} */
    public function update(string $id, array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/comment-rules/{$id}", $data);
    }

    /** DELETE /api/v1/comment-rules/{id} */
    public function delete(string $id): ?array
    {
        return $this->http->request('DELETE', "/api/v1/comment-rules/{$id}");
    }

}
