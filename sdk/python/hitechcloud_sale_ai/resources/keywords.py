"""Từ khoá tự trả lời (keywords)"""

from .base import BaseResource


class KeywordsResource(BaseResource):
    """Từ khoá tự trả lời (keywords)"""

    def list(self, params=None):
        """GET /api/v1/keywords"""
        return self._http.request("GET", "/api/v1/keywords", params=params)

    def get(self, id):
        """GET /api/v1/keywords/{id}"""
        return self._http.request("GET", f"/api/v1/keywords/{id}")

    def create(self, data=None):
        """POST /api/v1/keywords"""
        return self._http.request("POST", "/api/v1/keywords", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/keywords/{id}"""
        return self._http.request("PATCH", f"/api/v1/keywords/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/keywords/{id}"""
        return self._http.request("DELETE", f"/api/v1/keywords/{id}")

    def test(self, data=None):
        """POST /api/v1/keywords/test"""
        return self._http.request("POST", "/api/v1/keywords/test", json_data=data)
