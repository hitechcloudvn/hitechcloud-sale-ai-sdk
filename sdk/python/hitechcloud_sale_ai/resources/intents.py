"""Ý định AI (intents)"""

from .base import BaseResource


class IntentsResource(BaseResource):
    """Ý định AI (intents)"""

    def list(self, params=None):
        """GET /api/v1/ai/intents"""
        return self._http.request("GET", "/api/v1/ai/intents", params=params)

    def get(self, id):
        """GET /api/v1/ai/intents/{id}"""
        return self._http.request("GET", f"/api/v1/ai/intents/{id}")

    def create(self, data=None):
        """POST /api/v1/ai/intents"""
        return self._http.request("POST", "/api/v1/ai/intents", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/ai/intents/{id}"""
        return self._http.request("PATCH", f"/api/v1/ai/intents/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/ai/intents/{id}"""
        return self._http.request("DELETE", f"/api/v1/ai/intents/{id}")
