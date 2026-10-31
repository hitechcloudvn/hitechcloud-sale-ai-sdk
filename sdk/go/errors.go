package saleai

import "fmt"

// APIError is the base error returned by the Sale AI API.
type APIError struct {
	StatusCode int
	Code       string
	Message    string
	Details    interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("sale-ai api error %d (%s): %s", e.StatusCode, e.Code, e.Message)
}

// AuthenticationError represents 401 responses.
type AuthenticationError struct{ APIError }

// AuthorizationError represents 403 responses.
type AuthorizationError struct{ APIError }

// NotFoundError represents 404 responses.
type NotFoundError struct{ APIError }

// ConflictError represents 409 responses.
type ConflictError struct{ APIError }

// ValidationError represents 422 responses.
type ValidationError struct{ APIError }

// RateLimitError represents 429 responses.
type RateLimitError struct {
	APIError
	RetryAfter int
}

// ChannelError represents 502 responses from the messaging channel.
type ChannelError struct{ APIError }

// ServerError represents 5xx responses.
type ServerError struct{ APIError }

func mapError(status int, body map[string]interface{}, retryAfter int) error {
	code, message := "", "API error"
	var details interface{}
	if raw, ok := body["error"].(map[string]interface{}); ok {
		if c, ok := raw["code"].(string); ok {
			code = c
		}
		if m, ok := raw["message"].(string); ok {
			message = m
		}
		details = raw["details"]
	}
	base := APIError{StatusCode: status, Code: code, Message: message, Details: details}
	switch status {
	case 401:
		return &AuthenticationError{base}
	case 403:
		return &AuthorizationError{base}
	case 404:
		return &NotFoundError{base}
	case 409:
		return &ConflictError{base}
	case 422:
		return &ValidationError{base}
	case 429:
		return &RateLimitError{APIError: base, RetryAfter: retryAfter}
	case 502:
		return &ChannelError{base}
	default:
		if status >= 500 {
			return &ServerError{base}
		}
		return &base
	}
}
