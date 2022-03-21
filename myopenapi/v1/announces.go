package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// CreateChannelAnnounces 创建子频道公告
func (o *openAPI) CreateChannelAnnounces(ctx context.Context, channelID string,
	announce *mydto.ChannelAnnouncesToCreate) (*mydto.Announces, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("channel_id", channelID).
		SetBody(announce).
		Post(o.getURL(channelAnnouncesURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.Announces), nil
}

// DeleteChannelAnnounces 删除子频道公告,会校验 messageID
func (o *openAPI) DeleteChannelAnnounces(ctx context.Context, channelID, messageID string) error {
	_, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Delete(o.getURL(channelAnnounceURI))
	return err
}

// CleanChannelAnnounces 删除子频道公告,不校验 messageID
func (o *openAPI) CleanChannelAnnounces(ctx context.Context, channelID string) error {
	_, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", "all").
		Delete(o.getURL(channelAnnounceURI))
	return err
}

// CreateGuildAnnounces 创建频道全局公告
func (o *openAPI) CreateGuildAnnounces(ctx context.Context, guildID string,
	announce *mydto.GuildAnnouncesToCreate) (*mydto.Announces, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("guild_id", guildID).
		SetBody(announce).
		Post(o.getURL(guildAnnouncesURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.Announces), nil
}

// DeleteGuildAnnounces 删除频道全局公告,会校验 messageID
func (o *openAPI) DeleteGuildAnnounces(ctx context.Context, guildID, messageID string) error {
	_, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("message_id", messageID).
		Delete(o.getURL(guildAnnounceURI))
	return err
}

// CleanGuildAnnounces 删除道全局公告,不校验 messageID
func (o *openAPI) CleanGuildAnnounces(ctx context.Context, guildID string) error {
	_, err := o.request(ctx).
		SetResult(mydto.Announces{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("message_id", "all").
		Delete(o.getURL(guildAnnounceURI))
	return err
}
