package saleai

import "fmt"

// SubscribersResource - Người theo dõi (subscribers), nhãn, trường tuỳ chỉnh, ghi chú
type SubscribersResource struct {
	client *HTTPClient
}

// List - Liệt kê người theo dõi
func (r *SubscribersResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/subscribers", nil, params, nil)
}

// Get - Lấy chi tiết một người theo dõi theo ID
func (r *SubscribersResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/subscribers/%s", id), nil, nil, nil)
}

// Create - Tạo mới một người theo dõi
func (r *SubscribersResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/subscribers", data, nil, nil)
}

// Update - Cập nhật một phần người theo dõi
func (r *SubscribersResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/subscribers/%s", id), data, nil, nil)
}

// Delete - Xoá một người theo dõi
func (r *SubscribersResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/subscribers/%s", id), nil, nil, nil)
}

// AddTags - Gắn nhãn cho người theo dõi
func (r *SubscribersResource) AddTags(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/tags", id), data, nil, nil)
}

// RemoveTag - Gỡ nhãn khỏi người theo dõi
func (r *SubscribersResource) RemoveTag(id string, tag string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/subscribers/%s/tags/%s", id, tag), nil, nil, nil)
}

// SetCustomFields - Cập nhật trường tuỳ chỉnh
func (r *SubscribersResource) SetCustomFields(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/subscribers/%s/custom-fields", id), data, nil, nil)
}

// AddNote - Thêm ghi chú cho người theo dõi
func (r *SubscribersResource) AddNote(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/notes", id), data, nil, nil)
}

// Block - Chặn người theo dõi
func (r *SubscribersResource) Block(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/block", id), nil, nil, nil)
}

// Unblock - Bỏ chặn người theo dõi
func (r *SubscribersResource) Unblock(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/unblock", id), nil, nil, nil)
}

// ListMessages - Lấy lịch sử tin nhắn
func (r *SubscribersResource) ListMessages(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/subscribers/%s/messages", id), nil, params, nil)
}

// SendMessage - Gửi tin theo người theo dõi
func (r *SubscribersResource) SendMessage(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/messages", id), data, nil, nil)
}

// MarkSeen - Đánh dấu đã xem
func (r *SubscribersResource) MarkSeen(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/mark-seen", id), nil, nil, nil)
}

// Typing - Chỉ báo đang soạn tin
func (r *SubscribersResource) Typing(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/subscribers/%s/typing", id), data, nil, nil)
}
