package v1

import (
	"XMrobot/mybotgo/errs"
	"XMrobot/mybotgo/mydto"
	"context"
	"encoding/json"
)

// Me 拉取当前用户的信息
func (o *openAPI) Me(ctx context.Context) (*mydto.User, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.User{}).
		Get(o.getURL(userMeURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.User), nil
}

// MeGuilds 拉取当前用户加入的频道列表
func (o *openAPI) MeGuilds(ctx context.Context, pager *mydto.GuildPager) ([]*mydto.Guild, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(userMeGuildsURI))
	if err != nil {
		return nil, err
	}

	guilds := make([]*mydto.Guild, 0)
	if err := json.Unmarshal(resp.Body(), &guilds); err != nil {
		return nil, err
	}

	return guilds, nil
}
