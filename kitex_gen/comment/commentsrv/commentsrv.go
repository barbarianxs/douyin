// Code generated by Kitex v0.4.4. DO NOT EDIT.

package commentsrv

import (
	"context"
	comment "github.com/YANGJUNYAN0715/douyin/tree/zhao/kitex_gen/comment"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return commentSrvServiceInfo
}

var commentSrvServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "CommentSrv"
	handlerType := (*comment.CommentSrv)(nil)
	methods := map[string]kitex.MethodInfo{
		"CommentAction": kitex.NewMethodInfo(commentActionHandler, newCommentSrvCommentActionArgs, newCommentSrvCommentActionResult, false),
		"CommentList":   kitex.NewMethodInfo(commentListHandler, newCommentSrvCommentListArgs, newCommentSrvCommentListResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "comment",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.4.4",
		Extra:           extra,
	}
	return svcInfo
}

func commentActionHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentSrvCommentActionArgs)
	realResult := result.(*comment.CommentSrvCommentActionResult)
	success, err := handler.(comment.CommentSrv).CommentAction(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentSrvCommentActionArgs() interface{} {
	return comment.NewCommentSrvCommentActionArgs()
}

func newCommentSrvCommentActionResult() interface{} {
	return comment.NewCommentSrvCommentActionResult()
}

func commentListHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*comment.CommentSrvCommentListArgs)
	realResult := result.(*comment.CommentSrvCommentListResult)
	success, err := handler.(comment.CommentSrv).CommentList(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newCommentSrvCommentListArgs() interface{} {
	return comment.NewCommentSrvCommentListArgs()
}

func newCommentSrvCommentListResult() interface{} {
	return comment.NewCommentSrvCommentListResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) CommentAction(ctx context.Context, req *comment.DouyinCommentActionRequest) (r *comment.DouyinCommentActionResponse, err error) {
	var _args comment.CommentSrvCommentActionArgs
	_args.Req = req
	var _result comment.CommentSrvCommentActionResult
	if err = p.c.Call(ctx, "CommentAction", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) CommentList(ctx context.Context, req *comment.DouyinCommentListRequest) (r *comment.DouyinCommentListResponse, err error) {
	var _args comment.CommentSrvCommentListArgs
	_args.Req = req
	var _result comment.CommentSrvCommentListResult
	if err = p.c.Call(ctx, "CommentList", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}