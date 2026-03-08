# MCP 协议规范

## 端点

| 类型 | 路径 |
|------|------|
| Streamable HTTP | `/mcp` |
| Server-Sent Events | `/sse` |

## MCP Tools

### tools.list

列出可用工具。

### tools.call

调用工具:

#### get_chatlogs

获取聊天记录

**参数:**
- `time`: 时间范围字符串
- `talker`: 聊天对象标识
- `limit`: 返回数量
- `offset`: 分页偏移

**返回:**
- 消息数组，包含时间、发送者、内容等字段

#### get_contacts

获取联系人列表

#### get_chatrooms

获取群聊列表

#### get_sessions

获取会话列表

## 集成客户端

| 客户端 | 连接方式 |
|--------|----------|
| ChatWise | SSE: `http://127.0.0.1:5030/sse` |
| Cherry Studio | SSE: `http://127.0.0.1:5030/sse` |
| Claude Desktop | mcp-proxy 转发 |
| Monica Code | mcp-proxy 转发 |
