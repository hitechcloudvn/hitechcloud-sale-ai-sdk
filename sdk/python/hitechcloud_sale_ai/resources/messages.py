"""Gửi tin nhắn (text/image/file/quick-reply/button/carousel)."""

from .base import BaseResource


class MessagesResource(BaseResource):
    """Send messages through a bot. Requires the page access token.

    The page access token can be set once on the client
    (``page_access_token=...``) or passed per call.
    """

    def _send(self, bot_id, message, page_access_token=None):
        headers = {"X-Page-Access-Token": page_access_token} if page_access_token else None
        return self._http.request("POST", f"/api/v1/bots/{bot_id}/messages",
                                  json_data=message, headers=headers)

    def send(self, bot_id, message, page_access_token=None):
        """POST /api/v1/bots/{bot_id}/messages - send a raw message payload."""
        return self._send(bot_id, message, page_access_token)

    def send_text(self, bot_id, subscriber_id, text, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "text",
                                   "text": text}, page_access_token)

    def send_image(self, bot_id, subscriber_id, url, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "image",
                                   "url": url}, page_access_token)

    def send_file(self, bot_id, subscriber_id, url, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "file",
                                   "url": url}, page_access_token)

    def send_quick_reply(self, bot_id, subscriber_id, text, quick_replies, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "quick_reply",
                                   "text": text, "quick_replies": quick_replies}, page_access_token)

    def send_button(self, bot_id, subscriber_id, text, buttons, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "button",
                                   "text": text, "buttons": buttons}, page_access_token)

    def send_carousel(self, bot_id, subscriber_id, elements, page_access_token=None):
        return self._send(bot_id, {"subscriber_id": subscriber_id, "type": "carousel",
                                   "elements": elements}, page_access_token)
