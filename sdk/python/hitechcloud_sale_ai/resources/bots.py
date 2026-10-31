"""Bot & kênh kết nối (Facebook/Instagram/Zalo/Website)"""

from .base import BaseResource


class BotsResource(BaseResource):
    """Bot & kênh kết nối (Facebook/Instagram/Zalo/Website)"""

    def list(self, params=None):
        """GET /api/v1/bots"""
        return self._http.request("GET", "/api/v1/bots", params=params)

    def get(self, id):
        """GET /api/v1/bots/{id}"""
        return self._http.request("GET", f"/api/v1/bots/{id}")

    def create(self, data=None):
        """POST /api/v1/bots"""
        return self._http.request("POST", "/api/v1/bots", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/bots/{id}"""
        return self._http.request("PATCH", f"/api/v1/bots/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/bots/{id}"""
        return self._http.request("DELETE", f"/api/v1/bots/{id}")

    def connect(self, bot_id, data=None):
        """POST /api/v1/bots/{bot_id}/connect"""
        return self._http.request("POST", f"/api/v1/bots/{bot_id}/connect", json_data=data)

    def disconnect(self, bot_id):
        """POST /api/v1/bots/{bot_id}/disconnect"""
        return self._http.request("POST", f"/api/v1/bots/{bot_id}/disconnect")

    def update_settings(self, bot_id, data=None):
        """PATCH /api/v1/bots/{bot_id}/settings"""
        return self._http.request("PATCH", f"/api/v1/bots/{bot_id}/settings", json_data=data)

    def refresh_access_token(self, bot_id, data=None):
        """POST /api/v1/bots/{bot_id}/access-token/refresh"""
        return self._http.request("POST", f"/api/v1/bots/{bot_id}/access-token/refresh", json_data=data)
