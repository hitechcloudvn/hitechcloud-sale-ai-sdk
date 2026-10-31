"""Vai trò & quyền (roles)"""

from .base import BaseResource


class RolesResource(BaseResource):
    """Vai trò & quyền (roles)"""

    def list(self):
        """GET /api/v1/roles"""
        return self._http.request("GET", "/api/v1/roles")

    def get(self, id):
        """GET /api/v1/roles/{id}"""
        return self._http.request("GET", f"/api/v1/roles/{id}")

    def create(self, data=None):
        """POST /api/v1/roles"""
        return self._http.request("POST", "/api/v1/roles", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/roles/{id}"""
        return self._http.request("PATCH", f"/api/v1/roles/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/roles/{id}"""
        return self._http.request("DELETE", f"/api/v1/roles/{id}")
