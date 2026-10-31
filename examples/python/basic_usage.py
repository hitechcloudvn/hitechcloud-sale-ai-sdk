#!/usr/bin/env python3
"""Ví dụ cơ bản - HiTechCloud Sale AI Python SDK.

    export SALE_AI_TOKEN=your-token
    export SALE_AI_PAGE_ACCESS_TOKEN=page-token
"""
import os

from hitechcloud_sale_ai import SaleAI, verify_signature

client = SaleAI(
    token=os.environ["SALE_AI_TOKEN"],
    page_access_token=os.environ.get("SALE_AI_PAGE_ACCESS_TOKEN"),
)

# Liệt kê bot
bots = client.bots.list()
for bot in bots.get("data", []):
    print(bot["id"], bot["name"], bot["channel"])

# Tạo & xuất bản kịch bản
flow = client.flows.create(data={"bot_id": "bot_7c1a9f2e", "name": "Chao mung", "trigger_type": "welcome"})
client.flows.publish(flow["data"]["id"])

# Gửi carousel sản phẩm
client.messages.send_carousel(
    "bot_7c1a9f2e", "sub_2a9b8c7d",
    elements=[{
        "title": "Ao thun basic", "subtitle": "199.000d",
        "image_url": "https://cdn.hitechcloud.vn/prod/at-basic-01.jpg",
        "buttons": [{"type": "postback", "title": "Dat ngay", "payload": "BUY_prod_a1f2e3d4"}],
    }],
)

# Báo cáo tổng quan
print(client.analytics.overview(params={"bot_id": "bot_7c1a9f2e", "from": "2026-10-01", "to": "2026-10-31"}))

# Xác minh chữ ký webhook
raw = b'{"event":"message.received"}'
from hitechcloud_sale_ai import compute_signature
sig = compute_signature(raw, "whsec_demo")
assert verify_signature(raw, sig, "whsec_demo")
