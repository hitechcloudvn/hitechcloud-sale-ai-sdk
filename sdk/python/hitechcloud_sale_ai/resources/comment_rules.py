"""Quy tắc trả lời bình luận (comment rules)"""

from .base import BaseResource


class CommentRulesResource(BaseResource):
    """Quy tắc trả lời bình luận (comment rules)"""

    def list(self, params=None):
        """GET /api/v1/comment-rules"""
        return self._http.request("GET", "/api/v1/comment-rules", params=params)

    def get(self, id):
        """GET /api/v1/comment-rules/{id}"""
        return self._http.request("GET", f"/api/v1/comment-rules/{id}")

    def create(self, data=None):
        """POST /api/v1/comment-rules"""
        return self._http.request("POST", "/api/v1/comment-rules", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/comment-rules/{id}"""
        return self._http.request("PATCH", f"/api/v1/comment-rules/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/comment-rules/{id}"""
        return self._http.request("DELETE", f"/api/v1/comment-rules/{id}")
