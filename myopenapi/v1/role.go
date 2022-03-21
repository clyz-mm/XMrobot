package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
	"fmt"
)

func (o *openAPI) Roles(ctx context.Context, guildID string) (*mydto.GuildRoles, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.GuildRoles{}).
		SetPathParam("guild_id", guildID).
		Get(o.getURL(rolesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.GuildRoles), nil
}

func (o *openAPI) PostRole(ctx context.Context, guildID string, role *mydto.Role) (*mydto.UpdateResult, error) {
	if role.Color == 0 {
		role.Color = mydto.DefaultColor
	}
	// openapi 上修改哪个字段，就需要传递哪个 filter
	filter := &mydto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &mydto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	fmt.Sprint(body)
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetResult(mydto.UpdateResult{}).
		SetBody(body).
		Post(o.getURL(rolesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.UpdateResult), nil
}

func (o *openAPI) PatchRole(ctx context.Context,
	guildID string, roleID mydto.RoleID, role *mydto.Role) (*mydto.UpdateResult, error) {
	if role.Color == 0 {
		role.Color = mydto.DefaultColor
	}
	filter := &mydto.UpdateRoleFilter{
		Name:  1,
		Color: 1,
		Hoist: 1,
	}
	body := &mydto.UpdateRole{
		GuildID: guildID,
		Filter:  filter,
		Update:  role,
	}
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetResult(mydto.UpdateResult{}).
		SetBody(body).
		Patch(o.getURL(roleURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.UpdateResult), nil
}

func (o *openAPI) DeleteRole(ctx context.Context, guildID string, roleID mydto.RoleID) error {
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		Delete(o.getURL(roleURI))
	return err
}
