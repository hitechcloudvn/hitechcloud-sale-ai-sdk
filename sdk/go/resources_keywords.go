package saleai

import "fmt"

// KeywordsResource - Từ khoá tự trả lời (keywords)
type KeywordsResource struct {
	client *HTTPClient
}

// List - Liệt kê từ khoá
func (r *KeywordsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/keywords", nil, params, nil)
}

// Get - Lấy chi tiết một từ khoá theo ID
func (r *KeywordsResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/keywords/%s", id), nil, nil, nil)
}

// Create - Tạo mới một từ khoá
func (r *KeywordsResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/keywords", data, nil, nil)
}

// Update - Cập nhật một phần từ khoá
func (r *KeywordsResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/keywords/%s", id), data, nil, nil)
}

// Delete - Xoá một từ khoá
func (r *KeywordsResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/keywords/%s", id), nil, nil, nil)
}

// Test - Kiểm thử khớp từ khoá cho một câu nhập
func (r *KeywordsResource) Test(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/keywords/test", data, nil, nil)
}
