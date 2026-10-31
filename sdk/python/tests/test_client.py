import hitechcloud_sale_ai as sdk


def test_client_has_resources():
    client = sdk.SaleAI(token="t")
    for name in ["bots", "flows", "keywords", "sequences", "broadcasts",
                 "subscribers", "tags", "segments", "products", "orders",
                 "comment_rules", "widgets", "intents", "knowledge_base",
                 "users", "roles", "webhooks", "comments", "ai",
                 "conversations", "analytics", "settings", "messages"]:
        assert hasattr(client, name), name


def test_webhook_verify():
    body = b'{"event":"message.received"}'
    secret = "whsec_test"
    sig = sdk.compute_signature(body, secret)
    assert sdk.verify_signature(body, sig, secret)
    assert not sdk.verify_signature(body, "sha256=bad", secret)
