"""Đăng ký webhook (webhooks)"""

from .base import BaseResource


class WebhooksResource(BaseResource):
    """Đăng ký webhook (webhooks)"""

    def list(self):
        """GET /api/v1/webhooks"""
        return self._http.request("GET", "/api/v1/webhooks")

    def get(self, id):
        """GET /api/v1/webhooks/{id}"""
        return self._http.request("GET", f"/api/v1/webhooks/{id}")

    def create(self, data=None):
        """POST /api/v1/webhooks"""
        return self._http.request("POST", "/api/v1/webhooks", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/webhooks/{id}"""
        return self._http.request("PATCH", f"/api/v1/webhooks/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/webhooks/{id}"""
        return self._http.request("DELETE", f"/api/v1/webhooks/{id}")
