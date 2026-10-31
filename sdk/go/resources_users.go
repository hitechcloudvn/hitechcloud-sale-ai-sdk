package saleai

import "fmt"

// UsersResource - Người dùng / nhân sự (team)
type UsersResource struct {
	client *HTTPClient
}

// List - Liệt kê người dùng
func (r *UsersResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/users", nil, params, nil)
}

// Get - Lấy chi tiết một người dùng theo ID
func (r *UsersResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/users/%s", id), nil, nil, nil)
}

// Create - Tạo mới một người dùng
func (r *UsersResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/users", data, nil, nil)
}

// Update - Cập nhật một phần người dùng
func (r *UsersResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/users/%s", id), data, nil, nil)
}

// Delete - Xoá một người dùng
func (r *UsersResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/users/%s", id), nil, nil, nil)
}
