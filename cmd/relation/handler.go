package main

import (
	relation "github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/relation"
	// "github.com/YANGJUNYAN0715/douyin/tree/li/pkg/consts"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/relation/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/li/cmd/relation/service"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"

	"context"
	// "log"
)

// RelationServiceImpl implements the last service interface defined in the IDL.
type RelationServiceImpl struct{}

// RelationAction implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.DouyinRelationActionRequest) (resp *relation.DouyinRelationActionResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationActionResponse)
	// user:= ctx.Value(consts.IdentityKey)
	// user, _ := ctx.Get(consts.IdentityKey)
	// user.id 是在api的rpc中通过解析token获取到的
	if req.UserId == 0 || req.ToUserId == 0 {
		resp = pack.BuildDouyinRelationActionResponse(errno.UserIDErr)
		return resp, nil
	}
	if req.UserId == req.ToUserId {
		resp = pack.BuildDouyinRelationActionResponse(errno.FollowSelfErr)
		return resp, nil
	}

	if req.ActionType < 1 || req.ActionType > 2 {
		resp = pack.BuildDouyinRelationActionResponse(errno.ActionTypeErr)
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)

	if err != nil {
		resp = pack.BuildDouyinRelationActionResponse(err)
		return resp, nil
	}
	resp = pack.BuildDouyinRelationActionResponse(errno.Success)

	return resp, nil
}

// RelationFollowList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.DouyinRelationFollowListRequest) (resp *relation.DouyinRelationFollowListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFollowListResponse)
	// user:= ctx.Value(consts.IdentityKey)
	// user, _ := ctx.Get(consts.IdentityKey)
	if req.UserId == 0 {
		resp = pack.BuildDouyinRelationFollowListResponse(errno.UserIDErr)
		return resp, nil
	}

	users, err := service.NewRelationListService(ctx).RelationFollowList(req)
	if err != nil {
		resp = pack.BuildDouyinRelationFollowListResponse(err)
		return resp, nil
	}
	resp = pack.BuildDouyinRelationFollowListResponse(errno.Success)
	resp.UserList = users
	// log.Println("***relation-handler.go***")
	// log.Println(users)
	return resp, nil
}

// RelationFollowerList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.DouyinRelationFollowerListRequest) (resp *relation.DouyinRelationFollowerListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFollowerListResponse)

	if req.UserId == 0 {
		resp = pack.BuildDouyinRelationFollowerListResponse(errno.UserIDErr)
		return resp, nil
	}

	users, err := service.NewRelationListService(ctx).RelationFollowerList(req)
	if err != nil {
		resp = pack.BuildDouyinRelationFollowerListResponse(err)
		return resp, nil
	}
	resp = pack.BuildDouyinRelationFollowerListResponse(errno.Success)
	resp.UserList = users

	return resp, nil
}

// RelationFriendList implements the RelationServiceImpl interface.
func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.DouyinRelationFriendListRequest) (resp *relation.DouyinRelationFriendListResponse, err error) {
	// TODO: Your code here...
	resp = new(relation.DouyinRelationFriendListResponse)
	// user:= ctx.Value(consts.IdentityKey)
	if req.UserId == 0 {
		resp = pack.BuildDouyinRelationFriendListResponse(errno.UserIDErr)
		return resp, nil
	}

	users, err := service.NewRelationListService(ctx).RelationFriendList(req)
	if err != nil {
		resp = pack.BuildDouyinRelationFriendListResponse(err)
		return resp, nil
	}
	resp = pack.BuildDouyinRelationFriendListResponse(errno.Success)
	resp.UserList = users

	return resp, nil
}
