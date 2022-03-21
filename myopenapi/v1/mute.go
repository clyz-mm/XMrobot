package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// GuildMute 频道禁言
func (o *openAPI) GuildMute(ctx context.Context, guildID string, mute *mydto.UpdateGuildMute) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetBody(mute).
		Patch(o.getURL(guildMuteURI))
	if err != nil {
		return err
	}
	return nil
}

// MemberMute 频道指定成员禁言
func (o *openAPI) MemberMute(ctx context.Context, guildID, userID string,
	mute *mydto.UpdateGuildMute) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		SetBody(mute).
		Patch(o.getURL(guildMembersMuteURI))
	if err != nil {
		return err
	}
	return nil
}
