<?php

namespace HiTechCloud\SaleAI\Resources;

/** Cài đặt tài khoản (settings) */
class Settings extends BaseResource
{
    /** GET /api/v1/settings */
    public function get(): ?array
    {
        return $this->http->request('GET', "/api/v1/settings");
    }

    /** PATCH /api/v1/settings */
    public function update(array $data = []): ?array
    {
        return $this->http->request('PATCH', "/api/v1/settings", $data);
    }

}
