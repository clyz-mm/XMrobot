package mybotgo

import (
	"XMrobot/local"
	"XMrobot/mybotgo/mydto"
	"XMrobot/mytoken"
)

// defaultSessionManager 默认实现的 session manager 为单机版本
// 如果业务要自行实现分布式的 session 管理，则实现 SessionManger 后替换掉 defaultSessionManager
var defaultSessionManager SessionManager = local.New()

// SessionManager 接口，管理session
type SessionManager interface {
	// Start 启动连接，默认使用 apInfo 中的 shards 作为 shard 数量，如果有需要自己指定 shard 数，请修 apInfo 中的信息
	Start(apInfo *mydto.WebsocketAP, token *mytoken.MyToken, intents *mydto.Intent) error
}
