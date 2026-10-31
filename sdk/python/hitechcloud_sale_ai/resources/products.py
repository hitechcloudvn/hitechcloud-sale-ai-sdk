"""Sản phẩm (products)"""

from .base import BaseResource


class ProductsResource(BaseResource):
    """Sản phẩm (products)"""

    def list(self, params=None):
        """GET /api/v1/products"""
        return self._http.request("GET", "/api/v1/products", params=params)

    def get(self, id):
        """GET /api/v1/products/{id}"""
        return self._http.request("GET", f"/api/v1/products/{id}")

    def create(self, data=None):
        """POST /api/v1/products"""
        return self._http.request("POST", "/api/v1/products", json_data=data)

    def update(self, id, data=None):
        """PATCH /api/v1/products/{id}"""
        return self._http.request("PATCH", f"/api/v1/products/{id}", json_data=data)

    def delete(self, id):
        """DELETE /api/v1/products/{id}"""
        return self._http.request("DELETE", f"/api/v1/products/{id}")
