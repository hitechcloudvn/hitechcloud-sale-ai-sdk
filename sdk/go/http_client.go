package saleai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const userAgent = "HiTechCloud-SaleAI-Go/" + Version

// HTTPClient performs HTTP requests with retry + rate-limit handling.
type HTTPClient struct {
	baseURL         string
	httpClient      *http.Client
	token           string
	pageAccessToken string
	maxRetries      int
}

func newHTTPClient(baseURL, token, pageAccessToken string, timeout time.Duration, maxRetries int) *HTTPClient {
	return &HTTPClient{
		baseURL:         baseURL,
		httpClient:      &http.Client{Timeout: timeout},
		token:           token,
		pageAccessToken: pageAccessToken,
		maxRetries:      maxRetries,
	}
}

// SetToken sets the account access token.
func (c *HTTPClient) SetToken(token string) { c.token = token }

// ClearToken removes the account access token.
func (c *HTTPClient) ClearToken() { c.token = "" }

// SetPageAccessToken sets the default page access token.
func (c *HTTPClient) SetPageAccessToken(token string) { c.pageAccessToken = token }

// Request performs an HTTP request and decodes the JSON response.
func (c *HTTPClient) Request(method, path string, body interface{}, params map[string]string, headers map[string]string) (map[string]interface{}, error) {
	u := c.baseURL + path
	if len(params) > 0 {
		q := url.Values{}
		for k, v := range params {
			if v != "" {
				q.Set(k, v)
			}
		}
		if enc := q.Encode(); enc != "" {
			u += "?" + enc
		}
	}

	var lastErr error
	for attempt := 0; attempt <= c.maxRetries; attempt++ {
		var reader io.Reader
		if body != nil {
			raw, err := json.Marshal(body)
			if err != nil {
				return nil, fmt.Errorf("marshal body: %w", err)
			}
			reader = bytes.NewReader(raw)
		}

		req, err := http.NewRequest(method, u, reader)
		if err != nil {
			return nil, fmt.Errorf("new request: %w", err)
		}
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Accept", "application/json")
		req.Header.Set("User-Agent", userAgent)
		if c.token != "" {
			req.Header.Set("Authorization", "Bearer "+c.token)
		}
		if c.pageAccessToken != "" {
			req.Header.Set("X-Page-Access-Token", c.pageAccessToken)
		}
		for k, v := range headers {
			if v != "" {
				req.Header.Set(k, v)
			}
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			if attempt < c.maxRetries {
				time.Sleep(backoff(attempt))
				continue
			}
			return nil, fmt.Errorf("request failed: %w", err)
		}

		respBody, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 429 {
			retryAfter := parseRetryAfter(resp)
			if attempt < c.maxRetries {
				time.Sleep(time.Duration(retryAfter) * time.Second)
				continue
			}
			return nil, mapError(resp.StatusCode, decode(respBody), retryAfter)
		}
		if resp.StatusCode >= 500 && attempt < c.maxRetries {
			time.Sleep(backoff(attempt))
			continue
		}

		var result map[string]interface{}
		if len(respBody) > 0 {
			if err := json.Unmarshal(respBody, &result); err != nil && resp.StatusCode < 400 {
				return nil, fmt.Errorf("decode response: %w", err)
			}
		}
		if resp.StatusCode >= 400 {
			return nil, mapError(resp.StatusCode, result, parseRetryAfter(resp))
		}
		return result, nil
	}
	return nil, lastErr
}

func backoff(attempt int) time.Duration {
	return time.Duration(math.Pow(2, float64(attempt))) * time.Second
}

func parseRetryAfter(resp *http.Response) int {
	if v := resp.Header.Get("Retry-After"); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return 2
}

func decode(b []byte) map[string]interface{} {
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	return m
}

// Get performs a GET request.
func (c *HTTPClient) Get(path string, params map[string]string, headers map[string]string) (map[string]interface{}, error) {
	return c.Request("GET", path, nil, params, headers)
}

// Post performs a POST request.
func (c *HTTPClient) Post(path string, body interface{}, headers map[string]string) (map[string]interface{}, error) {
	return c.Request("POST", path, body, nil, headers)
}

// Patch performs a PATCH request.
func (c *HTTPClient) Patch(path string, body interface{}, headers map[string]string) (map[string]interface{}, error) {
	return c.Request("PATCH", path, body, nil, headers)
}

// Delete performs a DELETE request.
func (c *HTTPClient) Delete(path string, headers map[string]string) (map[string]interface{}, error) {
	return c.Request("DELETE", path, nil, nil, headers)
}
