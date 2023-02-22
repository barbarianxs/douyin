package rpc

import (
	"context"
	"log"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/user/userservice"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func Init() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName,
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
	userClient = c
}


func Info(ctx context.Context, req *user.UserInfoRequest) (*user.User, error) {
	resp, err := userClient.UserInfo(ctx, req)
	if err != nil {
		return resp.User, err
	}
	if resp.StatusCode != 0 {
		return resp.User, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	log.Println(resp.User)
	return resp.User, nil
}