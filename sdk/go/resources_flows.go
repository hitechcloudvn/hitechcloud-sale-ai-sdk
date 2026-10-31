package saleai

import "fmt"

// FlowsResource - Kịch bản hội thoại (flows) và khối (blocks)
type FlowsResource struct {
	client *HTTPClient
}

// List - Liệt kê kịch bản
func (r *FlowsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/flows", nil, params, nil)
}

// Get - Lấy chi tiết một kịch bản theo ID
func (r *FlowsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/flows/%s", id), nil, nil, nil)
}

// Create - Tạo mới một kịch bản
func (r *FlowsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/flows", data, nil, nil)
}

// Update - Cập nhật một phần kịch bản
func (r *FlowsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/flows/%s", id), data, nil, nil)
}

// Delete - Xoá một kịch bản
func (r *FlowsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/flows/%s", id), nil, nil, nil)
}

// Publish - Xuất bản kịch bản
func (r *FlowsResource) Publish(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/flows/%s/publish", id), nil, nil, nil)
}

// Duplicate - Nhân bản kịch bản
func (r *FlowsResource) Duplicate(id string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/flows/%s/duplicate", id), nil, nil, nil)
}

// Trigger - Kích hoạt kịch bản cho một người theo dõi
func (r *FlowsResource) Trigger(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/flows/%s/trigger", id), data, nil, nil)
}

// ListBlocks - Liệt kê khối của kịch bản
func (r *FlowsResource) ListBlocks(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/flows/%s/blocks", id), nil, nil, nil)
}

// AddBlock - Thêm khối vào kịch bản
func (r *FlowsResource) AddBlock(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/flows/%s/blocks", id), data, nil, nil)
}

// UpdateBlock - Cập nhật khối kịch bản
func (r *FlowsResource) UpdateBlock(id string, blockId string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/flows/%s/blocks/%s", id, blockId), data, nil, nil)
}

// DeleteBlock - Xoá khối kịch bản
func (r *FlowsResource) DeleteBlock(id string, blockId string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/flows/%s/blocks/%s", id, blockId), nil, nil, nil)
}
