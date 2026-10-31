"""Người dùng / nhân sự (team)"""

from .base import BaseResource


class UsersResource(BaseResource):
    """Người dùng / nhân sự (team)"""

    def list(self, params=None):
        """GET /api/v1/users"""
        return self._http.request("GET", "/api/v1/users", params=params)

    def get(self, id):
        """GET /api/v1/users/{id}"""
        return self._http.request("GET", f"/api/v1/users/{id}")

    def create(self, data=None):
        """POST /api/v1/users"""
        return self._http.request("POST", "/api/v1/users", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/users/{id}"""
        return self._http.request("PATCH", f"/api/v1/users/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/users/{id}"""
        return self._http.request("DELETE", f"/api/v1/users/{id}")
