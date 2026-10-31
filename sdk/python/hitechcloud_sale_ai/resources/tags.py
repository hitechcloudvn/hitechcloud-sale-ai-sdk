"""Nhãn (tags)"""

from .base import BaseResource


class TagsResource(BaseResource):
    """Nhãn (tags)"""

    def list(self, params=None):
        """GET /api/v1/tags"""
        return self._http.request("GET", "/api/v1/tags", params=params)

    def get(self, id):
        """GET /api/v1/tags/{id}"""
        return self._http.request("GET", f"/api/v1/tags/{id}")

    def create(self, data=None):
        """POST /api/v1/tags"""
        return self._http.request("POST", "/api/v1/tags", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/tags/{id}"""
        return self._http.request("PATCH", f"/api/v1/tags/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/tags/{id}"""
        return self._http.request("DELETE", f"/api/v1/tags/{id}")
