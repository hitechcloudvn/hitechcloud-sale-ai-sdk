"""Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú"""

from .base import BaseResource


class SubscribersResource(BaseResource):
    """Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú"""

    def list(self, params=None):
        """GET /api/v1/subscribers"""
        return self._http.request("GET", "/api/v1/subscribers", params=params)

    def get(self, id):
        """GET /api/v1/subscribers/{id}"""
        return self._http.request("GET", f"/api/v1/subscribers/{id}")

    def create(self, data=None):
        """POST /api/v1/subscribers"""
        return self._http.request("POST", "/api/v1/subscribers", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/subscribers/{id}"""
        return self._http.request("PATCH", f"/api/v1/subscribers/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/subscribers/{id}"""
        return self._http.request("DELETE", f"/api/v1/subscribers/{id}")

    def add_tags(self, id, data=None):
        """POST /api/v1/subscribers/{id}/tags"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/tags", json_data=data)

    def remove_tag(self, id, tag):
        """DELETE /api/v1/subscribers/{id}/tags/{tag}"""
        return self._http.request("DELETE", f"/api/v1/subscribers/{id}/tags/{tag}")

    def set_custom_fields(self, id, data=None):
        """PATCH /api/v1/subscribers/{id}/custom-fields"""
        return self._http.request("PATCH", f"/api/v1/subscribers/{id}/custom-fields", json_data=data)

    def add_note(self, id, data=None):
        """POST /api/v1/subscribers/{id}/notes"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/notes", json_data=data)

    def block(self, id):
        """POST /api/v1/subscribers/{id}/block"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/block")

    def unblock(self, id):
        """POST /api/v1/subscribers/{id}/unblock"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/unblock")

    def list_messages(self, id, params=None):
        """GET /api/v1/subscribers/{id}/messages"""
        return self._http.request("GET", f"/api/v1/subscribers/{id}/messages", params=params)

    def send_message(self, id, data=None):
        """POST /api/v1/subscribers/{id}/messages"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/messages", json_data=data)

    def mark_seen(self, id):
        """POST /api/v1/subscribers/{id}/mark-seen"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/mark-seen")

    def typing(self, id, data=None):
        """POST /api/v1/subscribers/{id}/typing"""
        return self._http.request("POST", f"/api/v1/subscribers/{id}/typing", json_data=data)
