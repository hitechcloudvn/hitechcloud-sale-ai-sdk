package saleai

import "fmt"

// CommentsResource - Thao tác bình luận trên kênh (comments)
type CommentsResource struct {
	client *HTTPClient
}

// Hide - Ẩn một bình luận
func (r *CommentsResource) Hide(commentId string, pageAccessToken string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/comments/%s/hide", commentId), nil, nil, map[string]string{"X-Page-Access-Token": pageAccessToken})
}

// Like - Thích một bình luận
func (r *CommentsResource) Like(commentId string, pageAccessToken string) (map[string]interface{}, error) {
	return r.client.Request("POST", fmt.Sprintf("/api/v1/comments/%s/like", commentId), nil, nil, map[string]string{"X-Page-Access-Token": pageAccessToken})
}
