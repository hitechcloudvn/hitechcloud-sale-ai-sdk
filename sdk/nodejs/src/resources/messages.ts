import { BaseResource } from "./base";

/** Gửi tin nhắn (text/image/file/quick-reply/button/carousel). */
export class MessagesResource extends BaseResource {
  private sendRaw(botId: string, message: Record<string, unknown>, pat?: string): Promise<unknown> {
    const headers = pat ? { "X-Page-Access-Token": pat } : undefined;
    return this.http.request("POST", `/api/v1/bots/${botId}/messages`, message, undefined, headers);
  }

  /** POST /api/v1/bots/{bot_id}/messages - send a raw message payload. */
  send(botId: string, message: Record<string, unknown>, pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, message, pageAccessToken);
  }

  sendText(botId: string, subscriberId: string, text: string, pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "text", text }, pageAccessToken);
  }

  sendImage(botId: string, subscriberId: string, url: string, pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "image", url }, pageAccessToken);
  }

  sendFile(botId: string, subscriberId: string, url: string, pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "file", url }, pageAccessToken);
  }

  sendQuickReply(botId: string, subscriberId: string, text: string, quickReplies: string[], pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "quick_reply", text, quick_replies: quickReplies }, pageAccessToken);
  }

  sendButton(botId: string, subscriberId: string, text: string, buttons: Record<string, unknown>[], pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "button", text, buttons }, pageAccessToken);
  }

  sendCarousel(botId: string, subscriberId: string, elements: Record<string, unknown>[], pageAccessToken?: string): Promise<unknown> {
    return this.sendRaw(botId, { subscriber_id: subscriberId, type: "carousel", elements }, pageAccessToken);
  }
}
