package websocket

import (
	"XMrobot/mybotgo/mydto"
)

// WebSocket 需要实现的接口
type WebSocket interface {
	// New 创建一个新的ws实例，需要传递 session 对象
	New(session mydto.Session) WebSocket
	// Connect 连接到 wss 地址
	Connect() error
	// Identify 鉴权连接
	Identify() error
	// Session 拉取 session 信息，包括 token，shard，seq 等
	Session() *mydto.Session
	// Resume 重连
	Resume() error
	// Listening 监听websocket事件
	Listening() error
	// Write 发送数据
	Write(message *mydto.WSPayload) error
	// Close 关闭连接
	Close()
}
