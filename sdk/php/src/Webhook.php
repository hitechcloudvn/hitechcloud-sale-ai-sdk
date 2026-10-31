<?php

namespace HiTechCloud\SaleAI;

class Webhook
{
    public const SIGNATURE_HEADER = 'X-SaleAI-Signature';

    public static function computeSignature(string $rawBody, string $secret): string
    {
        return 'sha256=' . hash_hmac('sha256', $rawBody, $secret);
    }

    public static function verifySignature(string $rawBody, ?string $signature, string $secret): bool
    {
        if ($signature === null || $signature === '') {
            return false;
        }
        $expected = self::computeSignature($rawBody, $secret);
        return hash_equals($expected, $signature);
    }
}
