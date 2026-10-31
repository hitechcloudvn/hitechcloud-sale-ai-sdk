package saleai

import "fmt"

// ConversationsResource - Hộp thư đa kênh (conversations / inbox)
type ConversationsResource struct {
	client *HTTPClient
}

// List - Liệt kê hội thoại
func (r *ConversationsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/conversations", nil, params, nil)
}

// Get - Chi tiết hội thoại
func (r *ConversationsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/conversations/%s", id), nil, nil, nil)
}

// Assign - Gán nhân viên cho hội thoại
func (r *ConversationsResource) Assign(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/conversations/%s/assign", id), data, nil, nil)
}

// AddNote - Thêm ghi chú hội thoại
func (r *ConversationsResource) AddNote(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/conversations/%s/notes", id), data, nil, nil)
}

// AddLabels - Gắn nhãn hội thoại
func (r *ConversationsResource) AddLabels(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/conversations/%s/labels", id), data, nil, nil)
}

// Close - Đóng hội thoại
func (r *ConversationsResource) Close(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/conversations/%s/close", id), nil, nil, nil)
}
