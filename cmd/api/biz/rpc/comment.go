package rpc

import (
	"context"

	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment/commentsrv"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/errno"
	"github.com/YANGJUNYAN0715/douyin/tree/zhao/pkg/mw"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var cmtClient commentsrv.Client

func initComment() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.ApiServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := commentsrv.NewClient(
		consts.CommentServiceName,
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
	cmtClient = c
}

// Register create user info
func CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (*comment.Comment, error) {
	resp, err := cmtClient.CommentAction(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.Comment, nil
}

// Login check user info
func CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) ([]*comment.Comment, error) {
	resp, err := cmtClient.CommentList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.CommentList, nil
}
