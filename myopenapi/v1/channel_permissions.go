package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
	"fmt"
	"strconv"
)

// ChannelPermissions 获取指定子频道的权限
func (o *openAPI) ChannelPermissions(ctx context.Context, channelID, userID string) (*mydto.ChannelPermissions, error) {
	rsp, err := o.request(ctx).
		SetResult(mydto.ChannelPermissions{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("user_id", userID).
		Get(o.getURL(channelPermissionsURI))
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*mydto.ChannelPermissions), nil
}

// ChannelRolesPermissions 获取指定子频道身份组的权限
func (o *openAPI) ChannelRolesPermissions(ctx context.Context,
	channelID, roleID string) (*mydto.ChannelRolesPermissions, error) {
	rsp, err := o.request(ctx).
		SetResult(mydto.ChannelRolesPermissions{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("role_id", roleID).
		Get(o.getURL(channelRolesPermissionsURI))
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*mydto.ChannelRolesPermissions), nil
}

// PutChannelPermissions 修改指定子频道的权限
func (o *openAPI) PutChannelPermissions(ctx context.Context, channelID, userID string,
	p *mydto.UpdateChannelPermissions) error {
	if p.Add != "" {
		if _, err := strconv.ParseUint(p.Add, 10, 64); err != nil {
			return fmt.Errorf("invalid parameter add: %v", err)
		}
	}
	if p.Remove != "" {
		if _, err := strconv.ParseUint(p.Remove, 10, 64); err != nil {
			return fmt.Errorf("invalid parameter remove: %v", err)
		}
	}
	_, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("user_id", userID).
		SetBody(p).
		Put(o.getURL(channelPermissionsURI))
	return err
}

// PutChannelRolesPermissions 修改指定子频道的权限
func (o *openAPI) PutChannelRolesPermissions(ctx context.Context, channelID, roleID string,
	p *mydto.UpdateChannelPermissions) error {
	if p.Add != "" {
		if _, err := strconv.ParseUint(p.Add, 10, 64); err != nil {
			return fmt.Errorf("invalid parameter add: %v", err)
		}
	}
	if p.Remove != "" {
		if _, err := strconv.ParseUint(p.Remove, 10, 64); err != nil {
			return fmt.Errorf("invalid parameter remove: %v", err)
		}
	}
	_, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("role_id", roleID).
		SetBody(p).
		Put(o.getURL(channelRolesPermissionsURI))
	return err
}
