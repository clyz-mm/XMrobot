package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
	"strconv"
)

// ListSchedules 查询某个子频道下，since开始的当天的日程列表。若since为0，默认返回当天的日程列表
func (o *openAPI) ListSchedules(ctx context.Context, channelID string, since uint64) ([]*mydto.Schedule, error) {
	rsp, err := o.request(ctx).
		SetResult([]*mydto.Schedule{}).
		SetPathParam("channel_id", channelID).
		SetQueryParam("since", strconv.FormatUint(since, 10)).
		Get(o.getURL(schedulesURI))
	if err != nil {
		return nil, err
	}
	return *rsp.Result().(*[]*mydto.Schedule), nil
}

// GetSchedule 获取单个日程信息
func (o *openAPI) GetSchedule(ctx context.Context, channelID, scheduleID string) (*mydto.Schedule, error) {
	rsp, err := o.request(ctx).
		SetResult(mydto.Schedule{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("schedule_id", scheduleID).
		Get(o.getURL(scheduleURI))
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*mydto.Schedule), nil
}

// CreateSchedule 创建日程
func (o *openAPI) CreateSchedule(ctx context.Context, channelID string, schedule *mydto.Schedule) (*mydto.Schedule, error) {
	rsp, err := o.request(ctx).
		SetResult(mydto.Schedule{}).
		SetPathParam("channel_id", channelID).
		SetBody(mydto.ScheduleWrapper{Schedule: schedule}).
		Post(o.getURL(schedulesURI))
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*mydto.Schedule), nil
}

// ModifySchedule 修改日程
func (o *openAPI) ModifySchedule(ctx context.Context,
	channelID, scheduleID string, schedule *mydto.Schedule) (*mydto.Schedule, error) {
	rsp, err := o.request(ctx).
		SetResult(mydto.Schedule{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("schedule_id", scheduleID).
		SetBody(mydto.ScheduleWrapper{Schedule: schedule}).
		Patch(o.getURL(scheduleURI))
	if err != nil {
		return nil, err
	}
	return rsp.Result().(*mydto.Schedule), nil
}

// DeleteSchedule 删除日程
func (o *openAPI) DeleteSchedule(ctx context.Context, channelID, scheduleID string) error {
	_, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("schedule_id", scheduleID).
		Delete(o.getURL(scheduleURI))
	return err
}
