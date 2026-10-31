"""Báo cáo & phân tích (analytics)"""

from .base import BaseResource


class AnalyticsResource(BaseResource):
    """Báo cáo & phân tích (analytics)"""

    def overview(self, params=None):
        """GET /api/v1/analytics/overview"""
        return self._http.request("GET", "/api/v1/analytics/overview", params=params)

    def bot(self, bot_id, params=None):
        """GET /api/v1/analytics/bots/{bot_id}"""
        return self._http.request("GET", f"/api/v1/analytics/bots/{bot_id}", params=params)

    def flow(self, id, params=None):
        """GET /api/v1/analytics/flows/{id}"""
        return self._http.request("GET", f"/api/v1/analytics/flows/{id}", params=params)

    def broadcast(self, id):
        """GET /api/v1/analytics/broadcasts/{id}"""
        return self._http.request("GET", f"/api/v1/analytics/broadcasts/{id}")

    def subscriber_growth(self, params=None):
        """GET /api/v1/analytics/subscriber-growth"""
        return self._http.request("GET", "/api/v1/analytics/subscriber-growth", params=params)
