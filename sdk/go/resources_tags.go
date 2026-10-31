package saleai

import "fmt"

// TagsResource - Nhãn (tags)
type TagsResource struct {
	client *HTTPClient
}

// List - Liệt kê nhãn
func (r *TagsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/tags", nil, params, nil)
}

// Get - Lấy chi tiết một nhãn theo ID
func (r *TagsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/tags/%s", id), nil, nil, nil)
}

// Create - Tạo mới một nhãn
func (r *TagsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/tags", data, nil, nil)
}

// Update - Cập nhật một phần nhãn
func (r *TagsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/tags/%s", id), data, nil, nil)
}

// Delete - Xoá một nhãn
func (r *TagsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/tags/%s", id), nil, nil, nil)
}
