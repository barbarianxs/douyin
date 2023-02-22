
package rpc

import (
	"context"
	// "log"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/interact/interactservice"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var interactClient interactservice.Client


func initinteract() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.InteractServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := interactservice.NewClient(
		consts.InteractServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware),
		client.WithInstanceMW(mw.ClientMiddleware),
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.InteractServiceName}),
	)
	if err != nil {
		panic(err)
	}
	interactClient = c
}


// FavoriteAction check interact info
func FavoriteAction(ctx context.Context, req *interact.FavoriteActionRequest) error {
	// log.Println("===============================================================================")
	resp, err := interactClient.FavoriteAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return nil
}

// FavoriteList query list of note info
func FavoriteList(ctx context.Context, req *interact.FavoriteListRequest) ([]*interact.Video, error) {
	resp, err := interactClient.FavoriteList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.VideoList, nil
}

// Register create user info
func CommentAction(ctx context.Context, req *interact.CommentActionRequest) (*interact.Comment, error) {
	resp, err := interactClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.Comment, nil
}

// Login check user info
func CommentList(ctx context.Context, req *interact.CommentListRequest) ([]*interact.Comment, error) {
	resp, err := interactClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, resp.StatusMsg)
	}
	return resp.CommentList, nil
}
