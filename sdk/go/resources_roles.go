package saleai

import "fmt"

// RolesResource - Vai trò & quyền (roles)
type RolesResource struct {
	client *HTTPClient
}

// List - Liệt kê vai trò
func (r *RolesResource) List() (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/roles", nil, nil, nil)
}

// Get - Lấy chi tiết một vai trò theo ID
func (r *RolesResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/roles/%s", id), nil, nil, nil)
}

// Create - Tạo mới một vai trò
func (r *RolesResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/roles", data, nil, nil)
}

// Update - Cập nhật một phần vai trò
func (r *RolesResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/roles/%s", id), data, nil, nil)
}

// Delete - Xoá một vai trò
func (r *RolesResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/roles/%s", id), nil, nil, nil)
}
