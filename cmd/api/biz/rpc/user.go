package rpc

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/user"
	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/user/userservice"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
	"log"
)

var userClient userservice.Client

func initUser() {
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

// Register create user info
func Register(ctx context.Context, req *user.DouyinUserRegisterRequest) error {
	resp, err := userClient.Register(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// Login check user info
func Login(ctx context.Context, req *user.DouyinUserRegisterRequest) (int64, error) {
	resp, err := userClient.Login(ctx, req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 0 {
		return 0, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.UserId, nil
}

func Info(ctx context.Context, req *user.DouyinUserRequest) (*user.User, error) {
	resp, err := userClient.GetUserById(ctx, req)
	if err != nil {
		return resp.User, err
	}
	if resp.StatusCode != 0 {
		return resp.User, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	log.Println("***api-rpc***")
	log.Println(resp.User)
	return resp.User, nil
}
