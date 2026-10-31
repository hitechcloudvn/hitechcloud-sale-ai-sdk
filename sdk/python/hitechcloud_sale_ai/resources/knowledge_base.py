"""Kho tri thức AI (knowledge base)"""

from .base import BaseResource


class KnowledgeBaseResource(BaseResource):
    """Kho tri thức AI (knowledge base)"""

    def list(self, params=None):
        """GET /api/v1/ai/knowledge-base"""
        return self._http.request("GET", "/api/v1/ai/knowledge-base", params=params)

    def get(self, id):
        """GET /api/v1/ai/knowledge-base/{id}"""
        return self._http.request("GET", f"/api/v1/ai/knowledge-base/{id}")

    def create(self, data=None):
        """POST /api/v1/ai/knowledge-base"""
        return self._http.request("POST", "/api/v1/ai/knowledge-base", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/ai/knowledge-base/{id}"""
        return self._http.request("PATCH", f"/api/v1/ai/knowledge-base/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/ai/knowledge-base/{id}"""
        return self._http.request("DELETE", f"/api/v1/ai/knowledge-base/{id}")
