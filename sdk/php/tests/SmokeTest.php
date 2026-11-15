<?php

// EN: Smoke tests - keep the suite green so CI can gate releases.
// VI: Smoke test - giu bo test xanh de CI kiem soat viec phat hanh.

namespace HiTechCloud\SaleAI\Tests;

use PHPUnit\Framework\TestCase;

final class SmokeTest extends TestCase
{
    public function testMainClientIsLoadable(): void
    {
        $this->assertTrue(class_exists(\HiTechCloud\SaleAI\SaleAI::class));
    }

    public function testExceptionHierarchyIsLoadable(): void
    {
        $this->assertTrue(class_exists(\HiTechCloud\SaleAI\SaleAIException::class));
        $this->assertTrue(class_exists(\HiTechCloud\SaleAI\AuthenticationException::class));
        $this->assertTrue(is_subclass_of(\HiTechCloud\SaleAI\AuthenticationException::class, \HiTechCloud\SaleAI\SaleAIException::class));
    }
}
