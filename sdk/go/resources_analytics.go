package saleai

import "fmt"

// AnalyticsResource - Báo cáo & phân tích (analytics)
type AnalyticsResource struct {
	client *HTTPClient
}

// Overview - Báo cáo tổng quan
func (r *AnalyticsResource) Overview(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/analytics/overview", nil, params, nil)
}

// Bot - Báo cáo theo bot
func (r *AnalyticsResource) Bot(botId string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/analytics/bots/%s", botId), nil, params, nil)
}

// Flow - Báo cáo theo kịch bản
func (r *AnalyticsResource) Flow(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/analytics/flows/%s", id), nil, params, nil)
}

// Broadcast - Báo cáo theo broadcast
func (r *AnalyticsResource) Broadcast(id string) (map[string]interface{}, error) {
	return r.client.Request("GET", fmt.Sprintf("/api/v1/analytics/broadcasts/%s", id), nil, nil, nil)
}

// SubscriberGrowth - Báo cáo tăng trưởng người theo dõi
func (r *AnalyticsResource) SubscriberGrowth(params map[string]string) (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/analytics/subscriber-growth", nil, params, nil)
}
