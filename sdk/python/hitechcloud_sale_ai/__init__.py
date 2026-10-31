"""
HiTechCloud Sale AI SDK for Python

Nền tảng chatbot bán hàng & tự động hoá hội thoại đa kênh.

    from hitechcloud_sale_ai import SaleAI
    client = SaleAI(token="YOUR_TOKEN")
    bots = client.bots.list()
"""

__version__ = "1.0.0"

from .client import SaleAI
from .webhook import verify_signature, compute_signature, SIGNATURE_HEADER
from .exceptions import (
    SaleAIError, AuthenticationError, AuthorizationError, NotFoundError,
    ConflictError, ValidationError, RateLimitError, ChannelError, ServerError,
)

__all__ = [
    "SaleAI", "verify_signature", "compute_signature", "SIGNATURE_HEADER",
    "SaleAIError", "AuthenticationError", "AuthorizationError", "NotFoundError",
    "ConflictError", "ValidationError", "RateLimitError", "ChannelError", "ServerError",
]
