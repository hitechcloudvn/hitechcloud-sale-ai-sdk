package saleai

import "fmt"

// MessagesResource - Gửi tin nhắn (text/image/file/quick-reply/button/carousel).
type MessagesResource struct {
	client *HTTPClient
}

func (r *MessagesResource) sendRaw(botID string, message map[string]interface{}, pat string) (map[string]interface{}, error) {
	var headers map[string]string
	if pat != "" {
		headers = map[string]string{"X-Page-Access-Token": pat}
	}
	return r.client.Request("POST", fmt.Sprintf("/api/v1/bots/%s/messages", botID), message, nil, headers)
}

// Send sends a raw message payload.
func (r *MessagesResource) Send(botID string, message map[string]interface{}, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, message, pageAccessToken)
}

// SendText sends a text message.
func (r *MessagesResource) SendText(botID, subscriberID, text, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "text", "text": text}, pageAccessToken)
}

// SendImage sends an image message.
func (r *MessagesResource) SendImage(botID, subscriberID, url, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "image", "url": url}, pageAccessToken)
}

// SendFile sends a file/audio/video message.
func (r *MessagesResource) SendFile(botID, subscriberID, url, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "file", "url": url}, pageAccessToken)
}

// SendQuickReply sends a quick-reply message.
func (r *MessagesResource) SendQuickReply(botID, subscriberID, text string, quickReplies []string, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "quick_reply", "text": text, "quick_replies": quickReplies}, pageAccessToken)
}

// SendButton sends a button template message.
func (r *MessagesResource) SendButton(botID, subscriberID, text string, buttons []map[string]interface{}, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "button", "text": text, "buttons": buttons}, pageAccessToken)
}

// SendCarousel sends a carousel / generic template message.
func (r *MessagesResource) SendCarousel(botID, subscriberID string, elements []map[string]interface{}, pageAccessToken string) (map[string]interface{}, error) {
	return r.sendRaw(botID, map[string]interface{}{"subscriber_id": subscriberID, "type": "carousel", "elements": elements}, pageAccessToken)
}
