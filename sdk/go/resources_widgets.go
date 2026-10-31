package saleai

import "fmt"

// WidgetsResource - Công cụ tăng trưởng (widgets)
type WidgetsResource struct {
	client *HTTPClient
}

// List - Liệt kê công cụ tăng trưởng
func (r *WidgetsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/widgets", nil, params, nil)
}

// Get - Lấy chi tiết một công cụ tăng trưởng theo ID
func (r *WidgetsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/widgets/%s", id), nil, nil, nil)
}

// Create - Tạo mới một công cụ tăng trưởng
func (r *WidgetsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/widgets", data, nil, nil)
}

// Update - Cập nhật một phần công cụ tăng trưởng
func (r *WidgetsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/widgets/%s", id), data, nil, nil)
}

// Delete - Xoá một công cụ tăng trưởng
func (r *WidgetsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/widgets/%s", id), nil, nil, nil)
}
