package main

import (
	"context"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/main/cmd/relation/service"
	"github.com/YANGJUNYAN0715/douyin/tree/main/kitex_gen/relation"
	"log"

	"github.com/YANGJUNYAN0715/douyin/tree/main/pkg/errno"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	// TODO: Your code here...
	// TODO: Your code here...
	resp = new (relation.RelationActionResponse)
	// user:= ctx.Value(consts.IdentityKey)
	// user, _ := ctx.Get(consts.IdentityKey)	
	// user.id 是在api的rpc中通过解析token获取到的
	if req.UserId==0 ||req.ToUserId==0{
		// resp = pack.BuildBaseResp(errno.UserIDErr)
		resp.StatusCode = pack.BuildBaseResp(errno.UserIDErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.UserIDErr).StatusMsg
		return resp,nil
	}
	if req.UserId ==req.ToUserId{
		// resp = pack.BuildBaseResp(errno.FollowSelfErr)
		resp.StatusCode = pack.BuildBaseResp(errno.ActionTypeErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ActionTypeErr).StatusMsg
		return resp,nil
	}

	if req.ActionType <1 || req.ActionType >2{
		// resp = pack.BuildBaseResp(errno.ActionTypeErr)
		resp.StatusCode = pack.BuildBaseResp(errno.ActionTypeErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ActionTypeErr).StatusMsg
		return resp,nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)

	if err != nil {
		// resp = pack.BuildBaseResp(err)
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}
	// resp = pack.BuildBaseResp(errno.Success)
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new (relation.RelationFollowListResponse)
	// user:= ctx.Value(consts.IdentityKey)
	// user, _ := ctx.Get(consts.IdentityKey)
	if req.UserId==0 {
		// resp = pack.BuildBaseResp(errno.UserIDErr)
		resp.StatusCode = pack.BuildBaseResp(errno.UserIDErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.UserIDErr).StatusMsg
		return resp,nil
	}
	
	users,err := service.NewRelationListService(ctx).RelationFollowList(req)
	if err != nil{
		// resp = pack.BuildBaseResp(err)
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp,nil
	}
	// resp = pack.BuildBaseResp(errno.Success)
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.UserList = users
	// log.Println("***relation-handler.go***")
	// log.Println(users)
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new (relation.RelationFollowerListResponse)

	if req.UserId==0 {
		// resp = pack.BuildBaseResp(errno.UserIDErr)
		resp.StatusCode = pack.BuildBaseResp(errno.UserIDErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.UserIDErr).StatusMsg
		return resp,nil
	}
	

	users,err := service.NewRelationListService(ctx).RelationFollowerList(req)
	if err != nil{
		// resp = pack.BuildBaseResp(err)
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp,nil
	}
	// resp = pack.BuildBaseResp(errno.Success)
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.UserList = users

	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error) {
	// TODO: Your code here...
	resp = new (relation.RelationFriendListResponse)
	// user:= ctx.Value(consts.IdentityKey)
	if req.UserId==0 {
		// resp = pack.BuildBaseResp(errno.UserIDErr)
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp,nil
	}
	
	users,err := service.NewRelationListService(ctx).RelationFriendList(req)
	if err != nil{
		// resp = pack.BuildBaseResp(err)
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp,nil
	}
	// resp = pack.BuildBaseResp(errno.Success)
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.UserList = users

	return resp, nil
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
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
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

	err = service.NewActionMsgService(ctx).ActionMsg(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		log.Println(resp.StatusCode, "---------------******************----------", resp.StatusMsg)
		return resp, nil
	}
	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	
	return resp, nil
}
