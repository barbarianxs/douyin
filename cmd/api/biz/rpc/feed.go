package rpc

import (
	"context"
	// "log"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/feed"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/feed/feedservice"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var feedClient feedservice.Client

func initFeed() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := feedservice.NewClient(
		consts.FeedServiceName,
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
	feedClient = c
}

func GetUserFeed(ctx context.Context, req *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	resp, err = feedClient.GetUserFeed(ctx, req)
	// if err != nil {
	// 	return nil, err
	// }
	// if resp.StatusCode != 0 {
	// 	return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	// }
	return resp, nil
}