package saleai

import "fmt"

// ProductsResource - Sản phẩm (products)
type ProductsResource struct {
	client *HTTPClient
}

// List - Liệt kê sản phẩm
func (r *ProductsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/products", nil, params, nil)
}

// Get - Lấy chi tiết một sản phẩm theo ID
func (r *ProductsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/products/%s", id), nil, nil, nil)
}

// Create - Tạo mới một sản phẩm
func (r *ProductsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/products", data, nil, nil)
}

// Update - Cập nhật một phần sản phẩm
func (r *ProductsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/products/%s", id), data, nil, nil)
}

// Delete - Xoá một sản phẩm
func (r *ProductsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/products/%s", id), nil, nil, nil)
}
