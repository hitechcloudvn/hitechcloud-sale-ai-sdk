# Phân trang & Giới hạn tốc độ

## Phân trang

Mọi endpoint danh sách dùng `page` và `per_page`, trả về khối `meta.pagination`.

| Tham số | Kiểu | Mặc định | Mô tả |
|---------|------|----------|-------|
| `page` | integer | 1 | Số thứ tự trang |
| `per_page` | integer | 20 | Số bản ghi mỗi trang (tối đa 100) |

```json
{
  "meta": {
    "pagination": {
      "total": 128, "per_page": 20, "current_page": 1,
      "last_page": 7, "has_more": true
    }
  }
}
```

### Duyệt hết dữ liệu (Python)

```python
page = 1
while True:
    res = client.subscribers.list(params={"bot_id": "bot_7c1a9f2e", "page": page, "per_page": 100})
    for sub in res["data"]:
        print(sub["id"])
    if not res["meta"]["pagination"]["has_more"]:
        break
    page += 1
```

## Giới hạn tốc độ

Mỗi access token có hạn mức gọi API theo phút. Thông tin trả trong header:

| Header | Ý nghĩa |
|--------|---------|
| `X-RateLimit-Limit` | Số request tối đa trong cửa sổ |
| `X-RateLimit-Remaining` | Số request còn lại |
| `X-RateLimit-Reset` | Thời điểm (epoch) cửa sổ reset |
| `Retry-After` | Số giây cần chờ khi bị `429` |

Khi vượt hạn mức, API trả `429 RATE_LIMITED`. Tất cả SDK trong bộ này **tự động
thử lại** theo `Retry-After` (và backoff luỹ thừa cho lỗi 5xx). Gửi broadcast
lớn qua endpoint broadcast thay vì lặp lại lời gọi gửi tin.
