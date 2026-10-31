"""Kịch bản hội thoại (flows) và khối (blocks)"""

from .base import BaseResource


class FlowsResource(BaseResource):
    """Kịch bản hội thoại (flows) và khối (blocks)"""

    def list(self, params=None):
        """GET /api/v1/flows"""
        return self._http.request("GET", "/api/v1/flows", params=params)

    def get(self, id):
        """GET /api/v1/flows/{id}"""
        return self._http.request("GET", f"/api/v1/flows/{id}")

    def create(self, data=None):
        """POST /api/v1/flows"""
        return self._http.request("POST", "/api/v1/flows", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/flows/{id}"""
        return self._http.request("PATCH", f"/api/v1/flows/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/flows/{id}"""
        return self._http.request("DELETE", f"/api/v1/flows/{id}")

    def publish(self, id):
        """POST /api/v1/flows/{id}/publish"""
        return self._http.request("POST", f"/api/v1/flows/{id}/publish")

    def duplicate(self, id):
        """POST /api/v1/flows/{id}/duplicate"""
        return self._http.request("POST", f"/api/v1/flows/{id}/duplicate")

    def trigger(self, id, data=None):
        """POST /api/v1/flows/{id}/trigger"""
        return self._http.request("POST", f"/api/v1/flows/{id}/trigger", json_data=data)

    def list_blocks(self, id):
        """GET /api/v1/flows/{id}/blocks"""
        return self._http.request("GET", f"/api/v1/flows/{id}/blocks")

    def add_block(self, id, data=None):
        """POST /api/v1/flows/{id}/blocks"""
        return self._http.request("POST", f"/api/v1/flows/{id}/blocks", json_data=data)

    def update_block(self, id, block_id, data=None):
        """PATCH /api/v1/flows/{id}/blocks/{block_id}"""
        return self._http.request("PATCH", f"/api/v1/flows/{id}/blocks/{block_id}", json_data=data)

    def delete_block(self, id, block_id):
        """DELETE /api/v1/flows/{id}/blocks/{block_id}"""
        return self._http.request("DELETE", f"/api/v1/flows/{id}/blocks/{block_id}")
