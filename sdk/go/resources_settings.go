package saleai

// SettingsResource - Cài đặt tài khoản (settings)
type SettingsResource struct {
	client *HTTPClient
}

// Get - Lấy cài đặt tài khoản
func (r *SettingsResource) Get() (map[string]interface{}, error) {
	return r.client.Request("GET", "/api/v1/settings", nil, nil, nil)
}

// Update - Cập nhật cài đặt tài khoản
func (r *SettingsResource) Update(data map[string]interface{}) (map[string]interface{}, error) {
	return r.client.Request("PATCH", "/api/v1/settings", data, nil, nil)
}
