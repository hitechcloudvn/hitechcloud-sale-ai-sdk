package saleai

import "fmt"

// CommentRulesResource - Quy tắc trả lời bình luận (comment rules)
type CommentRulesResource struct {
	client *HTTPClient
}

// List - Liệt kê quy tắc bình luận
func (r *CommentRulesResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/comment-rules", nil, params, nil)
}

// Get - Lấy chi tiết một quy tắc bình luận theo ID
func (r *CommentRulesResource) Get(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/comment-rules/%s", id), nil, nil, nil)
}

// Create - Tạo mới một quy tắc bình luận
func (r *CommentRulesResource) Create(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/comment-rules", data, nil, nil)
}

// Update - Cập nhật một phần quy tắc bình luận
func (r *CommentRulesResource) Update(id string, data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", fmt.Sprintf("/api/v1/comment-rules/%s", id), data, nil, nil)
}

// Delete - Xoá một quy tắc bình luận
func (r *CommentRulesResource) Delete(id string) (map[string]interface{}, error) {
	return r.client.Request("DELETE", fmt.Sprintf("/api/v1/comment-rules/%s", id), nil, nil, nil)
}
