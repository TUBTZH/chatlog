# 数据库解密规范

## 支持版本

- 微信 3.x
- 微信 4.x

## 密钥获取

### macOS

需要临时关闭 SIP (系统完整性保护):

1. 进入恢复模式
2. 执行 `csrutil disable`
3. 重启后使用 chatlog 获取密钥
4. 可选择重新开启 SIP

### Windows / macOS

```bash
chatlog key
```

获取:
- Data Key (数据库解密密钥)
- Image Key (图片解密密钥)

## 自动解密

开启自动解密后，程序监控微信数据库目录，新消息写入时自动解密。

### TUI 模式

菜单: `解密数据` -> `开启自动解密`

### Server 模式

配置文件添加 `autoDecrypt: true`

## 数据库文件

| 类型 | 模式 | 文件 |
|------|------|------|
| 会话 | 3.x/4.x | `session.db` |
| 消息 | 3.x | `MSG*.db` |
| 消息 | 4.x | `Message/*.db` |
| 联系人 | 3.x/4.x | `contact.db` |
| 群聊 | 3.x/4.x | `MicroMsg.db` |

## 输出

解密后数据库存放位置: `~/.chatlog/decrypted/`
