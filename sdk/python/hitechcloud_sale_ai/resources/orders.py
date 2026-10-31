"""Đơn hàng (orders)"""

from .base import BaseResource


class OrdersResource(BaseResource):
    """Đơn hàng (orders)"""

    def list(self, params=None):
        """GET /api/v1/orders"""
        return self._http.request("GET", "/api/v1/orders", params=params)

    def get(self, id):
        """GET /api/v1/orders/{id}"""
        return self._http.request("GET", f"/api/v1/orders/{id}")

    def create(self, data=None):
        """POST /api/v1/orders"""
        return self._http.request("POST", "/api/v1/orders", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/orders/{id}"""
        return self._http.request("PATCH", f"/api/v1/orders/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/orders/{id}"""
        return self._http.request("DELETE", f"/api/v1/orders/{id}")

    def create_from_conversation(self, data=None):
        """POST /api/v1/orders/from-conversation"""
        return self._http.request("POST", "/api/v1/orders/from-conversation", json_data=data)

    def update_status(self, id, data=None):
        """POST /api/v1/orders/{id}/status"""
        return self._http.request("POST", f"/api/v1/orders/{id}/status", json_data=data)
