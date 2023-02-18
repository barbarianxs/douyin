package pack

import (
	"errors"

	"github.com/YANGJUNYAN0715/douyin/tree/li/kitex_gen/relation"
	"github.com/YANGJUNYAN0715/douyin/tree/li/pkg/errno"
)

// 接口 打包 DouyinRelationActionResponse
func BuildDouyinRelationActionResponse(err error) *relation.DouyinRelationActionResponse {
	if err == nil {
		return DouyinRelationActionResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinRelationActionResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinRelationActionResponse(s)
}

func DouyinRelationActionResponse(err errno.ErrNo) *relation.DouyinRelationActionResponse {
	return &relation.DouyinRelationActionResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// 接口 打包 DouyinRelationFollowListResponse
func BuildDouyinRelationFollowListResponse(err error) *relation.DouyinRelationFollowListResponse {
	if err == nil {
		return DouyinRelationFollowListResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinRelationFollowListResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinRelationFollowListResponse(s)
}

func DouyinRelationFollowListResponse(err errno.ErrNo) *relation.DouyinRelationFollowListResponse {
	return &relation.DouyinRelationFollowListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// 接口 打包 DouyinRelationFollowerListResponse
func BuildDouyinRelationFollowerListResponse(err error) *relation.DouyinRelationFollowerListResponse {
	if err == nil {
		return DouyinRelationFollowerListResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinRelationFollowerListResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinRelationFollowerListResponse(s)
}

func DouyinRelationFollowerListResponse(err errno.ErrNo) *relation.DouyinRelationFollowerListResponse {
	return &relation.DouyinRelationFollowerListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}

// 接口 打包 DouyinRelationFriendListResponse
func BuildDouyinRelationFriendListResponse(err error) *relation.DouyinRelationFriendListResponse {
	if err == nil {
		return DouyinRelationFriendListResponse(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return DouyinRelationFriendListResponse(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return DouyinRelationFriendListResponse(s)
}

func DouyinRelationFriendListResponse(err errno.ErrNo) *relation.DouyinRelationFriendListResponse {
	return &relation.DouyinRelationFriendListResponse{StatusCode: int32(err.ErrCode), StatusMsg: err.ErrMsg}
}
