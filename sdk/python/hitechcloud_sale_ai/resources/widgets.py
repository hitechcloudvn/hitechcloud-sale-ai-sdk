"""Công cụ tăng trưởng (widgets)"""

from .base import BaseResource


class WidgetsResource(BaseResource):
    """Công cụ tăng trưởng (widgets)"""

    def list(self, params=None):
        """GET /api/v1/widgets"""
        return self._http.request("GET", "/api/v1/widgets", params=params)

    def get(self, id):
        """GET /api/v1/widgets/{id}"""
        return self._http.request("GET", f"/api/v1/widgets/{id}")

    def create(self, data=None):
        """POST /api/v1/widgets"""
        return self._http.request("POST", "/api/v1/widgets", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/widgets/{id}"""
        return self._http.request("PATCH", f"/api/v1/widgets/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/widgets/{id}"""
        return self._http.request("DELETE", f"/api/v1/widgets/{id}")
