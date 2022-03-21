package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// GetAPIPermissions 获取频道可用权限列表
func (o *openAPI) GetAPIPermissions(ctx context.Context, guildID string) (*mydto.APIPermissions, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.APIPermissions{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(apiPermissionURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.APIPermissions), nil
}

// RequireAPIPermissions 创建频道 API 接口权限授权链接
func (o *openAPI) RequireAPIPermissions(ctx context.Context,
	guildID string, demand *mydto.APIPermissionDemandToCreate) (*mydto.APIPermissionDemand, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.APIPermissionDemand{}).
		SetPathParam("guild_id", guildID).
		SetBody(demand).
		Post(o.getURL(apiPermissionDemandURI))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*mydto.APIPermissionDemand), nil
}
