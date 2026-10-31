<?php

namespace HiTechCloud\SaleAI;

class SaleAI
{
    public const PRODUCTION_URL = 'https://sale-ai.hitechcloud.vn';
    public const SANDBOX_URL = 'https://sale-ai-sandbox.hitechcloud.vn';

    private HttpClient $http;

    public Resources\Bots $bots;
    public Resources\Flows $flows;
    public Resources\Keywords $keywords;
    public Resources\Sequences $sequences;
    public Resources\Broadcasts $broadcasts;
    public Resources\Subscribers $subscribers;
    public Resources\Tags $tags;
    public Resources\Segments $segments;
    public Resources\Products $products;
    public Resources\Orders $orders;
    public Resources\CommentRules $commentRules;
    public Resources\Widgets $widgets;
    public Resources\Intents $intents;
    public Resources\KnowledgeBase $knowledgeBase;
    public Resources\Users $users;
    public Resources\Roles $roles;
    public Resources\Webhooks $webhooks;
    public Resources\Comments $comments;
    public Resources\Ai $ai;
    public Resources\Conversations $conversations;
    public Resources\Analytics $analytics;
    public Resources\Settings $settings;
    public Resources\Messages $messages;

    public function __construct(
        ?string $token = null,
        ?string $pageAccessToken = null,
        string $baseUrl = self::PRODUCTION_URL,
        int $timeout = 30,
        int $maxRetries = 3
    ) {
        $this->http = new HttpClient($baseUrl, $token, $pageAccessToken, $timeout, $maxRetries);
        $this->bots = new Resources\Bots($this->http);
        $this->flows = new Resources\Flows($this->http);
        $this->keywords = new Resources\Keywords($this->http);
        $this->sequences = new Resources\Sequences($this->http);
        $this->broadcasts = new Resources\Broadcasts($this->http);
        $this->subscribers = new Resources\Subscribers($this->http);
        $this->tags = new Resources\Tags($this->http);
        $this->segments = new Resources\Segments($this->http);
        $this->products = new Resources\Products($this->http);
        $this->orders = new Resources\Orders($this->http);
        $this->commentRules = new Resources\CommentRules($this->http);
        $this->widgets = new Resources\Widgets($this->http);
        $this->intents = new Resources\Intents($this->http);
        $this->knowledgeBase = new Resources\KnowledgeBase($this->http);
        $this->users = new Resources\Users($this->http);
        $this->roles = new Resources\Roles($this->http);
        $this->webhooks = new Resources\Webhooks($this->http);
        $this->comments = new Resources\Comments($this->http);
        $this->ai = new Resources\Ai($this->http);
        $this->conversations = new Resources\Conversations($this->http);
        $this->analytics = new Resources\Analytics($this->http);
        $this->settings = new Resources\Settings($this->http);
        $this->messages = new Resources\Messages($this->http);
    }

    public function setToken(string $token): void { $this->http->setToken($token); }
    public function setPageAccessToken(string $token): void { $this->http->setPageAccessToken($token); }
}
