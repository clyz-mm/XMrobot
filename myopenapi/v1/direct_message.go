package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// CreateDirectMessage 创建私信频道
func (o *openAPI) CreateDirectMessage(ctx context.Context, dm *mydto.DirectMessageToCreate) (*mydto.DirectMessage, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.DirectMessage{}).
		SetBody(dm).
		Post(o.getURL(userMeDMURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.DirectMessage), nil
}

// PostDirectMessage 在私信频道内发消息
func (o *openAPI) PostDirectMessage(ctx context.Context,
	dm *mydto.DirectMessage, msg *mydto.MessageToCreate) (*mydto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Message{}).
		SetPathParam("guild_id", dm.GuildID).
		SetBody(msg).
		Post(o.getURL(dmsURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.Message), nil
}

// RetractDMMessage 撤回私信消息
func (o *openAPI) RetractDMMessage(ctx context.Context, guildID, msgID string) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("message_id", string(msgID)).
		Delete(o.getURL(dmsMessageURI))
	return err
}
