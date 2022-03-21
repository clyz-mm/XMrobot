package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// Guild 拉取频道信息
func (o *openAPI) Guild(ctx context.Context, guildID string) (*mydto.Guild, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Guild{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(guildURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.Guild), nil
}
