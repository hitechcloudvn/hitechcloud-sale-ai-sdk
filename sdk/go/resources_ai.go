package saleai

// AiResource - Trợ lý AI bán hàng (config / train / suggest / handover)
type AiResource struct {
	client *HTTPClient
}

// UpdateConfig - Cập nhật cấu hình trợ lý AI
func (r *AiResource) UpdateConfig(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", "/api/v1/ai/config", data, nil, nil)
}

// Train - Huấn luyện lại trợ lý AI
func (r *AiResource) Train(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/ai/train", data, nil, nil)
}

// SuggestReply - Gợi ý câu trả lời cho nhân viên
func (r *AiResource) SuggestReply(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/ai/suggest-reply", data, nil, nil)
}

// Handover - Bàn giao hội thoại cho người thật
func (r *AiResource) Handover(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("POST", "/api/v1/ai/handover", data, nil, nil)
}
