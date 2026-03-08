# Webhook 规范

## 配置方式

### TUI 模式

在 `~/.chatlog/chatlog.json` 中配置:

```json
{
  "webhook": {
    "host": "localhost:5030",
    "items": [
      {
        "url": "http://localhost:8080/webhook",
        "talker": "wxid_123",
        "sender": "",
        "keyword": ""
      }
    ]
  }
}
```

### Server 模式

通过环境变量:

```bash
# 方式 1: JSON 字符串
CHATLOG_WEBHOOK='{"host":"localhost:5030","items":[{"url":"http://localhost:8080/proxy","talker":"wxid_123","sender":"","keyword":""}]}'

# 方式 2: 拆分变量
CHATLOG_WEBHOOK_HOST="localhost:5030"
CHATLOG_WEBHOOK_ITEMS='[{"url":"http://localhost:8080/proxy","talker":"wxid_123","sender":"","keyword":""}]'
```

## 配置字段

| 字段 | 必填 | 说明 |
|------|------|------|
| host | 是 | 消息中多媒体 URL 的 host |
| url | 是 | webhook 回调 URL |
| talker | 是 | 监控的聊天对象 |
| sender | 否 | 消息发送者筛选 |
| keyword | 否 | 关键词筛选 |

## 回调请求

**请求:**
```
POST <webhook-url>
Content-Type: application/json
```

**请求体:**
```json
{
  "keyword": "",
  "lastTime": "2025-08-27 00:00:00",
  "length": 1,
  "messages": [
    {
      "seq": 1756225000000,
      "time": "2025-08-27T00:00:00+08:00",
      "talker": "wxid_123",
      "talkerName": "",
      "isChatRoom": false,
      "sender": "wxid_123",
      "senderName": "Name",
      "isSelf": false,
      "type": 1,
      "subType": 0,
      "content": "测试消息",
      "contents": {
        "host": "localhost:5030"
      }
    }
  ],
  "sender": "",
  "talker": "wxid_123"
}
```

## 性能

- 本地服务消息回调延迟: ~13 秒
- 远程同步消息回调延迟: ~45 秒
