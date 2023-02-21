package main

import (
	"context"
	// "fmt"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/pack"
	"github.com/YANGJUNYAN0715/douyin/tree/guo/cmd/user/service"

	"github.com/YANGJUNYAN0715/douyin/tree/guo/kitex_gen/user"


	"github.com/YANGJUNYAN0715/douyin/tree/guo/pkg/errno"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// LoginUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LoginUser(ctx context.Context, req *user.LoginUserRequest) (resp *user.LoginUserResponse, err error) {
	// TODO: Your code here...

	resp = new(user.LoginUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	resp.UserId = uid

	return resp, nil

}

// LogoutUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) LogoutUser(ctx context.Context, req *user.LogoutUserRequest) (resp *user.LogoutUserResponse, err error) {
	// TODO: Your code here...
	return
}

// RegisterUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) RegisterUser(ctx context.Context, req *user.RegisterUserRequest) (resp *user.RegisterUserResponse, err error) {
	// TODO: Your code here...
	resp = new(user.RegisterUserResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}
	err = service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	// resp.UserId = uid

	return resp, nil
}

// UserInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserInfoRequest) (resp *user.UserInfoResponse, err error) {
	// TODO: Your code here...

	resp = new(user.UserInfoResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	user_info, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg

	resp.User = user_info

	return resp, nil
}

// PublishAction implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishAction(ctx context.Context, req *user.PublishActionRequest) (resp *user.PublishActionResponse, err error) {
	// TODO: Your code here...
	resp = new(user.PublishActionResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	err = service.NewPublishActionService(ctx).PublishAction(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg


	return resp, nil
}

// PublishList implements the UserServiceImpl interface.
func (s *UserServiceImpl) PublishList(ctx context.Context, req *user.PublishListRequest) (resp *user.PublishListResponse, err error) {
	// TODO: Your code here...
	resp = new(user.PublishListResponse)

	if err = req.IsValid(); err != nil {
		resp.StatusCode = pack.BuildBaseResp(errno.ParamErr).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(errno.ParamErr).StatusMsg
		return resp, nil
	}

	videos_list, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp.StatusCode = pack.BuildBaseResp(err).StatusCode
		resp.StatusMsg = pack.BuildBaseResp(err).StatusMsg
		return resp, nil
	}

	resp.StatusCode = pack.BuildBaseResp(errno.Success).StatusCode
	resp.StatusMsg = pack.BuildBaseResp(errno.Success).StatusMsg
	resp.VideoList = videos_list

	return resp, nil
}
