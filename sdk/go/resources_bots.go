package saleai

import "fmt"

// BotsResource - Bot & kênh kết nối (Facebook/Instagram/Zalo/Website)
type BotsResource struct {
	client *HTTPClient
}

// List - Liệt kê bot
func (r *BotsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/bots", nil, params, nil)
}

// Get - Lấy chi tiết một bot theo ID
func (r *BotsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/bots/%s", id), nil, nil, nil)
}

// Create - Tạo mới một bot
func (r *BotsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/bots", data, nil, nil)
}

// Update - Cập nhật một phần bot
func (r *BotsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/bots/%s", id), data, nil, nil)
}

// Delete - Xoá một bot
func (r *BotsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/bots/%s", id), nil, nil, nil)
}

// Connect - Kết nối bot với một kênh
func (r *BotsResource) Connect(botId string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/bots/%s/connect", botId), data, nil, nil)
}

// Disconnect - Ngắt kết nối bot khỏi kênh
func (r *BotsResource) Disconnect(botId string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/bots/%s/disconnect", botId), nil, nil, nil)
}

// UpdateSettings - Cập nhật cấu hình vận hành của bot
func (r *BotsResource) UpdateSettings(botId string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/bots/%s/settings", botId), data, nil, nil)
}

// RefreshAccessToken - Làm mới page access token của bot
func (r *BotsResource) RefreshAccessToken(botId string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/bots/%s/access-token/refresh", botId), data, nil, nil)
}
