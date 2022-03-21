package v1

import (
	"XMrobot/mybotgo/mydto"
	"context"
)

// WS 获取带分片 WSS 接入点
func (o *openAPI) WS(ctx context.Context, _ map[string]string, _ string) (*mydto.WebsocketAP, error) {
	resp, err := o.request(ctx).
		SetResult(mydto.WebsocketAP{}).
		Get(o.getURL(gatewayBotURI))
	if err != nil {
		return nil, err
	}

	return resp.Result().(*mydto.WebsocketAP), nil
}
