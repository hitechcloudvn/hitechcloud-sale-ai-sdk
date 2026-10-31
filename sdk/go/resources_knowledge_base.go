package saleai

import "fmt"

// KnowledgeBaseResource - Kho tri thức AI (knowledge base)
type KnowledgeBaseResource struct {
	client *HTTPClient
}

// List - Liệt kê mục kho tri thức
func (r *KnowledgeBaseResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/ai/knowledge-base", nil, params, nil)
}

// Get - Lấy chi tiết một mục kho tri thức theo ID
func (r *KnowledgeBaseResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/ai/knowledge-base/%s", id), nil, nil, nil)
}

// Create - Tạo mới một mục kho tri thức
func (r *KnowledgeBaseResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/ai/knowledge-base", data, nil, nil)
}

// Update - Cập nhật một phần mục kho tri thức
func (r *KnowledgeBaseResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/ai/knowledge-base/%s", id), data, nil, nil)
}

// Delete - Xoá một mục kho tri thức
func (r *KnowledgeBaseResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/ai/knowledge-base/%s", id), nil, nil, nil)
}
