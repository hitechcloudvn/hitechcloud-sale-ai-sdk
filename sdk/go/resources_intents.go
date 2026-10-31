package saleai

import "fmt"

// IntentsResource - Ý định AI (intents)
type IntentsResource struct {
	client *HTTPClient
}

// List - Liệt kê ý định AI
func (r *IntentsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/ai/intents", nil, params, nil)
}

// Get - Lấy chi tiết một ý định AI theo ID
func (r *IntentsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/ai/intents/%s", id), nil, nil, nil)
}

// Create - Tạo mới một ý định AI
func (r *IntentsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/ai/intents", data, nil, nil)
}

// Update - Cập nhật một phần ý định AI
func (r *IntentsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/ai/intents/%s", id), data, nil, nil)
}

// Delete - Xoá một ý định AI
func (r *IntentsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/ai/intents/%s", id), nil, nil, nil)
}
