"""Cài đặt tài khoản (settings)"""

from .base import BaseResource


class SettingsResource(BaseResource):
    """Cài đặt tài khoản (settings)"""

    def get(self):
        """GET /api/v1/settings"""
        return self._http.request("GET", "/api/v1/settings")

    def update(self, data=None):
        """PATCH /api/v1/settings"""
        return self._http.request("PATCH", "/api/v1/settings", json_data=data)
