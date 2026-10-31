<?php

namespace HiTechCloud\SaleAI\Resources;

use HiTechCloud\SaleAI\HttpClient;

class BaseResource
{
    protected HttpClient $http;

    public function __construct(HttpClient $http)
    {
        $this->http = $http;
    }
}
