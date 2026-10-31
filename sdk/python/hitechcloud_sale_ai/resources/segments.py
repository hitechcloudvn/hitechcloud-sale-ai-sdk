"""Phân khúc (segments)"""

from .base import BaseResource


class SegmentsResource(BaseResource):
    """Phân khúc (segments)"""

    def list(self, params=None):
        """GET /api/v1/segments"""
        return self._http.request("GET", "/api/v1/segments", params=params)

    def get(self, id):
        """GET /api/v1/segments/{id}"""
        return self._http.request("GET", f"/api/v1/segments/{id}")

    def create(self, data=None):
        """POST /api/v1/segments"""
        return self._http.request("POST", "/api/v1/segments", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/segments/{id}"""
        return self._http.request("PATCH", f"/api/v1/segments/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/segments/{id}"""
        return self._http.request("DELETE", f"/api/v1/segments/{id}")
