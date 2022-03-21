package v1

import (
	"XMrobot/mybotgo/errs"
	"XMrobot/mybotgo/mydto"
	"context"
	"encoding/json"
)

// MemberAddRole 添加成员角色
func (o *openAPI) MemberAddRole(
	ctx context.Context, guildID string, roleID mydto.RoleID, userID string,
	value *mydto.MemberAddRoleBody,
) error {
	if value == nil {
		value = new(mydto.MemberAddRoleBody)
	}
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetPathParam("user_id", userID).
		SetBody(value).
		Put(o.getURL(memberRoleURI))
	return err
}

// MemberDeleteRole 删除成员角色
func (o *openAPI) MemberDeleteRole(
	ctx context.Context, guildID string, roleID mydto.RoleID, userID string,
	value *mydto.MemberAddRoleBody,
) error {
	if value == nil {
		value = new(mydto.MemberAddRoleBody)
	}
	_, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetPathParam("role_id", string(roleID)).
		SetPathParam("user_id", userID).
		SetBody(value).
		Delete(o.getURL(memberRoleURI))
	return err
}

// GuildMember 拉取频道指定成员
func (o *openAPI) GuildMember(ctx context.Context, guildID, userID string) (*mydto.Member, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Member{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		Get(o.getURL(guildMemberURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.Member), nil
}

// GuildMembers 分页拉取频道内成员列表
func (o *openAPI) GuildMembers(
	ctx context.Context,
	guildID string, pager *mydto.GuildMembersPager,
) ([]*mydto.Member, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetPathParam("guild_id", guildID).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(guildMembersURI))
	if err != nil {
		return nil, err
	}

	members := make([]*mydto.Member, 0)
	if err := json.Unmarshal(resp.Body(), &members); err != nil {
		return nil, err
	}

	return members, nil
}

// DeleteGuildMember 将指定成员踢出频道
func (o *openAPI) DeleteGuildMember(ctx context.Context, guildID, userID string, opts ...mydto.MemberDeleteOption) error {
	opt := &mydto.MemberDeleteOpts{}
	for _, o := range opts {
		o(opt)
	}
	_, err := o.request(ctx).
		SetResult(mydto.Member{}).
		SetPathParam("guild_id", guildID).
		SetPathParam("user_id", userID).
		SetBody(opt).
		Delete(o.getURL(guildMemberURI))
	return err
}
