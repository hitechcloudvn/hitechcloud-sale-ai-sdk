"""Hộp thư đa kênh (conversations / inbox)"""

from .base import BaseResource


class ConversationsResource(BaseResource):
    """Hộp thư đa kênh (conversations / inbox)"""

    def list(self, params=None):
        """GET /api/v1/conversations"""
        return self._http.request("GET", "/api/v1/conversations", params=params)

    def get(self, id):
        """GET /api/v1/conversations/{id}"""
        return self._http.request("GET", f"/api/v1/conversations/{id}")

    def assign(self, id, data=None):
        """POST /api/v1/conversations/{id}/assign"""
        return self._http.request("POST", f"/api/v1/conversations/{id}/assign", json_data=data)

    def add_note(self, id, data=None):
        """POST /api/v1/conversations/{id}/notes"""
        return self._http.request("POST", f"/api/v1/conversations/{id}/notes", json_data=data)

    def add_labels(self, id, data=None):
        """POST /api/v1/conversations/{id}/labels"""
        return self._http.request("POST", f"/api/v1/conversations/{id}/labels", json_data=data)

    def close(self, id):
        """POST /api/v1/conversations/{id}/close"""
        return self._http.request("POST", f"/api/v1/conversations/{id}/close")
