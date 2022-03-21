package websocket

import (
	"XMrobot/mybotgo/mydto"
)

// DefaultHandlers 默认的 handler 结构，管理所有支持的 handler 类型
var DefaultHandlers struct {
	Ready       ReadyHandler
	ErrorNotify ErrorNotifyHandler
	Plain       PlainEventHandler

	Guild           GuildEventHandler
	GuildMember     GuildMemberEventHandler
	Channel         ChannelEventHandler
	Message         MessageEventHandler
	MessageReaction MessageReactionEventHandler
	ATMessage       ATMessageEventHandler
	DirectMessage   DirectMessageEventHandler
	Audio           AudioEventHandler
	MessageAudit    MessageAuditEventHandler
}

// ReadyHandler 可以处理 ws 的 ready 事件
type ReadyHandler func(event *mydto.WSPayload, data *mydto.WSReadyData)

// ErrorNotifyHandler 当 ws 连接发生错误的时候，会回调，方便使用方监控相关错误
// 比如 reconnect invalidSession 等错误，错误可以转换为 bot.Err
type ErrorNotifyHandler func(err error)

// PlainEventHandler 透传handler
type PlainEventHandler func(event *mydto.WSPayload, message []byte) error

// GuildEventHandler 频道事件handler
type GuildEventHandler func(event *mydto.WSPayload, data *mydto.WSGuildData) error

// GuildMemberEventHandler 频道成员事件 handler
type GuildMemberEventHandler func(event *mydto.WSPayload, data *mydto.WSGuildMemberData) error

// ChannelEventHandler 子频道事件 handler
type ChannelEventHandler func(event *mydto.WSPayload, data *mydto.WSChannelData) error

// MessageEventHandler 消息事件 handler
type MessageEventHandler func(event *mydto.WSPayload, data *mydto.WSMessageData) error

// MessageReactionEventHandler 表情表态事件 handler
type MessageReactionEventHandler func(event *mydto.WSPayload, data *mydto.WSMessageReactionData) error

// ATMessageEventHandler at 机器人消息事件 handler
type ATMessageEventHandler func(event *mydto.WSPayload, data *mydto.WSATMessageData) error

// DirectMessageEventHandler 私信消息事件 handler
type DirectMessageEventHandler func(event *mydto.WSPayload, data *mydto.WSDirectMessageData) error

// AudioEventHandler 音频机器人事件 handler
type AudioEventHandler func(event *mydto.WSPayload, data *mydto.WSAudioData) error

// MessageAuditEventHandler 消息审核事件 handler
type MessageAuditEventHandler func(event *mydto.WSPayload, data *mydto.WSMessageAuditData) error

// RegisterHandlers 注册事件回调，并返回 intent 用于 websocket 的鉴权
func RegisterHandlers(handlers ...interface{}) mydto.Intent {
	var i mydto.Intent
	for _, h := range handlers {
		switch handle := h.(type) {
		case ReadyHandler:
			DefaultHandlers.Ready = handle
		case ErrorNotifyHandler:
			DefaultHandlers.ErrorNotify = handle
		case PlainEventHandler:
			DefaultHandlers.Plain = handle
		case AudioEventHandler:
			DefaultHandlers.Audio = handle
			i = i | mydto.EventToIntent(
				mydto.EventAudioStart, mydto.EventAudioFinish,
				mydto.EventAudioOnMic, mydto.EventAudioOffMic,
			)
		default:
		}
	}
	i = i | registerRelationHandlers(i, handlers...)
	i = i | registerMessageHandlers(i, handlers...)

	return i
}

// registerRelationHandlers 注册频道关系链相关handlers
func registerRelationHandlers(i mydto.Intent, handlers ...interface{}) mydto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case GuildEventHandler:
			DefaultHandlers.Guild = handle
			i = i | mydto.EventToIntent(mydto.EventGuildCreate, mydto.EventGuildDelete, mydto.EventGuildUpdate)
		case GuildMemberEventHandler:
			DefaultHandlers.GuildMember = handle
			i = i | mydto.EventToIntent(mydto.EventGuildMemberAdd, mydto.EventGuildMemberRemove, mydto.EventGuildMemberUpdate)
		case ChannelEventHandler:
			DefaultHandlers.Channel = handle
			i = i | mydto.EventToIntent(mydto.EventChannelCreate, mydto.EventChannelDelete, mydto.EventChannelUpdate)
		default:
		}
	}
	return i
}

// registerMessageHandlers 注册消息相关的 handler
func registerMessageHandlers(i mydto.Intent, handlers ...interface{}) mydto.Intent {
	for _, h := range handlers {
		switch handle := h.(type) {
		case MessageEventHandler:
			DefaultHandlers.Message = handle
			i = i | mydto.EventToIntent(mydto.EventMessageCreate)
		case ATMessageEventHandler:
			DefaultHandlers.ATMessage = handle
			i = i | mydto.EventToIntent(mydto.EventAtMessageCreate)
		case DirectMessageEventHandler:
			DefaultHandlers.DirectMessage = handle
			i = i | mydto.EventToIntent(mydto.EventDirectMessageCreate)
		case MessageReactionEventHandler:
			DefaultHandlers.MessageReaction = handle
			i = i | mydto.EventToIntent(mydto.EventMessageReactionAdd, mydto.EventMessageReactionRemove)
		case MessageAuditEventHandler:
			DefaultHandlers.MessageAudit = handle
			i = i | mydto.EventToIntent(mydto.EventMessageAuditPass, mydto.EventMessageAuditReject)
		default:
		}
	}
	return i
}
