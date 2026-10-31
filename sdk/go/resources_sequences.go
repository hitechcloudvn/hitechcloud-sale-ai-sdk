package saleai

import "fmt"

// SequencesResource - Chuỗi tin hẹn giờ (sequences) và các bước
type SequencesResource struct {
	client *HTTPClient
}

// List - Liệt kê chuỗi tin
func (r *SequencesResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/sequences", nil, params, nil)
}

// Get - Lấy chi tiết một chuỗi tin theo ID
func (r *SequencesResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/sequences/%s", id), nil, nil, nil)
}

// Create - Tạo mới một chuỗi tin
func (r *SequencesResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/sequences", data, nil, nil)
}

// Update - Cập nhật một phần chuỗi tin
func (r *SequencesResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/sequences/%s", id), data, nil, nil)
}

// Delete - Xoá một chuỗi tin
func (r *SequencesResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/sequences/%s", id), nil, nil, nil)
}

// ListSteps - Liệt kê bước của chuỗi tin
func (r *SequencesResource) ListSteps(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/sequences/%s/steps", id), nil, nil, nil)
}

// AddStep - Thêm bước vào chuỗi tin
func (r *SequencesResource) AddStep(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/sequences/%s/steps", id), data, nil, nil)
}

// UpdateStep - Cập nhật bước chuỗi tin
func (r *SequencesResource) UpdateStep(id string, stepId string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/sequences/%s/steps/%s", id, stepId), data, nil, nil)
}

// DeleteStep - Xoá bước chuỗi tin
func (r *SequencesResource) DeleteStep(id string, stepId string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/sequences/%s/steps/%s", id, stepId), nil, nil, nil)
}

// Subscribe - Đăng ký người theo dõi vào chuỗi
func (r *SequencesResource) Subscribe(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/sequences/%s/subscribe", id), data, nil, nil)
}

// Unsubscribe - Huỷ đăng ký khỏi chuỗi
func (r *SequencesResource) Unsubscribe(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/sequences/%s/unsubscribe", id), data, nil, nil)
}
