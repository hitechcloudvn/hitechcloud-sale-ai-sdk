# API reference

> **English** | [Tiếng Việt](api-reference.vi.md)

---

Production host: `https://sale-ai.hitechcloud.vn` · Base path: `/api/v1`

Total SDK methods: **146** (excluding the typed message-sending helpers).

## Bots & connected channels (Facebook/Instagram/Zalo/Website) (`client.bots`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/bots` |
| `get` | GET | `/api/v1/bots/{id}` |
| `create` | POST | `/api/v1/bots` |
| `update` | PATCH | `/api/v1/bots/{id}` |
| `delete` | DELETE | `/api/v1/bots/{id}` |
| `connect` | POST | `/api/v1/bots/{bot_id}/connect` |
| `disconnect` | POST | `/api/v1/bots/{bot_id}/disconnect` |
| `updateSettings` | PATCH | `/api/v1/bots/{bot_id}/settings` |
| `refreshAccessToken` | POST | `/api/v1/bots/{bot_id}/access-token/refresh` |

## Conversation flows & blocks (`client.flows`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/flows` |
| `get` | GET | `/api/v1/flows/{id}` |
| `create` | POST | `/api/v1/flows` |
| `update` | PATCH | `/api/v1/flows/{id}` |
| `delete` | DELETE | `/api/v1/flows/{id}` |
| `publish` | POST | `/api/v1/flows/{id}/publish` |
| `duplicate` | POST | `/api/v1/flows/{id}/duplicate` |
| `trigger` | POST | `/api/v1/flows/{id}/trigger` |
| `listBlocks` | GET | `/api/v1/flows/{id}/blocks` |
| `addBlock` | POST | `/api/v1/flows/{id}/blocks` |
| `updateBlock` | PATCH | `/api/v1/flows/{id}/blocks/{block_id}` |
| `deleteBlock` | DELETE | `/api/v1/flows/{id}/blocks/{block_id}` |

## Auto-reply keywords (`client.keywords`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/keywords` |
| `get` | GET | `/api/v1/keywords/{id}` |
| `create` | POST | `/api/v1/keywords` |
| `update` | PATCH | `/api/v1/keywords/{id}` |
| `delete` | DELETE | `/api/v1/keywords/{id}` |
| `test` | POST | `/api/v1/keywords/test` |

## Message sequences & steps (`client.sequences`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/sequences` |
| `get` | GET | `/api/v1/sequences/{id}` |
| `create` | POST | `/api/v1/sequences` |
| `update` | PATCH | `/api/v1/sequences/{id}` |
| `delete` | DELETE | `/api/v1/sequences/{id}` |
| `listSteps` | GET | `/api/v1/sequences/{id}/steps` |
| `addStep` | POST | `/api/v1/sequences/{id}/steps` |
| `updateStep` | PATCH | `/api/v1/sequences/{id}/steps/{step_id}` |
| `deleteStep` | DELETE | `/api/v1/sequences/{id}/steps/{step_id}` |
| `subscribe` | POST | `/api/v1/sequences/{id}/subscribe` |
| `unsubscribe` | POST | `/api/v1/sequences/{id}/unsubscribe` |

## Broadcasts (`client.broadcasts`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/broadcasts` |
| `get` | GET | `/api/v1/broadcasts/{id}` |
| `create` | POST | `/api/v1/broadcasts` |
| `update` | PATCH | `/api/v1/broadcasts/{id}` |
| `delete` | DELETE | `/api/v1/broadcasts/{id}` |
| `schedule` | POST | `/api/v1/broadcasts/{id}/schedule` |
| `send` | POST | `/api/v1/broadcasts/{id}/send` |
| `status` | GET | `/api/v1/broadcasts/{id}/status` |
| `previewAudience` | POST | `/api/v1/broadcasts/{id}/audience/preview` |
| `cancel` | POST | `/api/v1/broadcasts/{id}/cancel` |

## Subscribers, tags, custom fields & notes (`client.subscribers`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/subscribers` |
| `get` | GET | `/api/v1/subscribers/{id}` |
| `create` | POST | `/api/v1/subscribers` |
| `update` | PATCH | `/api/v1/subscribers/{id}` |
| `delete` | DELETE | `/api/v1/subscribers/{id}` |
| `addTags` | POST | `/api/v1/subscribers/{id}/tags` |
| `removeTag` | DELETE | `/api/v1/subscribers/{id}/tags/{tag}` |
| `setCustomFields` | PATCH | `/api/v1/subscribers/{id}/custom-fields` |
| `addNote` | POST | `/api/v1/subscribers/{id}/notes` |
| `block` | POST | `/api/v1/subscribers/{id}/block` |
| `unblock` | POST | `/api/v1/subscribers/{id}/unblock` |
| `listMessages` | GET | `/api/v1/subscribers/{id}/messages` |
| `sendMessage` | POST | `/api/v1/subscribers/{id}/messages` |
| `markSeen` | POST | `/api/v1/subscribers/{id}/mark-seen` |
| `typing` | POST | `/api/v1/subscribers/{id}/typing` |

## Tags (`client.tags`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/tags` |
| `get` | GET | `/api/v1/tags/{id}` |
| `create` | POST | `/api/v1/tags` |
| `update` | PATCH | `/api/v1/tags/{id}` |
| `delete` | DELETE | `/api/v1/tags/{id}` |

## Segments (`client.segments`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/segments` |
| `get` | GET | `/api/v1/segments/{id}` |
| `create` | POST | `/api/v1/segments` |
| `update` | PATCH | `/api/v1/segments/{id}` |
| `delete` | DELETE | `/api/v1/segments/{id}` |

## Products (`client.products`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/products` |
| `get` | GET | `/api/v1/products/{id}` |
| `create` | POST | `/api/v1/products` |
| `update` | PATCH | `/api/v1/products/{id}` |
| `delete` | DELETE | `/api/v1/products/{id}` |

## Orders (`client.orders`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/orders` |
| `get` | GET | `/api/v1/orders/{id}` |
| `create` | POST | `/api/v1/orders` |
| `update` | PATCH | `/api/v1/orders/{id}` |
| `delete` | DELETE | `/api/v1/orders/{id}` |
| `createFromConversation` | POST | `/api/v1/orders/from-conversation` |
| `updateStatus` | POST | `/api/v1/orders/{id}/status` |

## Comment reply rules (`client.commentRules`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/comment-rules` |
| `get` | GET | `/api/v1/comment-rules/{id}` |
| `create` | POST | `/api/v1/comment-rules` |
| `update` | PATCH | `/api/v1/comment-rules/{id}` |
| `delete` | DELETE | `/api/v1/comment-rules/{id}` |

## Growth widgets (`client.widgets`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/widgets` |
| `get` | GET | `/api/v1/widgets/{id}` |
| `create` | POST | `/api/v1/widgets` |
| `update` | PATCH | `/api/v1/widgets/{id}` |
| `delete` | DELETE | `/api/v1/widgets/{id}` |

## AI intents (`client.intents`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/ai/intents` |
| `get` | GET | `/api/v1/ai/intents/{id}` |
| `create` | POST | `/api/v1/ai/intents` |
| `update` | PATCH | `/api/v1/ai/intents/{id}` |
| `delete` | DELETE | `/api/v1/ai/intents/{id}` |

## AI knowledge base (`client.knowledgeBase`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/ai/knowledge-base` |
| `get` | GET | `/api/v1/ai/knowledge-base/{id}` |
| `create` | POST | `/api/v1/ai/knowledge-base` |
| `update` | PATCH | `/api/v1/ai/knowledge-base/{id}` |
| `delete` | DELETE | `/api/v1/ai/knowledge-base/{id}` |

## Users / team (`client.users`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/users` |
| `get` | GET | `/api/v1/users/{id}` |
| `create` | POST | `/api/v1/users` |
| `update` | PATCH | `/api/v1/users/{id}` |
| `delete` | DELETE | `/api/v1/users/{id}` |

## Roles & permissions (`client.roles`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/roles` |
| `get` | GET | `/api/v1/roles/{id}` |
| `create` | POST | `/api/v1/roles` |
| `update` | PATCH | `/api/v1/roles/{id}` |
| `delete` | DELETE | `/api/v1/roles/{id}` |

## Webhook registration (`client.webhooks`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/webhooks` |
| `get` | GET | `/api/v1/webhooks/{id}` |
| `create` | POST | `/api/v1/webhooks` |
| `update` | PATCH | `/api/v1/webhooks/{id}` |
| `delete` | DELETE | `/api/v1/webhooks/{id}` |

## Channel comment operations (`client.comments`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `hide` | POST | `/api/v1/comments/{comment_id}/hide` |
| `like` | POST | `/api/v1/comments/{comment_id}/like` |

## Sales AI assistant (config / train / suggest / handover) (`client.ai`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `updateConfig` | PATCH | `/api/v1/ai/config` |
| `train` | POST | `/api/v1/ai/train` |
| `suggestReply` | POST | `/api/v1/ai/suggest-reply` |
| `handover` | POST | `/api/v1/ai/handover` |

## Unified inbox (conversations / inbox) (`client.conversations`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `list` | GET | `/api/v1/conversations` |
| `get` | GET | `/api/v1/conversations/{id}` |
| `assign` | POST | `/api/v1/conversations/{id}/assign` |
| `addNote` | POST | `/api/v1/conversations/{id}/notes` |
| `addLabels` | POST | `/api/v1/conversations/{id}/labels` |
| `close` | POST | `/api/v1/conversations/{id}/close` |

## Reports & analytics (`client.analytics`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `overview` | GET | `/api/v1/analytics/overview` |
| `bot` | GET | `/api/v1/analytics/bots/{bot_id}` |
| `flow` | GET | `/api/v1/analytics/flows/{id}` |
| `broadcast` | GET | `/api/v1/analytics/broadcasts/{id}` |
| `subscriberGrowth` | GET | `/api/v1/analytics/subscriber-growth` |

## Account settings (`client.settings`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `get` | GET | `/api/v1/settings` |
| `update` | PATCH | `/api/v1/settings` |

## Messaging (`client.messages`)

| Method | HTTP | Endpoint |
|--------|------|----------|
| `send` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendText` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendImage` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendFile` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendQuickReply` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendButton` | POST | `/api/v1/bots/{bot_id}/messages` |
| `sendCarousel` | POST | `/api/v1/bots/{bot_id}/messages` |

> All message-sending methods require `X-Page-Access-Token`.

## Webhook

Events are delivered via HTTP POST with the `X-SaleAI-Signature` (HMAC-SHA256)
signature header:

```
message.received, message.delivered, message.seen, subscriber.created, subscriber.updated, subscriber.unsubscribed, tag.added, keyword.matched, flow.completed, sequence.completed, handover.requested, handover.resolved, order.created, order.updated, order.paid, comment.received, broadcast.sent, broadcast.completed
```
