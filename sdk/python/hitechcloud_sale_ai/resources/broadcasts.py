"""Gửi hàng loạt (broadcasts)"""

from .base import BaseResource


class BroadcastsResource(BaseResource):
    """Gửi hàng loạt (broadcasts)"""

    def list(self, params=None):
        """GET /api/v1/broadcasts"""
        return self._http.request("GET", "/api/v1/broadcasts", params=params)

    def get(self, id):
        """GET /api/v1/broadcasts/{id}"""
        return self._http.request("GET", f"/api/v1/broadcasts/{id}")

    def create(self, data=None):
        """POST /api/v1/broadcasts"""
        return self._http.request("POST", "/api/v1/broadcasts", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/broadcasts/{id}"""
        return self._http.request("PATCH", f"/api/v1/broadcasts/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/broadcasts/{id}"""
        return self._http.request("DELETE", f"/api/v1/broadcasts/{id}")

    def schedule(self, id, data=None):
        """POST /api/v1/broadcasts/{id}/schedule"""
        return self._http.request("POST", f"/api/v1/broadcasts/{id}/schedule", json_data=data)

    def send(self, id):
        """POST /api/v1/broadcasts/{id}/send"""
        return self._http.request("POST", f"/api/v1/broadcasts/{id}/send")

    def status(self, id):
        """GET /api/v1/broadcasts/{id}/status"""
        return self._http.request("GET", f"/api/v1/broadcasts/{id}/status")

    def preview_audience(self, id, data=None):
        """POST /api/v1/broadcasts/{id}/audience/preview"""
        return self._http.request("POST", f"/api/v1/broadcasts/{id}/audience/preview", json_data=data)

    def cancel(self, id):
        """POST /api/v1/broadcasts/{id}/cancel"""
        return self._http.request("POST", f"/api/v1/broadcasts/{id}/cancel")
