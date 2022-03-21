package v1

import (
	"XMrobot/mybotgo/errs"
	"XMrobot/mybotgo/mydto"
	"context"
	"encoding/json"
)

// Message 拉取单条消息
func (o *openAPI) Message(ctx context.Context, channelID string, messageID string) (*mydto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Message{}).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", messageID).
		Get(o.getURL(messageURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.Message), nil
}

// Messages 拉取消息列表
func (o *openAPI) Messages(ctx context.Context, channelID string, pager *mydto.MessagesPager) ([]*mydto.Message, error) {
	if pager == nil {
		return nil, errs.ErrPagerIsNil
	}
	resp, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetQueryParams(pager.QueryParams()).
		Get(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	messages := make([]*mydto.Message, 0)
	if err := json.Unmarshal(resp.Body(), &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

// PostMessage 发消息
func (o *openAPI) PostMessage(ctx context.Context, channelID string, msg *mydto.MessageToCreate) (*mydto.Message, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.Message{}).
		SetPathParam("channel_id", channelID).
		SetBody(msg).
		Post(o.getURL(messagesURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.Message), nil
}

// RetractMessage 撤回消息
func (o *openAPI) RetractMessage(ctx context.Context, channelID, msgID string) error {
	_, err := o.request(ctx).
		SetPathParam("channel_id", channelID).
		SetPathParam("message_id", string(msgID)).
		Delete(o.getURL(messageURI))
	return err
}
