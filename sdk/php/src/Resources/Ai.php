<?php

namespace HiTechCloud\SaleAI\Resources;

/** Trợ lý AI bán hàng (config / train / suggest / handover) */
class Ai extends BaseResource
{
    /** PATCH /api/v1/ai/config */
    public function updateConfig(array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/ai/config", $data);
    }

    /** POST /api/v1/ai/train */
    public function train(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/ai/train", $data);
    }

    /** POST /api/v1/ai/suggest-reply */
    public function suggestReply(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/ai/suggest-reply", $data);
    }

    /** POST /api/v1/ai/handover */
    public function handover(array $data = []): ?array
    {
        return $this->http->request('POST', "/api/v1/ai/handover", $data);
    }

}
