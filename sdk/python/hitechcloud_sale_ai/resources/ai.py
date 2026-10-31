"""Trợ lý AI bán hàng (config / train / suggest / handover)"""

from .base import BaseResource


class AiResource(BaseResource):
    """Trợ lý AI bán hàng (config / train / suggest / handover)"""

    def update_config(self, data=None):
        """PATCH /api/v1/ai/config"""
        return self._http.request("PATCH", "/api/v1/ai/config", json_data=data)

    def train(self, data=None):
        """POST /api/v1/ai/train"""
        return self._http.request("POST", "/api/v1/ai/train", json_data=data)

    def suggest_reply(self, data=None):
        """POST /api/v1/ai/suggest-reply"""
        return self._http.request("POST", "/api/v1/ai/suggest-reply", json_data=data)

    def handover(self, data=None):
        """POST /api/v1/ai/handover"""
        return self._http.request("POST", "/api/v1/ai/handover", json_data=data)
