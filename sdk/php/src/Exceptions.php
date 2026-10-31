<?php

namespace HiTechCloud\SaleAI;

class SaleAIException extends \Exception
{
    protected ?string $errorCode;
    protected $details;

    public function __construct(string $message = "", int $code = 0, ?string $errorCode = null, $details = null)
    {
        parent::__construct($message, $code);
        $this->errorCode = $errorCode;
        $this->details = $details;
    }

    public function getErrorCode(): ?string { return $this->errorCode; }
    public function getDetails() { return $this->details; }
}

class AuthenticationException extends SaleAIException {}
class AuthorizationException extends SaleAIException {}
class NotFoundException extends SaleAIException {}
class ConflictException extends SaleAIException {}
class ValidationException extends SaleAIException {}
class ChannelException extends SaleAIException {}
class ServerException extends SaleAIException {}

class RateLimitException extends SaleAIException
{
    private int $retryAfter;

    public function __construct(string $message = "Rate limit exceeded", int $retryAfter = 2, ?string $errorCode = "RATE_LIMITED", $details = null)
    {
        parent::__construct($message, 429, $errorCode, $details);
        $this->retryAfter = $retryAfter;
    }

    public function getRetryAfter(): int { return $this->retryAfter; }
}
