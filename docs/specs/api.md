# API 规范

## 服务地址

默认: `http://127.0.0.1:5030`

## 接口列表

### 聊天记录查询

```
GET /api/v1/chatlog
```

**参数:**
| 参数 | 必填 | 说明 |
|------|------|------|
| time | 是 | 时间范围 |
| talker | 是 | 聊天对象标识 (wxid/群ID/备注名/昵称) |
| sender | 否 | 发送者筛选 |
| keyword | 否 | 关键词筛选 (正则表达式) |
| limit | 否 | 返回数量 (默认 5000) |
| offset | 否 | 分页偏移 |
| format | 否 | 输出格式: json/csv/plain (默认 json) |

**返回顺序:** 按时间正序（最早的消息在前），方便分页查询

**时间格式:**
- 单个日期: `2023-01-01` (查询当天 00:00-23:59)
- 日期范围: `2023-01-01~2023-01-31`
- 相对时间: `last-7d`, `last-30d`
- 时间戳: `1609459200`

**示例:**
```bash
# 查询 2026-03-09 群聊消息（默认返回当天全部消息）
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom"

# 查询并限制返回数量
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom&limit=1000"

# 分页查询（查询更早的消息）
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom&limit=500&offset=500"

# 查询特定发送者的消息
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom&sender=wxid_xxx"

# 关键词筛选
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom&keyword=正则表达式"

# 输出 CSV 格式
curl "http://127.0.0.1:5030/api/v1/chatlog?time=2026-03-09&talker=21042627603@chatroom&format=csv"
```

### 联系人列表

```
GET /api/v1/contact
```

**参数:**
| 参数 | 必填 | 说明 |
|------|------|------|
| keyword | 否 | 搜索关键词 (支持模糊匹配) |
| limit | 否 | 返回数量 |
| offset | 否 | 分页偏移 |
| format | 否 | 输出格式 |

**示例:**
```bash
# 查询所有联系人
curl "http://127.0.0.1:5030/api/v1/contact"

# 搜索联系人
curl "http://127.0.0.1:5030/api/v1/contact?keyword=张三"
```

### 群聊列表

```
GET /api/v1/chatroom
```

**参数:**
| 参数 | 必填 | 说明 |
|------|------|------|
| keyword | 否 | 搜索关键词 (支持模糊匹配群名) |
| limit | 否 | 返回数量 |
| offset | 否 | 分页偏移 |
| format | 否 | 输出格式 |

**示例:**
```bash
# 查询所有群聊
curl "http://127.0.0.1:5030/api/v1/chatroom"

# 搜索群聊
curl "http://127.0.0.1:5030/api/v1/chatroom?keyword=4403"
```

### 会话列表

```
GET /api/v1/session
```

**参数:**
| 参数 | 必填 | 说明 |
|------|------|------|
| keyword | 否 | 搜索关键词 |
| limit | 否 | 返回数量 |
| offset | 否 | 分页偏移 |
| format | 否 | 输出格式 |

### 多媒体内容

| 类型 | 路径 |
|------|------|
| 图片 | `/image/<id>` |
| 视频 | `/video/<id>` |
| 语音 | `/voice/<id>` |
| 文件 | `/file/<id>` |
| 数据文件 | `/data/<relative path>` |

**说明:**
- 图片/视频/文件请求返回 302 跳转到实际多媒体 URL
- 语音直接返回 MP3（实时转码 SILK）
- 数据文件直接返回，对加密图片实时解密
