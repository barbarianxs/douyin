
package rpc

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/message/messageservice"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var messageClient messageservice.Client

func initMessage() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := messageservice.NewClient(
		consts.MessageServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.ApiServiceName}),
	)
	if err != nil {
		panic(err)
	}
	messageClient = c
}

// MessageAction create note info
func MessageAction(ctx context.Context, req *message.MessageActionRequest) error {
	resp, err := messageClient.MessageAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.BaseResp.StatusCode != 0 {
		return errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return nil
}

// ChatMessage query list of note info
func MessageChat(ctx context.Context, req *message.MessageChatRequest) ([]*message.Message, int64, error) {
	resp, err := messageClient.MessageChat(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.BaseResp.StatusCode, resp.BaseResp.StatusMessage)
	}
	return resp.Messages, nil
}