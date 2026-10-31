<?php

namespace HiTechCloud\SaleAI\Resources;

/** Thao tác bình luận trên kênh (comments) */
class Comments extends BaseResource
{
    /** POST /api/v1/comments/{comment_id}/hide */
    public function hide(string $commentId, ?string $pageAccessToken = null): ?array
    {
        return $this->http->request('POST', "/api/v1/comments/{$commentId}/hide", null, null, $pageAccessToken ? ['X-Page-Access-Token' => $pageAccessToken] : []);
    }

    /** POST /api/v1/comments/{comment_id}/like */
    public function like(string $commentId, ?string $pageAccessToken = null): ?array
    {
        return $this->http->request('POST', "/api/v1/comments/{$commentId}/like", null, null, $pageAccessToken ? ['X-Page-Access-Token' => $pageAccessToken] : []);
    }

}
