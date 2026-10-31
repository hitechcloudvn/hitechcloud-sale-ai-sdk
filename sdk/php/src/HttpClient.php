<?php

namespace HiTechCloud\SaleAI;

use GuzzleHttp\Client;
use GuzzleHttp\Exception\RequestException;
use GuzzleHttp\Exception\ConnectException;

class HttpClient
{
    private Client $client;
    private ?string $token;
    private ?string $pageAccessToken;
    private int $maxRetries;

    public function __construct(string $baseUrl, ?string $token = null, ?string $pageAccessToken = null, int $timeout = 30, int $maxRetries = 3)
    {
        $this->token = $token;
        $this->pageAccessToken = $pageAccessToken;
        $this->maxRetries = $maxRetries;
        $this->client = new Client([
            'base_uri' => rtrim($baseUrl, '/') . '/',
            'timeout' => $timeout,
            'http_errors' => false,
            'headers' => [
                'Content-Type' => 'application/json',
                'Accept' => 'application/json',
                'User-Agent' => 'HiTechCloud-SaleAI-PHP/1.0.0',
            ],
        ]);
    }

    public function setToken(string $token): void { $this->token = $token; }
    public function clearToken(): void { $this->token = null; }
    public function setPageAccessToken(string $token): void { $this->pageAccessToken = $token; }

    private function headers(array $extra = []): array
    {
        $headers = [];
        if ($this->token) {
            $headers['Authorization'] = 'Bearer ' . $this->token;
        }
        if ($this->pageAccessToken) {
            $headers['X-Page-Access-Token'] = $this->pageAccessToken;
        }
        foreach ($extra as $k => $v) {
            if ($v !== null) {
                $headers[$k] = $v;
            }
        }
        return $headers;
    }

    public function request(string $method, string $path, array $data = null, array $params = null, array $headers = []): ?array
    {
        $path = ltrim($path, '/');
        $lastException = null;

        for ($attempt = 0; $attempt <= $this->maxRetries; $attempt++) {
            try {
                $options = ['headers' => $this->headers($headers)];
                if (!empty($data)) {
                    $options['json'] = $data;
                }
                if (!empty($params)) {
                    $options['query'] = $params;
                }

                $response = $this->client->request($method, $path, $options);
                $status = $response->getStatusCode();

                if ($status === 429) {
                    $retryAfter = (int) ($response->getHeader('Retry-After')[0] ?? 2);
                    if ($attempt < $this->maxRetries) {
                        sleep($retryAfter);
                        continue;
                    }
                    throw new RateLimitException('Rate limit exceeded', $retryAfter);
                }

                if ($status >= 500 && $attempt < $this->maxRetries) {
                    sleep(2 ** $attempt);
                    continue;
                }

                $raw = $response->getBody()->getContents();
                $body = $raw === '' ? null : json_decode($raw, true);

                if ($status >= 400) {
                    $this->raise($status, $body, $response);
                }

                return $body;
            } catch (ConnectException $e) {
                $lastException = $e;
                if ($attempt < $this->maxRetries) {
                    sleep(2 ** $attempt);
                    continue;
                }
                throw new SaleAIException('Request failed: ' . $e->getMessage());
            }
        }

        throw new SaleAIException('Request failed: ' . ($lastException ? $lastException->getMessage() : 'unknown'));
    }

    private function raise(int $status, $body, $response): void
    {
        $err = is_array($body) && isset($body['error']) ? $body['error'] : [];
        $message = $err['message'] ?? "HTTP {$status}";
        $code = $err['code'] ?? null;
        $details = $err['details'] ?? null;

        switch ($status) {
            case 401: throw new AuthenticationException($message, $status, $code, $details);
            case 403: throw new AuthorizationException($message, $status, $code, $details);
            case 404: throw new NotFoundException($message, $status, $code, $details);
            case 409: throw new ConflictException($message, $status, $code, $details);
            case 422: throw new ValidationException($message, $status, $code, $details);
            case 429:
                $retryAfter = (int) ($response->getHeader('Retry-After')[0] ?? 2);
                throw new RateLimitException($message, $retryAfter, $code, $details);
            case 502: throw new ChannelException($message, $status, $code, $details);
            default:
                if ($status >= 500) {
                    throw new ServerException($message, $status, $code, $details);
                }
                throw new SaleAIException($message, $status, $code, $details);
        }
    }

    public function get(string $path, array $params = null, array $headers = []): ?array
    {
        return $this->request('GET', $path, null, $params, $headers);
    }

    public function post(string $path, array $data = null, array $headers = []): ?array
    {
        return $this->request('POST', $path, $data, null, $headers);
    }

    public function patch(string $path, array $data = null, array $headers = []): ?array
    {
        return $this->request('PATCH', $path, $data, null, $headers);
    }

    public function delete(string $path, array $headers = []): ?array
    {
        return $this->request('DELETE', $path, null, null, $headers);
    }
}
