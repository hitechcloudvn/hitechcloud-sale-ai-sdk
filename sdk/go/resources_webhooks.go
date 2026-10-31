package saleai

import "fmt"

// WebhooksResource - Đăng ký webhook (webhooks)
type WebhooksResource struct {
	client *HTTPClient
}

// List - Liệt kê webhook
func (r *WebhooksResource) List() (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/webhooks", nil, nil, nil)
}

// Get - Lấy chi tiết một webhook theo ID
func (r *WebhooksResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/webhooks/%s", id), nil, nil, nil)
}

// Create - Tạo mới một webhook
func (r *WebhooksResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/webhooks", data, nil, nil)
}

// Update - Cập nhật một phần webhook
func (r *WebhooksResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/webhooks/%s", id), data, nil, nil)
}

// Delete - Xoá một webhook
func (r *WebhooksResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/webhooks/%s", id), nil, nil, nil)
}
