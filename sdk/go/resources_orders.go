package saleai

import "fmt"

// OrdersResource - Đơn hàng (orders)
type OrdersResource struct {
	client *HTTPClient
}

// List - Liệt kê đơn hàng
func (r *OrdersResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/orders", nil, params, nil)
}

// Get - Lấy chi tiết một đơn hàng theo ID
func (r *OrdersResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/orders/%s", id), nil, nil, nil)
}

// Create - Tạo mới một đơn hàng
func (r *OrdersResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/orders", data, nil, nil)
}

// Update - Cập nhật một phần đơn hàng
func (r *OrdersResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/orders/%s", id), data, nil, nil)
}

// Delete - Xoá một đơn hàng
func (r *OrdersResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/orders/%s", id), nil, nil, nil)
}

// CreateFromConversation - Tạo đơn từ hội thoại
func (r *OrdersResource) CreateFromConversation(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/orders/from-conversation", data, nil, nil)
}

// UpdateStatus - Cập nhật trạng thái đơn hàng
func (r *OrdersResource) UpdateStatus(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/orders/%s/status", id), data, nil, nil)
}
