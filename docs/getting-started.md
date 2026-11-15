# Bắt đầu nhanh

HiTechCloud Sale AI giúp bán hàng và tự động hoá hội thoại trực tiếp trong tin
nhắn trên nhiều kênh (Facebook Messenger, Instagram, Zalo, Website).

## 1. Lấy access token

1. Đăng nhập cổng quản trị tại `https://sale-ai.hitechcloud.vn`.
2. Vào **Cấu hình → API & Tích hợp → Access Token**.
3. Tạo token mới, chọn phạm vi (scope) và sao chép (chỉ hiển thị một lần).
4. Lưu token trong biến môi trường, không commit lên kho mã.

## 2. Cài đặt SDK

| Ngôn ngữ | Lệnh |
|----------|------|
| Node.js / TS | `npm install @hitechcloudvn/sale-ai-sdk` |
| Python | `pip install hitechcloud-sale-ai` |
| PHP | `composer require hitechcloudvn/sale-ai-sdk` |
| Go | `go get github.com/hitechcloudvn/hitechcloud-sale-ai-sdk/go` |
| Java | `vn.hitechcloud:sale-ai-sdk` (Maven) |

## 3. Gọi API đầu tiên

```python
from hitechcloud_sale_ai import SaleAI

client = SaleAI(token="YOUR_TOKEN")
bots = client.bots.list()
print(bots["data"])
```

## 4. Gửi tin nhắn đầu tiên

Các endpoint gửi tin cần thêm header `X-Page-Access-Token` (token trang của bot):

```python
client.messages.send_text(
    "bot_7c1a9f2e", "sub_2a9b8c7d", "Chao ban!",
    page_access_token="PAGE_ACCESS_TOKEN",
)
```

## 5. Môi trường Sandbox

Dùng `base_url="https://sale-ai-sandbox.hitechcloud.vn"` để thử nghiệm. Trên sandbox, tin nhắn KHÔNG gửi ra kênh
thật nhưng webhook vẫn phát sự kiện thật để bạn kiểm thử trình nhận.
