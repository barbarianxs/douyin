package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/relation/service"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/relation"


	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageChat(ctx context.Context, req *relation.MessageChatRequest) (resp *relation.MessageChatResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.MessageChatResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	messages, err := service.NewChatMsgService(ctx).MGetChatMsg(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	resp.Messages = messages
	resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
	return resp, nil
}

// MessageAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) MessageAction(ctx context.Context, req *relation.MessageActionRequest) (resp *relation.MessageActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.MessageActionResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	err = service.NewActionMsgService(ctx).MGetActionMsg(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
	return resp, nil
}
