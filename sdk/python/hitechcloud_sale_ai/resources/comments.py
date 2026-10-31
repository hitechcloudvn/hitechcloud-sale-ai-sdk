"""Thao tác bình luận trên kênh (comments)"""

from .base import BaseResource


class CommentsResource(BaseResource):
    """Thao tác bình luận trên kênh (comments)"""

    def hide(self, comment_id, page_access_token=None):
        """POST /api/v1/comments/{comment_id}/hide"""
        return self._http.request("POST", f"/api/v1/comments/{comment_id}/hide", headers={"X-Page-Access-Token": page_access_token})

    def like(self, comment_id, page_access_token=None):
        """POST /api/v1/comments/{comment_id}/like"""
        return self._http.request("POST", f"/api/v1/comments/{comment_id}/like", headers={"X-Page-Access-Token": page_access_token})
