package client

import (
	"XMrobot/mybotgo/mydto"
	"XMrobot/mybotgo/websocket"
	"encoding/json"

	"github.com/tidwall/gjson" // 由于回包的 d 类型不确定，gjson 用于从回包json中提取 d 并进行针对性的解析
)

var eventParseFuncMap = map[mydto.OPCode]map[mydto.EventType]eventParseFunc{
	mydto.WSDispatchEvent: {
		mydto.EventGuildCreate: guildHandler,
		mydto.EventGuildUpdate: guildHandler,
		mydto.EventGuildDelete: guildHandler,

		mydto.EventChannelCreate: channelHandler,
		mydto.EventChannelUpdate: channelHandler,
		mydto.EventChannelDelete: channelHandler,

		mydto.EventGuildMemberAdd:    guildMemberHandler,
		mydto.EventGuildMemberUpdate: guildMemberHandler,
		mydto.EventGuildMemberRemove: guildMemberHandler,

		mydto.EventMessageCreate: messageHandler,

		mydto.EventMessageReactionAdd:    messageReactionHandler,
		mydto.EventMessageReactionRemove: messageReactionHandler,

		mydto.EventAtMessageCreate:     atMessageHandler,
		mydto.EventDirectMessageCreate: directMessageHandler,

		mydto.EventAudioStart:  audioHandler,
		mydto.EventAudioFinish: audioHandler,
		mydto.EventAudioOnMic:  audioHandler,
		mydto.EventAudioOffMic: audioHandler,

		mydto.EventMessageAuditPass:   messageAuditHandler,
		mydto.EventMessageAuditReject: messageAuditHandler,
	},
}

type eventParseFunc func(event *mydto.WSPayload, message []byte) error

func parseAndHandle(event *mydto.WSPayload) error {
	// 指定类型的 handler
	if h, ok := eventParseFuncMap[event.OPCode][event.Type]; ok {
		return h(event, event.RawMessage)
	}
	// 透传handler，如果未注册具体类型的 handler，会统一投递到这个 handler
	if websocket.DefaultHandlers.Plain != nil {
		return websocket.DefaultHandlers.Plain(event, event.RawMessage)
	}
	return nil
}

func guildHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSGuildData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.Guild != nil {
		return websocket.DefaultHandlers.Guild(event, data)
	}
	return nil
}

func channelHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSChannelData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.Channel != nil {
		return websocket.DefaultHandlers.Channel(event, data)
	}
	return nil
}

func guildMemberHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSGuildMemberData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.GuildMember != nil {
		return websocket.DefaultHandlers.GuildMember(event, data)
	}
	return nil
}

func messageHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.Message != nil {
		return websocket.DefaultHandlers.Message(event, data)
	}
	return nil
}

func messageReactionHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSMessageReactionData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.MessageReaction != nil {
		return websocket.DefaultHandlers.MessageReaction(event, data)
	}
	return nil
}

func atMessageHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSATMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.ATMessage != nil {
		return websocket.DefaultHandlers.ATMessage(event, data)
	}
	return nil
}

func directMessageHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSDirectMessageData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.DirectMessage != nil {
		return websocket.DefaultHandlers.DirectMessage(event, data)
	}
	return nil
}

func audioHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSAudioData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.Audio != nil {
		return websocket.DefaultHandlers.Audio(event, data)
	}
	return nil
}

func parseData(message []byte, target interface{}) error {
	data := gjson.Get(string(message), "d")
	return json.Unmarshal([]byte(data.String()), target)
}

func messageAuditHandler(event *mydto.WSPayload, message []byte) error {
	data := &mydto.WSMessageAuditData{}
	if err := parseData(message, data); err != nil {
		return err
	}
	if websocket.DefaultHandlers.MessageAudit != nil {
		return websocket.DefaultHandlers.MessageAudit(event, data)
	}
	return nil
}
