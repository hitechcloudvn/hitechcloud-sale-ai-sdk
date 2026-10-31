"""Webhook signature verification for HiTechCloud Sale AI."""

import hashlib
import hmac

SIGNATURE_HEADER = "X-SaleAI-Signature"


def compute_signature(raw_body, secret):
    """Return the expected ``sha256=<hex>`` signature for a raw request body."""
    if isinstance(raw_body, str):
        raw_body = raw_body.encode("utf-8")
    if isinstance(secret, str):
        secret = secret.encode("utf-8")
    digest = hmac.new(secret, raw_body, hashlib.sha256).hexdigest()
    return "sha256=" + digest


def verify_signature(raw_body, signature, secret):
    """Constant-time verify the ``X-SaleAI-Signature`` header (HMAC-SHA256)."""
    if not signature:
        return False
    expected = compute_signature(raw_body, secret)
    return hmac.compare_digest(expected, signature)
