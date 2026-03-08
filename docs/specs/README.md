# Chatlog Specification Index

本目录包含 Chatlog 项目的功能规范文档。采用按需加载方式，仅在需要时引入相关 spec。

## 规范索引

### 基础规范
- [API 规范](./specs/api.md) - HTTP API 接口定义
- [MCP 协议](./specs/mcp_protocol.md) - MCP 集成规范
- [Webhook 规范](./specs/webhook.md) - 消息回调规范

### 平台规范
- [macOS 规范](./specs/platform_darwin.md) - macOS 特有功能
- [Windows 规范](./specs/platform_windows.md) - Windows 特有功能

### 功能规范
- [数据库解密](./specs/decrypt.md) - 数据库解密流程
- [多媒体处理](./specs/media.md) - 图片/语音/视频处理
- [多账号管理](./specs/multi_account.md) - 多账号切换

## 使用方式

在代码开发或 AI 辅助时，按需加载相关规范：

```bash
# 查看 API 规范
cat docs/specs/api.md

# 查看 MCP 规范
cat docs/specs/mcp_protocol.md
```

## 添加新规范

新增功能时，在 `specs/` 目录下创建对应的 Markdown 规范文件，并在本索引中注册。
