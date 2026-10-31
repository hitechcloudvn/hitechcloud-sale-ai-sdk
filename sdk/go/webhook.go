package saleai

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// SignatureHeader is the webhook signature header name.
const SignatureHeader = "X-SaleAI-Signature"

// ComputeSignature returns the expected "sha256=<hex>" signature for a body.
func ComputeSignature(rawBody []byte, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write(rawBody)
	return "sha256=" + hex.EncodeToString(mac.Sum(nil))
}

// VerifyWebhookSignature constant-time verifies the X-SaleAI-Signature header.
func VerifyWebhookSignature(rawBody []byte, signature, secret string) bool {
	if signature == "" {
		return false
	}
	expected := ComputeSignature(rawBody, secret)
	return hmac.Equal([]byte(expected), []byte(signature))
}
