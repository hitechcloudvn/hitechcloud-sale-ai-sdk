package saleai

import "fmt"

// BroadcastsResource - Gửi hàng loạt (broadcasts)
type BroadcastsResource struct {
	client *HTTPClient
}

// List - Liệt kê chiến dịch broadcast
func (r *BroadcastsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/broadcasts", nil, params, nil)
}

// Get - Lấy chi tiết một chiến dịch broadcast theo ID
func (r *BroadcastsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/broadcasts/%s", id), nil, nil, nil)
}

// Create - Tạo mới một chiến dịch broadcast
func (r *BroadcastsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/broadcasts", data, nil, nil)
}

// Update - Cập nhật một phần chiến dịch broadcast
func (r *BroadcastsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/broadcasts/%s", id), data, nil, nil)
}

// Delete - Xoá một chiến dịch broadcast
func (r *BroadcastsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/broadcasts/%s", id), nil, nil, nil)
}

// Schedule - Lên lịch gửi broadcast
func (r *BroadcastsResource) Schedule(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/broadcasts/%s/schedule", id), data, nil, nil)
}

// Send - Gửi broadcast ngay
func (r *BroadcastsResource) Send(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/broadcasts/%s/send", id), nil, nil, nil)
}

// Status - Lấy trạng thái gửi broadcast
func (r *BroadcastsResource) Status(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/broadcasts/%s/status", id), nil, nil, nil)
}

// PreviewAudience - Xem trước đối tượng broadcast
func (r *BroadcastsResource) PreviewAudience(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/broadcasts/%s/audience/preview", id), data, nil, nil)
}

// Cancel - Huỷ chiến dịch broadcast
func (r *BroadcastsResource) Cancel(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/broadcasts/%s/cancel", id), nil, nil, nil)
}
