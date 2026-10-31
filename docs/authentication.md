# Xác thực

HiTechCloud Sale AI dùng hai lớp xác thực.

## 1. Access token (Bearer)

Bắt buộc cho mọi request, gửi trong header `Authorization`:

```
Authorization: Bearer YOUR_TOKEN
```

Trong SDK: truyền `token` khi khởi tạo client, hoặc gọi `setToken()`.

## 2. Page access token theo bot

Các endpoint tác động trực tiếp lên kênh (gửi tin, ẩn/thích bình luận) cần thêm
header `X-Page-Access-Token` chứa token trang của bot tương ứng. Token này được
lưu mã hoá phía máy chủ và không bao giờ trả về đầy đủ.

```
X-Page-Access-Token: PAGE_ACCESS_TOKEN
```

Trong SDK, bạn có thể:

- Đặt mặc định cho client: `SaleAI(token=..., page_access_token=...)`.
- Hoặc truyền theo từng lời gọi: `client.messages.send_text(..., page_access_token=...)`.

## Mã lỗi xác thực

| HTTP | Mã lỗi | Ý nghĩa |
|------|--------|---------|
| 401 | `UNAUTHORIZED` | Thiếu hoặc sai access token |
| 401 | `TOKEN_EXPIRED` | Token đã hết hạn |
| 403 | `FORBIDDEN` | Không có quyền với tài nguyên/bot |
| 403 | `PAGE_TOKEN_REQUIRED` | Thiếu `X-Page-Access-Token` cho thao tác kênh |

> Luôn gọi API từ máy chủ. Không để lộ access token hay page access token trong
> ứng dụng phía người dùng.
