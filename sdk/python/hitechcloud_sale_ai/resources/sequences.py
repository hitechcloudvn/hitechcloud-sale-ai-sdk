"""Chuỗi tin hẹn giờ (sequences) và các bước"""

from .base import BaseResource


class SequencesResource(BaseResource):
    """Chuỗi tin hẹn giờ (sequences) và các bước"""

    def list(self, params=None):
        """GET /api/v1/sequences"""
        return self._http.request("GET", "/api/v1/sequences", params=params)

    def get(self, id):
        """GET /api/v1/sequences/{id}"""
        return self._http.request("GET", f"/api/v1/sequences/{id}")

    def create(self, data=None):
        """POST /api/v1/sequences"""
        return self._http.request("POST", "/api/v1/sequences", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/sequences/{id}"""
        return self._http.request("PATCH", f"/api/v1/sequences/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/sequences/{id}"""
        return self._http.request("DELETE", f"/api/v1/sequences/{id}")

    def list_steps(self, id):
        """GET /api/v1/sequences/{id}/steps"""
        return self._http.request("GET", f"/api/v1/sequences/{id}/steps")

    def add_step(self, id, data=None):
        """POST /api/v1/sequences/{id}/steps"""
        return self._http.request("POST", f"/api/v1/sequences/{id}/steps", json_data=data)

    def update_step(self, id, step_id, data=None):
        """PATCH /api/v1/sequences/{id}/steps/{step_id}"""
        return self._http.request("PATCH", f"/api/v1/sequences/{id}/steps/{step_id}", json_data=data)

    def delete_step(self, id, step_id):
        """DELETE /api/v1/sequences/{id}/steps/{step_id}"""
        return self._http.request("DELETE", f"/api/v1/sequences/{id}/steps/{step_id}")

    def subscribe(self, id, data=None):
        """POST /api/v1/sequences/{id}/subscribe"""
        return self._http.request("POST", f"/api/v1/sequences/{id}/subscribe", json_data=data)

    def unsubscribe(self, id, data=None):
        """POST /api/v1/sequences/{id}/unsubscribe"""
        return self._http.request("POST", f"/api/v1/sequences/{id}/unsubscribe", json_data=data)
