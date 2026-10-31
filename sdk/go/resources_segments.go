package saleai

import "fmt"

// SegmentsResource - Phân khúc (segments)
type SegmentsResource struct {
	client *HTTPClient
}

// List - Liệt kê phân khúc
func (r *SegmentsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/segments", nil, params, nil)
}

// Get - Lấy chi tiết một phân khúc theo ID
func (r *SegmentsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/segments/%s", id), nil, nil, nil)
}

// Create - Tạo mới một phân khúc
func (r *SegmentsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/segments", data, nil, nil)
}

// Update - Cập nhật một phần phân khúc
func (r *SegmentsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/segments/%s", id), data, nil, nil)
}

// Delete - Xoá một phân khúc
func (r *SegmentsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/segments/%s", id), nil, nil, nil)
}
